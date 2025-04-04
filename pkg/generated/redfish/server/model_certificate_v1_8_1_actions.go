// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CertificateV181Actions - The available actions for this resource.
type CertificateV181Actions struct {
	CertificateRekey CertificateV181Rekey `json:"#Certificate.Rekey,omitempty"`

	CertificateRenew CertificateV181Renew `json:"#Certificate.Renew,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertCertificateV181ActionsRequired checks if the required fields are not zero-ed
func AssertCertificateV181ActionsRequired(obj CertificateV181Actions) error {
	if err := AssertCertificateV181RekeyRequired(obj.CertificateRekey); err != nil {
		return err
	}
	if err := AssertCertificateV181RenewRequired(obj.CertificateRenew); err != nil {
		return err
	}
	return nil
}

// AssertCertificateV181ActionsConstraints checks if the values respects the defined constraints
func AssertCertificateV181ActionsConstraints(obj CertificateV181Actions) error {
	if err := AssertCertificateV181RekeyConstraints(obj.CertificateRekey); err != nil {
		return err
	}
	if err := AssertCertificateV181RenewConstraints(obj.CertificateRenew); err != nil {
		return err
	}
	return nil
}
