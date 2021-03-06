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

// ListFlavorsDetailOKBody list flavors detail o k body
// swagger:model listFlavorsDetailOKBody
type ListFlavorsDetailOKBody struct {

	// flavors
	// Required: true
	Flavors FlavorsDetail `json:"flavors"`
}

// Validate validates this list flavors detail o k body
func (m *ListFlavorsDetailOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlavors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListFlavorsDetailOKBody) validateFlavors(formats strfmt.Registry) error {

	if err := validate.Required("flavors", "body", m.Flavors); err != nil {
		return err
	}

	if err := m.Flavors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("flavors")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListFlavorsDetailOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListFlavorsDetailOKBody) UnmarshalBinary(b []byte) error {
	var res ListFlavorsDetailOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
