// Code generated by go-swagger; DO NOT EDIT.

package provider_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "yodlee-golang-client/models"
)

// EditCredentialsOrRefreshProviderAccountReader is a Reader for the EditCredentialsOrRefreshProviderAccount structure.
type EditCredentialsOrRefreshProviderAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditCredentialsOrRefreshProviderAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditCredentialsOrRefreshProviderAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEditCredentialsOrRefreshProviderAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEditCredentialsOrRefreshProviderAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEditCredentialsOrRefreshProviderAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditCredentialsOrRefreshProviderAccountOK creates a EditCredentialsOrRefreshProviderAccountOK with default headers values
func NewEditCredentialsOrRefreshProviderAccountOK() *EditCredentialsOrRefreshProviderAccountOK {
	return &EditCredentialsOrRefreshProviderAccountOK{}
}

/*EditCredentialsOrRefreshProviderAccountOK handles this case with default header values.

OK
*/
type EditCredentialsOrRefreshProviderAccountOK struct {
	Payload *models.UpdatedProviderAccountResponse
}

func (o *EditCredentialsOrRefreshProviderAccountOK) Error() string {
	return fmt.Sprintf("[PUT /providerAccounts][%d] editCredentialsOrRefreshProviderAccountOK  %+v", 200, o.Payload)
}

func (o *EditCredentialsOrRefreshProviderAccountOK) GetPayload() *models.UpdatedProviderAccountResponse {
	return o.Payload
}

func (o *EditCredentialsOrRefreshProviderAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdatedProviderAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditCredentialsOrRefreshProviderAccountBadRequest creates a EditCredentialsOrRefreshProviderAccountBadRequest with default headers values
func NewEditCredentialsOrRefreshProviderAccountBadRequest() *EditCredentialsOrRefreshProviderAccountBadRequest {
	return &EditCredentialsOrRefreshProviderAccountBadRequest{}
}

/*EditCredentialsOrRefreshProviderAccountBadRequest handles this case with default header values.

Y805 : Multiple providerAccountId not supported for updating credentials<br>Y800 : Invalid value for credentialsParam<br>Y400 : id and value in credentialsParam are mandatory<br>Y806 : Invalid input<br>Y823 : Credentials are not applicable for real estate aggregated / manual accounts<br>
*/
type EditCredentialsOrRefreshProviderAccountBadRequest struct {
	Payload *models.YodleeError
}

func (o *EditCredentialsOrRefreshProviderAccountBadRequest) Error() string {
	return fmt.Sprintf("[PUT /providerAccounts][%d] editCredentialsOrRefreshProviderAccountBadRequest  %+v", 400, o.Payload)
}

func (o *EditCredentialsOrRefreshProviderAccountBadRequest) GetPayload() *models.YodleeError {
	return o.Payload
}

func (o *EditCredentialsOrRefreshProviderAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.YodleeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditCredentialsOrRefreshProviderAccountUnauthorized creates a EditCredentialsOrRefreshProviderAccountUnauthorized with default headers values
func NewEditCredentialsOrRefreshProviderAccountUnauthorized() *EditCredentialsOrRefreshProviderAccountUnauthorized {
	return &EditCredentialsOrRefreshProviderAccountUnauthorized{}
}

/*EditCredentialsOrRefreshProviderAccountUnauthorized handles this case with default header values.

Unauthorized
*/
type EditCredentialsOrRefreshProviderAccountUnauthorized struct {
}

func (o *EditCredentialsOrRefreshProviderAccountUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /providerAccounts][%d] editCredentialsOrRefreshProviderAccountUnauthorized ", 401)
}

func (o *EditCredentialsOrRefreshProviderAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditCredentialsOrRefreshProviderAccountNotFound creates a EditCredentialsOrRefreshProviderAccountNotFound with default headers values
func NewEditCredentialsOrRefreshProviderAccountNotFound() *EditCredentialsOrRefreshProviderAccountNotFound {
	return &EditCredentialsOrRefreshProviderAccountNotFound{}
}

/*EditCredentialsOrRefreshProviderAccountNotFound handles this case with default header values.

Not Found
*/
type EditCredentialsOrRefreshProviderAccountNotFound struct {
}

func (o *EditCredentialsOrRefreshProviderAccountNotFound) Error() string {
	return fmt.Sprintf("[PUT /providerAccounts][%d] editCredentialsOrRefreshProviderAccountNotFound ", 404)
}

func (o *EditCredentialsOrRefreshProviderAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
