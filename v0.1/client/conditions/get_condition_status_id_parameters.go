// Code generated by go-swagger; DO NOT EDIT.

package conditions

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

// NewGetConditionStatusIDParams creates a new GetConditionStatusIDParams object
// with the default values initialized.
func NewGetConditionStatusIDParams() *GetConditionStatusIDParams {
	var ()
	return &GetConditionStatusIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConditionStatusIDParamsWithTimeout creates a new GetConditionStatusIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConditionStatusIDParamsWithTimeout(timeout time.Duration) *GetConditionStatusIDParams {
	var ()
	return &GetConditionStatusIDParams{

		timeout: timeout,
	}
}

// NewGetConditionStatusIDParamsWithContext creates a new GetConditionStatusIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConditionStatusIDParamsWithContext(ctx context.Context) *GetConditionStatusIDParams {
	var ()
	return &GetConditionStatusIDParams{

		Context: ctx,
	}
}

// NewGetConditionStatusIDParamsWithHTTPClient creates a new GetConditionStatusIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConditionStatusIDParamsWithHTTPClient(client *http.Client) *GetConditionStatusIDParams {
	var ()
	return &GetConditionStatusIDParams{
		HTTPClient: client,
	}
}

/*GetConditionStatusIDParams contains all the parameters to send to the API endpoint
for the get condition status ID operation typically these are written to a http.Request
*/
type GetConditionStatusIDParams struct {

	/*ConditionID
	  Condition identifier

	*/
	ConditionID string
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

// WithTimeout adds the timeout to the get condition status ID params
func (o *GetConditionStatusIDParams) WithTimeout(timeout time.Duration) *GetConditionStatusIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get condition status ID params
func (o *GetConditionStatusIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get condition status ID params
func (o *GetConditionStatusIDParams) WithContext(ctx context.Context) *GetConditionStatusIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get condition status ID params
func (o *GetConditionStatusIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get condition status ID params
func (o *GetConditionStatusIDParams) WithHTTPClient(client *http.Client) *GetConditionStatusIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get condition status ID params
func (o *GetConditionStatusIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConditionID adds the conditionID to the get condition status ID params
func (o *GetConditionStatusIDParams) WithConditionID(conditionID string) *GetConditionStatusIDParams {
	o.SetConditionID(conditionID)
	return o
}

// SetConditionID adds the conditionId to the get condition status ID params
func (o *GetConditionStatusIDParams) SetConditionID(conditionID string) {
	o.ConditionID = conditionID
}

// WithOrganization adds the organization to the get condition status ID params
func (o *GetConditionStatusIDParams) WithOrganization(organization string) *GetConditionStatusIDParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get condition status ID params
func (o *GetConditionStatusIDParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the get condition status ID params
func (o *GetConditionStatusIDParams) WithProject(project string) *GetConditionStatusIDParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get condition status ID params
func (o *GetConditionStatusIDParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetConditionStatusIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param condition-id
	if err := r.SetPathParam("condition-id", o.ConditionID); err != nil {
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
