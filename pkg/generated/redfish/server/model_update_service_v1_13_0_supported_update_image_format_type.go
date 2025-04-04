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

type UpdateServiceV1130SupportedUpdateImageFormatType string

// List of UpdateServiceV1130SupportedUpdateImageFormatType
const (
	UPDATESERVICEV1130SUPPORTEDUPDATEIMAGEFORMATTYPE_PLDMV1_0       UpdateServiceV1130SupportedUpdateImageFormatType = "PLDMv1_0"
	UPDATESERVICEV1130SUPPORTEDUPDATEIMAGEFORMATTYPE_PLDMV1_1       UpdateServiceV1130SupportedUpdateImageFormatType = "PLDMv1_1"
	UPDATESERVICEV1130SUPPORTEDUPDATEIMAGEFORMATTYPE_PLDMV1_2       UpdateServiceV1130SupportedUpdateImageFormatType = "PLDMv1_2"
	UPDATESERVICEV1130SUPPORTEDUPDATEIMAGEFORMATTYPE_PLDMV1_3       UpdateServiceV1130SupportedUpdateImageFormatType = "PLDMv1_3"
	UPDATESERVICEV1130SUPPORTEDUPDATEIMAGEFORMATTYPE_UEFI_CAPSULE   UpdateServiceV1130SupportedUpdateImageFormatType = "UEFICapsule"
	UPDATESERVICEV1130SUPPORTEDUPDATEIMAGEFORMATTYPE_VENDOR_DEFINED UpdateServiceV1130SupportedUpdateImageFormatType = "VendorDefined"
)

// AllowedUpdateServiceV1130SupportedUpdateImageFormatTypeEnumValues is all the allowed values of UpdateServiceV1130SupportedUpdateImageFormatType enum
var AllowedUpdateServiceV1130SupportedUpdateImageFormatTypeEnumValues = []UpdateServiceV1130SupportedUpdateImageFormatType{
	"PLDMv1_0",
	"PLDMv1_1",
	"PLDMv1_2",
	"PLDMv1_3",
	"UEFICapsule",
	"VendorDefined",
}

// validUpdateServiceV1130SupportedUpdateImageFormatTypeEnumValue provides a map of UpdateServiceV1130SupportedUpdateImageFormatTypes for fast verification of use input
var validUpdateServiceV1130SupportedUpdateImageFormatTypeEnumValues = map[UpdateServiceV1130SupportedUpdateImageFormatType]struct{}{
	"PLDMv1_0":      {},
	"PLDMv1_1":      {},
	"PLDMv1_2":      {},
	"PLDMv1_3":      {},
	"UEFICapsule":   {},
	"VendorDefined": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpdateServiceV1130SupportedUpdateImageFormatType) IsValid() bool {
	_, ok := validUpdateServiceV1130SupportedUpdateImageFormatTypeEnumValues[v]
	return ok
}

// NewUpdateServiceV1130SupportedUpdateImageFormatTypeFromValue returns a pointer to a valid UpdateServiceV1130SupportedUpdateImageFormatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpdateServiceV1130SupportedUpdateImageFormatTypeFromValue(v string) (UpdateServiceV1130SupportedUpdateImageFormatType, error) {
	ev := UpdateServiceV1130SupportedUpdateImageFormatType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for UpdateServiceV1130SupportedUpdateImageFormatType: valid values are %v", v, AllowedUpdateServiceV1130SupportedUpdateImageFormatTypeEnumValues)
}

// AssertUpdateServiceV1130SupportedUpdateImageFormatTypeRequired checks if the required fields are not zero-ed
func AssertUpdateServiceV1130SupportedUpdateImageFormatTypeRequired(obj UpdateServiceV1130SupportedUpdateImageFormatType) error {
	return nil
}

// AssertUpdateServiceV1130SupportedUpdateImageFormatTypeConstraints checks if the values respects the defined constraints
func AssertUpdateServiceV1130SupportedUpdateImageFormatTypeConstraints(obj UpdateServiceV1130SupportedUpdateImageFormatType) error {
	return nil
}
