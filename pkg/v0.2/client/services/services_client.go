// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the services client
type API interface {
	/*
	   ListServices lists services

	   Returns all reporting services for a project*/
	ListServices(ctx context.Context, params *ListServicesParams) (*ListServicesOK, error)
}

// New creates a new services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
ListServices lists services

Returns all reporting services for a project
*/
func (a *Client) ListServices(ctx context.Context, params *ListServicesParams) (*ListServicesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listServices",
		Method:             "GET",
		PathPattern:        "/{organization}/projects/{project}/directory/services",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListServicesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListServicesOK), nil

}
