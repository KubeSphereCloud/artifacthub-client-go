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

// NewGetFalcoRulesVersionDetailsParams creates a new GetFalcoRulesVersionDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFalcoRulesVersionDetailsParams() *GetFalcoRulesVersionDetailsParams {
	return &GetFalcoRulesVersionDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFalcoRulesVersionDetailsParamsWithTimeout creates a new GetFalcoRulesVersionDetailsParams object
// with the ability to set a timeout on a request.
func NewGetFalcoRulesVersionDetailsParamsWithTimeout(timeout time.Duration) *GetFalcoRulesVersionDetailsParams {
	return &GetFalcoRulesVersionDetailsParams{
		timeout: timeout,
	}
}

// NewGetFalcoRulesVersionDetailsParamsWithContext creates a new GetFalcoRulesVersionDetailsParams object
// with the ability to set a context for a request.
func NewGetFalcoRulesVersionDetailsParamsWithContext(ctx context.Context) *GetFalcoRulesVersionDetailsParams {
	return &GetFalcoRulesVersionDetailsParams{
		Context: ctx,
	}
}

// NewGetFalcoRulesVersionDetailsParamsWithHTTPClient creates a new GetFalcoRulesVersionDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFalcoRulesVersionDetailsParamsWithHTTPClient(client *http.Client) *GetFalcoRulesVersionDetailsParams {
	return &GetFalcoRulesVersionDetailsParams{
		HTTPClient: client,
	}
}

/*
GetFalcoRulesVersionDetailsParams contains all the parameters to send to the API endpoint

	for the get falco rules version details operation.

	Typically these are written to a http.Request.
*/
type GetFalcoRulesVersionDetailsParams struct {

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

// WithDefaults hydrates default values in the get falco rules version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFalcoRulesVersionDetailsParams) WithDefaults() *GetFalcoRulesVersionDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get falco rules version details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFalcoRulesVersionDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) WithTimeout(timeout time.Duration) *GetFalcoRulesVersionDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) WithContext(ctx context.Context) *GetFalcoRulesVersionDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) WithHTTPClient(client *http.Client) *GetFalcoRulesVersionDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) WithPackageName(packageName string) *GetFalcoRulesVersionDetailsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoName adds the repoName to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) WithRepoName(repoName string) *GetFalcoRulesVersionDetailsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WithVersion adds the version to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) WithVersion(version string) *GetFalcoRulesVersionDetailsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get falco rules version details params
func (o *GetFalcoRulesVersionDetailsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetFalcoRulesVersionDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
