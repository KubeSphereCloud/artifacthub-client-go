// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OptOutBody OptOutBody
//
// swagger:model OptOutBody
type OptOutBody struct {

	// event kind
	// Required: true
	EventKind *EventKindID `json:"event_kind"`

	// repository id
	// Required: true
	// Format: uuid
	RepositoryID *strfmt.UUID `json:"repository_id"`
}

// Validate validates this opt out body
func (m *OptOutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OptOutBody) validateEventKind(formats strfmt.Registry) error {

	if err := validate.Required("event_kind", "body", m.EventKind); err != nil {
		return err
	}

	if err := validate.Required("event_kind", "body", m.EventKind); err != nil {
		return err
	}

	if m.EventKind != nil {
		if err := m.EventKind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_kind")
			}
			return err
		}
	}

	return nil
}

func (m *OptOutBody) validateRepositoryID(formats strfmt.Registry) error {

	if err := validate.Required("repository_id", "body", m.RepositoryID); err != nil {
		return err
	}

	if err := validate.FormatOf("repository_id", "body", "uuid", m.RepositoryID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this opt out body based on the context it is used
func (m *OptOutBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OptOutBody) contextValidateEventKind(ctx context.Context, formats strfmt.Registry) error {

	if m.EventKind != nil {
		if err := m.EventKind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_kind")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OptOutBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OptOutBody) UnmarshalBinary(b []byte) error {
	var res OptOutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
