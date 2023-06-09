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

// NewGetChartValuesParams creates a new GetChartValuesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetChartValuesParams() *GetChartValuesParams {
	return &GetChartValuesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetChartValuesParamsWithTimeout creates a new GetChartValuesParams object
// with the ability to set a timeout on a request.
func NewGetChartValuesParamsWithTimeout(timeout time.Duration) *GetChartValuesParams {
	return &GetChartValuesParams{
		timeout: timeout,
	}
}

// NewGetChartValuesParamsWithContext creates a new GetChartValuesParams object
// with the ability to set a context for a request.
func NewGetChartValuesParamsWithContext(ctx context.Context) *GetChartValuesParams {
	return &GetChartValuesParams{
		Context: ctx,
	}
}

// NewGetChartValuesParamsWithHTTPClient creates a new GetChartValuesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetChartValuesParamsWithHTTPClient(client *http.Client) *GetChartValuesParams {
	return &GetChartValuesParams{
		HTTPClient: client,
	}
}

/*
GetChartValuesParams contains all the parameters to send to the API endpoint

	for the get chart values operation.

	Typically these are written to a http.Request.
*/
type GetChartValuesParams struct {

	/* PackageID.

	   Package ID

	   Format: uuid
	*/
	PackageID strfmt.UUID

	/* Version.

	   Package version
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get chart values params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChartValuesParams) WithDefaults() *GetChartValuesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get chart values params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChartValuesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get chart values params
func (o *GetChartValuesParams) WithTimeout(timeout time.Duration) *GetChartValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chart values params
func (o *GetChartValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chart values params
func (o *GetChartValuesParams) WithContext(ctx context.Context) *GetChartValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chart values params
func (o *GetChartValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get chart values params
func (o *GetChartValuesParams) WithHTTPClient(client *http.Client) *GetChartValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get chart values params
func (o *GetChartValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the get chart values params
func (o *GetChartValuesParams) WithPackageID(packageID strfmt.UUID) *GetChartValuesParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the get chart values params
func (o *GetChartValuesParams) SetPackageID(packageID strfmt.UUID) {
	o.PackageID = packageID
}

// WithVersion adds the version to the get chart values params
func (o *GetChartValuesParams) WithVersion(version string) *GetChartValuesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get chart values params
func (o *GetChartValuesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetChartValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packageID
	if err := r.SetPathParam("packageID", o.PackageID.String()); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
