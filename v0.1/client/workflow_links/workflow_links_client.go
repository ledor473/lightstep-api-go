// Code generated by go-swagger; DO NOT EDIT.

package workflow_links

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workflow links API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow links API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateWorkflowLink(params *CreateWorkflowLinkParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkflowLinkOK, error)

	DeleteWorkflowLink(params *DeleteWorkflowLinkParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWorkflowLinkNoContent, error)

	GetWorkflowLink(params *GetWorkflowLinkParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowLinkOK, error)

	ListWorkflowLinks(params *ListWorkflowLinksParams, authInfo runtime.ClientAuthInfoWriter) (*ListWorkflowLinksOK, error)

	PatchWorkflowLink(params *PatchWorkflowLinkParams, authInfo runtime.ClientAuthInfoWriter) (*PatchWorkflowLinkOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateWorkflowLink creates workflow link

  Creates a new workflow link. Links within a project must have a unique combination of Name and URL. Admin or Member privileges are required to created workflow links.
*/
func (a *Client) CreateWorkflowLink(params *CreateWorkflowLinkParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkflowLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateWorkflowLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createWorkflowLink",
		Method:             "POST",
		PathPattern:        "/{organization}/projects/{project}/wf_links",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateWorkflowLinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateWorkflowLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createWorkflowLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteWorkflowLink deletes workflow link

  Deletes an existing workflow link.
*/
func (a *Client) DeleteWorkflowLink(params *DeleteWorkflowLinkParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWorkflowLinkNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkflowLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWorkflowLink",
		Method:             "DELETE",
		PathPattern:        "/{organization}/projects/{project}/wf_links/{link-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteWorkflowLinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteWorkflowLinkNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteWorkflowLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkflowLink gets workflow link

  Returns information on a specific workflow link definition within a project.
*/
func (a *Client) GetWorkflowLink(params *GetWorkflowLinkParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkflowLink",
		Method:             "GET",
		PathPattern:        "/{organization}/projects/{project}/wf_links/{link-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkflowLinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkflowLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkflowLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListWorkflowLinks lists workflow links

  Returns information on all workflow link definitions within a project
*/
func (a *Client) ListWorkflowLinks(params *ListWorkflowLinksParams, authInfo runtime.ClientAuthInfoWriter) (*ListWorkflowLinksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListWorkflowLinksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listWorkflowLinks",
		Method:             "GET",
		PathPattern:        "/{organization}/projects/{project}/wf_links",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListWorkflowLinksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListWorkflowLinksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listWorkflowLinks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchWorkflowLink updates workflow link

  Updates the workflow link with a new name or URL (if applicable), or replaces the set of rules on the workflow link. If a non-empty parameter (i.e., name, URL, or Rules) is provided, the field will will be overwritten with the new value. Links within a project must have a unique combination of name and URL. Admin or Member privileges are required to created workflow links.
*/
func (a *Client) PatchWorkflowLink(params *PatchWorkflowLinkParams, authInfo runtime.ClientAuthInfoWriter) (*PatchWorkflowLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchWorkflowLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchWorkflowLink",
		Method:             "PATCH",
		PathPattern:        "/{organization}/projects/{project}/wf_links/{link-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchWorkflowLinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchWorkflowLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchWorkflowLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
