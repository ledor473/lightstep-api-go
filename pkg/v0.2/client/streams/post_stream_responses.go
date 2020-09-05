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

// PostStreamReader is a Reader for the PostStream structure.
type PostStreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostStreamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostStreamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostStreamOK creates a PostStreamOK with default headers values
func NewPostStreamOK() *PostStreamOK {
	return &PostStreamOK{}
}

/*PostStreamOK handles this case with default header values.

The stream was created (or updated) successfully
*/
type PostStreamOK struct {
	Payload response.PostStream
}

func (o *PostStreamOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/streams][%d] postStreamOK  %+v", 200, o.Payload)
}

func (o *PostStreamOK) GetPayload() response.PostStream {
	return o.Payload
}

func (o *PostStreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStreamUnauthorized creates a PostStreamUnauthorized with default headers values
func NewPostStreamUnauthorized() *PostStreamUnauthorized {
	return &PostStreamUnauthorized{}
}

/*PostStreamUnauthorized handles this case with default header values.

The API Key does not provide access to this resource, or the organization name does not exist
*/
type PostStreamUnauthorized struct {
}

func (o *PostStreamUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/streams][%d] postStreamUnauthorized ", 401)
}

func (o *PostStreamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStreamNotFound creates a PostStreamNotFound with default headers values
func NewPostStreamNotFound() *PostStreamNotFound {
	return &PostStreamNotFound{}
}

/*PostStreamNotFound handles this case with default header values.

Project name is not found
*/
type PostStreamNotFound struct {
}

func (o *PostStreamNotFound) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/streams][%d] postStreamNotFound ", 404)
}

func (o *PostStreamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
