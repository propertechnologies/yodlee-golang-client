// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DerivedHoldingsLinks DerivedHoldingsLinks
// swagger:model DerivedHoldingsLinks
type DerivedHoldingsLinks struct {

	// holdings
	// Read Only: true
	Holdings string `json:"holdings,omitempty"`
}

// Validate validates this derived holdings links
func (m *DerivedHoldingsLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DerivedHoldingsLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DerivedHoldingsLinks) UnmarshalBinary(b []byte) error {
	var res DerivedHoldingsLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
