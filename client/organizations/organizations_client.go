// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AcceptOrganizationInvitation(params *AcceptOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcceptOrganizationInvitationNoContent, error)

	AddOrganization(params *AddOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrganizationCreated, error)

	AddOrganizationMember(params *AddOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrganizationMemberCreated, error)

	DeleteOrganization(params *DeleteOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationNoContent, error)

	DeleteOrganizationMember(params *DeleteOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationMemberNoContent, error)

	GetAllowedActions(params *GetAllowedActionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllowedActionsOK, error)

	GetOrganizationAuthPolicy(params *GetOrganizationAuthPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationAuthPolicyOK, error)

	GetOrganizationMembers(params *GetOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationMembersOK, error)

	GetOrganizationProfile(params *GetOrganizationProfileParams, opts ...ClientOption) (*GetOrganizationProfileOK, error)

	GetUserOrganizations(params *GetUserOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserOrganizationsOK, error)

	UpdateOrganizationAuthPolicy(params *UpdateOrganizationAuthPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrganizationAuthPolicyNoContent, error)

	UpdateOrganizationProfile(params *UpdateOrganizationProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrganizationProfileNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AcceptOrganizationInvitation accepts organization invitation

Confirm user's membership to an organization
*/
func (a *Client) AcceptOrganizationInvitation(params *AcceptOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AcceptOrganizationInvitationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcceptOrganizationInvitationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "acceptOrganizationInvitation",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/accept-invitation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptOrganizationInvitationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AcceptOrganizationInvitationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for acceptOrganizationInvitation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddOrganization adds organization

Register new organization
*/
func (a *Client) AddOrganization(params *AddOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrganizationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addOrganization",
		Method:             "POST",
		PathPattern:        "/orgs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddOrganizationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddOrganizationMember adds organization member

Add a new member to the organization
*/
func (a *Client) AddOrganizationMember(params *AddOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrganizationMemberCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrganizationMemberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addOrganizationMember",
		Method:             "POST",
		PathPattern:        "/orgs/{orgName}/member/{userAlias}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddOrganizationMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddOrganizationMemberCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addOrganizationMember: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteOrganization deletes organization

Delete organization
*/
func (a *Client) DeleteOrganization(params *DeleteOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteOrganization",
		Method:             "DELETE",
		PathPattern:        "/orgs/{orgName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOrganizationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteOrganizationMember deletes organization member

Delete a member from the organization
*/
func (a *Client) DeleteOrganizationMember(params *DeleteOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationMemberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationMemberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteOrganizationMember",
		Method:             "DELETE",
		PathPattern:        "/orgs/{orgName}/member/{userAlias}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOrganizationMemberNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOrganizationMember: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllowedActions gets allowed actions

Get actions which user is allowed to perform in the provided organization
*/
func (a *Client) GetAllowedActions(params *GetAllowedActionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllowedActionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllowedActionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllowedActions",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/user-allowed-actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllowedActionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllowedActionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllowedActions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrganizationAuthPolicy gets organization auth policy

Get organization's authorization policy
*/
func (a *Client) GetOrganizationAuthPolicy(params *GetOrganizationAuthPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationAuthPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationAuthPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationAuthPolicy",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/authorization-policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationAuthPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationAuthPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationAuthPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrganizationMembers gets organization members

Get organization members
*/
func (a *Client) GetOrganizationMembers(params *GetOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationMembers",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationMembers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrganizationProfile gets organization profile

Get organization profile
*/
func (a *Client) GetOrganizationProfile(params *GetOrganizationProfileParams, opts ...ClientOption) (*GetOrganizationProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationProfile",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserOrganizations gets user organizations

Get organizations the user belongs to
*/
func (a *Client) GetUserOrganizations(params *GetUserOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserOrganizationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserOrganizations",
		Method:             "GET",
		PathPattern:        "/orgs/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserOrganizationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserOrganizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserOrganizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrganizationAuthPolicy updates organization auth policy

Update organization's authorization policy
*/
func (a *Client) UpdateOrganizationAuthPolicy(params *UpdateOrganizationAuthPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrganizationAuthPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationAuthPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrganizationAuthPolicy",
		Method:             "PUT",
		PathPattern:        "/orgs/{orgName}/authorization-policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrganizationAuthPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOrganizationAuthPolicyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrganizationAuthPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrganizationProfile updates organization profile

Updates organization profile
*/
func (a *Client) UpdateOrganizationProfile(params *UpdateOrganizationProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrganizationProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrganizationProfile",
		Method:             "PUT",
		PathPattern:        "/orgs/{orgName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrganizationProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOrganizationProfileNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrganizationProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}