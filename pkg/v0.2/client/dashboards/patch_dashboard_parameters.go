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

// NewPatchDashboardParams creates a new PatchDashboardParams object
// with the default values initialized.
func NewPatchDashboardParams() *PatchDashboardParams {
	var ()
	return &PatchDashboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDashboardParamsWithTimeout creates a new PatchDashboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchDashboardParamsWithTimeout(timeout time.Duration) *PatchDashboardParams {
	var ()
	return &PatchDashboardParams{

		timeout: timeout,
	}
}

// NewPatchDashboardParamsWithContext creates a new PatchDashboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchDashboardParamsWithContext(ctx context.Context) *PatchDashboardParams {
	var ()
	return &PatchDashboardParams{

		Context: ctx,
	}
}

// NewPatchDashboardParamsWithHTTPClient creates a new PatchDashboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchDashboardParamsWithHTTPClient(client *http.Client) *PatchDashboardParams {
	var ()
	return &PatchDashboardParams{
		HTTPClient: client,
	}
}

/*PatchDashboardParams contains all the parameters to send to the API endpoint
for the patch dashboard operation typically these are written to a http.Request
*/
type PatchDashboardParams struct {

	/*DashboardID
	  Identifier of the dashboard to modify

	*/
	DashboardID string
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

// WithTimeout adds the timeout to the patch dashboard params
func (o *PatchDashboardParams) WithTimeout(timeout time.Duration) *PatchDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch dashboard params
func (o *PatchDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch dashboard params
func (o *PatchDashboardParams) WithContext(ctx context.Context) *PatchDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch dashboard params
func (o *PatchDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch dashboard params
func (o *PatchDashboardParams) WithHTTPClient(client *http.Client) *PatchDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch dashboard params
func (o *PatchDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardID adds the dashboardID to the patch dashboard params
func (o *PatchDashboardParams) WithDashboardID(dashboardID string) *PatchDashboardParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the patch dashboard params
func (o *PatchDashboardParams) SetDashboardID(dashboardID string) {
	o.DashboardID = dashboardID
}

// WithData adds the data to the patch dashboard params
func (o *PatchDashboardParams) WithData(data *models.DashboardRequestBody) *PatchDashboardParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the patch dashboard params
func (o *PatchDashboardParams) SetData(data *models.DashboardRequestBody) {
	o.Data = data
}

// WithOrganization adds the organization to the patch dashboard params
func (o *PatchDashboardParams) WithOrganization(organization string) *PatchDashboardParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the patch dashboard params
func (o *PatchDashboardParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the patch dashboard params
func (o *PatchDashboardParams) WithProject(project string) *PatchDashboardParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch dashboard params
func (o *PatchDashboardParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dashboard-id
	if err := r.SetPathParam("dashboard-id", o.DashboardID); err != nil {
		return err
	}

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
