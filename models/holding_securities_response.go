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

// HoldingSecuritiesResponse HoldingSecuritiesResponse
// swagger:model HoldingSecuritiesResponse
type HoldingSecuritiesResponse struct {

	// holding
	// Read Only: true
	Holding []*SecurityHolding `json:"holding"`
}

// Validate validates this holding securities response
func (m *HoldingSecuritiesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHolding(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HoldingSecuritiesResponse) validateHolding(formats strfmt.Registry) error {

	if swag.IsZero(m.Holding) { // not required
		return nil
	}

	for i := 0; i < len(m.Holding); i++ {
		if swag.IsZero(m.Holding[i]) { // not required
			continue
		}

		if m.Holding[i] != nil {
			if err := m.Holding[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("holding" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HoldingSecuritiesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HoldingSecuritiesResponse) UnmarshalBinary(b []byte) error {
	var res HoldingSecuritiesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
