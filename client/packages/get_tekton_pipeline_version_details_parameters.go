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

// NewGetTektonPipelineVersionDetailsParams creates a new GetTektonPipelineVersionDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTektonPipelineVersionDetailsParams() *GetTektonPipelineVersionDetailsParams {
	return &GetTektonPipelineVersionDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTektonPipelineVersionDetailsParamsWithTimeout creates a new GetTektonPipelineVersionDetailsParams object
// with the ability to set a timeout on a request.
func NewGetTektonPipelineVersionDetailsParamsWithTimeout(timeout time.Duration) *GetTektonPipelineVersionDetailsParams {
	return &GetTektonPipelineVersionDetailsParams{
		timeout: timeout,
	}
}

// NewGetTektonPipelineVersionDetailsParamsWithContext creates a new GetTektonPipelineVersionDetailsParams object
// with the ability to set a context for a request.
func NewGetTektonPipelineVersionDetailsParamsWithContext(ctx context.Context) *GetTektonPipelineVersionDetailsParams {
	return &GetTektonPipelineVersionDetailsParams{
		Context: ctx,
	}
}

// NewGetTektonPipelineVersionDetailsParamsWithHTTPClient creates a new GetTektonPipelineVersionDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTektonPipelineVersionDetailsParamsWithHTTPClient(client *http.Client) *GetTektonPipelineVersionDetailsParams {
	return &GetTektonPipelineVersionDetailsParams{
		HTTPClient: client,
	}
}

/*
GetTektonPipelineVersionDetailsParams contains all the parameters to send to the API endpoint

	for the get tekton pipeline version details operation.

	Typically these are written to a http.Request.
*/
type GetTektonPipelineVersionDetailsParams struct {

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

// WithDefaults hydrates default values in the get tekton pipeline version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTektonPipelineVersionDetailsParams) WithDefaults() *GetTektonPipelineVersionDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tekton pipeline version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTektonPipelineVersionDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) WithTimeout(timeout time.Duration) *GetTektonPipelineVersionDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) WithContext(ctx context.Context) *GetTektonPipelineVersionDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) WithHTTPClient(client *http.Client) *GetTektonPipelineVersionDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) WithPackageName(packageName string) *GetTektonPipelineVersionDetailsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoName adds the repoName to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) WithRepoName(repoName string) *GetTektonPipelineVersionDetailsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithVersion adds the version to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) WithVersion(version string) *GetTektonPipelineVersionDetailsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get tekton pipeline version details params
func (o *GetTektonPipelineVersionDetailsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetTektonPipelineVersionDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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