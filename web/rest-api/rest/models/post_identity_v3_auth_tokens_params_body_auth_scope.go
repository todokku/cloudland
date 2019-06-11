// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostIdentityV3AuthTokensParamsBodyAuthScope post identity v3 auth tokens params body auth scope
// swagger:model postIdentityV3AuthTokensParamsBodyAuthScope
type PostIdentityV3AuthTokensParamsBodyAuthScope struct {

	// project
	Project *PostIdentityV3AuthTokensParamsBodyAuthScopeProject `json:"project,omitempty"`
}

// Validate validates this post identity v3 auth tokens params body auth scope
func (m *PostIdentityV3AuthTokensParamsBodyAuthScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostIdentityV3AuthTokensParamsBodyAuthScope) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostIdentityV3AuthTokensParamsBodyAuthScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostIdentityV3AuthTokensParamsBodyAuthScope) UnmarshalBinary(b []byte) error {
	var res PostIdentityV3AuthTokensParamsBodyAuthScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
