// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateUserRegistration UpdateUserRegistration
// swagger:model UpdateUserRegistration
type UpdateUserRegistration struct {

	// address
	Address *UserAddress `json:"address,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// name
	Name *Name `json:"name,omitempty"`

	// preferences
	Preferences *UserRequestPreferences `json:"preferences,omitempty"`

	// segment name
	SegmentName string `json:"segmentName,omitempty"`
}

// Validate validates this update user registration
func (m *UpdateUserRegistration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUserRegistration) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateUserRegistration) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateUserRegistration) validatePreferences(formats strfmt.Registry) error {

	if swag.IsZero(m.Preferences) { // not required
		return nil
	}

	if m.Preferences != nil {
		if err := m.Preferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preferences")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUserRegistration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserRegistration) UnmarshalBinary(b []byte) error {
	var res UpdateUserRegistration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
