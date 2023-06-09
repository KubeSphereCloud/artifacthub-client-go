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

// NewAddProductionUsageParams creates a new AddProductionUsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddProductionUsageParams() *AddProductionUsageParams {
	return &AddProductionUsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddProductionUsageParamsWithTimeout creates a new AddProductionUsageParams object
// with the ability to set a timeout on a request.
func NewAddProductionUsageParamsWithTimeout(timeout time.Duration) *AddProductionUsageParams {
	return &AddProductionUsageParams{
		timeout: timeout,
	}
}

// NewAddProductionUsageParamsWithContext creates a new AddProductionUsageParams object
// with the ability to set a context for a request.
func NewAddProductionUsageParamsWithContext(ctx context.Context) *AddProductionUsageParams {
	return &AddProductionUsageParams{
		Context: ctx,
	}
}

// NewAddProductionUsageParamsWithHTTPClient creates a new AddProductionUsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddProductionUsageParamsWithHTTPClient(client *http.Client) *AddProductionUsageParams {
	return &AddProductionUsageParams{
		HTTPClient: client,
	}
}

/*
AddProductionUsageParams contains all the parameters to send to the API endpoint

	for the add production usage operation.

	Typically these are written to a http.Request.
*/
type AddProductionUsageParams struct {

	/* OrgName.

	   Organization name
	*/
	OrgName string

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

// WithDefaults hydrates default values in the add production usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddProductionUsageParams) WithDefaults() *AddProductionUsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add production usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddProductionUsageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add production usage params
func (o *AddProductionUsageParams) WithTimeout(timeout time.Duration) *AddProductionUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add production usage params
func (o *AddProductionUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add production usage params
func (o *AddProductionUsageParams) WithContext(ctx context.Context) *AddProductionUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add production usage params
func (o *AddProductionUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add production usage params
func (o *AddProductionUsageParams) WithHTTPClient(client *http.Client) *AddProductionUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add production usage params
func (o *AddProductionUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgName adds the orgName to the add production usage params
func (o *AddProductionUsageParams) WithOrgName(orgName string) *AddProductionUsageParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the add production usage params
func (o *AddProductionUsageParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithPackageName adds the packageName to the add production usage params
func (o *AddProductionUsageParams) WithPackageName(packageName string) *AddProductionUsageParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the add production usage params
func (o *AddProductionUsageParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoKindParam adds the repoKindParam to the add production usage params
func (o *AddProductionUsageParams) WithRepoKindParam(repoKindParam string) *AddProductionUsageParams {
	o.SetRepoKindParam(repoKindParam)
	return o
}

// SetRepoKindParam adds the repoKindParam to the add production usage params
func (o *AddProductionUsageParams) SetRepoKindParam(repoKindParam string) {
	o.RepoKindParam = repoKindParam
}

// WithRepoName adds the repoName to the add production usage params
func (o *AddProductionUsageParams) WithRepoName(repoName string) *AddProductionUsageParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the add production usage params
func (o *AddProductionUsageParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *AddProductionUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

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
