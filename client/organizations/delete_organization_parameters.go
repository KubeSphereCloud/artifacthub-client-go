// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewDeleteOrganizationParams creates a new DeleteOrganizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationParams() *DeleteOrganizationParams {
	return &DeleteOrganizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationParamsWithTimeout creates a new DeleteOrganizationParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationParamsWithTimeout(timeout time.Duration) *DeleteOrganizationParams {
	return &DeleteOrganizationParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationParamsWithContext creates a new DeleteOrganizationParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationParamsWithContext(ctx context.Context) *DeleteOrganizationParams {
	return &DeleteOrganizationParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationParamsWithHTTPClient creates a new DeleteOrganizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationParamsWithHTTPClient(client *http.Client) *DeleteOrganizationParams {
	return &DeleteOrganizationParams{
		HTTPClient: client,
	}
}

/*
DeleteOrganizationParams contains all the parameters to send to the API endpoint

	for the delete organization operation.

	Typically these are written to a http.Request.
*/
type DeleteOrganizationParams struct {

	/* OrgName.

	   Organization name
	*/
	OrgName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationParams) WithDefaults() *DeleteOrganizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organization params
func (o *DeleteOrganizationParams) WithTimeout(timeout time.Duration) *DeleteOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization params
func (o *DeleteOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization params
func (o *DeleteOrganizationParams) WithContext(ctx context.Context) *DeleteOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization params
func (o *DeleteOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization params
func (o *DeleteOrganizationParams) WithHTTPClient(client *http.Client) *DeleteOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization params
func (o *DeleteOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgName adds the orgName to the delete organization params
func (o *DeleteOrganizationParams) WithOrgName(orgName string) *DeleteOrganizationParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the delete organization params
func (o *DeleteOrganizationParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
