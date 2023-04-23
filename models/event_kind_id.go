// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EventKindID EventKindId
//
// Event kind:
//   - `0` - New package release
//   - `1` - Security alerts
//   - `2` - Repository tracking errors
//   - `4` - Repository scanning errors
//
// swagger:model EventKindId
type EventKindID int32

// for schema
var eventKindIdEnum []interface{}

func init() {
	var res []EventKindID
	if err := json.Unmarshal([]byte(`["0","1","2","4"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventKindIdEnum = append(eventKindIdEnum, v)
	}
}

func (m EventKindID) validateEventKindIDEnum(path, location string, value EventKindID) error {
	if err := validate.EnumCase(path, location, value, eventKindIdEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this event kind Id
func (m EventKindID) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEventKindIDEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this event kind Id based on context it is used
func (m EventKindID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}