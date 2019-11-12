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

// StatementResponse StatementResponse
// swagger:model StatementResponse
type StatementResponse struct {

	// statement
	// Read Only: true
	Statement []*Statement `json:"statement"`
}

// Validate validates this statement response
func (m *StatementResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatement(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatementResponse) validateStatement(formats strfmt.Registry) error {

	if swag.IsZero(m.Statement) { // not required
		return nil
	}

	for i := 0; i < len(m.Statement); i++ {
		if swag.IsZero(m.Statement[i]) { // not required
			continue
		}

		if m.Statement[i] != nil {
			if err := m.Statement[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statement" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatementResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatementResponse) UnmarshalBinary(b []byte) error {
	var res StatementResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
