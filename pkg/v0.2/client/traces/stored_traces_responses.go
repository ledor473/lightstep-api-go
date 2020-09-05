// Code generated by go-swagger; DO NOT EDIT.

package traces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StoredTracesReader is a Reader for the StoredTraces structure.
type StoredTracesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StoredTracesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStoredTracesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStoredTracesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStoredTracesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStoredTracesOK creates a StoredTracesOK with default headers values
func NewStoredTracesOK() *StoredTracesOK {
	return &StoredTracesOK{}
}

/*StoredTracesOK handles this case with default header values.

JSON representation of a stored trace
*/
type StoredTracesOK struct {
}

func (o *StoredTracesOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/stored-traces][%d] storedTracesOK ", 200)
}

func (o *StoredTracesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStoredTracesBadRequest creates a StoredTracesBadRequest with default headers values
func NewStoredTracesBadRequest() *StoredTracesBadRequest {
	return &StoredTracesBadRequest{}
}

/*StoredTracesBadRequest handles this case with default header values.

Missing required parameter
*/
type StoredTracesBadRequest struct {
}

func (o *StoredTracesBadRequest) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/stored-traces][%d] storedTracesBadRequest ", 400)
}

func (o *StoredTracesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStoredTracesNotFound creates a StoredTracesNotFound with default headers values
func NewStoredTracesNotFound() *StoredTracesNotFound {
	return &StoredTracesNotFound{}
}

/*StoredTracesNotFound handles this case with default header values.

No stored traces found
*/
type StoredTracesNotFound struct {
}

func (o *StoredTracesNotFound) Error() string {
	return fmt.Sprintf("[GET /{organization}/projects/{project}/stored-traces][%d] storedTracesNotFound ", 404)
}

func (o *StoredTracesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
