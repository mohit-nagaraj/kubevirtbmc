// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AggregationServiceV102Reset - This action is used to reset a set of resources.  For example this could be a list of computer systems.
type AggregationServiceV102Reset struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertAggregationServiceV102ResetRequired checks if the required fields are not zero-ed
func AssertAggregationServiceV102ResetRequired(obj AggregationServiceV102Reset) error {
	return nil
}

// AssertAggregationServiceV102ResetConstraints checks if the values respects the defined constraints
func AssertAggregationServiceV102ResetConstraints(obj AggregationServiceV102Reset) error {
	return nil
}
