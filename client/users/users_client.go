// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetUserProfile(params *GetUserProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserProfileOK, error)

	RegisterUser(params *RegisterUserParams, opts ...ClientOption) (*RegisterUserCreated, error)

	ResetPassword(params *ResetPasswordParams, opts ...ClientOption) (*ResetPasswordNoContent, error)

	ResetPasswordCode(params *ResetPasswordCodeParams, opts ...ClientOption) (*ResetPasswordCodeCreated, error)

	UpdateUserPassword(params *UpdateUserPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserPasswordNoContent, error)

	UpdateUserProfile(params *UpdateUserProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserProfileNoContent, error)

	VerifyEmail(params *VerifyEmailParams, opts ...ClientOption) (*VerifyEmailNoContent, error)

	VerifyPasswordCode(params *VerifyPasswordCodeParams, opts ...ClientOption) (*VerifyPasswordCodeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetUserProfile gets user profile

Get user's profile
*/
func (a *Client) GetUserProfile(params *GetUserProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserProfile",
		Method:             "GET",
		PathPattern:        "/users/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserProfileReader{formats: a.formats},
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
	success, ok := result.(*GetUserProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RegisterUser registers user

Register a new user
*/
func (a *Client) RegisterUser(params *RegisterUserParams, opts ...ClientOption) (*RegisterUserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "registerUser",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegisterUserReader{formats: a.formats},
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
	success, ok := result.(*RegisterUserCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for registerUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResetPassword resets password

Reset the user's password
*/
func (a *Client) ResetPassword(params *ResetPasswordParams, opts ...ClientOption) (*ResetPasswordNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetPassword",
		Method:             "PUT",
		PathPattern:        "/users/reset-password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResetPasswordReader{formats: a.formats},
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
	success, ok := result.(*ResetPasswordNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResetPasswordCode resets password code

Register a code to reset the password
*/
func (a *Client) ResetPasswordCode(params *ResetPasswordCodeParams, opts ...ClientOption) (*ResetPasswordCodeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetPasswordCodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetPasswordCode",
		Method:             "POST",
		PathPattern:        "/users/password-reset-code",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResetPasswordCodeReader{formats: a.formats},
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
	success, ok := result.(*ResetPasswordCodeCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetPasswordCode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserPassword updates user password

Update user's password
*/
func (a *Client) UpdateUserPassword(params *UpdateUserPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserPasswordNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserPassword",
		Method:             "PUT",
		PathPattern:        "/users/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserPasswordReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserPasswordNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserProfile updates user profile

Update user's profile
*/
func (a *Client) UpdateUserProfile(params *UpdateUserProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserProfile",
		Method:             "PUT",
		PathPattern:        "/users/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserProfileReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserProfileNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VerifyEmail verifies email

Verify user's email address
*/
func (a *Client) VerifyEmail(params *VerifyEmailParams, opts ...ClientOption) (*VerifyEmailNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyEmailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "verifyEmail",
		Method:             "POST",
		PathPattern:        "/users/verify-email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VerifyEmailReader{formats: a.formats},
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
	success, ok := result.(*VerifyEmailNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for verifyEmail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VerifyPasswordCode verifies password code

Verify a reset password code
*/
func (a *Client) VerifyPasswordCode(params *VerifyPasswordCodeParams, opts ...ClientOption) (*VerifyPasswordCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyPasswordCodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "verifyPasswordCode",
		Method:             "POST",
		PathPattern:        "/users/verify-password-reset-code",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VerifyPasswordCodeReader{formats: a.formats},
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
	success, ok := result.(*VerifyPasswordCodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for verifyPasswordCode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
