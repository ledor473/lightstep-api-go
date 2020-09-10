// Code generated by go-swagger; DO NOT EDIT.

package conditions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ledor473/lightstep-api-go/pkg/v0.2/response"
)

// GetConditionReader is a Reader for the GetCondition structure.
type GetConditionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConditionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConditionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetConditionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConditionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConditionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConditionOK creates a GetConditionOK with default headers values
func NewGetConditionOK() *GetConditionOK {
	return &GetConditionOK{}
}

/*GetConditionOK handles this case with default header values.

JSON-formatted data about the given condition
*/
type GetConditionOK struct {
	Payload response.GetCondition
}

func (o *GetConditionOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/conditions/{condition-id}][%d] getConditionOK  %+v", 200, o.Payload)
}

func (o *GetConditionOK) GetPayload() response.GetCondition {
	return o.Payload
}

func (o *GetConditionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConditionUnauthorized creates a GetConditionUnauthorized with default headers values
func NewGetConditionUnauthorized() *GetConditionUnauthorized {
	return &GetConditionUnauthorized{}
}

/*GetConditionUnauthorized handles this case with default header values.

The API Key does not provide access to this resource, or the organization name does not exist
*/
type GetConditionUnauthorized struct {
}

func (o *GetConditionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/conditions/{condition-id}][%d] getConditionUnauthorized ", 401)
}

func (o *GetConditionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConditionNotFound creates a GetConditionNotFound with default headers values
func NewGetConditionNotFound() *GetConditionNotFound {
	return &GetConditionNotFound{}
}

/*GetConditionNotFound handles this case with default header values.

Project name is not found
*/
type GetConditionNotFound struct {
}

func (o *GetConditionNotFound) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/conditions/{condition-id}][%d] getConditionNotFound ", 404)
}

func (o *GetConditionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConditionInternalServerError creates a GetConditionInternalServerError with default headers values
func NewGetConditionInternalServerError() *GetConditionInternalServerError {
	return &GetConditionInternalServerError{}
}

/*GetConditionInternalServerError handles this case with default header values.

The condition identifier is not valid
*/
type GetConditionInternalServerError struct {
}

func (o *GetConditionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/conditions/{condition-id}][%d] getConditionInternalServerError ", 500)
}

func (o *GetConditionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
