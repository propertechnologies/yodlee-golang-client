// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VerificationTransaction VerificationTransaction
// swagger:model VerificationTransaction
type VerificationTransaction struct {

	// amount
	Amount *Money `json:"amount,omitempty"`

	// base type
	// Enum: [CREDIT DEBIT]
	BaseType string `json:"baseType,omitempty"`
}

// Validate validates this verification transaction
func (m *VerificationTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VerificationTransaction) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

var verificationTransactionTypeBaseTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREDIT","DEBIT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		verificationTransactionTypeBaseTypePropEnum = append(verificationTransactionTypeBaseTypePropEnum, v)
	}
}

const (

	// VerificationTransactionBaseTypeCREDIT captures enum value "CREDIT"
	VerificationTransactionBaseTypeCREDIT string = "CREDIT"

	// VerificationTransactionBaseTypeDEBIT captures enum value "DEBIT"
	VerificationTransactionBaseTypeDEBIT string = "DEBIT"
)

// prop value enum
func (m *VerificationTransaction) validateBaseTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, verificationTransactionTypeBaseTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VerificationTransaction) validateBaseType(formats strfmt.Registry) error {

	if swag.IsZero(m.BaseType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBaseTypeEnum("baseType", "body", m.BaseType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VerificationTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VerificationTransaction) UnmarshalBinary(b []byte) error {
	var res VerificationTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
