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

// NewGetContainerImageVersionDetailsParams creates a new GetContainerImageVersionDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContainerImageVersionDetailsParams() *GetContainerImageVersionDetailsParams {
	return &GetContainerImageVersionDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContainerImageVersionDetailsParamsWithTimeout creates a new GetContainerImageVersionDetailsParams object
// with the ability to set a timeout on a request.
func NewGetContainerImageVersionDetailsParamsWithTimeout(timeout time.Duration) *GetContainerImageVersionDetailsParams {
	return &GetContainerImageVersionDetailsParams{
		timeout: timeout,
	}
}

// NewGetContainerImageVersionDetailsParamsWithContext creates a new GetContainerImageVersionDetailsParams object
// with the ability to set a context for a request.
func NewGetContainerImageVersionDetailsParamsWithContext(ctx context.Context) *GetContainerImageVersionDetailsParams {
	return &GetContainerImageVersionDetailsParams{
		Context: ctx,
	}
}

// NewGetContainerImageVersionDetailsParamsWithHTTPClient creates a new GetContainerImageVersionDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContainerImageVersionDetailsParamsWithHTTPClient(client *http.Client) *GetContainerImageVersionDetailsParams {
	return &GetContainerImageVersionDetailsParams{
		HTTPClient: client,
	}
}

/*
GetContainerImageVersionDetailsParams contains all the parameters to send to the API endpoint

	for the get container image version details operation.

	Typically these are written to a http.Request.
*/
type GetContainerImageVersionDetailsParams struct {

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

// WithDefaults hydrates default values in the get container image version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContainerImageVersionDetailsParams) WithDefaults() *GetContainerImageVersionDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get container image version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContainerImageVersionDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) WithTimeout(timeout time.Duration) *GetContainerImageVersionDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) WithContext(ctx context.Context) *GetContainerImageVersionDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) WithHTTPClient(client *http.Client) *GetContainerImageVersionDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) WithPackageName(packageName string) *GetContainerImageVersionDetailsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoName adds the repoName to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) WithRepoName(repoName string) *GetContainerImageVersionDetailsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithVersion adds the version to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) WithVersion(version string) *GetContainerImageVersionDetailsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get container image version details params
func (o *GetContainerImageVersionDetailsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetContainerImageVersionDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
