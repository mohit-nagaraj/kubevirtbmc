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

type CircuitBreakerStates string

// List of CircuitBreakerStates
const (
	CIRCUITBREAKERSTATES_NORMAL  CircuitBreakerStates = "Normal"
	CIRCUITBREAKERSTATES_TRIPPED CircuitBreakerStates = "Tripped"
	CIRCUITBREAKERSTATES_OFF     CircuitBreakerStates = "Off"
)

// AllowedCircuitBreakerStatesEnumValues is all the allowed values of CircuitBreakerStates enum
var AllowedCircuitBreakerStatesEnumValues = []CircuitBreakerStates{
	"Normal",
	"Tripped",
	"Off",
}

// validCircuitBreakerStatesEnumValue provides a map of CircuitBreakerStatess for fast verification of use input
var validCircuitBreakerStatesEnumValues = map[CircuitBreakerStates]struct{}{
	"Normal":  {},
	"Tripped": {},
	"Off":     {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CircuitBreakerStates) IsValid() bool {
	_, ok := validCircuitBreakerStatesEnumValues[v]
	return ok
}

// NewCircuitBreakerStatesFromValue returns a pointer to a valid CircuitBreakerStates
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCircuitBreakerStatesFromValue(v string) (CircuitBreakerStates, error) {
	ev := CircuitBreakerStates(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for CircuitBreakerStates: valid values are %v", v, AllowedCircuitBreakerStatesEnumValues)
}

// AssertCircuitBreakerStatesRequired checks if the required fields are not zero-ed
func AssertCircuitBreakerStatesRequired(obj CircuitBreakerStates) error {
	return nil
}

// AssertCircuitBreakerStatesConstraints checks if the values respects the defined constraints
func AssertCircuitBreakerStatesConstraints(obj CircuitBreakerStates) error {
	return nil
}
