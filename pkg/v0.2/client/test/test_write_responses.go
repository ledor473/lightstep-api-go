// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// TestWriteReader is a Reader for the TestWrite structure.
type TestWriteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestWriteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestWriteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTestWriteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestWriteOK creates a TestWriteOK with default headers values
func NewTestWriteOK() *TestWriteOK {
	return &TestWriteOK{}
}

/*TestWriteOK handles this case with default header values.

OK!
*/
type TestWriteOK struct {
	Payload interface{}
}

func (o *TestWriteOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/test][%d] testWriteOK  %+v", 200, o.Payload)
}

func (o *TestWriteOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TestWriteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestWriteUnauthorized creates a TestWriteUnauthorized with default headers values
func NewTestWriteUnauthorized() *TestWriteUnauthorized {
	return &TestWriteUnauthorized{}
}

/*TestWriteUnauthorized handles this case with default header values.

Unauthorized
*/
type TestWriteUnauthorized struct {
	Payload interface{}
}

func (o *TestWriteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{organization}/test][%d] testWriteUnauthorized  %+v", 401, o.Payload)
}

func (o *TestWriteUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *TestWriteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}