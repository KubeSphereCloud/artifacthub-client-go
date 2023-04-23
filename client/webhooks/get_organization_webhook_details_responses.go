// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// GetOrganizationWebhookDetailsReader is a Reader for the GetOrganizationWebhookDetails structure.
type GetOrganizationWebhookDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationWebhookDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationWebhookDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrganizationWebhookDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationWebhookDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationWebhookDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationWebhookDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationWebhookDetailsOK creates a GetOrganizationWebhookDetailsOK with default headers values
func NewGetOrganizationWebhookDetailsOK() *GetOrganizationWebhookDetailsOK {
	return &GetOrganizationWebhookDetailsOK{}
}

/*
GetOrganizationWebhookDetailsOK describes a response with status code 200, with default header values.

GetOrganizationWebhookDetailsOK get organization webhook details o k
*/
type GetOrganizationWebhookDetailsOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this get organization webhook details o k response has a 2xx status code
func (o *GetOrganizationWebhookDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization webhook details o k response has a 3xx status code
func (o *GetOrganizationWebhookDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization webhook details o k response has a 4xx status code
func (o *GetOrganizationWebhookDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization webhook details o k response has a 5xx status code
func (o *GetOrganizationWebhookDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization webhook details o k response a status code equal to that given
func (o *GetOrganizationWebhookDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationWebhookDetailsOK) Error() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationWebhookDetailsOK) String() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationWebhookDetailsOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *GetOrganizationWebhookDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationWebhookDetailsUnauthorized creates a GetOrganizationWebhookDetailsUnauthorized with default headers values
func NewGetOrganizationWebhookDetailsUnauthorized() *GetOrganizationWebhookDetailsUnauthorized {
	return &GetOrganizationWebhookDetailsUnauthorized{}
}

/*
GetOrganizationWebhookDetailsUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type GetOrganizationWebhookDetailsUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization webhook details unauthorized response has a 2xx status code
func (o *GetOrganizationWebhookDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization webhook details unauthorized response has a 3xx status code
func (o *GetOrganizationWebhookDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization webhook details unauthorized response has a 4xx status code
func (o *GetOrganizationWebhookDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization webhook details unauthorized response has a 5xx status code
func (o *GetOrganizationWebhookDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization webhook details unauthorized response a status code equal to that given
func (o *GetOrganizationWebhookDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOrganizationWebhookDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationWebhookDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationWebhookDetailsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationWebhookDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationWebhookDetailsForbidden creates a GetOrganizationWebhookDetailsForbidden with default headers values
func NewGetOrganizationWebhookDetailsForbidden() *GetOrganizationWebhookDetailsForbidden {
	return &GetOrganizationWebhookDetailsForbidden{}
}

/*
GetOrganizationWebhookDetailsForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the requested operation
*/
type GetOrganizationWebhookDetailsForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization webhook details forbidden response has a 2xx status code
func (o *GetOrganizationWebhookDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization webhook details forbidden response has a 3xx status code
func (o *GetOrganizationWebhookDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization webhook details forbidden response has a 4xx status code
func (o *GetOrganizationWebhookDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization webhook details forbidden response has a 5xx status code
func (o *GetOrganizationWebhookDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization webhook details forbidden response a status code equal to that given
func (o *GetOrganizationWebhookDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrganizationWebhookDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationWebhookDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationWebhookDetailsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationWebhookDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationWebhookDetailsTooManyRequests creates a GetOrganizationWebhookDetailsTooManyRequests with default headers values
func NewGetOrganizationWebhookDetailsTooManyRequests() *GetOrganizationWebhookDetailsTooManyRequests {
	return &GetOrganizationWebhookDetailsTooManyRequests{}
}

/*
GetOrganizationWebhookDetailsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetOrganizationWebhookDetailsTooManyRequests struct {
}

// IsSuccess returns true when this get organization webhook details too many requests response has a 2xx status code
func (o *GetOrganizationWebhookDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization webhook details too many requests response has a 3xx status code
func (o *GetOrganizationWebhookDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization webhook details too many requests response has a 4xx status code
func (o *GetOrganizationWebhookDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organization webhook details too many requests response has a 5xx status code
func (o *GetOrganizationWebhookDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization webhook details too many requests response a status code equal to that given
func (o *GetOrganizationWebhookDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrganizationWebhookDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsTooManyRequests ", 429)
}

func (o *GetOrganizationWebhookDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsTooManyRequests ", 429)
}

func (o *GetOrganizationWebhookDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationWebhookDetailsInternalServerError creates a GetOrganizationWebhookDetailsInternalServerError with default headers values
func NewGetOrganizationWebhookDetailsInternalServerError() *GetOrganizationWebhookDetailsInternalServerError {
	return &GetOrganizationWebhookDetailsInternalServerError{}
}

/*
GetOrganizationWebhookDetailsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetOrganizationWebhookDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get organization webhook details internal server error response has a 2xx status code
func (o *GetOrganizationWebhookDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organization webhook details internal server error response has a 3xx status code
func (o *GetOrganizationWebhookDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization webhook details internal server error response has a 4xx status code
func (o *GetOrganizationWebhookDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization webhook details internal server error response has a 5xx status code
func (o *GetOrganizationWebhookDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get organization webhook details internal server error response a status code equal to that given
func (o *GetOrganizationWebhookDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrganizationWebhookDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationWebhookDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /webhooks/org/{orgName}/{webhookID}][%d] getOrganizationWebhookDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationWebhookDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOrganizationWebhookDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
