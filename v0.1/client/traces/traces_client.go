// Code generated by go-swagger; DO NOT EDIT.

package traces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new traces API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for traces API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	StoredTraces(params *StoredTracesParams, authInfo runtime.ClientAuthInfoWriter) (*StoredTracesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  StoredTraces storeds traces

  Returns complete traces that have already been assembled and stored
*/
func (a *Client) StoredTraces(params *StoredTracesParams, authInfo runtime.ClientAuthInfoWriter) (*StoredTracesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStoredTracesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "storedTraces",
		Method:             "GET",
		PathPattern:        "/{organization}/projects/{project}/stored-traces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StoredTracesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StoredTracesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storedTraces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
