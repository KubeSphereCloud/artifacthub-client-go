// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewVerifyEmailParams creates a new VerifyEmailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVerifyEmailParams() *VerifyEmailParams {
	return &VerifyEmailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyEmailParamsWithTimeout creates a new VerifyEmailParams object
// with the ability to set a timeout on a request.
func NewVerifyEmailParamsWithTimeout(timeout time.Duration) *VerifyEmailParams {
	return &VerifyEmailParams{
		timeout: timeout,
	}
}

// NewVerifyEmailParamsWithContext creates a new VerifyEmailParams object
// with the ability to set a context for a request.
func NewVerifyEmailParamsWithContext(ctx context.Context) *VerifyEmailParams {
	return &VerifyEmailParams{
		Context: ctx,
	}
}

// NewVerifyEmailParamsWithHTTPClient creates a new VerifyEmailParams object
// with the ability to set a custom HTTPClient for a request.
func NewVerifyEmailParamsWithHTTPClient(client *http.Client) *VerifyEmailParams {
	return &VerifyEmailParams{
		HTTPClient: client,
	}
}

/*
VerifyEmailParams contains all the parameters to send to the API endpoint

	for the verify email operation.

	Typically these are written to a http.Request.
*/
type VerifyEmailParams struct {

	// ContentType.
	ContentType *string

	// Body.
	Body *models.UsersVerifyEmailRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the verify email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyEmailParams) WithDefaults() *VerifyEmailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the verify email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyEmailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the verify email params
func (o *VerifyEmailParams) WithTimeout(timeout time.Duration) *VerifyEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify email params
func (o *VerifyEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify email params
func (o *VerifyEmailParams) WithContext(ctx context.Context) *VerifyEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify email params
func (o *VerifyEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify email params
func (o *VerifyEmailParams) WithHTTPClient(client *http.Client) *VerifyEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify email params
func (o *VerifyEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the verify email params
func (o *VerifyEmailParams) WithContentType(contentType *string) *VerifyEmailParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the verify email params
func (o *VerifyEmailParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the verify email params
func (o *VerifyEmailParams) WithBody(body *models.UsersVerifyEmailRequest) *VerifyEmailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the verify email params
func (o *VerifyEmailParams) SetBody(body *models.UsersVerifyEmailRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
