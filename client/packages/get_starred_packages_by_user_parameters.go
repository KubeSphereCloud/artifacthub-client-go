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
	"github.com/go-openapi/swag"
)

// NewGetStarredPackagesByUserParams creates a new GetStarredPackagesByUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStarredPackagesByUserParams() *GetStarredPackagesByUserParams {
	return &GetStarredPackagesByUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStarredPackagesByUserParamsWithTimeout creates a new GetStarredPackagesByUserParams object
// with the ability to set a timeout on a request.
func NewGetStarredPackagesByUserParamsWithTimeout(timeout time.Duration) *GetStarredPackagesByUserParams {
	return &GetStarredPackagesByUserParams{
		timeout: timeout,
	}
}

// NewGetStarredPackagesByUserParamsWithContext creates a new GetStarredPackagesByUserParams object
// with the ability to set a context for a request.
func NewGetStarredPackagesByUserParamsWithContext(ctx context.Context) *GetStarredPackagesByUserParams {
	return &GetStarredPackagesByUserParams{
		Context: ctx,
	}
}

// NewGetStarredPackagesByUserParamsWithHTTPClient creates a new GetStarredPackagesByUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStarredPackagesByUserParamsWithHTTPClient(client *http.Client) *GetStarredPackagesByUserParams {
	return &GetStarredPackagesByUserParams{
		HTTPClient: client,
	}
}

/*
GetStarredPackagesByUserParams contains all the parameters to send to the API endpoint

	for the get starred packages by user operation.

	Typically these are written to a http.Request.
*/
type GetStarredPackagesByUserParams struct {

	/* Limit.

	   The number of items to return

	   Format: int32
	   Default: 20
	*/
	Limit *int32

	/* Offset.

	   The number of items to skip before starting to collect the result set

	   Format: int32
	*/
	Offset *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get starred packages by user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStarredPackagesByUserParams) WithDefaults() *GetStarredPackagesByUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get starred packages by user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStarredPackagesByUserParams) SetDefaults() {
	var (
		limitDefault = int32(20)

		offsetDefault = int32(0)
	)

	val := GetStarredPackagesByUserParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) WithTimeout(timeout time.Duration) *GetStarredPackagesByUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) WithContext(ctx context.Context) *GetStarredPackagesByUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) WithHTTPClient(client *http.Client) *GetStarredPackagesByUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) WithLimit(limit *int32) *GetStarredPackagesByUserParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) WithOffset(offset *int32) *GetStarredPackagesByUserParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get starred packages by user params
func (o *GetStarredPackagesByUserParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetStarredPackagesByUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
