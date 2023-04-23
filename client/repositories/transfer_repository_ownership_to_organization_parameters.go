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

// NewTransferRepositoryOwnershipToOrganizationParams creates a new TransferRepositoryOwnershipToOrganizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTransferRepositoryOwnershipToOrganizationParams() *TransferRepositoryOwnershipToOrganizationParams {
	return &TransferRepositoryOwnershipToOrganizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTransferRepositoryOwnershipToOrganizationParamsWithTimeout creates a new TransferRepositoryOwnershipToOrganizationParams object
// with the ability to set a timeout on a request.
func NewTransferRepositoryOwnershipToOrganizationParamsWithTimeout(timeout time.Duration) *TransferRepositoryOwnershipToOrganizationParams {
	return &TransferRepositoryOwnershipToOrganizationParams{
		timeout: timeout,
	}
}

// NewTransferRepositoryOwnershipToOrganizationParamsWithContext creates a new TransferRepositoryOwnershipToOrganizationParams object
// with the ability to set a context for a request.
func NewTransferRepositoryOwnershipToOrganizationParamsWithContext(ctx context.Context) *TransferRepositoryOwnershipToOrganizationParams {
	return &TransferRepositoryOwnershipToOrganizationParams{
		Context: ctx,
	}
}

// NewTransferRepositoryOwnershipToOrganizationParamsWithHTTPClient creates a new TransferRepositoryOwnershipToOrganizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewTransferRepositoryOwnershipToOrganizationParamsWithHTTPClient(client *http.Client) *TransferRepositoryOwnershipToOrganizationParams {
	return &TransferRepositoryOwnershipToOrganizationParams{
		HTTPClient: client,
	}
}

/*
TransferRepositoryOwnershipToOrganizationParams contains all the parameters to send to the API endpoint

	for the transfer repository ownership to organization operation.

	Typically these are written to a http.Request.
*/
type TransferRepositoryOwnershipToOrganizationParams struct {

	/* Org.

	   The org to transfer or from claiming the repository
	*/
	Org *string

	/* RepoName.

	   Repository name
	*/
	RepoName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the transfer repository ownership to organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferRepositoryOwnershipToOrganizationParams) WithDefaults() *TransferRepositoryOwnershipToOrganizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the transfer repository ownership to organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransferRepositoryOwnershipToOrganizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) WithTimeout(timeout time.Duration) *TransferRepositoryOwnershipToOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) WithContext(ctx context.Context) *TransferRepositoryOwnershipToOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) WithHTTPClient(client *http.Client) *TransferRepositoryOwnershipToOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) WithOrg(org *string) *TransferRepositoryOwnershipToOrganizationParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) SetOrg(org *string) {
	o.Org = org
}

// WithRepoName adds the repoName to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) WithRepoName(repoName string) *TransferRepositoryOwnershipToOrganizationParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the transfer repository ownership to organization params
func (o *TransferRepositoryOwnershipToOrganizationParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *TransferRepositoryOwnershipToOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param repoName
	if err := r.SetPathParam("repoName", o.RepoName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}