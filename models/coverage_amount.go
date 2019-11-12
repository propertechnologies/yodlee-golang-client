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

// CoverageAmount CoverageAmount
// swagger:model CoverageAmount
type CoverageAmount struct {

	// The maximum amount that will be paid to an individual or an entity for a covered loss<br><br><b>Aggregated / Manual</b>: Aggregated<br><b>Applicable containers</b>: insurance,investment<br><b>Endpoints</b>:<ul><li>GET accounts</li><li>GET accounts/{accountId}</li></ul>
	// Read Only: true
	Cover *Money `json:"cover,omitempty"`

	// The type of coverage limit indicates if the coverage is in-network or out-of-network.<br><br><b>Aggregated / Manual</b>: Aggregated<br><b>Applicable containers</b>: insurance,investment<br><b>Endpoints</b>:<ul><li>GET accounts</li><li>GET accounts/{accountId}</li></ul><b>Applicable Values:</b><br>
	// * IN_NETWORK: Indicates the doctor or facility providing the health care services has negotiated a contracted rate with the health insurance company.<br>
	// * OUT_NETWORK: Indicates the doctor or facility providing health care services has no contract with the health insurance company.<br>
	// Read Only: true
	// Enum: [IN_NETWORK OUT_NETWORK]
	LimitType string `json:"limitType,omitempty"`

	// The amount the insurance company paid for the incurred medical expenses.<br><br><b>Aggregated / Manual</b>: Aggregated<br><b>Applicable containers</b>: insurance,investment<br><b>Endpoints</b>:<ul><li>GET accounts</li><li>GET accounts/{accountId}</li></ul>
	// Read Only: true
	Met *Money `json:"met,omitempty"`

	// The type of coverage provided to an individual or an entity.<br><br><b>Aggregated / Manual</b>: Aggregated<br><b>Applicable containers</b>: insurance,investment<br><b>Endpoints</b>:<ul><li>GET accounts</li><li>GET accounts/{accountId}</li></ul><b>Applicable Values:</b><br>
	// * DEDUCTIBLE: The amount paid for covered health care services before the health insurance plan starts paying.<br>
	// * OUT_OF_POCKET: The maximum amount the insurer has to pay for covered health care services in a plan year.<br>
	// * ANNUAL_BENEFIT: A cap on the benefits the insurance company will pay in a year while the insurer is enrolled in a particular health insurance plan.<br>
	// * MAX_BENEFIT: The maximum amount the insurance company pays for nonessential healthcare services.<br>
	// * COVERAGE_AMOUNT: The maximum amount payable in the event of a claim by the policyholder.<br>
	// * MONTHLY_BENEFIT: As part of the income protection cover, the monthly amount provided as financial support in the event of sickness or injury.<br>
	// * OTHER: Any coverage type other than what has been listed here.<br>
	// Read Only: true
	// Enum: [DEDUCTIBLE OUT_OF_POCKET ANNUAL_BENEFIT MAX_BENEFIT COVERAGE_AMOUNT MONTHLY_BENEFIT OTHER]
	Type string `json:"type,omitempty"`

	// The type of coverage unit indicates if the coverage is for an individual or a family.<br><br><b>Aggregated / Manual</b>: Aggregated<br><b>Applicable containers</b>: insurance,investment<br><b>Endpoints</b>:<ul><li>GET accounts</li><li>GET accounts/{accountId}</li></ul><b>Applicable Values:</b><br>
	// * PER_FAMILY: Indicates the health insurance coverage is for the complete family.<br>
	// * PER_MEMBER: Indicates the health insurance coverage is for an individual.<br>
	// Read Only: true
	// Enum: [PER_FAMILY PER_MEMBER]
	UnitType string `json:"unitType,omitempty"`
}

// Validate validates this coverage amount
func (m *CoverageAmount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCover(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CoverageAmount) validateCover(formats strfmt.Registry) error {

	if swag.IsZero(m.Cover) { // not required
		return nil
	}

	if m.Cover != nil {
		if err := m.Cover.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cover")
			}
			return err
		}
	}

	return nil
}

var coverageAmountTypeLimitTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IN_NETWORK","OUT_NETWORK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		coverageAmountTypeLimitTypePropEnum = append(coverageAmountTypeLimitTypePropEnum, v)
	}
}

const (

	// CoverageAmountLimitTypeINNETWORK captures enum value "IN_NETWORK"
	CoverageAmountLimitTypeINNETWORK string = "IN_NETWORK"

	// CoverageAmountLimitTypeOUTNETWORK captures enum value "OUT_NETWORK"
	CoverageAmountLimitTypeOUTNETWORK string = "OUT_NETWORK"
)

// prop value enum
func (m *CoverageAmount) validateLimitTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, coverageAmountTypeLimitTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CoverageAmount) validateLimitType(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitType) { // not required
		return nil
	}

	// value enum
	if err := m.validateLimitTypeEnum("limitType", "body", m.LimitType); err != nil {
		return err
	}

	return nil
}

func (m *CoverageAmount) validateMet(formats strfmt.Registry) error {

	if swag.IsZero(m.Met) { // not required
		return nil
	}

	if m.Met != nil {
		if err := m.Met.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("met")
			}
			return err
		}
	}

	return nil
}

var coverageAmountTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEDUCTIBLE","OUT_OF_POCKET","ANNUAL_BENEFIT","MAX_BENEFIT","COVERAGE_AMOUNT","MONTHLY_BENEFIT","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		coverageAmountTypeTypePropEnum = append(coverageAmountTypeTypePropEnum, v)
	}
}

const (

	// CoverageAmountTypeDEDUCTIBLE captures enum value "DEDUCTIBLE"
	CoverageAmountTypeDEDUCTIBLE string = "DEDUCTIBLE"

	// CoverageAmountTypeOUTOFPOCKET captures enum value "OUT_OF_POCKET"
	CoverageAmountTypeOUTOFPOCKET string = "OUT_OF_POCKET"

	// CoverageAmountTypeANNUALBENEFIT captures enum value "ANNUAL_BENEFIT"
	CoverageAmountTypeANNUALBENEFIT string = "ANNUAL_BENEFIT"

	// CoverageAmountTypeMAXBENEFIT captures enum value "MAX_BENEFIT"
	CoverageAmountTypeMAXBENEFIT string = "MAX_BENEFIT"

	// CoverageAmountTypeCOVERAGEAMOUNT captures enum value "COVERAGE_AMOUNT"
	CoverageAmountTypeCOVERAGEAMOUNT string = "COVERAGE_AMOUNT"

	// CoverageAmountTypeMONTHLYBENEFIT captures enum value "MONTHLY_BENEFIT"
	CoverageAmountTypeMONTHLYBENEFIT string = "MONTHLY_BENEFIT"

	// CoverageAmountTypeOTHER captures enum value "OTHER"
	CoverageAmountTypeOTHER string = "OTHER"
)

// prop value enum
func (m *CoverageAmount) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, coverageAmountTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CoverageAmount) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

var coverageAmountTypeUnitTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PER_FAMILY","PER_MEMBER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		coverageAmountTypeUnitTypePropEnum = append(coverageAmountTypeUnitTypePropEnum, v)
	}
}

const (

	// CoverageAmountUnitTypePERFAMILY captures enum value "PER_FAMILY"
	CoverageAmountUnitTypePERFAMILY string = "PER_FAMILY"

	// CoverageAmountUnitTypePERMEMBER captures enum value "PER_MEMBER"
	CoverageAmountUnitTypePERMEMBER string = "PER_MEMBER"
)

// prop value enum
func (m *CoverageAmount) validateUnitTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, coverageAmountTypeUnitTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CoverageAmount) validateUnitType(formats strfmt.Registry) error {

	if swag.IsZero(m.UnitType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnitTypeEnum("unitType", "body", m.UnitType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CoverageAmount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoverageAmount) UnmarshalBinary(b []byte) error {
	var res CoverageAmount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
