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

// NewGetTinkerbellActionsVersionDetailsParams creates a new GetTinkerbellActionsVersionDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTinkerbellActionsVersionDetailsParams() *GetTinkerbellActionsVersionDetailsParams {
	return &GetTinkerbellActionsVersionDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTinkerbellActionsVersionDetailsParamsWithTimeout creates a new GetTinkerbellActionsVersionDetailsParams object
// with the ability to set a timeout on a request.
func NewGetTinkerbellActionsVersionDetailsParamsWithTimeout(timeout time.Duration) *GetTinkerbellActionsVersionDetailsParams {
	return &GetTinkerbellActionsVersionDetailsParams{
		timeout: timeout,
	}
}

// NewGetTinkerbellActionsVersionDetailsParamsWithContext creates a new GetTinkerbellActionsVersionDetailsParams object
// with the ability to set a context for a request.
func NewGetTinkerbellActionsVersionDetailsParamsWithContext(ctx context.Context) *GetTinkerbellActionsVersionDetailsParams {
	return &GetTinkerbellActionsVersionDetailsParams{
		Context: ctx,
	}
}

// NewGetTinkerbellActionsVersionDetailsParamsWithHTTPClient creates a new GetTinkerbellActionsVersionDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTinkerbellActionsVersionDetailsParamsWithHTTPClient(client *http.Client) *GetTinkerbellActionsVersionDetailsParams {
	return &GetTinkerbellActionsVersionDetailsParams{
		HTTPClient: client,
	}
}

/*
GetTinkerbellActionsVersionDetailsParams contains all the parameters to send to the API endpoint

	for the get tinkerbell actions version details operation.

	Typically these are written to a http.Request.
*/
type GetTinkerbellActionsVersionDetailsParams struct {

	/* PackageName.

	   Package name
	*/
	PackageName string

	/* RepoName.

	   Repository name
	*/
	RepoName string

	/* Version.

	   Package version
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tinkerbell actions version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTinkerbellActionsVersionDetailsParams) WithDefaults() *GetTinkerbellActionsVersionDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tinkerbell actions version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTinkerbellActionsVersionDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) WithTimeout(timeout time.Duration) *GetTinkerbellActionsVersionDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) WithContext(ctx context.Context) *GetTinkerbellActionsVersionDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) WithHTTPClient(client *http.Client) *GetTinkerbellActionsVersionDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) WithPackageName(packageName string) *GetTinkerbellActionsVersionDetailsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoName adds the repoName to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) WithRepoName(repoName string) *GetTinkerbellActionsVersionDetailsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithVersion adds the version to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) WithVersion(version string) *GetTinkerbellActionsVersionDetailsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get tinkerbell actions version details params
func (o *GetTinkerbellActionsVersionDetailsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetTinkerbellActionsVersionDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
	}

	// path param repoName
	if err := r.SetPathParam("repoName", o.RepoName); err != nil {
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
