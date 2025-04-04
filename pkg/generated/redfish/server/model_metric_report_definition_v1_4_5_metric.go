// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MetricReportDefinitionV145Metric - Specifies a set of metrics to include in the metric report.  Calculation parameters, if present, are applied to the metrics prior to being included in the metric report.
type MetricReportDefinitionV145Metric struct {

	// The duration over which the function is computed.
	CollectionDuration *string `json:"CollectionDuration,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	CollectionFunction MetricReportDefinitionV145CalculationAlgorithmEnum `json:"CollectionFunction,omitempty"`

	CollectionTimeScope MetricReportDefinitionV145CollectionTimeScope `json:"CollectionTimeScope,omitempty"`

	// The metric definition identifier that contains the metric properties to include in the metric report.
	MetricId *string `json:"MetricId,omitempty"`

	// The list of URIs with wildcards and property identifiers to include in the metric report.  If a URI has wildcards, the wildcards are substituted as specified in the Wildcards property.
	MetricProperties []*string `json:"MetricProperties,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertMetricReportDefinitionV145MetricRequired checks if the required fields are not zero-ed
func AssertMetricReportDefinitionV145MetricRequired(obj MetricReportDefinitionV145Metric) error {
	return nil
}

// AssertMetricReportDefinitionV145MetricConstraints checks if the values respects the defined constraints
func AssertMetricReportDefinitionV145MetricConstraints(obj MetricReportDefinitionV145Metric) error {
	return nil
}
