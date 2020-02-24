// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateSnapshotReader is a Reader for the CreateSnapshot structure.
type CreateSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateSnapshotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateSnapshotTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSnapshotOK creates a CreateSnapshotOK with default headers values
func NewCreateSnapshotOK() *CreateSnapshotOK {
	return &CreateSnapshotOK{}
}

/*CreateSnapshotOK handles this case with default header values.

The snapshot was created successfully
*/
type CreateSnapshotOK struct {
	Payload interface{}
}

func (o *CreateSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/snapshots][%d] createSnapshotOK  %+v", 200, o.Payload)
}

func (o *CreateSnapshotOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotBadRequest creates a CreateSnapshotBadRequest with default headers values
func NewCreateSnapshotBadRequest() *CreateSnapshotBadRequest {
	return &CreateSnapshotBadRequest{}
}

/*CreateSnapshotBadRequest handles this case with default header values.

Query parameters invalid
*/
type CreateSnapshotBadRequest struct {
	Payload interface{}
}

func (o *CreateSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/snapshots][%d] createSnapshotBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSnapshotBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotUnauthorized creates a CreateSnapshotUnauthorized with default headers values
func NewCreateSnapshotUnauthorized() *CreateSnapshotUnauthorized {
	return &CreateSnapshotUnauthorized{}
}

/*CreateSnapshotUnauthorized handles this case with default header values.

The API Key does not provide access to this resource, or the organization name does not exist
*/
type CreateSnapshotUnauthorized struct {
	Payload interface{}
}

func (o *CreateSnapshotUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/snapshots][%d] createSnapshotUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateSnapshotUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateSnapshotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotNotFound creates a CreateSnapshotNotFound with default headers values
func NewCreateSnapshotNotFound() *CreateSnapshotNotFound {
	return &CreateSnapshotNotFound{}
}

/*CreateSnapshotNotFound handles this case with default header values.

Project name is not found
*/
type CreateSnapshotNotFound struct {
	Payload interface{}
}

func (o *CreateSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/snapshots][%d] createSnapshotNotFound  %+v", 404, o.Payload)
}

func (o *CreateSnapshotNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotTooManyRequests creates a CreateSnapshotTooManyRequests with default headers values
func NewCreateSnapshotTooManyRequests() *CreateSnapshotTooManyRequests {
	return &CreateSnapshotTooManyRequests{}
}

/*CreateSnapshotTooManyRequests handles this case with default header values.

Snapshot creation exceeded rate limit
*/
type CreateSnapshotTooManyRequests struct {
	Payload interface{}
}

func (o *CreateSnapshotTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/snapshots][%d] createSnapshotTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateSnapshotTooManyRequests) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateSnapshotTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
