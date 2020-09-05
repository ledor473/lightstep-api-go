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

	"github.com/ledor473/lightstep-api-go/pkg/v0.2/models"
)

// NewPatchStreamParams creates a new PatchStreamParams object
// with the default values initialized.
func NewPatchStreamParams() *PatchStreamParams {
	var ()
	return &PatchStreamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchStreamParamsWithTimeout creates a new PatchStreamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchStreamParamsWithTimeout(timeout time.Duration) *PatchStreamParams {
	var ()
	return &PatchStreamParams{

		timeout: timeout,
	}
}

// NewPatchStreamParamsWithContext creates a new PatchStreamParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchStreamParamsWithContext(ctx context.Context) *PatchStreamParams {
	var ()
	return &PatchStreamParams{

		Context: ctx,
	}
}

// NewPatchStreamParamsWithHTTPClient creates a new PatchStreamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchStreamParamsWithHTTPClient(client *http.Client) *PatchStreamParams {
	var ()
	return &PatchStreamParams{
		HTTPClient: client,
	}
}

/*PatchStreamParams contains all the parameters to send to the API endpoint
for the patch stream operation typically these are written to a http.Request
*/
type PatchStreamParams struct {

	/*Data
	  Stream definition

	*/
	Data *models.CreateOrUpdateBody
	/*Organization
	  Name of the customer organization

	*/
	Organization string
	/*Project
	  Name of the customer project

	*/
	Project string
	/*StreamID
	  Identifier of the stream to modify

	*/
	StreamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch stream params
func (o *PatchStreamParams) WithTimeout(timeout time.Duration) *PatchStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch stream params
func (o *PatchStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch stream params
func (o *PatchStreamParams) WithContext(ctx context.Context) *PatchStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch stream params
func (o *PatchStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch stream params
func (o *PatchStreamParams) WithHTTPClient(client *http.Client) *PatchStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch stream params
func (o *PatchStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the patch stream params
func (o *PatchStreamParams) WithData(data *models.CreateOrUpdateBody) *PatchStreamParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the patch stream params
func (o *PatchStreamParams) SetData(data *models.CreateOrUpdateBody) {
	o.Data = data
}

// WithOrganization adds the organization to the patch stream params
func (o *PatchStreamParams) WithOrganization(organization string) *PatchStreamParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the patch stream params
func (o *PatchStreamParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the patch stream params
func (o *PatchStreamParams) WithProject(project string) *PatchStreamParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch stream params
func (o *PatchStreamParams) SetProject(project string) {
	o.Project = project
}

// WithStreamID adds the streamID to the patch stream params
func (o *PatchStreamParams) WithStreamID(streamID string) *PatchStreamParams {
	o.SetStreamID(streamID)
	return o
}

// SetStreamID adds the streamId to the patch stream params
func (o *PatchStreamParams) SetStreamID(streamID string) {
	o.StreamID = streamID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

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
