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

// NewGetKnativeClientPluginsVersionDetailsParams creates a new GetKnativeClientPluginsVersionDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKnativeClientPluginsVersionDetailsParams() *GetKnativeClientPluginsVersionDetailsParams {
	return &GetKnativeClientPluginsVersionDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnativeClientPluginsVersionDetailsParamsWithTimeout creates a new GetKnativeClientPluginsVersionDetailsParams object
// with the ability to set a timeout on a request.
func NewGetKnativeClientPluginsVersionDetailsParamsWithTimeout(timeout time.Duration) *GetKnativeClientPluginsVersionDetailsParams {
	return &GetKnativeClientPluginsVersionDetailsParams{
		timeout: timeout,
	}
}

// NewGetKnativeClientPluginsVersionDetailsParamsWithContext creates a new GetKnativeClientPluginsVersionDetailsParams object
// with the ability to set a context for a request.
func NewGetKnativeClientPluginsVersionDetailsParamsWithContext(ctx context.Context) *GetKnativeClientPluginsVersionDetailsParams {
	return &GetKnativeClientPluginsVersionDetailsParams{
		Context: ctx,
	}
}

// NewGetKnativeClientPluginsVersionDetailsParamsWithHTTPClient creates a new GetKnativeClientPluginsVersionDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKnativeClientPluginsVersionDetailsParamsWithHTTPClient(client *http.Client) *GetKnativeClientPluginsVersionDetailsParams {
	return &GetKnativeClientPluginsVersionDetailsParams{
		HTTPClient: client,
	}
}

/*
GetKnativeClientPluginsVersionDetailsParams contains all the parameters to send to the API endpoint

	for the get knative client plugins version details operation.

	Typically these are written to a http.Request.
*/
type GetKnativeClientPluginsVersionDetailsParams struct {

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

// WithDefaults hydrates default values in the get knative client plugins version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnativeClientPluginsVersionDetailsParams) WithDefaults() *GetKnativeClientPluginsVersionDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get knative client plugins version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnativeClientPluginsVersionDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) WithTimeout(timeout time.Duration) *GetKnativeClientPluginsVersionDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) WithContext(ctx context.Context) *GetKnativeClientPluginsVersionDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) WithHTTPClient(client *http.Client) *GetKnativeClientPluginsVersionDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) WithPackageName(packageName string) *GetKnativeClientPluginsVersionDetailsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoName adds the repoName to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) WithRepoName(repoName string) *GetKnativeClientPluginsVersionDetailsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithVersion adds the version to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) WithVersion(version string) *GetKnativeClientPluginsVersionDetailsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get knative client plugins version details params
func (o *GetKnativeClientPluginsVersionDetailsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnativeClientPluginsVersionDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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