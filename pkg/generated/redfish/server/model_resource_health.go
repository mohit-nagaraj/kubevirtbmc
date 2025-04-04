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

type ResourceHealth string

// List of ResourceHealth
const (
	RESOURCEHEALTH_OK       ResourceHealth = "OK"
	RESOURCEHEALTH_WARNING  ResourceHealth = "Warning"
	RESOURCEHEALTH_CRITICAL ResourceHealth = "Critical"
)

// AllowedResourceHealthEnumValues is all the allowed values of ResourceHealth enum
var AllowedResourceHealthEnumValues = []ResourceHealth{
	"OK",
	"Warning",
	"Critical",
}

// validResourceHealthEnumValue provides a map of ResourceHealths for fast verification of use input
var validResourceHealthEnumValues = map[ResourceHealth]struct{}{
	"OK":       {},
	"Warning":  {},
	"Critical": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceHealth) IsValid() bool {
	_, ok := validResourceHealthEnumValues[v]
	return ok
}

// NewResourceHealthFromValue returns a pointer to a valid ResourceHealth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceHealthFromValue(v string) (ResourceHealth, error) {
	ev := ResourceHealth(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ResourceHealth: valid values are %v", v, AllowedResourceHealthEnumValues)
}

// AssertResourceHealthRequired checks if the required fields are not zero-ed
func AssertResourceHealthRequired(obj ResourceHealth) error {
	return nil
}

// AssertResourceHealthConstraints checks if the values respects the defined constraints
func AssertResourceHealthConstraints(obj ResourceHealth) error {
	return nil
}
