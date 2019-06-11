// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostIdentityV3AuthTokensParamsBodyAuthScopeProject post identity v3 auth tokens params body auth scope project
// swagger:model postIdentityV3AuthTokensParamsBodyAuthScopeProject
type PostIdentityV3AuthTokensParamsBodyAuthScopeProject struct {

	// domain
	Domain *PostIdentityV3AuthTokensParamsBodyAuthScopeProjectDomain `json:"domain,omitempty"`

	// name
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this post identity v3 auth tokens params body auth scope project
func (m *PostIdentityV3AuthTokensParamsBodyAuthScopeProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostIdentityV3AuthTokensParamsBodyAuthScopeProject) validateDomain(formats strfmt.Registry) error {

	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if m.Domain != nil {
		if err := m.Domain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

func (m *PostIdentityV3AuthTokensParamsBodyAuthScopeProject) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostIdentityV3AuthTokensParamsBodyAuthScopeProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostIdentityV3AuthTokensParamsBodyAuthScopeProject) UnmarshalBinary(b []byte) error {
	var res PostIdentityV3AuthTokensParamsBodyAuthScopeProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
