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

// ComponentIntegrityV122ComponentIntegrity - The ComponentIntegrity resource provides critical and pertinent security information about a specific device, system, software element, or other managed entity.
type ComponentIntegrityV122ComponentIntegrity struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions ComponentIntegrityV122Actions `json:"Actions,omitempty"`

	// An indication of whether security protocols are enabled for the component.
	ComponentIntegrityEnabled bool `json:"ComponentIntegrityEnabled,omitempty"`

	ComponentIntegrityType ComponentIntegrityV122ComponentIntegrityType `json:"ComponentIntegrityType"`

	// The version of the security technology.
	ComponentIntegrityTypeVersion string `json:"ComponentIntegrityTypeVersion"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The date and time when information for the component was last updated.
	LastUpdated *time.Time `json:"LastUpdated,omitempty"`

	Links ComponentIntegrityV122Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	SPDM ComponentIntegrityV122SpdMinfo `json:"SPDM,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	TPM ComponentIntegrityV122TpMinfo `json:"TPM,omitempty"`

	// The link to the component whose integrity that this resource reports.
	TargetComponentURI string `json:"TargetComponentURI"`
}

// AssertComponentIntegrityV122ComponentIntegrityRequired checks if the required fields are not zero-ed
func AssertComponentIntegrityV122ComponentIntegrityRequired(obj ComponentIntegrityV122ComponentIntegrity) error {
	elements := map[string]interface{}{
		"@odata.id":                     obj.OdataId,
		"@odata.type":                   obj.OdataType,
		"ComponentIntegrityType":        obj.ComponentIntegrityType,
		"ComponentIntegrityTypeVersion": obj.ComponentIntegrityTypeVersion,
		"Id":                            obj.Id,
		"Name":                          obj.Name,
		"TargetComponentURI":            obj.TargetComponentURI,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertComponentIntegrityV122ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122SpdMinfoRequired(obj.SPDM); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122TpMinfoRequired(obj.TPM); err != nil {
		return err
	}
	return nil
}

// AssertComponentIntegrityV122ComponentIntegrityConstraints checks if the values respects the defined constraints
func AssertComponentIntegrityV122ComponentIntegrityConstraints(obj ComponentIntegrityV122ComponentIntegrity) error {
	if err := AssertComponentIntegrityV122ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122SpdMinfoConstraints(obj.SPDM); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122TpMinfoConstraints(obj.TPM); err != nil {
		return err
	}
	return nil
}
