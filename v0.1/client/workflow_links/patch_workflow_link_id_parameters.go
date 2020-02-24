// Code generated by go-swagger; DO NOT EDIT.

package workflow_links

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ledor473/lightstep-api-go/v0.1/models"
)

// NewPatchWorkflowLinkIDParams creates a new PatchWorkflowLinkIDParams object
// with the default values initialized.
func NewPatchWorkflowLinkIDParams() *PatchWorkflowLinkIDParams {
	var ()
	return &PatchWorkflowLinkIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWorkflowLinkIDParamsWithTimeout creates a new PatchWorkflowLinkIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWorkflowLinkIDParamsWithTimeout(timeout time.Duration) *PatchWorkflowLinkIDParams {
	var ()
	return &PatchWorkflowLinkIDParams{

		timeout: timeout,
	}
}

// NewPatchWorkflowLinkIDParamsWithContext creates a new PatchWorkflowLinkIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWorkflowLinkIDParamsWithContext(ctx context.Context) *PatchWorkflowLinkIDParams {
	var ()
	return &PatchWorkflowLinkIDParams{

		Context: ctx,
	}
}

// NewPatchWorkflowLinkIDParamsWithHTTPClient creates a new PatchWorkflowLinkIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWorkflowLinkIDParamsWithHTTPClient(client *http.Client) *PatchWorkflowLinkIDParams {
	var ()
	return &PatchWorkflowLinkIDParams{
		HTTPClient: client,
	}
}

/*PatchWorkflowLinkIDParams contains all the parameters to send to the API endpoint
for the patch workflow link ID operation typically these are written to a http.Request
*/
type PatchWorkflowLinkIDParams struct {

	/*Data
	  Workflow Link definition

	*/
	Data *models.ExternalLinkRequestBody
	/*LinkID
	  Identifier of the workflow link to modify

	*/
	LinkID string
	/*Organization
	  Name of the customer organization

	*/
	Organization string
	/*Project
	  Name of the customer project

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) WithTimeout(timeout time.Duration) *PatchWorkflowLinkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) WithContext(ctx context.Context) *PatchWorkflowLinkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) WithHTTPClient(client *http.Client) *PatchWorkflowLinkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) WithData(data *models.ExternalLinkRequestBody) *PatchWorkflowLinkIDParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) SetData(data *models.ExternalLinkRequestBody) {
	o.Data = data
}

// WithLinkID adds the linkID to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) WithLinkID(linkID string) *PatchWorkflowLinkIDParams {
	o.SetLinkID(linkID)
	return o
}

// SetLinkID adds the linkId to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) SetLinkID(linkID string) {
	o.LinkID = linkID
}

// WithOrganization adds the organization to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) WithOrganization(organization string) *PatchWorkflowLinkIDParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) WithProject(project string) *PatchWorkflowLinkIDParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch workflow link ID params
func (o *PatchWorkflowLinkIDParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWorkflowLinkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param link-id
	if err := r.SetPathParam("link-id", o.LinkID); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
