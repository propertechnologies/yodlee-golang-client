// Code generated by go-swagger; DO NOT EDIT.

package cobrand

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CobrandLogoutReader is a Reader for the CobrandLogout structure.
type CobrandLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CobrandLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCobrandLogoutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCobrandLogoutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCobrandLogoutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCobrandLogoutNoContent creates a CobrandLogoutNoContent with default headers values
func NewCobrandLogoutNoContent() *CobrandLogoutNoContent {
	return &CobrandLogoutNoContent{}
}

/*CobrandLogoutNoContent handles this case with default header values.

Logout successful
*/
type CobrandLogoutNoContent struct {
}

func (o *CobrandLogoutNoContent) Error() string {
	return fmt.Sprintf("[POST /cobrand/logout][%d] cobrandLogoutNoContent ", 204)
}

func (o *CobrandLogoutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCobrandLogoutUnauthorized creates a CobrandLogoutUnauthorized with default headers values
func NewCobrandLogoutUnauthorized() *CobrandLogoutUnauthorized {
	return &CobrandLogoutUnauthorized{}
}

/*CobrandLogoutUnauthorized handles this case with default header values.

Unauthorized
*/
type CobrandLogoutUnauthorized struct {
}

func (o *CobrandLogoutUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cobrand/logout][%d] cobrandLogoutUnauthorized ", 401)
}

func (o *CobrandLogoutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCobrandLogoutNotFound creates a CobrandLogoutNotFound with default headers values
func NewCobrandLogoutNotFound() *CobrandLogoutNotFound {
	return &CobrandLogoutNotFound{}
}

/*CobrandLogoutNotFound handles this case with default header values.

Not Found
*/
type CobrandLogoutNotFound struct {
}

func (o *CobrandLogoutNotFound) Error() string {
	return fmt.Sprintf("[POST /cobrand/logout][%d] cobrandLogoutNotFound ", 404)
}

func (o *CobrandLogoutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
