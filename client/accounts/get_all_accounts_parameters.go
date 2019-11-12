// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllAccountsParams creates a new GetAllAccountsParams object
// with the default values initialized.
func NewGetAllAccountsParams() *GetAllAccountsParams {
	var ()
	return &GetAllAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllAccountsParamsWithTimeout creates a new GetAllAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllAccountsParamsWithTimeout(timeout time.Duration) *GetAllAccountsParams {
	var ()
	return &GetAllAccountsParams{

		timeout: timeout,
	}
}

// NewGetAllAccountsParamsWithContext creates a new GetAllAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllAccountsParamsWithContext(ctx context.Context) *GetAllAccountsParams {
	var ()
	return &GetAllAccountsParams{

		Context: ctx,
	}
}

// NewGetAllAccountsParamsWithHTTPClient creates a new GetAllAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllAccountsParamsWithHTTPClient(client *http.Client) *GetAllAccountsParams {
	var ()
	return &GetAllAccountsParams{
		HTTPClient: client,
	}
}

/*GetAllAccountsParams contains all the parameters to send to the API endpoint
for the get all accounts operation typically these are written to a http.Request
*/
type GetAllAccountsParams struct {

	/*AccountID
	  Comma separated accountIds.

	*/
	AccountID *string
	/*Container
	  bank/creditCard/investment/insurance/loan/reward/bill/realEstate/otherAssets/otherLiabilities

	*/
	Container *string
	/*Include
	  profile, holder, fullAccountNumber, paymentProfile, autoRefresh

	*/
	Include *string
	/*ProviderAccountID
	  providerAccountId

	*/
	ProviderAccountID *string
	/*RequestID
	  The unique identifier that returns contextual data

	*/
	RequestID *string
	/*Status
	  ACTIVE/INACTIVE/TO_BE_CLOSED/CLOSED

	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all accounts params
func (o *GetAllAccountsParams) WithTimeout(timeout time.Duration) *GetAllAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all accounts params
func (o *GetAllAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all accounts params
func (o *GetAllAccountsParams) WithContext(ctx context.Context) *GetAllAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all accounts params
func (o *GetAllAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all accounts params
func (o *GetAllAccountsParams) WithHTTPClient(client *http.Client) *GetAllAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all accounts params
func (o *GetAllAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get all accounts params
func (o *GetAllAccountsParams) WithAccountID(accountID *string) *GetAllAccountsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get all accounts params
func (o *GetAllAccountsParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithContainer adds the container to the get all accounts params
func (o *GetAllAccountsParams) WithContainer(container *string) *GetAllAccountsParams {
	o.SetContainer(container)
	return o
}

// SetContainer adds the container to the get all accounts params
func (o *GetAllAccountsParams) SetContainer(container *string) {
	o.Container = container
}

// WithInclude adds the include to the get all accounts params
func (o *GetAllAccountsParams) WithInclude(include *string) *GetAllAccountsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get all accounts params
func (o *GetAllAccountsParams) SetInclude(include *string) {
	o.Include = include
}

// WithProviderAccountID adds the providerAccountID to the get all accounts params
func (o *GetAllAccountsParams) WithProviderAccountID(providerAccountID *string) *GetAllAccountsParams {
	o.SetProviderAccountID(providerAccountID)
	return o
}

// SetProviderAccountID adds the providerAccountId to the get all accounts params
func (o *GetAllAccountsParams) SetProviderAccountID(providerAccountID *string) {
	o.ProviderAccountID = providerAccountID
}

// WithRequestID adds the requestID to the get all accounts params
func (o *GetAllAccountsParams) WithRequestID(requestID *string) *GetAllAccountsParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get all accounts params
func (o *GetAllAccountsParams) SetRequestID(requestID *string) {
	o.RequestID = requestID
}

// WithStatus adds the status to the get all accounts params
func (o *GetAllAccountsParams) WithStatus(status *string) *GetAllAccountsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get all accounts params
func (o *GetAllAccountsParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID string
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}

	}

	if o.Container != nil {

		// query param container
		var qrContainer string
		if o.Container != nil {
			qrContainer = *o.Container
		}
		qContainer := qrContainer
		if qContainer != "" {
			if err := r.SetQueryParam("container", qContainer); err != nil {
				return err
			}
		}

	}

	if o.Include != nil {

		// query param include
		var qrInclude string
		if o.Include != nil {
			qrInclude = *o.Include
		}
		qInclude := qrInclude
		if qInclude != "" {
			if err := r.SetQueryParam("include", qInclude); err != nil {
				return err
			}
		}

	}

	if o.ProviderAccountID != nil {

		// query param providerAccountId
		var qrProviderAccountID string
		if o.ProviderAccountID != nil {
			qrProviderAccountID = *o.ProviderAccountID
		}
		qProviderAccountID := qrProviderAccountID
		if qProviderAccountID != "" {
			if err := r.SetQueryParam("providerAccountId", qProviderAccountID); err != nil {
				return err
			}
		}

	}

	if o.RequestID != nil {

		// query param requestId
		var qrRequestID string
		if o.RequestID != nil {
			qrRequestID = *o.RequestID
		}
		qRequestID := qrRequestID
		if qRequestID != "" {
			if err := r.SetQueryParam("requestId", qRequestID); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
