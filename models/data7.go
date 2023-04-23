// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Data7 Data7
//
// swagger:model Data7
type Data7 struct {

	// kyverno category
	// Example: OpenShift
	KyvernoCategory string `json:"kyverno/category,omitempty"`

	// kyverno kubernetes version
	// Example: 1.21
	KyvernoKubernetesVersion string `json:"kyverno/kubernetesVersion,omitempty"`

	// kyverno subject
	// Example: Pod, Volume
	KyvernoSubject string `json:"kyverno/subject,omitempty"`

	// kyverno version
	// Example: 1.5.0
	KyvernoVersion string `json:"kyverno/version,omitempty"`

	// policy
	// Example: apiVersion: kyverno.io/v1 kind: ClusterPolicy metadata:\n  name: add-certificates-volume\n  annotations:\n    policies.kyverno.io/title: Add Certificates as a Volume\n    policies.kyverno.io/category: Sample\n    policies.kyverno.io/subject: Pod,Volume
	Policy string `json:"policy,omitempty"`
}

// Validate validates this data7
func (m *Data7) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this data7 based on context it is used
func (m *Data7) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Data7) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Data7) UnmarshalBinary(b []byte) error {
	var res Data7
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}