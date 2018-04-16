// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddRepository adds repository to watch
*/
func (a *Client) AddRepository(params *AddRepositoryParams, authInfo runtime.ClientAuthInfoWriter) (*AddRepositoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRepositoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "add_repository",
		Method:             "POST",
		PathPattern:        "/repositories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddRepositoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddRepositoryOK), nil

}

/*
GetSystemPruneCandidates gets list of candidates for pruning
*/
func (a *Client) GetSystemPruneCandidates(params *GetSystemPruneCandidatesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSystemPruneCandidatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemPruneCandidatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_system_prune_candidates",
		Method:             "GET",
		PathPattern:        "/system/prune/{resourcetype}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSystemPruneCandidatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSystemPruneCandidatesOK), nil

}

/*
GetSystemPruneResourcetypes gets list of resources that can be pruned
*/
func (a *Client) GetSystemPruneResourcetypes(params *GetSystemPruneResourcetypesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSystemPruneResourcetypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemPruneResourcetypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_system_prune_resourcetypes",
		Method:             "GET",
		PathPattern:        "/system/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSystemPruneResourcetypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSystemPruneResourcetypesOK), nil

}

/*
PostSystemPruneCandidates performs pruning on input resource name
*/
func (a *Client) PostSystemPruneCandidates(params *PostSystemPruneCandidatesParams, authInfo runtime.ClientAuthInfoWriter) (*PostSystemPruneCandidatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSystemPruneCandidatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post_system_prune_candidates",
		Method:             "POST",
		PathPattern:        "/system/prune/{resourcetype}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostSystemPruneCandidatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSystemPruneCandidatesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}