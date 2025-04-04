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

type ThermalV172ReadingUnits string

// List of ThermalV172ReadingUnits
const (
	THERMALV172READINGUNITS_RPM     ThermalV172ReadingUnits = "RPM"
	THERMALV172READINGUNITS_PERCENT ThermalV172ReadingUnits = "Percent"
)

// AllowedThermalV172ReadingUnitsEnumValues is all the allowed values of ThermalV172ReadingUnits enum
var AllowedThermalV172ReadingUnitsEnumValues = []ThermalV172ReadingUnits{
	"RPM",
	"Percent",
}

// validThermalV172ReadingUnitsEnumValue provides a map of ThermalV172ReadingUnitss for fast verification of use input
var validThermalV172ReadingUnitsEnumValues = map[ThermalV172ReadingUnits]struct{}{
	"RPM":     {},
	"Percent": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ThermalV172ReadingUnits) IsValid() bool {
	_, ok := validThermalV172ReadingUnitsEnumValues[v]
	return ok
}

// NewThermalV172ReadingUnitsFromValue returns a pointer to a valid ThermalV172ReadingUnits
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewThermalV172ReadingUnitsFromValue(v string) (ThermalV172ReadingUnits, error) {
	ev := ThermalV172ReadingUnits(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ThermalV172ReadingUnits: valid values are %v", v, AllowedThermalV172ReadingUnitsEnumValues)
}

// AssertThermalV172ReadingUnitsRequired checks if the required fields are not zero-ed
func AssertThermalV172ReadingUnitsRequired(obj ThermalV172ReadingUnits) error {
	return nil
}

// AssertThermalV172ReadingUnitsConstraints checks if the values respects the defined constraints
func AssertThermalV172ReadingUnitsConstraints(obj ThermalV172ReadingUnits) error {
	return nil
}
