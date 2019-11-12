// Code generated by go-swagger; DO NOT EDIT.

package data_extracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "yodlee-golang-client/models"
)

// GetDataExtractsEventsReader is a Reader for the GetDataExtractsEvents structure.
type GetDataExtractsEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataExtractsEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataExtractsEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDataExtractsEventsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDataExtractsEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDataExtractsEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDataExtractsEventsOK creates a GetDataExtractsEventsOK with default headers values
func NewGetDataExtractsEventsOK() *GetDataExtractsEventsOK {
	return &GetDataExtractsEventsOK{}
}

/*GetDataExtractsEventsOK handles this case with default header values.

OK
*/
type GetDataExtractsEventsOK struct {
	Payload *models.DataExtractsEventResponse
}

func (o *GetDataExtractsEventsOK) Error() string {
	return fmt.Sprintf("[GET /dataExtracts/events][%d] getDataExtractsEventsOK  %+v", 200, o.Payload)
}

func (o *GetDataExtractsEventsOK) GetPayload() *models.DataExtractsEventResponse {
	return o.Payload
}

func (o *GetDataExtractsEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataExtractsEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataExtractsEventsBadRequest creates a GetDataExtractsEventsBadRequest with default headers values
func NewGetDataExtractsEventsBadRequest() *GetDataExtractsEventsBadRequest {
	return &GetDataExtractsEventsBadRequest{}
}

/*GetDataExtractsEventsBadRequest handles this case with default header values.

Y800 : Invalid value for fromDate.fromDate cannot be greater than current time<br>Y800 : Invalid value for toDate.toDate cannot be greater than current time<br>Y800 : Invalid value for fromDate or toDate.fromDate and toDate cannot be older than 7 days<br>Y800 : Invalid value for fromDate.fromDate cannot be greater than toDate.
*/
type GetDataExtractsEventsBadRequest struct {
	Payload *models.YodleeError
}

func (o *GetDataExtractsEventsBadRequest) Error() string {
	return fmt.Sprintf("[GET /dataExtracts/events][%d] getDataExtractsEventsBadRequest  %+v", 400, o.Payload)
}

func (o *GetDataExtractsEventsBadRequest) GetPayload() *models.YodleeError {
	return o.Payload
}

func (o *GetDataExtractsEventsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.YodleeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataExtractsEventsUnauthorized creates a GetDataExtractsEventsUnauthorized with default headers values
func NewGetDataExtractsEventsUnauthorized() *GetDataExtractsEventsUnauthorized {
	return &GetDataExtractsEventsUnauthorized{}
}

/*GetDataExtractsEventsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDataExtractsEventsUnauthorized struct {
}

func (o *GetDataExtractsEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dataExtracts/events][%d] getDataExtractsEventsUnauthorized ", 401)
}

func (o *GetDataExtractsEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDataExtractsEventsNotFound creates a GetDataExtractsEventsNotFound with default headers values
func NewGetDataExtractsEventsNotFound() *GetDataExtractsEventsNotFound {
	return &GetDataExtractsEventsNotFound{}
}

/*GetDataExtractsEventsNotFound handles this case with default header values.

Not Found
*/
type GetDataExtractsEventsNotFound struct {
}

func (o *GetDataExtractsEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /dataExtracts/events][%d] getDataExtractsEventsNotFound ", 404)
}

func (o *GetDataExtractsEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
