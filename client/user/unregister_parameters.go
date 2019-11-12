// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUnregisterParams creates a new UnregisterParams object
// with the default values initialized.
func NewUnregisterParams() *UnregisterParams {

	return &UnregisterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnregisterParamsWithTimeout creates a new UnregisterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnregisterParamsWithTimeout(timeout time.Duration) *UnregisterParams {

	return &UnregisterParams{

		timeout: timeout,
	}
}

// NewUnregisterParamsWithContext creates a new UnregisterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnregisterParamsWithContext(ctx context.Context) *UnregisterParams {

	return &UnregisterParams{

		Context: ctx,
	}
}

// NewUnregisterParamsWithHTTPClient creates a new UnregisterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnregisterParamsWithHTTPClient(client *http.Client) *UnregisterParams {

	return &UnregisterParams{
		HTTPClient: client,
	}
}

/*UnregisterParams contains all the parameters to send to the API endpoint
for the unregister operation typically these are written to a http.Request
*/
type UnregisterParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unregister params
func (o *UnregisterParams) WithTimeout(timeout time.Duration) *UnregisterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unregister params
func (o *UnregisterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unregister params
func (o *UnregisterParams) WithContext(ctx context.Context) *UnregisterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unregister params
func (o *UnregisterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unregister params
func (o *UnregisterParams) WithHTTPClient(client *http.Client) *UnregisterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unregister params
func (o *UnregisterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UnregisterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
