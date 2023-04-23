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

// NewGetProductionUsageParams creates a new GetProductionUsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProductionUsageParams() *GetProductionUsageParams {
	return &GetProductionUsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductionUsageParamsWithTimeout creates a new GetProductionUsageParams object
// with the ability to set a timeout on a request.
func NewGetProductionUsageParamsWithTimeout(timeout time.Duration) *GetProductionUsageParams {
	return &GetProductionUsageParams{
		timeout: timeout,
	}
}

// NewGetProductionUsageParamsWithContext creates a new GetProductionUsageParams object
// with the ability to set a context for a request.
func NewGetProductionUsageParamsWithContext(ctx context.Context) *GetProductionUsageParams {
	return &GetProductionUsageParams{
		Context: ctx,
	}
}

// NewGetProductionUsageParamsWithHTTPClient creates a new GetProductionUsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProductionUsageParamsWithHTTPClient(client *http.Client) *GetProductionUsageParams {
	return &GetProductionUsageParams{
		HTTPClient: client,
	}
}

/*
GetProductionUsageParams contains all the parameters to send to the API endpoint

	for the get production usage operation.

	Typically these are written to a http.Request.
*/
type GetProductionUsageParams struct {

	/* PackageName.

	   Package name
	*/
	PackageName string

	/* RepoKindParam.

	   Package kind name
	*/
	RepoKindParam string

	/* RepoName.

	   Repository name
	*/
	RepoName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get production usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductionUsageParams) WithDefaults() *GetProductionUsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get production usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductionUsageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get production usage params
func (o *GetProductionUsageParams) WithTimeout(timeout time.Duration) *GetProductionUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get production usage params
func (o *GetProductionUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get production usage params
func (o *GetProductionUsageParams) WithContext(ctx context.Context) *GetProductionUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get production usage params
func (o *GetProductionUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get production usage params
func (o *GetProductionUsageParams) WithHTTPClient(client *http.Client) *GetProductionUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get production usage params
func (o *GetProductionUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get production usage params
func (o *GetProductionUsageParams) WithPackageName(packageName string) *GetProductionUsageParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get production usage params
func (o *GetProductionUsageParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoKindParam adds the repoKindParam to the get production usage params
func (o *GetProductionUsageParams) WithRepoKindParam(repoKindParam string) *GetProductionUsageParams {
	o.SetRepoKindParam(repoKindParam)
	return o
}

// SetRepoKindParam adds the repoKindParam to the get production usage params
func (o *GetProductionUsageParams) SetRepoKindParam(repoKindParam string) {
	o.RepoKindParam = repoKindParam
}

// WithRepoName adds the repoName to the get production usage params
func (o *GetProductionUsageParams) WithRepoName(repoName string) *GetProductionUsageParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get production usage params
func (o *GetProductionUsageParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductionUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
	}

	// path param repoKindParam
	if err := r.SetPathParam("repoKindParam", o.RepoKindParam); err != nil {
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
