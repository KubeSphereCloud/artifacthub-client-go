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
)

// NewClaimRepositoryOwnershipFromOrganizationParams creates a new ClaimRepositoryOwnershipFromOrganizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClaimRepositoryOwnershipFromOrganizationParams() *ClaimRepositoryOwnershipFromOrganizationParams {
	return &ClaimRepositoryOwnershipFromOrganizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClaimRepositoryOwnershipFromOrganizationParamsWithTimeout creates a new ClaimRepositoryOwnershipFromOrganizationParams object
// with the ability to set a timeout on a request.
func NewClaimRepositoryOwnershipFromOrganizationParamsWithTimeout(timeout time.Duration) *ClaimRepositoryOwnershipFromOrganizationParams {
	return &ClaimRepositoryOwnershipFromOrganizationParams{
		timeout: timeout,
	}
}

// NewClaimRepositoryOwnershipFromOrganizationParamsWithContext creates a new ClaimRepositoryOwnershipFromOrganizationParams object
// with the ability to set a context for a request.
func NewClaimRepositoryOwnershipFromOrganizationParamsWithContext(ctx context.Context) *ClaimRepositoryOwnershipFromOrganizationParams {
	return &ClaimRepositoryOwnershipFromOrganizationParams{
		Context: ctx,
	}
}

// NewClaimRepositoryOwnershipFromOrganizationParamsWithHTTPClient creates a new ClaimRepositoryOwnershipFromOrganizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewClaimRepositoryOwnershipFromOrganizationParamsWithHTTPClient(client *http.Client) *ClaimRepositoryOwnershipFromOrganizationParams {
	return &ClaimRepositoryOwnershipFromOrganizationParams{
		HTTPClient: client,
	}
}

/*
ClaimRepositoryOwnershipFromOrganizationParams contains all the parameters to send to the API endpoint

	for the claim repository ownership from organization operation.

	Typically these are written to a http.Request.
*/
type ClaimRepositoryOwnershipFromOrganizationParams struct {

	/* Org.

	   The org to transfer or from claiming the repository
	*/
	Org *string

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

// WithDefaults hydrates default values in the claim repository ownership from organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClaimRepositoryOwnershipFromOrganizationParams) WithDefaults() *ClaimRepositoryOwnershipFromOrganizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the claim repository ownership from organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClaimRepositoryOwnershipFromOrganizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) WithTimeout(timeout time.Duration) *ClaimRepositoryOwnershipFromOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) WithContext(ctx context.Context) *ClaimRepositoryOwnershipFromOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) WithHTTPClient(client *http.Client) *ClaimRepositoryOwnershipFromOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) WithOrg(org *string) *ClaimRepositoryOwnershipFromOrganizationParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) SetOrg(org *string) {
	o.Org = org
}

// WithOrgName adds the orgName to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) WithOrgName(orgName string) *ClaimRepositoryOwnershipFromOrganizationParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithRepoName adds the repoName to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) WithRepoName(repoName string) *ClaimRepositoryOwnershipFromOrganizationParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the claim repository ownership from organization params
func (o *ClaimRepositoryOwnershipFromOrganizationParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *ClaimRepositoryOwnershipFromOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Org != nil {

		// query param org
		var qrOrg string

		if o.Org != nil {
			qrOrg = *o.Org
		}
		qOrg := qrOrg
		if qOrg != "" {

			if err := r.SetQueryParam("org", qOrg); err != nil {
				return err
			}
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
