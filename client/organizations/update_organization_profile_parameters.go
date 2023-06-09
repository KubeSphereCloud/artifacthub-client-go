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

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// NewUpdateOrganizationProfileParams creates a new UpdateOrganizationProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationProfileParams() *UpdateOrganizationProfileParams {
	return &UpdateOrganizationProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationProfileParamsWithTimeout creates a new UpdateOrganizationProfileParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationProfileParamsWithTimeout(timeout time.Duration) *UpdateOrganizationProfileParams {
	return &UpdateOrganizationProfileParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationProfileParamsWithContext creates a new UpdateOrganizationProfileParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationProfileParamsWithContext(ctx context.Context) *UpdateOrganizationProfileParams {
	return &UpdateOrganizationProfileParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationProfileParamsWithHTTPClient creates a new UpdateOrganizationProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationProfileParamsWithHTTPClient(client *http.Client) *UpdateOrganizationProfileParams {
	return &UpdateOrganizationProfileParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationProfileParams contains all the parameters to send to the API endpoint

	for the update organization profile operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationProfileParams struct {

	// ContentType.
	ContentType *string

	// Body.
	Body *models.OrganizationSummary

	/* OrgName.

	   Organization name
	*/
	OrgName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationProfileParams) WithDefaults() *UpdateOrganizationProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization profile params
func (o *UpdateOrganizationProfileParams) WithTimeout(timeout time.Duration) *UpdateOrganizationProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization profile params
func (o *UpdateOrganizationProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization profile params
func (o *UpdateOrganizationProfileParams) WithContext(ctx context.Context) *UpdateOrganizationProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization profile params
func (o *UpdateOrganizationProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization profile params
func (o *UpdateOrganizationProfileParams) WithHTTPClient(client *http.Client) *UpdateOrganizationProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization profile params
func (o *UpdateOrganizationProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the update organization profile params
func (o *UpdateOrganizationProfileParams) WithContentType(contentType *string) *UpdateOrganizationProfileParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the update organization profile params
func (o *UpdateOrganizationProfileParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithBody adds the body to the update organization profile params
func (o *UpdateOrganizationProfileParams) WithBody(body *models.OrganizationSummary) *UpdateOrganizationProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update organization profile params
func (o *UpdateOrganizationProfileParams) SetBody(body *models.OrganizationSummary) {
	o.Body = body
}

// WithOrgName adds the orgName to the update organization profile params
func (o *UpdateOrganizationProfileParams) WithOrgName(orgName string) *UpdateOrganizationProfileParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the update organization profile params
func (o *UpdateOrganizationProfileParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
