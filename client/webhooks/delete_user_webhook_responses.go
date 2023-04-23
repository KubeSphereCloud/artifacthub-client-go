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

// DeleteUserWebhookReader is a Reader for the DeleteUserWebhook structure.
type DeleteUserWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserWebhookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUserWebhookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserWebhookTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserWebhookInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserWebhookNoContent creates a DeleteUserWebhookNoContent with default headers values
func NewDeleteUserWebhookNoContent() *DeleteUserWebhookNoContent {
	return &DeleteUserWebhookNoContent{}
}

/*
DeleteUserWebhookNoContent describes a response with status code 204, with default header values.

The request has succeeded, no content returned
*/
type DeleteUserWebhookNoContent struct {
}

// IsSuccess returns true when this delete user webhook no content response has a 2xx status code
func (o *DeleteUserWebhookNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user webhook no content response has a 3xx status code
func (o *DeleteUserWebhookNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user webhook no content response has a 4xx status code
func (o *DeleteUserWebhookNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user webhook no content response has a 5xx status code
func (o *DeleteUserWebhookNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user webhook no content response a status code equal to that given
func (o *DeleteUserWebhookNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteUserWebhookNoContent) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/user/{webhookID}][%d] deleteUserWebhookNoContent ", 204)
}

func (o *DeleteUserWebhookNoContent) String() string {
	return fmt.Sprintf("[DELETE /webhooks/user/{webhookID}][%d] deleteUserWebhookNoContent ", 204)
}

func (o *DeleteUserWebhookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserWebhookUnauthorized creates a DeleteUserWebhookUnauthorized with default headers values
func NewDeleteUserWebhookUnauthorized() *DeleteUserWebhookUnauthorized {
	return &DeleteUserWebhookUnauthorized{}
}

/*
DeleteUserWebhookUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type DeleteUserWebhookUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete user webhook unauthorized response has a 2xx status code
func (o *DeleteUserWebhookUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user webhook unauthorized response has a 3xx status code
func (o *DeleteUserWebhookUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user webhook unauthorized response has a 4xx status code
func (o *DeleteUserWebhookUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user webhook unauthorized response has a 5xx status code
func (o *DeleteUserWebhookUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user webhook unauthorized response a status code equal to that given
func (o *DeleteUserWebhookUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteUserWebhookUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/user/{webhookID}][%d] deleteUserWebhookUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserWebhookUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /webhooks/user/{webhookID}][%d] deleteUserWebhookUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserWebhookUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserWebhookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserWebhookTooManyRequests creates a DeleteUserWebhookTooManyRequests with default headers values
func NewDeleteUserWebhookTooManyRequests() *DeleteUserWebhookTooManyRequests {
	return &DeleteUserWebhookTooManyRequests{}
}

/*
DeleteUserWebhookTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type DeleteUserWebhookTooManyRequests struct {
}

// IsSuccess returns true when this delete user webhook too many requests response has a 2xx status code
func (o *DeleteUserWebhookTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user webhook too many requests response has a 3xx status code
func (o *DeleteUserWebhookTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user webhook too many requests response has a 4xx status code
func (o *DeleteUserWebhookTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user webhook too many requests response has a 5xx status code
func (o *DeleteUserWebhookTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user webhook too many requests response a status code equal to that given
func (o *DeleteUserWebhookTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteUserWebhookTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/user/{webhookID}][%d] deleteUserWebhookTooManyRequests ", 429)
}

func (o *DeleteUserWebhookTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /webhooks/user/{webhookID}][%d] deleteUserWebhookTooManyRequests ", 429)
}

func (o *DeleteUserWebhookTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserWebhookInternalServerError creates a DeleteUserWebhookInternalServerError with default headers values
func NewDeleteUserWebhookInternalServerError() *DeleteUserWebhookInternalServerError {
	return &DeleteUserWebhookInternalServerError{}
}

/*
DeleteUserWebhookInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type DeleteUserWebhookInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete user webhook internal server error response has a 2xx status code
func (o *DeleteUserWebhookInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user webhook internal server error response has a 3xx status code
func (o *DeleteUserWebhookInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user webhook internal server error response has a 4xx status code
func (o *DeleteUserWebhookInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user webhook internal server error response has a 5xx status code
func (o *DeleteUserWebhookInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete user webhook internal server error response a status code equal to that given
func (o *DeleteUserWebhookInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteUserWebhookInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/user/{webhookID}][%d] deleteUserWebhookInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserWebhookInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /webhooks/user/{webhookID}][%d] deleteUserWebhookInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserWebhookInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserWebhookInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
