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

	strfmt "github.com/go-openapi/strfmt"
)

// NewListStreamsParams creates a new ListStreamsParams object
// with the default values initialized.
func NewListStreamsParams() *ListStreamsParams {
	var ()
	return &ListStreamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListStreamsParamsWithTimeout creates a new ListStreamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListStreamsParamsWithTimeout(timeout time.Duration) *ListStreamsParams {
	var ()
	return &ListStreamsParams{

		timeout: timeout,
	}
}

// NewListStreamsParamsWithContext creates a new ListStreamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListStreamsParamsWithContext(ctx context.Context) *ListStreamsParams {
	var ()
	return &ListStreamsParams{

		Context: ctx,
	}
}

// NewListStreamsParamsWithHTTPClient creates a new ListStreamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListStreamsParamsWithHTTPClient(client *http.Client) *ListStreamsParams {
	var ()
	return &ListStreamsParams{
		HTTPClient: client,
	}
}

/*ListStreamsParams contains all the parameters to send to the API endpoint
for the list streams operation typically these are written to a http.Request
*/
type ListStreamsParams struct {

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

// WithTimeout adds the timeout to the list streams params
func (o *ListStreamsParams) WithTimeout(timeout time.Duration) *ListStreamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list streams params
func (o *ListStreamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list streams params
func (o *ListStreamsParams) WithContext(ctx context.Context) *ListStreamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list streams params
func (o *ListStreamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list streams params
func (o *ListStreamsParams) WithHTTPClient(client *http.Client) *ListStreamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list streams params
func (o *ListStreamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganization adds the organization to the list streams params
func (o *ListStreamsParams) WithOrganization(organization string) *ListStreamsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the list streams params
func (o *ListStreamsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the list streams params
func (o *ListStreamsParams) WithProject(project string) *ListStreamsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the list streams params
func (o *ListStreamsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *ListStreamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
