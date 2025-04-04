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

type LogEntryV1170LogDiagnosticDataTypes string

// List of LogEntryV1170LogDiagnosticDataTypes
const (
	LOGENTRYV1170LOGDIAGNOSTICDATATYPES_MANAGER      LogEntryV1170LogDiagnosticDataTypes = "Manager"
	LOGENTRYV1170LOGDIAGNOSTICDATATYPES_PRE_OS       LogEntryV1170LogDiagnosticDataTypes = "PreOS"
	LOGENTRYV1170LOGDIAGNOSTICDATATYPES_OS           LogEntryV1170LogDiagnosticDataTypes = "OS"
	LOGENTRYV1170LOGDIAGNOSTICDATATYPES_OEM          LogEntryV1170LogDiagnosticDataTypes = "OEM"
	LOGENTRYV1170LOGDIAGNOSTICDATATYPES_CPER         LogEntryV1170LogDiagnosticDataTypes = "CPER"
	LOGENTRYV1170LOGDIAGNOSTICDATATYPES_CPER_SECTION LogEntryV1170LogDiagnosticDataTypes = "CPERSection"
)

// AllowedLogEntryV1170LogDiagnosticDataTypesEnumValues is all the allowed values of LogEntryV1170LogDiagnosticDataTypes enum
var AllowedLogEntryV1170LogDiagnosticDataTypesEnumValues = []LogEntryV1170LogDiagnosticDataTypes{
	"Manager",
	"PreOS",
	"OS",
	"OEM",
	"CPER",
	"CPERSection",
}

// validLogEntryV1170LogDiagnosticDataTypesEnumValue provides a map of LogEntryV1170LogDiagnosticDataTypess for fast verification of use input
var validLogEntryV1170LogDiagnosticDataTypesEnumValues = map[LogEntryV1170LogDiagnosticDataTypes]struct{}{
	"Manager":     {},
	"PreOS":       {},
	"OS":          {},
	"OEM":         {},
	"CPER":        {},
	"CPERSection": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogEntryV1170LogDiagnosticDataTypes) IsValid() bool {
	_, ok := validLogEntryV1170LogDiagnosticDataTypesEnumValues[v]
	return ok
}

// NewLogEntryV1170LogDiagnosticDataTypesFromValue returns a pointer to a valid LogEntryV1170LogDiagnosticDataTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogEntryV1170LogDiagnosticDataTypesFromValue(v string) (LogEntryV1170LogDiagnosticDataTypes, error) {
	ev := LogEntryV1170LogDiagnosticDataTypes(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for LogEntryV1170LogDiagnosticDataTypes: valid values are %v", v, AllowedLogEntryV1170LogDiagnosticDataTypesEnumValues)
}

// AssertLogEntryV1170LogDiagnosticDataTypesRequired checks if the required fields are not zero-ed
func AssertLogEntryV1170LogDiagnosticDataTypesRequired(obj LogEntryV1170LogDiagnosticDataTypes) error {
	return nil
}

// AssertLogEntryV1170LogDiagnosticDataTypesConstraints checks if the values respects the defined constraints
func AssertLogEntryV1170LogDiagnosticDataTypesConstraints(obj LogEntryV1170LogDiagnosticDataTypes) error {
	return nil
}
