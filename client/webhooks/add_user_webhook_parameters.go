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

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// NewAddUserWebhookParams creates a new AddUserWebhookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddUserWebhookParams() *AddUserWebhookParams {
	return &AddUserWebhookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddUserWebhookParamsWithTimeout creates a new AddUserWebhookParams object
// with the ability to set a timeout on a request.
func NewAddUserWebhookParamsWithTimeout(timeout time.Duration) *AddUserWebhookParams {
	return &AddUserWebhookParams{
		timeout: timeout,
	}
}

// NewAddUserWebhookParamsWithContext creates a new AddUserWebhookParams object
// with the ability to set a context for a request.
func NewAddUserWebhookParamsWithContext(ctx context.Context) *AddUserWebhookParams {
	return &AddUserWebhookParams{
		Context: ctx,
	}
}

// NewAddUserWebhookParamsWithHTTPClient creates a new AddUserWebhookParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddUserWebhookParamsWithHTTPClient(client *http.Client) *AddUserWebhookParams {
	return &AddUserWebhookParams{
		HTTPClient: client,
	}
}

/*
AddUserWebhookParams contains all the parameters to send to the API endpoint

	for the add user webhook operation.

	Typically these are written to a http.Request.
*/
type AddUserWebhookParams struct {

	// ContentType.
	ContentType *string

	/* Body.

	   Webhook body
	*/
	Body *models.WebhookSummaryWithPackages

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add user webhook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUserWebhookParams) WithDefaults() *AddUserWebhookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add user webhook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUserWebhookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add user webhook params
func (o *AddUserWebhookParams) WithTimeout(timeout time.Duration) *AddUserWebhookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add user webhook params
func (o *AddUserWebhookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add user webhook params
func (o *AddUserWebhookParams) WithContext(ctx context.Context) *AddUserWebhookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add user webhook params
func (o *AddUserWebhookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add user webhook params
func (o *AddUserWebhookParams) WithHTTPClient(client *http.Client) *AddUserWebhookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add user webhook params
func (o *AddUserWebhookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the add user webhook params
func (o *AddUserWebhookParams) WithContentType(contentType *string) *AddUserWebhookParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the add user webhook params
func (o *AddUserWebhookParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the add user webhook params
func (o *AddUserWebhookParams) WithBody(body *models.WebhookSummaryWithPackages) *AddUserWebhookParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add user webhook params
func (o *AddUserWebhookParams) SetBody(body *models.WebhookSummaryWithPackages) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddUserWebhookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// header param Content-Type
		if err := r.SetHeaderParam("Content-Type", *o.ContentType); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
