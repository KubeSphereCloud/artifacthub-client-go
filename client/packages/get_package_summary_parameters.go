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

// NewGetPackageSummaryParams creates a new GetPackageSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPackageSummaryParams() *GetPackageSummaryParams {
	return &GetPackageSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackageSummaryParamsWithTimeout creates a new GetPackageSummaryParams object
// with the ability to set a timeout on a request.
func NewGetPackageSummaryParamsWithTimeout(timeout time.Duration) *GetPackageSummaryParams {
	return &GetPackageSummaryParams{
		timeout: timeout,
	}
}

// NewGetPackageSummaryParamsWithContext creates a new GetPackageSummaryParams object
// with the ability to set a context for a request.
func NewGetPackageSummaryParamsWithContext(ctx context.Context) *GetPackageSummaryParams {
	return &GetPackageSummaryParams{
		Context: ctx,
	}
}

// NewGetPackageSummaryParamsWithHTTPClient creates a new GetPackageSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPackageSummaryParamsWithHTTPClient(client *http.Client) *GetPackageSummaryParams {
	return &GetPackageSummaryParams{
		HTTPClient: client,
	}
}

/*
GetPackageSummaryParams contains all the parameters to send to the API endpoint

	for the get package summary operation.

	Typically these are written to a http.Request.
*/
type GetPackageSummaryParams struct {

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

// WithDefaults hydrates default values in the get package summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackageSummaryParams) WithDefaults() *GetPackageSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get package summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackageSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get package summary params
func (o *GetPackageSummaryParams) WithTimeout(timeout time.Duration) *GetPackageSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get package summary params
func (o *GetPackageSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get package summary params
func (o *GetPackageSummaryParams) WithContext(ctx context.Context) *GetPackageSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get package summary params
func (o *GetPackageSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get package summary params
func (o *GetPackageSummaryParams) WithHTTPClient(client *http.Client) *GetPackageSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get package summary params
func (o *GetPackageSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get package summary params
func (o *GetPackageSummaryParams) WithPackageName(packageName string) *GetPackageSummaryParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get package summary params
func (o *GetPackageSummaryParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoKindParam adds the repoKindParam to the get package summary params
func (o *GetPackageSummaryParams) WithRepoKindParam(repoKindParam string) *GetPackageSummaryParams {
	o.SetRepoKindParam(repoKindParam)
	return o
}

// SetRepoKindParam adds the repoKindParam to the get package summary params
func (o *GetPackageSummaryParams) SetRepoKindParam(repoKindParam string) {
	o.RepoKindParam = repoKindParam
}

// WithRepoName adds the repoName to the get package summary params
func (o *GetPackageSummaryParams) WithRepoName(repoName string) *GetPackageSummaryParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the get package summary params
func (o *GetPackageSummaryParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackageSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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