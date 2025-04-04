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
	"time"
)

// ControlV150ControlLoop - The details and coefficients used to operate a control loop.
type ControlV150ControlLoop struct {

	// The date and time that the control loop coefficients were changed.
	CoefficientUpdateTime *time.Time `json:"CoefficientUpdateTime,omitempty"`

	// The differential coefficient.
	Differential *float32 `json:"Differential,omitempty"`

	// The integral coefficient.
	Integral *float32 `json:"Integral,omitempty"`

	// The proportional coefficient.
	Proportional *float32 `json:"Proportional,omitempty"`
}

// AssertControlV150ControlLoopRequired checks if the required fields are not zero-ed
func AssertControlV150ControlLoopRequired(obj ControlV150ControlLoop) error {
	return nil
}

// AssertControlV150ControlLoopConstraints checks if the values respects the defined constraints
func AssertControlV150ControlLoopConstraints(obj ControlV150ControlLoop) error {
	return nil
}
