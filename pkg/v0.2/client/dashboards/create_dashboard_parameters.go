// Code generated by go-swagger; DO NOT EDIT.

package dashboards

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

// NewCreateDashboardParams creates a new CreateDashboardParams object
// with the default values initialized.
func NewCreateDashboardParams() *CreateDashboardParams {
	var ()
	return &CreateDashboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDashboardParamsWithTimeout creates a new CreateDashboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDashboardParamsWithTimeout(timeout time.Duration) *CreateDashboardParams {
	var ()
	return &CreateDashboardParams{

		timeout: timeout,
	}
}

// NewCreateDashboardParamsWithContext creates a new CreateDashboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDashboardParamsWithContext(ctx context.Context) *CreateDashboardParams {
	var ()
	return &CreateDashboardParams{

		Context: ctx,
	}
}

// NewCreateDashboardParamsWithHTTPClient creates a new CreateDashboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDashboardParamsWithHTTPClient(client *http.Client) *CreateDashboardParams {
	var ()
	return &CreateDashboardParams{
		HTTPClient: client,
	}
}

/*CreateDashboardParams contains all the parameters to send to the API endpoint
for the create dashboard operation typically these are written to a http.Request
*/
type CreateDashboardParams struct {

	/*Data
	  Dashboard definition

	*/
	Data *models.DashboardRequestBody
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

// WithTimeout adds the timeout to the create dashboard params
func (o *CreateDashboardParams) WithTimeout(timeout time.Duration) *CreateDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create dashboard params
func (o *CreateDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create dashboard params
func (o *CreateDashboardParams) WithContext(ctx context.Context) *CreateDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create dashboard params
func (o *CreateDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create dashboard params
func (o *CreateDashboardParams) WithHTTPClient(client *http.Client) *CreateDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create dashboard params
func (o *CreateDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create dashboard params
func (o *CreateDashboardParams) WithData(data *models.DashboardRequestBody) *CreateDashboardParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create dashboard params
func (o *CreateDashboardParams) SetData(data *models.DashboardRequestBody) {
	o.Data = data
}

// WithOrganization adds the organization to the create dashboard params
func (o *CreateDashboardParams) WithOrganization(organization string) *CreateDashboardParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the create dashboard params
func (o *CreateDashboardParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the create dashboard params
func (o *CreateDashboardParams) WithProject(project string) *CreateDashboardParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create dashboard params
func (o *CreateDashboardParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
