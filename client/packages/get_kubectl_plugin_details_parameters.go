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

// NewGetKubectlPluginDetailsParams creates a new GetKubectlPluginDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKubectlPluginDetailsParams() *GetKubectlPluginDetailsParams {
	return &GetKubectlPluginDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubectlPluginDetailsParamsWithTimeout creates a new GetKubectlPluginDetailsParams object
// with the ability to set a timeout on a request.
func NewGetKubectlPluginDetailsParamsWithTimeout(timeout time.Duration) *GetKubectlPluginDetailsParams {
	return &GetKubectlPluginDetailsParams{
		timeout: timeout,
	}
}

// NewGetKubectlPluginDetailsParamsWithContext creates a new GetKubectlPluginDetailsParams object
// with the ability to set a context for a request.
func NewGetKubectlPluginDetailsParamsWithContext(ctx context.Context) *GetKubectlPluginDetailsParams {
	return &GetKubectlPluginDetailsParams{
		Context: ctx,
	}
}

// NewGetKubectlPluginDetailsParamsWithHTTPClient creates a new GetKubectlPluginDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKubectlPluginDetailsParamsWithHTTPClient(client *http.Client) *GetKubectlPluginDetailsParams {
	return &GetKubectlPluginDetailsParams{
		HTTPClient: client,
	}
}

/*
GetKubectlPluginDetailsParams contains all the parameters to send to the API endpoint

	for the get kubectl plugin details operation.

	Typically these are written to a http.Request.
*/
type GetKubectlPluginDetailsParams struct {

	/* PackageName.

	   Package name
	*/
	PackageName string

	/* RepoName.

	   Repository name
	*/
	RepoName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kubectl plugin details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubectlPluginDetailsParams) WithDefaults() *GetKubectlPluginDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kubectl plugin details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKubectlPluginDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) WithTimeout(timeout time.Duration) *GetKubectlPluginDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) WithContext(ctx context.Context) *GetKubectlPluginDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) WithHTTPClient(client *http.Client) *GetKubectlPluginDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) WithPackageName(packageName string) *GetKubectlPluginDetailsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoName adds the repoName to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) WithRepoName(repoName string) *GetKubectlPluginDetailsParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get kubectl plugin details params
func (o *GetKubectlPluginDetailsParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubectlPluginDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}