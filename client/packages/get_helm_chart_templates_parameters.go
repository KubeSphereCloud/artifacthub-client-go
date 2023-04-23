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

// NewGetHelmChartTemplatesParams creates a new GetHelmChartTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHelmChartTemplatesParams() *GetHelmChartTemplatesParams {
	return &GetHelmChartTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHelmChartTemplatesParamsWithTimeout creates a new GetHelmChartTemplatesParams object
// with the ability to set a timeout on a request.
func NewGetHelmChartTemplatesParamsWithTimeout(timeout time.Duration) *GetHelmChartTemplatesParams {
	return &GetHelmChartTemplatesParams{
		timeout: timeout,
	}
}

// NewGetHelmChartTemplatesParamsWithContext creates a new GetHelmChartTemplatesParams object
// with the ability to set a context for a request.
func NewGetHelmChartTemplatesParamsWithContext(ctx context.Context) *GetHelmChartTemplatesParams {
	return &GetHelmChartTemplatesParams{
		Context: ctx,
	}
}

// NewGetHelmChartTemplatesParamsWithHTTPClient creates a new GetHelmChartTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHelmChartTemplatesParamsWithHTTPClient(client *http.Client) *GetHelmChartTemplatesParams {
	return &GetHelmChartTemplatesParams{
		HTTPClient: client,
	}
}

/*
GetHelmChartTemplatesParams contains all the parameters to send to the API endpoint

	for the get helm chart templates operation.

	Typically these are written to a http.Request.
*/
type GetHelmChartTemplatesParams struct {

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

// WithDefaults hydrates default values in the get helm chart templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHelmChartTemplatesParams) WithDefaults() *GetHelmChartTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get helm chart templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHelmChartTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) WithTimeout(timeout time.Duration) *GetHelmChartTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) WithContext(ctx context.Context) *GetHelmChartTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) WithHTTPClient(client *http.Client) *GetHelmChartTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) WithPackageID(packageID strfmt.UUID) *GetHelmChartTemplatesParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) SetPackageID(packageID strfmt.UUID) {
	o.PackageID = packageID
}

// WithVersion adds the version to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) WithVersion(version string) *GetHelmChartTemplatesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get helm chart templates params
func (o *GetHelmChartTemplatesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetHelmChartTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
