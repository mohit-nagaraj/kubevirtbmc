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

type NetworkDeviceFunctionV191AuthenticationMethod string

// List of NetworkDeviceFunctionV191AuthenticationMethod
const (
	NETWORKDEVICEFUNCTIONV191AUTHENTICATIONMETHOD_NONE        NetworkDeviceFunctionV191AuthenticationMethod = "None"
	NETWORKDEVICEFUNCTIONV191AUTHENTICATIONMETHOD_CHAP        NetworkDeviceFunctionV191AuthenticationMethod = "CHAP"
	NETWORKDEVICEFUNCTIONV191AUTHENTICATIONMETHOD_MUTUAL_CHAP NetworkDeviceFunctionV191AuthenticationMethod = "MutualCHAP"
)

// AllowedNetworkDeviceFunctionV191AuthenticationMethodEnumValues is all the allowed values of NetworkDeviceFunctionV191AuthenticationMethod enum
var AllowedNetworkDeviceFunctionV191AuthenticationMethodEnumValues = []NetworkDeviceFunctionV191AuthenticationMethod{
	"None",
	"CHAP",
	"MutualCHAP",
}

// validNetworkDeviceFunctionV191AuthenticationMethodEnumValue provides a map of NetworkDeviceFunctionV191AuthenticationMethods for fast verification of use input
var validNetworkDeviceFunctionV191AuthenticationMethodEnumValues = map[NetworkDeviceFunctionV191AuthenticationMethod]struct{}{
	"None":       {},
	"CHAP":       {},
	"MutualCHAP": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkDeviceFunctionV191AuthenticationMethod) IsValid() bool {
	_, ok := validNetworkDeviceFunctionV191AuthenticationMethodEnumValues[v]
	return ok
}

// NewNetworkDeviceFunctionV191AuthenticationMethodFromValue returns a pointer to a valid NetworkDeviceFunctionV191AuthenticationMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkDeviceFunctionV191AuthenticationMethodFromValue(v string) (NetworkDeviceFunctionV191AuthenticationMethod, error) {
	ev := NetworkDeviceFunctionV191AuthenticationMethod(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for NetworkDeviceFunctionV191AuthenticationMethod: valid values are %v", v, AllowedNetworkDeviceFunctionV191AuthenticationMethodEnumValues)
}

// AssertNetworkDeviceFunctionV191AuthenticationMethodRequired checks if the required fields are not zero-ed
func AssertNetworkDeviceFunctionV191AuthenticationMethodRequired(obj NetworkDeviceFunctionV191AuthenticationMethod) error {
	return nil
}

// AssertNetworkDeviceFunctionV191AuthenticationMethodConstraints checks if the values respects the defined constraints
func AssertNetworkDeviceFunctionV191AuthenticationMethodConstraints(obj NetworkDeviceFunctionV191AuthenticationMethod) error {
	return nil
}
