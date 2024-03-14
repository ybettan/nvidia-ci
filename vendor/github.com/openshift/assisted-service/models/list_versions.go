// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListVersions list versions
//
// swagger:model list-versions
type ListVersions struct {

	// release tag
	ReleaseTag string `json:"release_tag,omitempty"`

	// versions
	Versions Versions `json:"versions,omitempty"`
}

// Validate validates this list versions
func (m *ListVersions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListVersions) validateVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	if m.Versions != nil {
		if err := m.Versions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list versions based on the context it is used
func (m *ListVersions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListVersions) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Versions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("versions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("versions")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListVersions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListVersions) UnmarshalBinary(b []byte) error {
	var res ListVersions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
