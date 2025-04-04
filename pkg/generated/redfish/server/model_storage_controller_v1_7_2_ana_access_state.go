// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

import (
	"fmt"
)

type StorageControllerV172AnaAccessState string

// List of StorageControllerV172AnaAccessState
const (
	STORAGECONTROLLERV172ANAACCESSSTATE_OPTIMIZED       StorageControllerV172AnaAccessState = "Optimized"
	STORAGECONTROLLERV172ANAACCESSSTATE_NON_OPTIMIZED   StorageControllerV172AnaAccessState = "NonOptimized"
	STORAGECONTROLLERV172ANAACCESSSTATE_INACCESSIBLE    StorageControllerV172AnaAccessState = "Inaccessible"
	STORAGECONTROLLERV172ANAACCESSSTATE_PERSISTENT_LOSS StorageControllerV172AnaAccessState = "PersistentLoss"
)

// AllowedStorageControllerV172AnaAccessStateEnumValues is all the allowed values of StorageControllerV172AnaAccessState enum
var AllowedStorageControllerV172AnaAccessStateEnumValues = []StorageControllerV172AnaAccessState{
	"Optimized",
	"NonOptimized",
	"Inaccessible",
	"PersistentLoss",
}

// validStorageControllerV172AnaAccessStateEnumValue provides a map of StorageControllerV172AnaAccessStates for fast verification of use input
var validStorageControllerV172AnaAccessStateEnumValues = map[StorageControllerV172AnaAccessState]struct{}{
	"Optimized":      {},
	"NonOptimized":   {},
	"Inaccessible":   {},
	"PersistentLoss": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageControllerV172AnaAccessState) IsValid() bool {
	_, ok := validStorageControllerV172AnaAccessStateEnumValues[v]
	return ok
}

// NewStorageControllerV172AnaAccessStateFromValue returns a pointer to a valid StorageControllerV172AnaAccessState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageControllerV172AnaAccessStateFromValue(v string) (StorageControllerV172AnaAccessState, error) {
	ev := StorageControllerV172AnaAccessState(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for StorageControllerV172AnaAccessState: valid values are %v", v, AllowedStorageControllerV172AnaAccessStateEnumValues)
}

// AssertStorageControllerV172AnaAccessStateRequired checks if the required fields are not zero-ed
func AssertStorageControllerV172AnaAccessStateRequired(obj StorageControllerV172AnaAccessState) error {
	return nil
}

// AssertStorageControllerV172AnaAccessStateConstraints checks if the values respects the defined constraints
func AssertStorageControllerV172AnaAccessStateConstraints(obj StorageControllerV172AnaAccessState) error {
	return nil
}
