// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AddressPoolV124SystemMacRange - The Media Access Control (MAC) address range for the EVPN-based fabrics.
type AddressPoolV124SystemMacRange struct {

	// The lower system MAC address.
	Lower *string `json:"Lower,omitempty" validate:"regexp=^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"`

	// The upper system MAC address.
	Upper *string `json:"Upper,omitempty" validate:"regexp=^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"`
}

// AssertAddressPoolV124SystemMacRangeRequired checks if the required fields are not zero-ed
func AssertAddressPoolV124SystemMacRangeRequired(obj AddressPoolV124SystemMacRange) error {
	return nil
}

// AssertAddressPoolV124SystemMacRangeConstraints checks if the values respects the defined constraints
func AssertAddressPoolV124SystemMacRangeConstraints(obj AddressPoolV124SystemMacRange) error {
	return nil
}
