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

// AddOrganizationWebhookReader is a Reader for the AddOrganizationWebhook structure.
type AddOrganizationWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOrganizationWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddOrganizationWebhookCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddOrganizationWebhookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddOrganizationWebhookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddOrganizationWebhookForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddOrganizationWebhookTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddOrganizationWebhookInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddOrganizationWebhookCreated creates a AddOrganizationWebhookCreated with default headers values
func NewAddOrganizationWebhookCreated() *AddOrganizationWebhookCreated {
	return &AddOrganizationWebhookCreated{}
}

/*
AddOrganizationWebhookCreated describes a response with status code 201, with default header values.

The request has succeeded and has led to the creation of a resource
*/
type AddOrganizationWebhookCreated struct {
}

// IsSuccess returns true when this add organization webhook created response has a 2xx status code
func (o *AddOrganizationWebhookCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add organization webhook created response has a 3xx status code
func (o *AddOrganizationWebhookCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization webhook created response has a 4xx status code
func (o *AddOrganizationWebhookCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add organization webhook created response has a 5xx status code
func (o *AddOrganizationWebhookCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add organization webhook created response a status code equal to that given
func (o *AddOrganizationWebhookCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddOrganizationWebhookCreated) Error() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookCreated ", 201)
}

func (o *AddOrganizationWebhookCreated) String() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookCreated ", 201)
}

func (o *AddOrganizationWebhookCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddOrganizationWebhookBadRequest creates a AddOrganizationWebhookBadRequest with default headers values
func NewAddOrganizationWebhookBadRequest() *AddOrganizationWebhookBadRequest {
	return &AddOrganizationWebhookBadRequest{}
}

/*
AddOrganizationWebhookBadRequest describes a response with status code 400, with default header values.

The request sent was not valid
*/
type AddOrganizationWebhookBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this add organization webhook bad request response has a 2xx status code
func (o *AddOrganizationWebhookBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add organization webhook bad request response has a 3xx status code
func (o *AddOrganizationWebhookBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization webhook bad request response has a 4xx status code
func (o *AddOrganizationWebhookBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add organization webhook bad request response has a 5xx status code
func (o *AddOrganizationWebhookBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add organization webhook bad request response a status code equal to that given
func (o *AddOrganizationWebhookBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddOrganizationWebhookBadRequest) Error() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *AddOrganizationWebhookBadRequest) String() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *AddOrganizationWebhookBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrganizationWebhookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrganizationWebhookUnauthorized creates a AddOrganizationWebhookUnauthorized with default headers values
func NewAddOrganizationWebhookUnauthorized() *AddOrganizationWebhookUnauthorized {
	return &AddOrganizationWebhookUnauthorized{}
}

/*
AddOrganizationWebhookUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type AddOrganizationWebhookUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this add organization webhook unauthorized response has a 2xx status code
func (o *AddOrganizationWebhookUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add organization webhook unauthorized response has a 3xx status code
func (o *AddOrganizationWebhookUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization webhook unauthorized response has a 4xx status code
func (o *AddOrganizationWebhookUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add organization webhook unauthorized response has a 5xx status code
func (o *AddOrganizationWebhookUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add organization webhook unauthorized response a status code equal to that given
func (o *AddOrganizationWebhookUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AddOrganizationWebhookUnauthorized) Error() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookUnauthorized  %+v", 401, o.Payload)
}

func (o *AddOrganizationWebhookUnauthorized) String() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookUnauthorized  %+v", 401, o.Payload)
}

func (o *AddOrganizationWebhookUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrganizationWebhookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrganizationWebhookForbidden creates a AddOrganizationWebhookForbidden with default headers values
func NewAddOrganizationWebhookForbidden() *AddOrganizationWebhookForbidden {
	return &AddOrganizationWebhookForbidden{}
}

/*
AddOrganizationWebhookForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the requested operation
*/
type AddOrganizationWebhookForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this add organization webhook forbidden response has a 2xx status code
func (o *AddOrganizationWebhookForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add organization webhook forbidden response has a 3xx status code
func (o *AddOrganizationWebhookForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization webhook forbidden response has a 4xx status code
func (o *AddOrganizationWebhookForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add organization webhook forbidden response has a 5xx status code
func (o *AddOrganizationWebhookForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add organization webhook forbidden response a status code equal to that given
func (o *AddOrganizationWebhookForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AddOrganizationWebhookForbidden) Error() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookForbidden  %+v", 403, o.Payload)
}

func (o *AddOrganizationWebhookForbidden) String() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookForbidden  %+v", 403, o.Payload)
}

func (o *AddOrganizationWebhookForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrganizationWebhookForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrganizationWebhookTooManyRequests creates a AddOrganizationWebhookTooManyRequests with default headers values
func NewAddOrganizationWebhookTooManyRequests() *AddOrganizationWebhookTooManyRequests {
	return &AddOrganizationWebhookTooManyRequests{}
}

/*
AddOrganizationWebhookTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type AddOrganizationWebhookTooManyRequests struct {
}

// IsSuccess returns true when this add organization webhook too many requests response has a 2xx status code
func (o *AddOrganizationWebhookTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add organization webhook too many requests response has a 3xx status code
func (o *AddOrganizationWebhookTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization webhook too many requests response has a 4xx status code
func (o *AddOrganizationWebhookTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add organization webhook too many requests response has a 5xx status code
func (o *AddOrganizationWebhookTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add organization webhook too many requests response a status code equal to that given
func (o *AddOrganizationWebhookTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AddOrganizationWebhookTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookTooManyRequests ", 429)
}

func (o *AddOrganizationWebhookTooManyRequests) String() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookTooManyRequests ", 429)
}

func (o *AddOrganizationWebhookTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddOrganizationWebhookInternalServerError creates a AddOrganizationWebhookInternalServerError with default headers values
func NewAddOrganizationWebhookInternalServerError() *AddOrganizationWebhookInternalServerError {
	return &AddOrganizationWebhookInternalServerError{}
}

/*
AddOrganizationWebhookInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type AddOrganizationWebhookInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this add organization webhook internal server error response has a 2xx status code
func (o *AddOrganizationWebhookInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add organization webhook internal server error response has a 3xx status code
func (o *AddOrganizationWebhookInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization webhook internal server error response has a 4xx status code
func (o *AddOrganizationWebhookInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add organization webhook internal server error response has a 5xx status code
func (o *AddOrganizationWebhookInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add organization webhook internal server error response a status code equal to that given
func (o *AddOrganizationWebhookInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddOrganizationWebhookInternalServerError) Error() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookInternalServerError  %+v", 500, o.Payload)
}

func (o *AddOrganizationWebhookInternalServerError) String() string {
	return fmt.Sprintf("[POST /webhooks/org/{orgName}][%d] addOrganizationWebhookInternalServerError  %+v", 500, o.Payload)
}

func (o *AddOrganizationWebhookInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrganizationWebhookInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
