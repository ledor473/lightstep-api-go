// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteDashboardReader is a Reader for the DeleteDashboard structure.
type DeleteDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDashboardNoContent creates a DeleteDashboardNoContent with default headers values
func NewDeleteDashboardNoContent() *DeleteDashboardNoContent {
	return &DeleteDashboardNoContent{}
}

/*DeleteDashboardNoContent handles this case with default header values.

Dashboard was successfully deleted
*/
type DeleteDashboardNoContent struct {
	Payload interface{}
}

func (o *DeleteDashboardNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/projects/{project}/dashboard/{dashboard-id}][%d] deleteDashboardNoContent  %+v", 204, o.Payload)
}

func (o *DeleteDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardUnauthorized creates a DeleteDashboardUnauthorized with default headers values
func NewDeleteDashboardUnauthorized() *DeleteDashboardUnauthorized {
	return &DeleteDashboardUnauthorized{}
}

/*DeleteDashboardUnauthorized handles this case with default header values.

The API Key does not provide access to this resource, or the organization name does not exist
*/
type DeleteDashboardUnauthorized struct {
	Payload interface{}
}

func (o *DeleteDashboardUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/projects/{project}/dashboard/{dashboard-id}][%d] deleteDashboardUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteDashboardUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardNotFound creates a DeleteDashboardNotFound with default headers values
func NewDeleteDashboardNotFound() *DeleteDashboardNotFound {
	return &DeleteDashboardNotFound{}
}

/*DeleteDashboardNotFound handles this case with default header values.

Project name is not found
*/
type DeleteDashboardNotFound struct {
	Payload interface{}
}

func (o *DeleteDashboardNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/projects/{project}/dashboard/{dashboard-id}][%d] deleteDashboardNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardInternalServerError creates a DeleteDashboardInternalServerError with default headers values
func NewDeleteDashboardInternalServerError() *DeleteDashboardInternalServerError {
	return &DeleteDashboardInternalServerError{}
}

/*DeleteDashboardInternalServerError handles this case with default header values.

Dashboard identifier not found
*/
type DeleteDashboardInternalServerError struct {
	Payload interface{}
}

func (o *DeleteDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/projects/{project}/dashboard/{dashboard-id}][%d] deleteDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDashboardInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}