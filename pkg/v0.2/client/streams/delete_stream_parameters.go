// Code generated by go-swagger; DO NOT EDIT.

package streams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteStreamParams creates a new DeleteStreamParams object
// with the default values initialized.
func NewDeleteStreamParams() *DeleteStreamParams {
	var ()
	return &DeleteStreamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStreamParamsWithTimeout creates a new DeleteStreamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteStreamParamsWithTimeout(timeout time.Duration) *DeleteStreamParams {
	var ()
	return &DeleteStreamParams{

		timeout: timeout,
	}
}

// NewDeleteStreamParamsWithContext creates a new DeleteStreamParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteStreamParamsWithContext(ctx context.Context) *DeleteStreamParams {
	var ()
	return &DeleteStreamParams{

		Context: ctx,
	}
}

// NewDeleteStreamParamsWithHTTPClient creates a new DeleteStreamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteStreamParamsWithHTTPClient(client *http.Client) *DeleteStreamParams {
	var ()
	return &DeleteStreamParams{
		HTTPClient: client,
	}
}

/*DeleteStreamParams contains all the parameters to send to the API endpoint
for the delete stream operation typically these are written to a http.Request
*/
type DeleteStreamParams struct {

	/*Organization
	  Name of the customer organization

	*/
	Organization string
	/*Project
	  Name of the customer project

	*/
	Project string
	/*StreamID
	  Identifier of the stream to delete

	*/
	StreamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete stream params
func (o *DeleteStreamParams) WithTimeout(timeout time.Duration) *DeleteStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete stream params
func (o *DeleteStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete stream params
func (o *DeleteStreamParams) WithContext(ctx context.Context) *DeleteStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete stream params
func (o *DeleteStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete stream params
func (o *DeleteStreamParams) WithHTTPClient(client *http.Client) *DeleteStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete stream params
func (o *DeleteStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganization adds the organization to the delete stream params
func (o *DeleteStreamParams) WithOrganization(organization string) *DeleteStreamParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete stream params
func (o *DeleteStreamParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the delete stream params
func (o *DeleteStreamParams) WithProject(project string) *DeleteStreamParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete stream params
func (o *DeleteStreamParams) SetProject(project string) {
	o.Project = project
}

// WithStreamID adds the streamID to the delete stream params
func (o *DeleteStreamParams) WithStreamID(streamID string) *DeleteStreamParams {
	o.SetStreamID(streamID)
	return o
}

// SetStreamID adds the streamId to the delete stream params
func (o *DeleteStreamParams) SetStreamID(streamID string) {
	o.StreamID = streamID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param stream-id
	if err := r.SetPathParam("stream-id", o.StreamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
