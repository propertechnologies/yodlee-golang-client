// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UnregisterReader is a Reader for the Unregister structure.
type UnregisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnregisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUnregisterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUnregisterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnregisterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUnregisterNoContent creates a UnregisterNoContent with default headers values
func NewUnregisterNoContent() *UnregisterNoContent {
	return &UnregisterNoContent{}
}

/*UnregisterNoContent handles this case with default header values.

No Content
*/
type UnregisterNoContent struct {
}

func (o *UnregisterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user/unregister][%d] unregisterNoContent ", 204)
}

func (o *UnregisterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnregisterUnauthorized creates a UnregisterUnauthorized with default headers values
func NewUnregisterUnauthorized() *UnregisterUnauthorized {
	return &UnregisterUnauthorized{}
}

/*UnregisterUnauthorized handles this case with default header values.

Unauthorized
*/
type UnregisterUnauthorized struct {
}

func (o *UnregisterUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /user/unregister][%d] unregisterUnauthorized ", 401)
}

func (o *UnregisterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnregisterNotFound creates a UnregisterNotFound with default headers values
func NewUnregisterNotFound() *UnregisterNotFound {
	return &UnregisterNotFound{}
}

/*UnregisterNotFound handles this case with default header values.

Not Found
*/
type UnregisterNotFound struct {
}

func (o *UnregisterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user/unregister][%d] unregisterNotFound ", 404)
}

func (o *UnregisterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
