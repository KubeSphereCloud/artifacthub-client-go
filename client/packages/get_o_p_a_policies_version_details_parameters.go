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

// NewGetOPAPoliciesVersionDetailsParams creates a new GetOPAPoliciesVersionDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOPAPoliciesVersionDetailsParams() *GetOPAPoliciesVersionDetailsParams {
	return &GetOPAPoliciesVersionDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOPAPoliciesVersionDetailsParamsWithTimeout creates a new GetOPAPoliciesVersionDetailsParams object
// with the ability to set a timeout on a request.
func NewGetOPAPoliciesVersionDetailsParamsWithTimeout(timeout time.Duration) *GetOPAPoliciesVersionDetailsParams {
	return &GetOPAPoliciesVersionDetailsParams{
		timeout: timeout,
	}
}

// NewGetOPAPoliciesVersionDetailsParamsWithContext creates a new GetOPAPoliciesVersionDetailsParams object
// with the ability to set a context for a request.
func NewGetOPAPoliciesVersionDetailsParamsWithContext(ctx context.Context) *GetOPAPoliciesVersionDetailsParams {
	return &GetOPAPoliciesVersionDetailsParams{
		Context: ctx,
	}
}

// NewGetOPAPoliciesVersionDetailsParamsWithHTTPClient creates a new GetOPAPoliciesVersionDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOPAPoliciesVersionDetailsParamsWithHTTPClient(client *http.Client) *GetOPAPoliciesVersionDetailsParams {
	return &GetOPAPoliciesVersionDetailsParams{
		HTTPClient: client,
	}
}

/*
GetOPAPoliciesVersionDetailsParams contains all the parameters to send to the API endpoint

	for the get o p a policies version details operation.

	Typically these are written to a http.Request.
*/
type GetOPAPoliciesVersionDetailsParams struct {

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

// WithDefaults hydrates default values in the get o p a policies version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOPAPoliciesVersionDetailsParams) WithDefaults() *GetOPAPoliciesVersionDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get o p a policies version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOPAPoliciesVersionDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) WithTimeout(timeout time.Duration) *GetOPAPoliciesVersionDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) WithContext(ctx context.Context) *GetOPAPoliciesVersionDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) WithHTTPClient(client *http.Client) *GetOPAPoliciesVersionDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) WithPackageName(packageName string) *GetOPAPoliciesVersionDetailsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoName adds the repoName to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) WithRepoName(repoName string) *GetOPAPoliciesVersionDetailsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithVersion adds the version to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) WithVersion(version string) *GetOPAPoliciesVersionDetailsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get o p a policies version details params
func (o *GetOPAPoliciesVersionDetailsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetOPAPoliciesVersionDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
