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
	"github.com/go-openapi/strfmt"

	"github.com/ledor473/lightstep-api-go/pkg/v0.2/models"
)

// NewPostConditionParams creates a new PostConditionParams object
// with the default values initialized.
func NewPostConditionParams() *PostConditionParams {
	var ()
	return &PostConditionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostConditionParamsWithTimeout creates a new PostConditionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostConditionParamsWithTimeout(timeout time.Duration) *PostConditionParams {
	var ()
	return &PostConditionParams{

		timeout: timeout,
	}
}

// NewPostConditionParamsWithContext creates a new PostConditionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostConditionParamsWithContext(ctx context.Context) *PostConditionParams {
	var ()
	return &PostConditionParams{

		Context: ctx,
	}
}

// NewPostConditionParamsWithHTTPClient creates a new PostConditionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostConditionParamsWithHTTPClient(client *http.Client) *PostConditionParams {
	var ()
	return &PostConditionParams{
		HTTPClient: client,
	}
}

/*PostConditionParams contains all the parameters to send to the API endpoint
for the post condition operation typically these are written to a http.Request
*/
type PostConditionParams struct {

	/*Data
	  Condition definition

	*/
	Data *models.ConditionRequestBody
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

// WithTimeout adds the timeout to the post condition params
func (o *PostConditionParams) WithTimeout(timeout time.Duration) *PostConditionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post condition params
func (o *PostConditionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post condition params
func (o *PostConditionParams) WithContext(ctx context.Context) *PostConditionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post condition params
func (o *PostConditionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post condition params
func (o *PostConditionParams) WithHTTPClient(client *http.Client) *PostConditionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post condition params
func (o *PostConditionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the post condition params
func (o *PostConditionParams) WithData(data *models.ConditionRequestBody) *PostConditionParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the post condition params
func (o *PostConditionParams) SetData(data *models.ConditionRequestBody) {
	o.Data = data
}

// WithOrganization adds the organization to the post condition params
func (o *PostConditionParams) WithOrganization(organization string) *PostConditionParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the post condition params
func (o *PostConditionParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithProject adds the project to the post condition params
func (o *PostConditionParams) WithProject(project string) *PostConditionParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the post condition params
func (o *PostConditionParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PostConditionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
