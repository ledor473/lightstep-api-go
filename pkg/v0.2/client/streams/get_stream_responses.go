// Code generated by go-swagger; DO NOT EDIT.

package streams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ledor473/lightstep-api-go/pkg/v0.2/response"
)

// GetStreamReader is a Reader for the GetStream structure.
type GetStreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetStreamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStreamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStreamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStreamOK creates a GetStreamOK with default headers values
func NewGetStreamOK() *GetStreamOK {
	return &GetStreamOK{}
}

/*GetStreamOK handles this case with default header values.

JSON-formatted data about the given stream
*/
type GetStreamOK struct {
	Payload response.GetStream
}

func (o *GetStreamOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/streams/{stream-id}][%d] getStreamOK  %+v", 200, o.Payload)
}

func (o *GetStreamOK) GetPayload() response.GetStream {
	return o.Payload
}

func (o *GetStreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStreamUnauthorized creates a GetStreamUnauthorized with default headers values
func NewGetStreamUnauthorized() *GetStreamUnauthorized {
	return &GetStreamUnauthorized{}
}

/*GetStreamUnauthorized handles this case with default header values.

The API Key does not provide access to this resource, or the organization name does not exist
*/
type GetStreamUnauthorized struct {
}

func (o *GetStreamUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/streams/{stream-id}][%d] getStreamUnauthorized ", 401)
}

func (o *GetStreamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStreamNotFound creates a GetStreamNotFound with default headers values
func NewGetStreamNotFound() *GetStreamNotFound {
	return &GetStreamNotFound{}
}

/*GetStreamNotFound handles this case with default header values.

Project name is not found
*/
type GetStreamNotFound struct {
}

func (o *GetStreamNotFound) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/streams/{stream-id}][%d] getStreamNotFound ", 404)
}

func (o *GetStreamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStreamInternalServerError creates a GetStreamInternalServerError with default headers values
func NewGetStreamInternalServerError() *GetStreamInternalServerError {
	return &GetStreamInternalServerError{}
}

/*GetStreamInternalServerError handles this case with default header values.

The stream identifier is not valid
*/
type GetStreamInternalServerError struct {
}

func (o *GetStreamInternalServerError) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/streams/{stream-id}][%d] getStreamInternalServerError ", 500)
}

func (o *GetStreamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
