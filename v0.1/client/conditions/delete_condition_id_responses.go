// Code generated by go-swagger; DO NOT EDIT.

package conditions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteConditionIDReader is a Reader for the DeleteConditionID structure.
type DeleteConditionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConditionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConditionIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteConditionIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConditionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteConditionIDNoContent creates a DeleteConditionIDNoContent with default headers values
func NewDeleteConditionIDNoContent() *DeleteConditionIDNoContent {
	return &DeleteConditionIDNoContent{}
}

/*DeleteConditionIDNoContent handles this case with default header values.

Condition was successfully deleted
*/
type DeleteConditionIDNoContent struct {
}

func (o *DeleteConditionIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/projects/{project}/conditions/{condition-id}][%d] deleteConditionIdNoContent ", 204)
}

func (o *DeleteConditionIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConditionIDUnauthorized creates a DeleteConditionIDUnauthorized with default headers values
func NewDeleteConditionIDUnauthorized() *DeleteConditionIDUnauthorized {
	return &DeleteConditionIDUnauthorized{}
}

/*DeleteConditionIDUnauthorized handles this case with default header values.

The API Key does not provide access to this resource, or the organization name does not exist
*/
type DeleteConditionIDUnauthorized struct {
}

func (o *DeleteConditionIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/projects/{project}/conditions/{condition-id}][%d] deleteConditionIdUnauthorized ", 401)
}

func (o *DeleteConditionIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConditionIDNotFound creates a DeleteConditionIDNotFound with default headers values
func NewDeleteConditionIDNotFound() *DeleteConditionIDNotFound {
	return &DeleteConditionIDNotFound{}
}

/*DeleteConditionIDNotFound handles this case with default header values.

Project name is not found
*/
type DeleteConditionIDNotFound struct {
}

func (o *DeleteConditionIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/projects/{project}/conditions/{condition-id}][%d] deleteConditionIdNotFound ", 404)
}

func (o *DeleteConditionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
