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

	"github.com/ledor473/lightstep-api-go/v0.2/models"
)

// NewCreateWorkflowLinkParams creates a new CreateWorkflowLinkParams object
// with the default values initialized.
func NewCreateWorkflowLinkParams() *CreateWorkflowLinkParams {
	var ()
	return &CreateWorkflowLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWorkflowLinkParamsWithTimeout creates a new CreateWorkflowLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateWorkflowLinkParamsWithTimeout(timeout time.Duration) *CreateWorkflowLinkParams {
	var ()
	return &CreateWorkflowLinkParams{

		timeout: timeout,
	}
}

// NewCreateWorkflowLinkParamsWithContext creates a new CreateWorkflowLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateWorkflowLinkParamsWithContext(ctx context.Context) *CreateWorkflowLinkParams {
	var ()
	return &CreateWorkflowLinkParams{

		Context: ctx,
	}
}

// NewCreateWorkflowLinkParamsWithHTTPClient creates a new CreateWorkflowLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateWorkflowLinkParamsWithHTTPClient(client *http.Client) *CreateWorkflowLinkParams {
	var ()
	return &CreateWorkflowLinkParams{
		HTTPClient: client,
	}
}

/*CreateWorkflowLinkParams contains all the parameters to send to the API endpoint
for the create workflow link operation typically these are written to a http.Request
*/
type CreateWorkflowLinkParams struct {

	/*Data
	  Workflow Link definition

	*/
	Data *models.ExternalLinkRequestBody
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

// WithTimeout adds the timeout to the create workflow link params
func (o *CreateWorkflowLinkParams) WithTimeout(timeout time.Duration) *CreateWorkflowLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create workflow link params
func (o *CreateWorkflowLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create workflow link params
func (o *CreateWorkflowLinkParams) WithContext(ctx context.Context) *CreateWorkflowLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create workflow link params
func (o *CreateWorkflowLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create workflow link params
func (o *CreateWorkflowLinkParams) WithHTTPClient(client *http.Client) *CreateWorkflowLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create workflow link params
func (o *CreateWorkflowLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create workflow link params
func (o *CreateWorkflowLinkParams) WithData(data *models.ExternalLinkRequestBody) *CreateWorkflowLinkParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create workflow link params
func (o *CreateWorkflowLinkParams) SetData(data *models.ExternalLinkRequestBody) {
	o.Data = data
}

// WithOrganization adds the organization to the create workflow link params
func (o *CreateWorkflowLinkParams) WithOrganization(organization string) *CreateWorkflowLinkParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the create workflow link params
func (o *CreateWorkflowLinkParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the create workflow link params
func (o *CreateWorkflowLinkParams) WithProject(project string) *CreateWorkflowLinkParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create workflow link params
func (o *CreateWorkflowLinkParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWorkflowLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
