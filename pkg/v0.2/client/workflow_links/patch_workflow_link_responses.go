// Code generated by go-swagger; DO NOT EDIT.

package workflow_links

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ledor473/lightstep-api-go/pkg/v0.2/response"
)

// PatchWorkflowLinkReader is a Reader for the PatchWorkflowLink structure.
type PatchWorkflowLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkflowLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWorkflowLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchWorkflowLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchWorkflowLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchWorkflowLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchWorkflowLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchWorkflowLinkOK creates a PatchWorkflowLinkOK with default headers values
func NewPatchWorkflowLinkOK() *PatchWorkflowLinkOK {
	return &PatchWorkflowLinkOK{}
}

/*PatchWorkflowLinkOK handles this case with default header values.

The workflow link was updated successfully
*/
type PatchWorkflowLinkOK struct {
	Payload response.PatchWorkflowLink
}

func (o *PatchWorkflowLinkOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/wf_links/{link-id}][%d] patchWorkflowLinkOK  %+v", 200, o.Payload)
}

func (o *PatchWorkflowLinkOK) GetPayload() response.PatchWorkflowLink {
	return o.Payload
}

func (o *PatchWorkflowLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkflowLinkBadRequest creates a PatchWorkflowLinkBadRequest with default headers values
func NewPatchWorkflowLinkBadRequest() *PatchWorkflowLinkBadRequest {
	return &PatchWorkflowLinkBadRequest{}
}

/*PatchWorkflowLinkBadRequest handles this case with default header values.

Bad request - could not decode JSON request
*/
type PatchWorkflowLinkBadRequest struct {
}

func (o *PatchWorkflowLinkBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/wf_links/{link-id}][%d] patchWorkflowLinkBadRequest ", 400)
}

func (o *PatchWorkflowLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchWorkflowLinkUnauthorized creates a PatchWorkflowLinkUnauthorized with default headers values
func NewPatchWorkflowLinkUnauthorized() *PatchWorkflowLinkUnauthorized {
	return &PatchWorkflowLinkUnauthorized{}
}

/*PatchWorkflowLinkUnauthorized handles this case with default header values.

The API Key does not provide access to this resource
*/
type PatchWorkflowLinkUnauthorized struct {
}

func (o *PatchWorkflowLinkUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/wf_links/{link-id}][%d] patchWorkflowLinkUnauthorized ", 401)
}

func (o *PatchWorkflowLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchWorkflowLinkForbidden creates a PatchWorkflowLinkForbidden with default headers values
func NewPatchWorkflowLinkForbidden() *PatchWorkflowLinkForbidden {
	return &PatchWorkflowLinkForbidden{}
}

/*PatchWorkflowLinkForbidden handles this case with default header values.

Unsupported request to create resource - links within the same project cannot have the same name and URL.
*/
type PatchWorkflowLinkForbidden struct {
}

func (o *PatchWorkflowLinkForbidden) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/wf_links/{link-id}][%d] patchWorkflowLinkForbidden ", 403)
}

func (o *PatchWorkflowLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchWorkflowLinkNotFound creates a PatchWorkflowLinkNotFound with default headers values
func NewPatchWorkflowLinkNotFound() *PatchWorkflowLinkNotFound {
	return &PatchWorkflowLinkNotFound{}
}

/*PatchWorkflowLinkNotFound handles this case with default header values.

Organization, project name, or workflow link ID is not found
*/
type PatchWorkflowLinkNotFound struct {
}

func (o *PatchWorkflowLinkNotFound) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/wf_links/{link-id}][%d] patchWorkflowLinkNotFound ", 404)
}

func (o *PatchWorkflowLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
