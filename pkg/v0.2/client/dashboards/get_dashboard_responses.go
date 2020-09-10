// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ledor473/lightstep-api-go/pkg/v0.2/response"
)

// GetDashboardReader is a Reader for the GetDashboard structure.
type GetDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDashboardOK creates a GetDashboardOK with default headers values
func NewGetDashboardOK() *GetDashboardOK {
	return &GetDashboardOK{}
}

/*GetDashboardOK handles this case with default header values.

JSON-formatted metadata about the dashboard
*/
type GetDashboardOK struct {
	Payload response.GetDashboard
}

func (o *GetDashboardOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/dashboards/{dashboard-id}][%d] getDashboardOK  %+v", 200, o.Payload)
}

func (o *GetDashboardOK) GetPayload() response.GetDashboard {
	return o.Payload
}

func (o *GetDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardUnauthorized creates a GetDashboardUnauthorized with default headers values
func NewGetDashboardUnauthorized() *GetDashboardUnauthorized {
	return &GetDashboardUnauthorized{}
}

/*GetDashboardUnauthorized handles this case with default header values.

The API Key does not provide access to this resource, or the organization name does not exist
*/
type GetDashboardUnauthorized struct {
}

func (o *GetDashboardUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/dashboards/{dashboard-id}][%d] getDashboardUnauthorized ", 401)
}

func (o *GetDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDashboardNotFound creates a GetDashboardNotFound with default headers values
func NewGetDashboardNotFound() *GetDashboardNotFound {
	return &GetDashboardNotFound{}
}

/*GetDashboardNotFound handles this case with default header values.

Project name or dashboard ID is not found
*/
type GetDashboardNotFound struct {
}

func (o *GetDashboardNotFound) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/dashboards/{dashboard-id}][%d] getDashboardNotFound ", 404)
}

func (o *GetDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
