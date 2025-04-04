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

type DriveV1180StatusIndicator string

// List of DriveV1180StatusIndicator
const (
	DRIVEV1180STATUSINDICATOR_OK                          DriveV1180StatusIndicator = "OK"
	DRIVEV1180STATUSINDICATOR_FAIL                        DriveV1180StatusIndicator = "Fail"
	DRIVEV1180STATUSINDICATOR_REBUILD                     DriveV1180StatusIndicator = "Rebuild"
	DRIVEV1180STATUSINDICATOR_PREDICTIVE_FAILURE_ANALYSIS DriveV1180StatusIndicator = "PredictiveFailureAnalysis"
	DRIVEV1180STATUSINDICATOR_HOTSPARE                    DriveV1180StatusIndicator = "Hotspare"
	DRIVEV1180STATUSINDICATOR_IN_A_CRITICAL_ARRAY         DriveV1180StatusIndicator = "InACriticalArray"
	DRIVEV1180STATUSINDICATOR_IN_A_FAILED_ARRAY           DriveV1180StatusIndicator = "InAFailedArray"
)

// AllowedDriveV1180StatusIndicatorEnumValues is all the allowed values of DriveV1180StatusIndicator enum
var AllowedDriveV1180StatusIndicatorEnumValues = []DriveV1180StatusIndicator{
	"OK",
	"Fail",
	"Rebuild",
	"PredictiveFailureAnalysis",
	"Hotspare",
	"InACriticalArray",
	"InAFailedArray",
}

// validDriveV1180StatusIndicatorEnumValue provides a map of DriveV1180StatusIndicators for fast verification of use input
var validDriveV1180StatusIndicatorEnumValues = map[DriveV1180StatusIndicator]struct{}{
	"OK":                        {},
	"Fail":                      {},
	"Rebuild":                   {},
	"PredictiveFailureAnalysis": {},
	"Hotspare":                  {},
	"InACriticalArray":          {},
	"InAFailedArray":            {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DriveV1180StatusIndicator) IsValid() bool {
	_, ok := validDriveV1180StatusIndicatorEnumValues[v]
	return ok
}

// NewDriveV1180StatusIndicatorFromValue returns a pointer to a valid DriveV1180StatusIndicator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDriveV1180StatusIndicatorFromValue(v string) (DriveV1180StatusIndicator, error) {
	ev := DriveV1180StatusIndicator(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for DriveV1180StatusIndicator: valid values are %v", v, AllowedDriveV1180StatusIndicatorEnumValues)
}

// AssertDriveV1180StatusIndicatorRequired checks if the required fields are not zero-ed
func AssertDriveV1180StatusIndicatorRequired(obj DriveV1180StatusIndicator) error {
	return nil
}

// AssertDriveV1180StatusIndicatorConstraints checks if the values respects the defined constraints
func AssertDriveV1180StatusIndicatorConstraints(obj DriveV1180StatusIndicator) error {
	return nil
}
