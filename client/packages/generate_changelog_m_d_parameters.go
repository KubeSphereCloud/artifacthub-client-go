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

// NewGenerateChangelogMDParams creates a new GenerateChangelogMDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateChangelogMDParams() *GenerateChangelogMDParams {
	return &GenerateChangelogMDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateChangelogMDParamsWithTimeout creates a new GenerateChangelogMDParams object
// with the ability to set a timeout on a request.
func NewGenerateChangelogMDParamsWithTimeout(timeout time.Duration) *GenerateChangelogMDParams {
	return &GenerateChangelogMDParams{
		timeout: timeout,
	}
}

// NewGenerateChangelogMDParamsWithContext creates a new GenerateChangelogMDParams object
// with the ability to set a context for a request.
func NewGenerateChangelogMDParamsWithContext(ctx context.Context) *GenerateChangelogMDParams {
	return &GenerateChangelogMDParams{
		Context: ctx,
	}
}

// NewGenerateChangelogMDParamsWithHTTPClient creates a new GenerateChangelogMDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateChangelogMDParamsWithHTTPClient(client *http.Client) *GenerateChangelogMDParams {
	return &GenerateChangelogMDParams{
		HTTPClient: client,
	}
}

/*
GenerateChangelogMDParams contains all the parameters to send to the API endpoint

	for the generate changelog m d operation.

	Typically these are written to a http.Request.
*/
type GenerateChangelogMDParams struct {

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

// WithDefaults hydrates default values in the generate changelog m d params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateChangelogMDParams) WithDefaults() *GenerateChangelogMDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate changelog m d params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateChangelogMDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate changelog m d params
func (o *GenerateChangelogMDParams) WithTimeout(timeout time.Duration) *GenerateChangelogMDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate changelog m d params
func (o *GenerateChangelogMDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate changelog m d params
func (o *GenerateChangelogMDParams) WithContext(ctx context.Context) *GenerateChangelogMDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate changelog m d params
func (o *GenerateChangelogMDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate changelog m d params
func (o *GenerateChangelogMDParams) WithHTTPClient(client *http.Client) *GenerateChangelogMDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate changelog m d params
func (o *GenerateChangelogMDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the generate changelog m d params
func (o *GenerateChangelogMDParams) WithPackageName(packageName string) *GenerateChangelogMDParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the generate changelog m d params
func (o *GenerateChangelogMDParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithRepoKindParam adds the repoKindParam to the generate changelog m d params
func (o *GenerateChangelogMDParams) WithRepoKindParam(repoKindParam string) *GenerateChangelogMDParams {
	o.SetRepoKindParam(repoKindParam)
	return o
}

// SetRepoKindParam adds the repoKindParam to the generate changelog m d params
func (o *GenerateChangelogMDParams) SetRepoKindParam(repoKindParam string) {
	o.RepoKindParam = repoKindParam
}

// WithRepoName adds the repoName to the generate changelog m d params
func (o *GenerateChangelogMDParams) WithRepoName(repoName string) *GenerateChangelogMDParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the generate changelog m d params
func (o *GenerateChangelogMDParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateChangelogMDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
