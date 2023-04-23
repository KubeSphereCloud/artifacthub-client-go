// Code generated by go-swagger; DO NOT EDIT.

package repositories

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

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// NewUpdateOrganizationRepositoryParams creates a new UpdateOrganizationRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationRepositoryParams() *UpdateOrganizationRepositoryParams {
	return &UpdateOrganizationRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationRepositoryParamsWithTimeout creates a new UpdateOrganizationRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationRepositoryParamsWithTimeout(timeout time.Duration) *UpdateOrganizationRepositoryParams {
	return &UpdateOrganizationRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationRepositoryParamsWithContext creates a new UpdateOrganizationRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationRepositoryParamsWithContext(ctx context.Context) *UpdateOrganizationRepositoryParams {
	return &UpdateOrganizationRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationRepositoryParamsWithHTTPClient creates a new UpdateOrganizationRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationRepositoryParamsWithHTTPClient(client *http.Client) *UpdateOrganizationRepositoryParams {
	return &UpdateOrganizationRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationRepositoryParams contains all the parameters to send to the API endpoint

	for the update organization repository operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationRepositoryParams struct {

	// ContentType.
	ContentType *string

	/* Body.

	   Repository request body
	*/
	Body *models.RepositoryBody

	/* OrgName.

	   Organization name
	*/
	OrgName string

	/* RepoName.

	   Repository name
	*/
	RepoName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationRepositoryParams) WithDefaults() *UpdateOrganizationRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) WithTimeout(timeout time.Duration) *UpdateOrganizationRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) WithContext(ctx context.Context) *UpdateOrganizationRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) WithHTTPClient(client *http.Client) *UpdateOrganizationRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) WithContentType(contentType *string) *UpdateOrganizationRepositoryParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) WithBody(body *models.RepositoryBody) *UpdateOrganizationRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) SetBody(body *models.RepositoryBody) {
	o.Body = body
}

// WithOrgName adds the orgName to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) WithOrgName(orgName string) *UpdateOrganizationRepositoryParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithRepoName adds the repoName to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) WithRepoName(repoName string) *UpdateOrganizationRepositoryParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the update organization repository params
func (o *UpdateOrganizationRepositoryParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// header param Content-Type
		if err := r.SetHeaderParam("Content-Type", *o.ContentType); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
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
