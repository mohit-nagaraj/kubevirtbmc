// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// BatteryV122Reset - This action resets the battery.
type BatteryV122Reset struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertBatteryV122ResetRequired checks if the required fields are not zero-ed
func AssertBatteryV122ResetRequired(obj BatteryV122Reset) error {
	return nil
}

// AssertBatteryV122ResetConstraints checks if the values respects the defined constraints
func AssertBatteryV122ResetConstraints(obj BatteryV122Reset) error {
	return nil
}
