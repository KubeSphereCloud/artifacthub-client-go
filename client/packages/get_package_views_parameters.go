// Code generated by go-swagger; DO NOT EDIT.

package packages

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

// NewGetPackageViewsParams creates a new GetPackageViewsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPackageViewsParams() *GetPackageViewsParams {
	return &GetPackageViewsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackageViewsParamsWithTimeout creates a new GetPackageViewsParams object
// with the ability to set a timeout on a request.
func NewGetPackageViewsParamsWithTimeout(timeout time.Duration) *GetPackageViewsParams {
	return &GetPackageViewsParams{
		timeout: timeout,
	}
}

// NewGetPackageViewsParamsWithContext creates a new GetPackageViewsParams object
// with the ability to set a context for a request.
func NewGetPackageViewsParamsWithContext(ctx context.Context) *GetPackageViewsParams {
	return &GetPackageViewsParams{
		Context: ctx,
	}
}

// NewGetPackageViewsParamsWithHTTPClient creates a new GetPackageViewsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPackageViewsParamsWithHTTPClient(client *http.Client) *GetPackageViewsParams {
	return &GetPackageViewsParams{
		HTTPClient: client,
	}
}

/*
GetPackageViewsParams contains all the parameters to send to the API endpoint

	for the get package views operation.

	Typically these are written to a http.Request.
*/
type GetPackageViewsParams struct {

	/* PackageID.

	   Package ID

	   Format: uuid
	*/
	PackageID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get package views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackageViewsParams) WithDefaults() *GetPackageViewsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get package views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackageViewsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get package views params
func (o *GetPackageViewsParams) WithTimeout(timeout time.Duration) *GetPackageViewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get package views params
func (o *GetPackageViewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get package views params
func (o *GetPackageViewsParams) WithContext(ctx context.Context) *GetPackageViewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get package views params
func (o *GetPackageViewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get package views params
func (o *GetPackageViewsParams) WithHTTPClient(client *http.Client) *GetPackageViewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get package views params
func (o *GetPackageViewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the get package views params
func (o *GetPackageViewsParams) WithPackageID(packageID strfmt.UUID) *GetPackageViewsParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the get package views params
func (o *GetPackageViewsParams) SetPackageID(packageID strfmt.UUID) {
	o.PackageID = packageID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackageViewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packageID
	if err := r.SetPathParam("packageID", o.PackageID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
