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

// PostIdentityV3AuthTokensCreatedBody post identity v3 auth tokens created body
// swagger:model postIdentityV3AuthTokensCreatedBody
type PostIdentityV3AuthTokensCreatedBody struct {

	// token
	// Required: true
	Token *Token `json:"token"`
}

// Validate validates this post identity v3 auth tokens created body
func (m *PostIdentityV3AuthTokensCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostIdentityV3AuthTokensCreatedBody) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostIdentityV3AuthTokensCreatedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostIdentityV3AuthTokensCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostIdentityV3AuthTokensCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
