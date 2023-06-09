// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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
)

// NewGetUserWebhookDetailParams creates a new GetUserWebhookDetailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserWebhookDetailParams() *GetUserWebhookDetailParams {
	return &GetUserWebhookDetailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserWebhookDetailParamsWithTimeout creates a new GetUserWebhookDetailParams object
// with the ability to set a timeout on a request.
func NewGetUserWebhookDetailParamsWithTimeout(timeout time.Duration) *GetUserWebhookDetailParams {
	return &GetUserWebhookDetailParams{
		timeout: timeout,
	}
}

// NewGetUserWebhookDetailParamsWithContext creates a new GetUserWebhookDetailParams object
// with the ability to set a context for a request.
func NewGetUserWebhookDetailParamsWithContext(ctx context.Context) *GetUserWebhookDetailParams {
	return &GetUserWebhookDetailParams{
		Context: ctx,
	}
}

// NewGetUserWebhookDetailParamsWithHTTPClient creates a new GetUserWebhookDetailParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserWebhookDetailParamsWithHTTPClient(client *http.Client) *GetUserWebhookDetailParams {
	return &GetUserWebhookDetailParams{
		HTTPClient: client,
	}
}

/*
GetUserWebhookDetailParams contains all the parameters to send to the API endpoint

	for the get user webhook detail operation.

	Typically these are written to a http.Request.
*/
type GetUserWebhookDetailParams struct {

	/* WebhookID.

	   Webhook ID

	   Format: uuid
	*/
	WebhookID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user webhook detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserWebhookDetailParams) WithDefaults() *GetUserWebhookDetailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user webhook detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserWebhookDetailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user webhook detail params
func (o *GetUserWebhookDetailParams) WithTimeout(timeout time.Duration) *GetUserWebhookDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user webhook detail params
func (o *GetUserWebhookDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user webhook detail params
func (o *GetUserWebhookDetailParams) WithContext(ctx context.Context) *GetUserWebhookDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user webhook detail params
func (o *GetUserWebhookDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user webhook detail params
func (o *GetUserWebhookDetailParams) WithHTTPClient(client *http.Client) *GetUserWebhookDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user webhook detail params
func (o *GetUserWebhookDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWebhookID adds the webhookID to the get user webhook detail params
func (o *GetUserWebhookDetailParams) WithWebhookID(webhookID strfmt.UUID) *GetUserWebhookDetailParams {
	o.SetWebhookID(webhookID)
	return o
}

// SetWebhookID adds the webhookId to the get user webhook detail params
func (o *GetUserWebhookDetailParams) SetWebhookID(webhookID strfmt.UUID) {
	o.WebhookID = webhookID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserWebhookDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param webhookID
	if err := r.SetPathParam("webhookID", o.WebhookID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
