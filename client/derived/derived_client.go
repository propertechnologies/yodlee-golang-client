// Code generated by go-swagger; DO NOT EDIT.

package derived

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new derived API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for derived API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetHoldingSummary gets holding summary

The get holding summary service is used to get the summary of asset classifications for the user.<br>By default, accounts with status as ACTIVE and TO BE CLOSED will be considered.<br>If the include parameter value is passed as details then a summary with holdings and account information is returned.<br>
*/
func (a *Client) GetHoldingSummary(params *GetHoldingSummaryParams) (*GetHoldingSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHoldingSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHoldingSummary",
		Method:             "GET",
		PathPattern:        "/derived/holdingSummary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHoldingSummaryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHoldingSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHoldingSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNetworth gets networth summary

The get networth service is used to get the networth for the user.<br>If the include parameter value is passed as details then networth with historical balances is returned. <br>
*/
func (a *Client) GetNetworth(params *GetNetworthParams) (*GetNetworthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworth",
		Method:             "GET",
		PathPattern:        "/derived/networth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNetworthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTransactionSummary gets transaction summary

The transaction summary service provides the summary values of transactions for the given date range by category type, high-level categories, or system-defined categories.<br><br>Yodlee has the transaction data stored for a day, month, year and week per category as per the availability of user's data. If the include parameter value is passed as details, then summary details will be returned depending on the interval passed-monthly is the default.<br><br><b>Notes:</b><br>1.Details can be requested for only one system-defined category<br>2.Dates will not be respected for monthly, yearly, and weekly details<br>3.When monthly details are requested, only the fromDate and toDate month will be respected<br>4.When yearly details are requested, only the fromDate and toDate year will be respected<br>5.For weekly data points, details will be provided for every Sunday date available within the fromDate and toDate<br>
*/
func (a *Client) GetTransactionSummary(params *GetTransactionSummaryParams) (*GetTransactionSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTransactionSummary",
		Method:             "GET",
		PathPattern:        "/derived/transactionSummary",
		ProducesMediaTypes: []string{"application/json;charset=UTF-8"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTransactionSummaryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTransactionSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactionSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
