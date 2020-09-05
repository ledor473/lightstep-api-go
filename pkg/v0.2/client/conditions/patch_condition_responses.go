// Code generated by go-swagger; DO NOT EDIT.

package conditions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchConditionReader is a Reader for the PatchCondition structure.
type PatchConditionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConditionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConditionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConditionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConditionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConditionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConditionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConditionOK creates a PatchConditionOK with default headers values
func NewPatchConditionOK() *PatchConditionOK {
	return &PatchConditionOK{}
}

/*PatchConditionOK handles this case with default header values.

The condition was updated successfully
*/
type PatchConditionOK struct {
}

func (o *PatchConditionOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/conditions/{condition-id}][%d] patchConditionOK ", 200)
}

func (o *PatchConditionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConditionBadRequest creates a PatchConditionBadRequest with default headers values
func NewPatchConditionBadRequest() *PatchConditionBadRequest {
	return &PatchConditionBadRequest{}
}

/*PatchConditionBadRequest handles this case with default header values.

One or more parameter(s) are not valid
*/
type PatchConditionBadRequest struct {
}

func (o *PatchConditionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/conditions/{condition-id}][%d] patchConditionBadRequest ", 400)
}

func (o *PatchConditionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConditionUnauthorized creates a PatchConditionUnauthorized with default headers values
func NewPatchConditionUnauthorized() *PatchConditionUnauthorized {
	return &PatchConditionUnauthorized{}
}

/*PatchConditionUnauthorized handles this case with default header values.

The API Key does not provide access to this resource, or the organization name does not exist
*/
type PatchConditionUnauthorized struct {
}

func (o *PatchConditionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/conditions/{condition-id}][%d] patchConditionUnauthorized ", 401)
}

func (o *PatchConditionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConditionNotFound creates a PatchConditionNotFound with default headers values
func NewPatchConditionNotFound() *PatchConditionNotFound {
	return &PatchConditionNotFound{}
}

/*PatchConditionNotFound handles this case with default header values.

Project name is not found
*/
type PatchConditionNotFound struct {
}

func (o *PatchConditionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/conditions/{condition-id}][%d] patchConditionNotFound ", 404)
}

func (o *PatchConditionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConditionInternalServerError creates a PatchConditionInternalServerError with default headers values
func NewPatchConditionInternalServerError() *PatchConditionInternalServerError {
	return &PatchConditionInternalServerError{}
}

/*PatchConditionInternalServerError handles this case with default header values.

Condition identifier not found
*/
type PatchConditionInternalServerError struct {
}

func (o *PatchConditionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/projects/{project}/conditions/{condition-id}][%d] patchConditionInternalServerError ", 500)
}

func (o *PatchConditionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
