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

// UpdateOrganizationWebhookReader is a Reader for the UpdateOrganizationWebhook structure.
type UpdateOrganizationWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateOrganizationWebhookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrganizationWebhookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrganizationWebhookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrganizationWebhookForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateOrganizationWebhookTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateOrganizationWebhookInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationWebhookNoContent creates a UpdateOrganizationWebhookNoContent with default headers values
func NewUpdateOrganizationWebhookNoContent() *UpdateOrganizationWebhookNoContent {
	return &UpdateOrganizationWebhookNoContent{}
}

/*
UpdateOrganizationWebhookNoContent describes a response with status code 204, with default header values.

The request has succeeded, no content returned
*/
type UpdateOrganizationWebhookNoContent struct {
}

// IsSuccess returns true when this update organization webhook no content response has a 2xx status code
func (o *UpdateOrganizationWebhookNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization webhook no content response has a 3xx status code
func (o *UpdateOrganizationWebhookNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization webhook no content response has a 4xx status code
func (o *UpdateOrganizationWebhookNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization webhook no content response has a 5xx status code
func (o *UpdateOrganizationWebhookNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization webhook no content response a status code equal to that given
func (o *UpdateOrganizationWebhookNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateOrganizationWebhookNoContent) Error() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookNoContent ", 204)
}

func (o *UpdateOrganizationWebhookNoContent) String() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookNoContent ", 204)
}

func (o *UpdateOrganizationWebhookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationWebhookBadRequest creates a UpdateOrganizationWebhookBadRequest with default headers values
func NewUpdateOrganizationWebhookBadRequest() *UpdateOrganizationWebhookBadRequest {
	return &UpdateOrganizationWebhookBadRequest{}
}

/*
UpdateOrganizationWebhookBadRequest describes a response with status code 400, with default header values.

The request sent was not valid
*/
type UpdateOrganizationWebhookBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update organization webhook bad request response has a 2xx status code
func (o *UpdateOrganizationWebhookBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization webhook bad request response has a 3xx status code
func (o *UpdateOrganizationWebhookBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization webhook bad request response has a 4xx status code
func (o *UpdateOrganizationWebhookBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization webhook bad request response has a 5xx status code
func (o *UpdateOrganizationWebhookBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization webhook bad request response a status code equal to that given
func (o *UpdateOrganizationWebhookBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateOrganizationWebhookBadRequest) Error() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrganizationWebhookBadRequest) String() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrganizationWebhookBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOrganizationWebhookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationWebhookUnauthorized creates a UpdateOrganizationWebhookUnauthorized with default headers values
func NewUpdateOrganizationWebhookUnauthorized() *UpdateOrganizationWebhookUnauthorized {
	return &UpdateOrganizationWebhookUnauthorized{}
}

/*
UpdateOrganizationWebhookUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type UpdateOrganizationWebhookUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this update organization webhook unauthorized response has a 2xx status code
func (o *UpdateOrganizationWebhookUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization webhook unauthorized response has a 3xx status code
func (o *UpdateOrganizationWebhookUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization webhook unauthorized response has a 4xx status code
func (o *UpdateOrganizationWebhookUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization webhook unauthorized response has a 5xx status code
func (o *UpdateOrganizationWebhookUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization webhook unauthorized response a status code equal to that given
func (o *UpdateOrganizationWebhookUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateOrganizationWebhookUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrganizationWebhookUnauthorized) String() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrganizationWebhookUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOrganizationWebhookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationWebhookForbidden creates a UpdateOrganizationWebhookForbidden with default headers values
func NewUpdateOrganizationWebhookForbidden() *UpdateOrganizationWebhookForbidden {
	return &UpdateOrganizationWebhookForbidden{}
}

/*
UpdateOrganizationWebhookForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the requested operation
*/
type UpdateOrganizationWebhookForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this update organization webhook forbidden response has a 2xx status code
func (o *UpdateOrganizationWebhookForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization webhook forbidden response has a 3xx status code
func (o *UpdateOrganizationWebhookForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization webhook forbidden response has a 4xx status code
func (o *UpdateOrganizationWebhookForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization webhook forbidden response has a 5xx status code
func (o *UpdateOrganizationWebhookForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization webhook forbidden response a status code equal to that given
func (o *UpdateOrganizationWebhookForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateOrganizationWebhookForbidden) Error() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrganizationWebhookForbidden) String() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrganizationWebhookForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOrganizationWebhookForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationWebhookTooManyRequests creates a UpdateOrganizationWebhookTooManyRequests with default headers values
func NewUpdateOrganizationWebhookTooManyRequests() *UpdateOrganizationWebhookTooManyRequests {
	return &UpdateOrganizationWebhookTooManyRequests{}
}

/*
UpdateOrganizationWebhookTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type UpdateOrganizationWebhookTooManyRequests struct {
}

// IsSuccess returns true when this update organization webhook too many requests response has a 2xx status code
func (o *UpdateOrganizationWebhookTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization webhook too many requests response has a 3xx status code
func (o *UpdateOrganizationWebhookTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization webhook too many requests response has a 4xx status code
func (o *UpdateOrganizationWebhookTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization webhook too many requests response has a 5xx status code
func (o *UpdateOrganizationWebhookTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization webhook too many requests response a status code equal to that given
func (o *UpdateOrganizationWebhookTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateOrganizationWebhookTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookTooManyRequests ", 429)
}

func (o *UpdateOrganizationWebhookTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookTooManyRequests ", 429)
}

func (o *UpdateOrganizationWebhookTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOrganizationWebhookInternalServerError creates a UpdateOrganizationWebhookInternalServerError with default headers values
func NewUpdateOrganizationWebhookInternalServerError() *UpdateOrganizationWebhookInternalServerError {
	return &UpdateOrganizationWebhookInternalServerError{}
}

/*
UpdateOrganizationWebhookInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type UpdateOrganizationWebhookInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update organization webhook internal server error response has a 2xx status code
func (o *UpdateOrganizationWebhookInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization webhook internal server error response has a 3xx status code
func (o *UpdateOrganizationWebhookInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization webhook internal server error response has a 4xx status code
func (o *UpdateOrganizationWebhookInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization webhook internal server error response has a 5xx status code
func (o *UpdateOrganizationWebhookInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update organization webhook internal server error response a status code equal to that given
func (o *UpdateOrganizationWebhookInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateOrganizationWebhookInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOrganizationWebhookInternalServerError) String() string {
	return fmt.Sprintf("[PUT /webhooks/org/{orgName}/{webhookID}][%d] updateOrganizationWebhookInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOrganizationWebhookInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOrganizationWebhookInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}