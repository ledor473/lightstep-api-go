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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDashboardIDParams creates a new GetDashboardIDParams object
// with the default values initialized.
func NewGetDashboardIDParams() *GetDashboardIDParams {
	var ()
	return &GetDashboardIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardIDParamsWithTimeout creates a new GetDashboardIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDashboardIDParamsWithTimeout(timeout time.Duration) *GetDashboardIDParams {
	var ()
	return &GetDashboardIDParams{

		timeout: timeout,
	}
}

// NewGetDashboardIDParamsWithContext creates a new GetDashboardIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDashboardIDParamsWithContext(ctx context.Context) *GetDashboardIDParams {
	var ()
	return &GetDashboardIDParams{

		Context: ctx,
	}
}

// NewGetDashboardIDParamsWithHTTPClient creates a new GetDashboardIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDashboardIDParamsWithHTTPClient(client *http.Client) *GetDashboardIDParams {
	var ()
	return &GetDashboardIDParams{
		HTTPClient: client,
	}
}

/*GetDashboardIDParams contains all the parameters to send to the API endpoint
for the get dashboard ID operation typically these are written to a http.Request
*/
type GetDashboardIDParams struct {

	/*DashboardID
	  Dashboard identifier

	*/
	DashboardID string
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

// WithTimeout adds the timeout to the get dashboard ID params
func (o *GetDashboardIDParams) WithTimeout(timeout time.Duration) *GetDashboardIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard ID params
func (o *GetDashboardIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard ID params
func (o *GetDashboardIDParams) WithContext(ctx context.Context) *GetDashboardIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard ID params
func (o *GetDashboardIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard ID params
func (o *GetDashboardIDParams) WithHTTPClient(client *http.Client) *GetDashboardIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard ID params
func (o *GetDashboardIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardID adds the dashboardID to the get dashboard ID params
func (o *GetDashboardIDParams) WithDashboardID(dashboardID string) *GetDashboardIDParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the get dashboard ID params
func (o *GetDashboardIDParams) SetDashboardID(dashboardID string) {
	o.DashboardID = dashboardID
}

// WithOrganization adds the organization to the get dashboard ID params
func (o *GetDashboardIDParams) WithOrganization(organization string) *GetDashboardIDParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get dashboard ID params
func (o *GetDashboardIDParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get dashboard ID params
func (o *GetDashboardIDParams) WithProject(project string) *GetDashboardIDParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get dashboard ID params
func (o *GetDashboardIDParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dashboard-id
	if err := r.SetPathParam("dashboard-id", o.DashboardID); err != nil {
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
