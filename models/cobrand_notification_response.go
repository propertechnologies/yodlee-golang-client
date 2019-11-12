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

// CobrandNotificationResponse CobrandNotificationResponse
// swagger:model CobrandNotificationResponse
type CobrandNotificationResponse struct {

	// event
	// Read Only: true
	Event []*CreateCobrandNotificationEvent `json:"event"`
}

// Validate validates this cobrand notification response
func (m *CobrandNotificationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CobrandNotificationResponse) validateEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.Event) { // not required
		return nil
	}

	for i := 0; i < len(m.Event); i++ {
		if swag.IsZero(m.Event[i]) { // not required
			continue
		}

		if m.Event[i] != nil {
			if err := m.Event[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("event" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CobrandNotificationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CobrandNotificationResponse) UnmarshalBinary(b []byte) error {
	var res CobrandNotificationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
