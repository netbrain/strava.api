// Code generated by go-swagger; DO NOT EDIT.

package gears

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new gears API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gears API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetGearByID(params *GetGearByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetGearByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetGearByID gets equipment

  Returns an equipment using its identifier.
*/
func (a *Client) GetGearByID(params *GetGearByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetGearByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGearByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGearById",
		Method:             "GET",
		PathPattern:        "/gear/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGearByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGearByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetGearByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
