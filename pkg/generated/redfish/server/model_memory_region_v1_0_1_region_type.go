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

type MemoryRegionV101RegionType string

// List of MemoryRegionV101RegionType
const (
	MEMORYREGIONV101REGIONTYPE_STATIC  MemoryRegionV101RegionType = "Static"
	MEMORYREGIONV101REGIONTYPE_DYNAMIC MemoryRegionV101RegionType = "Dynamic"
)

// AllowedMemoryRegionV101RegionTypeEnumValues is all the allowed values of MemoryRegionV101RegionType enum
var AllowedMemoryRegionV101RegionTypeEnumValues = []MemoryRegionV101RegionType{
	"Static",
	"Dynamic",
}

// validMemoryRegionV101RegionTypeEnumValue provides a map of MemoryRegionV101RegionTypes for fast verification of use input
var validMemoryRegionV101RegionTypeEnumValues = map[MemoryRegionV101RegionType]struct{}{
	"Static":  {},
	"Dynamic": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MemoryRegionV101RegionType) IsValid() bool {
	_, ok := validMemoryRegionV101RegionTypeEnumValues[v]
	return ok
}

// NewMemoryRegionV101RegionTypeFromValue returns a pointer to a valid MemoryRegionV101RegionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMemoryRegionV101RegionTypeFromValue(v string) (MemoryRegionV101RegionType, error) {
	ev := MemoryRegionV101RegionType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for MemoryRegionV101RegionType: valid values are %v", v, AllowedMemoryRegionV101RegionTypeEnumValues)
}

// AssertMemoryRegionV101RegionTypeRequired checks if the required fields are not zero-ed
func AssertMemoryRegionV101RegionTypeRequired(obj MemoryRegionV101RegionType) error {
	return nil
}

// AssertMemoryRegionV101RegionTypeConstraints checks if the values respects the defined constraints
func AssertMemoryRegionV101RegionTypeConstraints(obj MemoryRegionV101RegionType) error {
	return nil
}
