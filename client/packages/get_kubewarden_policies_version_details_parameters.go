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

// NewGetKubewardenPoliciesVersionDetailsParams creates a new GetKubewardenPoliciesVersionDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubewardenPoliciesVersionDetailsParams() *GetKubewardenPoliciesVersionDetailsParams {
	return &GetKubewardenPoliciesVersionDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubewardenPoliciesVersionDetailsParamsWithTimeout creates a new GetKubewardenPoliciesVersionDetailsParams object
// with the ability to set a timeout on a request.
func NewGetKubewardenPoliciesVersionDetailsParamsWithTimeout(timeout time.Duration) *GetKubewardenPoliciesVersionDetailsParams {
	return &GetKubewardenPoliciesVersionDetailsParams{
		timeout: timeout,
	}
}

// NewGetKubewardenPoliciesVersionDetailsParamsWithContext creates a new GetKubewardenPoliciesVersionDetailsParams object
// with the ability to set a context for a request.
func NewGetKubewardenPoliciesVersionDetailsParamsWithContext(ctx context.Context) *GetKubewardenPoliciesVersionDetailsParams {
	return &GetKubewardenPoliciesVersionDetailsParams{
		Context: ctx,
	}
}

// NewGetKubewardenPoliciesVersionDetailsParamsWithHTTPClient creates a new GetKubewardenPoliciesVersionDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubewardenPoliciesVersionDetailsParamsWithHTTPClient(client *http.Client) *GetKubewardenPoliciesVersionDetailsParams {
	return &GetKubewardenPoliciesVersionDetailsParams{
		HTTPClient: client,
	}
}

/*
GetKubewardenPoliciesVersionDetailsParams contains all the parameters to send to the API endpoint

	for the get kubewarden policies version details operation.

	Typically these are written to a http.Request.
*/
type GetKubewardenPoliciesVersionDetailsParams struct {

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

// WithDefaults hydrates default values in the get kubewarden policies version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubewardenPoliciesVersionDetailsParams) WithDefaults() *GetKubewardenPoliciesVersionDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubewarden policies version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubewardenPoliciesVersionDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) WithTimeout(timeout time.Duration) *GetKubewardenPoliciesVersionDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) WithContext(ctx context.Context) *GetKubewardenPoliciesVersionDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) WithHTTPClient(client *http.Client) *GetKubewardenPoliciesVersionDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) WithPackageName(packageName string) *GetKubewardenPoliciesVersionDetailsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoName adds the repoName to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) WithRepoName(repoName string) *GetKubewardenPoliciesVersionDetailsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithVersion adds the version to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) WithVersion(version string) *GetKubewardenPoliciesVersionDetailsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get kubewarden policies version details params
func (o *GetKubewardenPoliciesVersionDetailsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubewardenPoliciesVersionDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
