// Code generated by go-swagger; DO NOT EDIT.

package workflow_links

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateWorkflowLinkIDReader is a Reader for the CreateWorkflowLinkID structure.
type CreateWorkflowLinkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWorkflowLinkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateWorkflowLinkIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateWorkflowLinkIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateWorkflowLinkIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateWorkflowLinkIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateWorkflowLinkIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateWorkflowLinkIDOK creates a CreateWorkflowLinkIDOK with default headers values
func NewCreateWorkflowLinkIDOK() *CreateWorkflowLinkIDOK {
	return &CreateWorkflowLinkIDOK{}
}

/*CreateWorkflowLinkIDOK handles this case with default header values.

The workflow link was created successfully
*/
type CreateWorkflowLinkIDOK struct {
}

func (o *CreateWorkflowLinkIDOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/wf_links][%d] createWorkflowLinkIdOK ", 200)
}

func (o *CreateWorkflowLinkIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkflowLinkIDBadRequest creates a CreateWorkflowLinkIDBadRequest with default headers values
func NewCreateWorkflowLinkIDBadRequest() *CreateWorkflowLinkIDBadRequest {
	return &CreateWorkflowLinkIDBadRequest{}
}

/*CreateWorkflowLinkIDBadRequest handles this case with default header values.

Bad request - could not decode JSON request, or a parameter (name, URL, or Rules) is missing
*/
type CreateWorkflowLinkIDBadRequest struct {
}

func (o *CreateWorkflowLinkIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/wf_links][%d] createWorkflowLinkIdBadRequest ", 400)
}

func (o *CreateWorkflowLinkIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkflowLinkIDUnauthorized creates a CreateWorkflowLinkIDUnauthorized with default headers values
func NewCreateWorkflowLinkIDUnauthorized() *CreateWorkflowLinkIDUnauthorized {
	return &CreateWorkflowLinkIDUnauthorized{}
}

/*CreateWorkflowLinkIDUnauthorized handles this case with default header values.

The API Key does not provide access to this resource
*/
type CreateWorkflowLinkIDUnauthorized struct {
}

func (o *CreateWorkflowLinkIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/wf_links][%d] createWorkflowLinkIdUnauthorized ", 401)
}

func (o *CreateWorkflowLinkIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkflowLinkIDForbidden creates a CreateWorkflowLinkIDForbidden with default headers values
func NewCreateWorkflowLinkIDForbidden() *CreateWorkflowLinkIDForbidden {
	return &CreateWorkflowLinkIDForbidden{}
}

/*CreateWorkflowLinkIDForbidden handles this case with default header values.

Unsupported request to create resource - links within the same project cannot have the same name and URL.
*/
type CreateWorkflowLinkIDForbidden struct {
}

func (o *CreateWorkflowLinkIDForbidden) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/wf_links][%d] createWorkflowLinkIdForbidden ", 403)
}

func (o *CreateWorkflowLinkIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkflowLinkIDNotFound creates a CreateWorkflowLinkIDNotFound with default headers values
func NewCreateWorkflowLinkIDNotFound() *CreateWorkflowLinkIDNotFound {
	return &CreateWorkflowLinkIDNotFound{}
}

/*CreateWorkflowLinkIDNotFound handles this case with default header values.

Organization or project not found
*/
type CreateWorkflowLinkIDNotFound struct {
}

func (o *CreateWorkflowLinkIDNotFound) Error() string {
	return fmt.Sprintf("[POST /{organization}/projects/{project}/wf_links][%d] createWorkflowLinkIdNotFound ", 404)
}

func (o *CreateWorkflowLinkIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
