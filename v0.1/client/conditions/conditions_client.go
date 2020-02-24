// Code generated by go-swagger; DO NOT EDIT.

package conditions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new conditions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for conditions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCondition(params *DeleteConditionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConditionNoContent, error)

	GetCondition(params *GetConditionParams, authInfo runtime.ClientAuthInfoWriter) (*GetConditionOK, error)

	GetConditionStatus(params *GetConditionStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetConditionStatusOK, error)

	ListConditions(params *ListConditionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListConditionsOK, error)

	ListConditionsForStream(params *ListConditionsForStreamParams, authInfo runtime.ClientAuthInfoWriter) (*ListConditionsForStreamOK, error)

	PatchCondition(params *PatchConditionParams, authInfo runtime.ClientAuthInfoWriter) (*PatchConditionOK, error)

	PostCondition(params *PostConditionParams, authInfo runtime.ClientAuthInfoWriter) (*PostConditionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteCondition deletes condition

  Deletes an existing condition
*/
func (a *Client) DeleteCondition(params *DeleteConditionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConditionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConditionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCondition",
		Method:             "DELETE",
		PathPattern:        "/{organization}/projects/{project}/conditions/{condition-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteConditionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteConditionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCondition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCondition gets condition

  Returns information about a specific condition
*/
func (a *Client) GetCondition(params *GetConditionParams, authInfo runtime.ClientAuthInfoWriter) (*GetConditionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConditionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCondition",
		Method:             "GET",
		PathPattern:        "/{organization}/projects/{project}/conditions/{condition-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConditionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConditionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCondition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConditionStatus gets condition status

  Returns status information about a specific condition
*/
func (a *Client) GetConditionStatus(params *GetConditionStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetConditionStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConditionStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConditionStatus",
		Method:             "GET",
		PathPattern:        "/{organization}/projects/{project}/conditions/{condition-id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConditionStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConditionStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConditionStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListConditions lists conditions

  Returns information about all conditions in a project
*/
func (a *Client) ListConditions(params *ListConditionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListConditionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConditionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listConditions",
		Method:             "GET",
		PathPattern:        "/{organization}/projects/{project}/conditions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListConditionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListConditionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listConditions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListConditionsForStream lists conditions for stream

  Returns information about all conditions in a specific stream
*/
func (a *Client) ListConditionsForStream(params *ListConditionsForStreamParams, authInfo runtime.ClientAuthInfoWriter) (*ListConditionsForStreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConditionsForStreamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listConditionsForStream",
		Method:             "GET",
		PathPattern:        "/{organization}/projects/{project}/searches/{stream-id}/conditions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListConditionsForStreamReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListConditionsForStreamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listConditionsForStream: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchCondition updates condition

  Modifies the settings for an existing condition
*/
func (a *Client) PatchCondition(params *PatchConditionParams, authInfo runtime.ClientAuthInfoWriter) (*PatchConditionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchConditionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchCondition",
		Method:             "PATCH",
		PathPattern:        "/{organization}/projects/{project}/conditions/{condition-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchConditionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchConditionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchCondition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCondition creates condition

  Creates a new condition
*/
func (a *Client) PostCondition(params *PostConditionParams, authInfo runtime.ClientAuthInfoWriter) (*PostConditionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostConditionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCondition",
		Method:             "POST",
		PathPattern:        "/{organization}/projects/{project}/conditions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostConditionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostConditionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postCondition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
