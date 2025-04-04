// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CableV122Actions - The available actions for this resource.
type CableV122Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertCableV122ActionsRequired checks if the required fields are not zero-ed
func AssertCableV122ActionsRequired(obj CableV122Actions) error {
	return nil
}

// AssertCableV122ActionsConstraints checks if the values respects the defined constraints
func AssertCableV122ActionsConstraints(obj CableV122Actions) error {
	return nil
}
