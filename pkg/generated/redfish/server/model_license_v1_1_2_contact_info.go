// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// LicenseV112ContactInfo - Contact information for this license.
type LicenseV112ContactInfo struct {

	// Name of this contact.
	ContactName *string `json:"ContactName,omitempty"`

	// Email address for this contact.
	EmailAddress *string `json:"EmailAddress,omitempty"`

	// Phone number for this contact.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
}

// AssertLicenseV112ContactInfoRequired checks if the required fields are not zero-ed
func AssertLicenseV112ContactInfoRequired(obj LicenseV112ContactInfo) error {
	return nil
}

// AssertLicenseV112ContactInfoConstraints checks if the values respects the defined constraints
func AssertLicenseV112ContactInfoConstraints(obj LicenseV112ContactInfo) error {
	return nil
}
