// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserAccessToken UserAccessToken
// swagger:model UserAccessToken
type UserAccessToken struct {

	// access tokens
	AccessTokens []*AccessTokens `json:"accessTokens"`
}

// Validate validates this user access token
func (m *UserAccessToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessTokens(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAccessToken) validateAccessTokens(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessTokens) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessTokens); i++ {
		if swag.IsZero(m.AccessTokens[i]) { // not required
			continue
		}

		if m.AccessTokens[i] != nil {
			if err := m.AccessTokens[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessTokens" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserAccessToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAccessToken) UnmarshalBinary(b []byte) error {
	var res UserAccessToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
