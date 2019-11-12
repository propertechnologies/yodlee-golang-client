// Code generated by go-swagger; DO NOT EDIT.

package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "yodlee-golang-client/models"
)

// GetProviderReader is a Reader for the GetProvider structure.
type GetProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProviderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProviderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProviderOK creates a GetProviderOK with default headers values
func NewGetProviderOK() *GetProviderOK {
	return &GetProviderOK{}
}

/*GetProviderOK handles this case with default header values.

OK
*/
type GetProviderOK struct {
	Payload *models.ProviderDetailResponse
}

func (o *GetProviderOK) Error() string {
	return fmt.Sprintf("[GET /providers/{providerId}][%d] getProviderOK  %+v", 200, o.Payload)
}

func (o *GetProviderOK) GetPayload() *models.ProviderDetailResponse {
	return o.Payload
}

func (o *GetProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProviderDetailResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProviderBadRequest creates a GetProviderBadRequest with default headers values
func NewGetProviderBadRequest() *GetProviderBadRequest {
	return &GetProviderBadRequest{}
}

/*GetProviderBadRequest handles this case with default header values.

Y800 : Invalid value for providerId
*/
type GetProviderBadRequest struct {
	Payload *models.YodleeError
}

func (o *GetProviderBadRequest) Error() string {
	return fmt.Sprintf("[GET /providers/{providerId}][%d] getProviderBadRequest  %+v", 400, o.Payload)
}

func (o *GetProviderBadRequest) GetPayload() *models.YodleeError {
	return o.Payload
}

func (o *GetProviderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.YodleeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProviderUnauthorized creates a GetProviderUnauthorized with default headers values
func NewGetProviderUnauthorized() *GetProviderUnauthorized {
	return &GetProviderUnauthorized{}
}

/*GetProviderUnauthorized handles this case with default header values.

Unauthorized
*/
type GetProviderUnauthorized struct {
}

func (o *GetProviderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /providers/{providerId}][%d] getProviderUnauthorized ", 401)
}

func (o *GetProviderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProviderNotFound creates a GetProviderNotFound with default headers values
func NewGetProviderNotFound() *GetProviderNotFound {
	return &GetProviderNotFound{}
}

/*GetProviderNotFound handles this case with default header values.

Not Found
*/
type GetProviderNotFound struct {
}

func (o *GetProviderNotFound) Error() string {
	return fmt.Sprintf("[GET /providers/{providerId}][%d] getProviderNotFound ", 404)
}

func (o *GetProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
