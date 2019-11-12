// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Cobrand Cobrand
// swagger:model Cobrand
type Cobrand struct {

	// cobrand login
	CobrandLogin string `json:"cobrandLogin,omitempty"`

	// cobrand password
	CobrandPassword string `json:"cobrandPassword,omitempty"`

	// The customer's locale that will be considered for the localization functionality.<br><br><b>Endpoints</b>:<ul><li>POST cobrand/login</li></ul>
	Locale string `json:"locale,omitempty"`
}

// Validate validates this cobrand
func (m *Cobrand) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cobrand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cobrand) UnmarshalBinary(b []byte) error {
	var res Cobrand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
