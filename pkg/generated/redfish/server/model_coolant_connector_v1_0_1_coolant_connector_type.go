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

type CoolantConnectorV101CoolantConnectorType string

// List of CoolantConnectorV101CoolantConnectorType
const (
	COOLANTCONNECTORV101COOLANTCONNECTORTYPE_PAIR   CoolantConnectorV101CoolantConnectorType = "Pair"
	COOLANTCONNECTORV101COOLANTCONNECTORTYPE_SUPPLY CoolantConnectorV101CoolantConnectorType = "Supply"
	COOLANTCONNECTORV101COOLANTCONNECTORTYPE_RETURN CoolantConnectorV101CoolantConnectorType = "Return"
	COOLANTCONNECTORV101COOLANTCONNECTORTYPE_INLINE CoolantConnectorV101CoolantConnectorType = "Inline"
	COOLANTCONNECTORV101COOLANTCONNECTORTYPE_CLOSED CoolantConnectorV101CoolantConnectorType = "Closed"
)

// AllowedCoolantConnectorV101CoolantConnectorTypeEnumValues is all the allowed values of CoolantConnectorV101CoolantConnectorType enum
var AllowedCoolantConnectorV101CoolantConnectorTypeEnumValues = []CoolantConnectorV101CoolantConnectorType{
	"Pair",
	"Supply",
	"Return",
	"Inline",
	"Closed",
}

// validCoolantConnectorV101CoolantConnectorTypeEnumValue provides a map of CoolantConnectorV101CoolantConnectorTypes for fast verification of use input
var validCoolantConnectorV101CoolantConnectorTypeEnumValues = map[CoolantConnectorV101CoolantConnectorType]struct{}{
	"Pair":   {},
	"Supply": {},
	"Return": {},
	"Inline": {},
	"Closed": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CoolantConnectorV101CoolantConnectorType) IsValid() bool {
	_, ok := validCoolantConnectorV101CoolantConnectorTypeEnumValues[v]
	return ok
}

// NewCoolantConnectorV101CoolantConnectorTypeFromValue returns a pointer to a valid CoolantConnectorV101CoolantConnectorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCoolantConnectorV101CoolantConnectorTypeFromValue(v string) (CoolantConnectorV101CoolantConnectorType, error) {
	ev := CoolantConnectorV101CoolantConnectorType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for CoolantConnectorV101CoolantConnectorType: valid values are %v", v, AllowedCoolantConnectorV101CoolantConnectorTypeEnumValues)
}

// AssertCoolantConnectorV101CoolantConnectorTypeRequired checks if the required fields are not zero-ed
func AssertCoolantConnectorV101CoolantConnectorTypeRequired(obj CoolantConnectorV101CoolantConnectorType) error {
	return nil
}

// AssertCoolantConnectorV101CoolantConnectorTypeConstraints checks if the values respects the defined constraints
func AssertCoolantConnectorV101CoolantConnectorTypeConstraints(obj CoolantConnectorV101CoolantConnectorType) error {
	return nil
}
