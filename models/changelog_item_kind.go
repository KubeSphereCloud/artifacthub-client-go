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

// ChangelogItemKind ChangelogItemKind
//
// Types of changes:
//   - `added` - New features
//   - `changed` - Changes in existing functionality
//   - `deprecated` - Soon-to-be removed features
//   - `removed` - Removed features
//   - `fixed` - Any bug fixed
//   - `security` - In case of vulnerabilities
//
// swagger:model ChangelogItemKind
type ChangelogItemKind string

func NewChangelogItemKind(value ChangelogItemKind) *ChangelogItemKind {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ChangelogItemKind.
func (m ChangelogItemKind) Pointer() *ChangelogItemKind {
	return &m
}

const (

	// ChangelogItemKindAdded captures enum value "added"
	ChangelogItemKindAdded ChangelogItemKind = "added"

	// ChangelogItemKindChanged captures enum value "changed"
	ChangelogItemKindChanged ChangelogItemKind = "changed"

	// ChangelogItemKindDeprecated captures enum value "deprecated"
	ChangelogItemKindDeprecated ChangelogItemKind = "deprecated"

	// ChangelogItemKindRemoved captures enum value "removed"
	ChangelogItemKindRemoved ChangelogItemKind = "removed"

	// ChangelogItemKindFixed captures enum value "fixed"
	ChangelogItemKindFixed ChangelogItemKind = "fixed"

	// ChangelogItemKindSecurity captures enum value "security"
	ChangelogItemKindSecurity ChangelogItemKind = "security"
)

// for schema
var changelogItemKindEnum []interface{}

func init() {
	var res []ChangelogItemKind
	if err := json.Unmarshal([]byte(`["added","changed","deprecated","removed","fixed","security"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changelogItemKindEnum = append(changelogItemKindEnum, v)
	}
}

func (m ChangelogItemKind) validateChangelogItemKindEnum(path, location string, value ChangelogItemKind) error {
	if err := validate.EnumCase(path, location, value, changelogItemKindEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this changelog item kind
func (m ChangelogItemKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateChangelogItemKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this changelog item kind based on context it is used
func (m ChangelogItemKind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
