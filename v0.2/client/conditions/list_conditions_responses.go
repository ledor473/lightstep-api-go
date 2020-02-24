// Code generated by go-swagger; DO NOT EDIT.

package conditions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ListConditionsReader is a Reader for the ListConditions structure.
type ListConditionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConditionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConditionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListConditionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListConditionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListConditionsOK creates a ListConditionsOK with default headers values
func NewListConditionsOK() *ListConditionsOK {
	return &ListConditionsOK{}
}

/*ListConditionsOK handles this case with default header values.

JSON-formatted metadata about all conditions in the project
*/
type ListConditionsOK struct {
	Payload interface{}
}

func (o *ListConditionsOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/conditions][%d] listConditionsOK  %+v", 200, o.Payload)
}

func (o *ListConditionsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ListConditionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConditionsUnauthorized creates a ListConditionsUnauthorized with default headers values
func NewListConditionsUnauthorized() *ListConditionsUnauthorized {
	return &ListConditionsUnauthorized{}
}

/*ListConditionsUnauthorized handles this case with default header values.

The API Key does not provide access to this resource, or the organization name does not exist
*/
type ListConditionsUnauthorized struct {
	Payload interface{}
}

func (o *ListConditionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/conditions][%d] listConditionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListConditionsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ListConditionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConditionsNotFound creates a ListConditionsNotFound with default headers values
func NewListConditionsNotFound() *ListConditionsNotFound {
	return &ListConditionsNotFound{}
}

/*ListConditionsNotFound handles this case with default header values.

Project name is not found
*/
type ListConditionsNotFound struct {
	Payload interface{}
}

func (o *ListConditionsNotFound) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/conditions][%d] listConditionsNotFound  %+v", 404, o.Payload)
}

func (o *ListConditionsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListConditionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
