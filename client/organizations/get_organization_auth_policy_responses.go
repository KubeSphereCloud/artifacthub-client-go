// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// GetOrganizationAuthPolicyReader is a Reader for the GetOrganizationAuthPolicy structure.
type GetOrganizationAuthPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationAuthPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationAuthPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrganizationAuthPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationAuthPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationAuthPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationAuthPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationAuthPolicyOK creates a GetOrganizationAuthPolicyOK with default headers values
func NewGetOrganizationAuthPolicyOK() *GetOrganizationAuthPolicyOK {
	return &GetOrganizationAuthPolicyOK{}
}

/*
GetOrganizationAuthPolicyOK describes a response with status code 200, with default header values.

GetOrganizationAuthPolicyOK get organization auth policy o k
*/
type GetOrganizationAuthPolicyOK struct {
	Payload *models.AuthorizationPolicy
}

// IsSuccess returns true when this get organization auth policy o k response has a 2xx status code
func (o *GetOrganizationAuthPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization auth policy o k response has a 3xx status code
func (o *GetOrganizationAuthPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization auth policy o k response has a 4xx status code
func (o *GetOrganizationAuthPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization auth policy o k response has a 5xx status code
func (o *GetOrganizationAuthPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization auth policy o k response a status code equal to that given
func (o *GetOrganizationAuthPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationAuthPolicyOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAuthPolicyOK) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAuthPolicyOK) GetPayload() *models.AuthorizationPolicy {
	return o.Payload
}

func (o *GetOrganizationAuthPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationAuthPolicyUnauthorized creates a GetOrganizationAuthPolicyUnauthorized with default headers values
func NewGetOrganizationAuthPolicyUnauthorized() *GetOrganizationAuthPolicyUnauthorized {
	return &GetOrganizationAuthPolicyUnauthorized{}
}

/*
GetOrganizationAuthPolicyUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type GetOrganizationAuthPolicyUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization auth policy unauthorized response has a 2xx status code
func (o *GetOrganizationAuthPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization auth policy unauthorized response has a 3xx status code
func (o *GetOrganizationAuthPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization auth policy unauthorized response has a 4xx status code
func (o *GetOrganizationAuthPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization auth policy unauthorized response has a 5xx status code
func (o *GetOrganizationAuthPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization auth policy unauthorized response a status code equal to that given
func (o *GetOrganizationAuthPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOrganizationAuthPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationAuthPolicyUnauthorized) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationAuthPolicyUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationAuthPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationAuthPolicyForbidden creates a GetOrganizationAuthPolicyForbidden with default headers values
func NewGetOrganizationAuthPolicyForbidden() *GetOrganizationAuthPolicyForbidden {
	return &GetOrganizationAuthPolicyForbidden{}
}

/*
GetOrganizationAuthPolicyForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the requested operation
*/
type GetOrganizationAuthPolicyForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization auth policy forbidden response has a 2xx status code
func (o *GetOrganizationAuthPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization auth policy forbidden response has a 3xx status code
func (o *GetOrganizationAuthPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization auth policy forbidden response has a 4xx status code
func (o *GetOrganizationAuthPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization auth policy forbidden response has a 5xx status code
func (o *GetOrganizationAuthPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization auth policy forbidden response a status code equal to that given
func (o *GetOrganizationAuthPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrganizationAuthPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationAuthPolicyForbidden) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationAuthPolicyForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationAuthPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationAuthPolicyTooManyRequests creates a GetOrganizationAuthPolicyTooManyRequests with default headers values
func NewGetOrganizationAuthPolicyTooManyRequests() *GetOrganizationAuthPolicyTooManyRequests {
	return &GetOrganizationAuthPolicyTooManyRequests{}
}

/*
GetOrganizationAuthPolicyTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetOrganizationAuthPolicyTooManyRequests struct {
}

// IsSuccess returns true when this get organization auth policy too many requests response has a 2xx status code
func (o *GetOrganizationAuthPolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization auth policy too many requests response has a 3xx status code
func (o *GetOrganizationAuthPolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization auth policy too many requests response has a 4xx status code
func (o *GetOrganizationAuthPolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization auth policy too many requests response has a 5xx status code
func (o *GetOrganizationAuthPolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization auth policy too many requests response a status code equal to that given
func (o *GetOrganizationAuthPolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrganizationAuthPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyTooManyRequests ", 429)
}

func (o *GetOrganizationAuthPolicyTooManyRequests) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyTooManyRequests ", 429)
}

func (o *GetOrganizationAuthPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationAuthPolicyInternalServerError creates a GetOrganizationAuthPolicyInternalServerError with default headers values
func NewGetOrganizationAuthPolicyInternalServerError() *GetOrganizationAuthPolicyInternalServerError {
	return &GetOrganizationAuthPolicyInternalServerError{}
}

/*
GetOrganizationAuthPolicyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetOrganizationAuthPolicyInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization auth policy internal server error response has a 2xx status code
func (o *GetOrganizationAuthPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization auth policy internal server error response has a 3xx status code
func (o *GetOrganizationAuthPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization auth policy internal server error response has a 4xx status code
func (o *GetOrganizationAuthPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization auth policy internal server error response has a 5xx status code
func (o *GetOrganizationAuthPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get organization auth policy internal server error response a status code equal to that given
func (o *GetOrganizationAuthPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrganizationAuthPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationAuthPolicyInternalServerError) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/authorization-policy][%d] getOrganizationAuthPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationAuthPolicyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationAuthPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
