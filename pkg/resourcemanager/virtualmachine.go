package resourcemanager

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"

	kubevirttypev1 "kubevirt.io/kubevirtbmc/pkg/generated/clientset/versioned/typed/core/v1"
	"kubevirt.io/kubevirtbmc/pkg/generated/redfish/server"
)

const (
	defaultManagerId        = "BMC"
	defaultManagerName      = "Manager"
	defaultComputerSystemId = "1"
)

var (
	powerStateMap = map[bool]server.ResourcePowerState{
		true:  server.RESOURCEPOWERSTATE_ON,
		false: server.RESOURCEPOWERSTATE_OFF,
	}
	bootSourceMap = map[BootDevice]server.ComputerSystemBootSource{
		BootDevicePxe: server.COMPUTERSYSTEMBOOTSOURCE_PXE,
		BootDeviceHdd: server.COMPUTERSYSTEMBOOTSOURCE_HDD,
	}
)

type KubeVirtClientInterface interface {
	VirtualMachines(namespace string) kubevirttypev1.VirtualMachineInterface
	VirtualMachineInstances(namespace string) kubevirttypev1.VirtualMachineInstanceInterface
}

type VirtualMachineResourceManager struct {
	ctx      context.Context
	kvClient KubeVirtClientInterface

	namespace string
	name      string

	computerSystem *ComputerSystemAdapter
	manager        *ManagerAdapter
}

func NewVirtualMachineResourceManager(
	ctx context.Context,
	kvClient KubeVirtClientInterface,
) *VirtualMachineResourceManager {
	return &VirtualMachineResourceManager{
		ctx:      ctx,
		kvClient: kvClient,
	}
}

func (m *VirtualMachineResourceManager) Initialize(namespace, name string) error {
	vm, err := m.kvClient.VirtualMachines(namespace).Get(m.ctx, name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	m.namespace = vm.Namespace
	m.name = vm.Name

	// Initialize computer system
	m.computerSystem = NewComputerSystem(
		defaultComputerSystemId,
		strings.Join([]string{vm.Namespace, vm.Name}, "/"),
		powerStateMap[vm.Status.Ready],
	)

	// Initialize manager
	m.manager = NewManager(defaultManagerId, defaultManagerName)

	// Build relationships
	var (
		oDataManager        ODataInterface = m.manager
		oDataComputerSystem ODataInterface = m.computerSystem
	)
	if err := oDataComputerSystem.ManagedBy(oDataManager); err != nil {
		return err
	}
	if err := oDataManager.Manage(oDataComputerSystem); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineResourceManager) GetComputerSystem() (ComputerSystemInterface, error) {
	if m.computerSystem == nil {
		return nil, fmt.Errorf("computer system not initialized")
	}

	// Update the power state just-in-time until we actually implement a control loop for it
	vm, err := m.kvClient.VirtualMachines(m.namespace).
		Get(m.ctx, m.name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	switch vm.Status.Ready {
	case true:
		m.computerSystem.SetPowerState(server.RESOURCEPOWERSTATE_ON)
	case false:
		m.computerSystem.SetPowerState(server.RESOURCEPOWERSTATE_OFF)
	}

	return m.computerSystem, nil
}

func (m *VirtualMachineResourceManager) GetManager() (ManagerInterface, error) {
	return m.manager, nil
}

func (m *VirtualMachineResourceManager) GetPowerStatus() (bool, error) {
	// TODO: Implement a control loop to keep the power state in sync, then we will be able to
	// return the power state from the intermediate object, i.e. ComputerSystem.
	//
	// ps := m.computerSystem.GetPowerState()
	// switch ps {
	// case server.RESOURCEPOWERSTATE_ON, server.RESOURCEPOWERSTATE_POWERING_ON:
	// 	return true, nil
	// case server.RESOURCEPOWERSTATE_OFF, server.RESOURCEPOWERSTATE_POWERING_OFF:
	// 	return false, nil
	// default:
	// 	return false, nil
	// }
	vm, err := m.kvClient.VirtualMachines(m.namespace).
		Get(m.ctx, m.name, metav1.GetOptions{})
	if err != nil {
		return false, err
	}

	return vm.Status.Ready, nil
}

func (m *VirtualMachineResourceManager) PowerOn() error {
	vm, err := m.kvClient.VirtualMachines(m.namespace).
		Get(m.ctx, m.name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if vm.Spec.RunStrategy == nil {
		running := func(b bool) *bool { return &b }(true)
		vm.Spec.Running = running
	} else {
		runStrategy := func(
			rs kubevirtv1.VirtualMachineRunStrategy,
		) *kubevirtv1.VirtualMachineRunStrategy {
			return &rs
		}(kubevirtv1.RunStrategyRerunOnFailure)
		vm.Spec.RunStrategy = runStrategy
	}
	if _, err := m.kvClient.VirtualMachines(m.namespace).
		Update(m.ctx, vm, metav1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (m *VirtualMachineResourceManager) PowerOff() error {
	vm, err := m.kvClient.VirtualMachines(m.namespace).
		Get(m.ctx, m.name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if vm.Spec.RunStrategy == nil {
		running := func(b bool) *bool { return &b }(false)
		vm.Spec.Running = running
	} else {
		runStrategy := func(rs kubevirtv1.VirtualMachineRunStrategy) *kubevirtv1.VirtualMachineRunStrategy {
			return &rs
		}(kubevirtv1.RunStrategyHalted)
		vm.Spec.RunStrategy = runStrategy
	}
	if _, err := m.kvClient.VirtualMachines(m.namespace).
		Update(m.ctx, vm, metav1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (m *VirtualMachineResourceManager) PowerCycle() error {
	return m.kvClient.VirtualMachineInstances(m.namespace).
		Delete(m.ctx, m.name, metav1.DeleteOptions{})
}

func (m *VirtualMachineResourceManager) SetBootDevice(bootDevice BootDevice) error {
	logrus.Info("SetBootDevice")
	vm, err := m.kvClient.VirtualMachines(m.namespace).
		Get(m.ctx, m.name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if vm.Spec.Template == nil {
		return fmt.Errorf("no template found")
	}

	for i, intf := range vm.Spec.Template.Spec.Domain.Devices.Interfaces {
		logrus.Infof("interface: %+v", intf)
		vm.Spec.Template.Spec.Domain.Devices.Interfaces[i].BootOrder = nil
	}
	for i, dev := range vm.Spec.Template.Spec.Domain.Devices.Disks {
		logrus.Infof("disk: %+v", dev)
		vm.Spec.Template.Spec.Domain.Devices.Disks[i].BootOrder = nil
	}

	var firstOrder uint = 1
	switch bootDevice {
	case BootDevicePxe:
		if vm.Spec.Template.Spec.Domain.Devices.Interfaces == nil {
			return fmt.Errorf("no interfaces found")
		}
		vm.Spec.Template.Spec.Domain.Devices.Interfaces[0].BootOrder = &firstOrder
		logrus.Infof("To be updated vm: %+v", vm.Spec.Template.Spec.Domain.Devices.Interfaces[0])
	case BootDeviceHdd:
		if vm.Spec.Template.Spec.Domain.Devices.Disks == nil {
			return fmt.Errorf("no disks found")
		}
		vm.Spec.Template.Spec.Domain.Devices.Disks[0].BootOrder = &firstOrder
		logrus.Infof("To be updated vm: %+v", vm.Spec.Template.Spec.Domain.Devices.Disks[0])
	}

	if _, err := m.kvClient.VirtualMachines(m.namespace).
		Update(m.ctx, vm, metav1.UpdateOptions{}); err != nil {
		logrus.Errorf("update vm error: %v", err)
		return err
	}

	if m.computerSystem == nil {
		logrus.Warn("computer system not initialized")
		return nil
	}
	m.computerSystem.SetBootOverride(bootSourceMap[bootDevice])

	return nil
}
