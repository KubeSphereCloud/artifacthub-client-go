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

// GetUserWebhookDetailReader is a Reader for the GetUserWebhookDetail structure.
type GetUserWebhookDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserWebhookDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserWebhookDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserWebhookDetailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserWebhookDetailTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserWebhookDetailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserWebhookDetailOK creates a GetUserWebhookDetailOK with default headers values
func NewGetUserWebhookDetailOK() *GetUserWebhookDetailOK {
	return &GetUserWebhookDetailOK{}
}

/*
GetUserWebhookDetailOK describes a response with status code 200, with default header values.

GetUserWebhookDetailOK get user webhook detail o k
*/
type GetUserWebhookDetailOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this get user webhook detail o k response has a 2xx status code
func (o *GetUserWebhookDetailOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user webhook detail o k response has a 3xx status code
func (o *GetUserWebhookDetailOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user webhook detail o k response has a 4xx status code
func (o *GetUserWebhookDetailOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user webhook detail o k response has a 5xx status code
func (o *GetUserWebhookDetailOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user webhook detail o k response a status code equal to that given
func (o *GetUserWebhookDetailOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserWebhookDetailOK) Error() string {
	return fmt.Sprintf("[GET /webhooks/user/{webhookID}][%d] getUserWebhookDetailOK  %+v", 200, o.Payload)
}

func (o *GetUserWebhookDetailOK) String() string {
	return fmt.Sprintf("[GET /webhooks/user/{webhookID}][%d] getUserWebhookDetailOK  %+v", 200, o.Payload)
}

func (o *GetUserWebhookDetailOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *GetUserWebhookDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserWebhookDetailUnauthorized creates a GetUserWebhookDetailUnauthorized with default headers values
func NewGetUserWebhookDetailUnauthorized() *GetUserWebhookDetailUnauthorized {
	return &GetUserWebhookDetailUnauthorized{}
}

/*
GetUserWebhookDetailUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type GetUserWebhookDetailUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user webhook detail unauthorized response has a 2xx status code
func (o *GetUserWebhookDetailUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user webhook detail unauthorized response has a 3xx status code
func (o *GetUserWebhookDetailUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user webhook detail unauthorized response has a 4xx status code
func (o *GetUserWebhookDetailUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user webhook detail unauthorized response has a 5xx status code
func (o *GetUserWebhookDetailUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user webhook detail unauthorized response a status code equal to that given
func (o *GetUserWebhookDetailUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserWebhookDetailUnauthorized) Error() string {
	return fmt.Sprintf("[GET /webhooks/user/{webhookID}][%d] getUserWebhookDetailUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserWebhookDetailUnauthorized) String() string {
	return fmt.Sprintf("[GET /webhooks/user/{webhookID}][%d] getUserWebhookDetailUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserWebhookDetailUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserWebhookDetailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserWebhookDetailTooManyRequests creates a GetUserWebhookDetailTooManyRequests with default headers values
func NewGetUserWebhookDetailTooManyRequests() *GetUserWebhookDetailTooManyRequests {
	return &GetUserWebhookDetailTooManyRequests{}
}

/*
GetUserWebhookDetailTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetUserWebhookDetailTooManyRequests struct {
}

// IsSuccess returns true when this get user webhook detail too many requests response has a 2xx status code
func (o *GetUserWebhookDetailTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user webhook detail too many requests response has a 3xx status code
func (o *GetUserWebhookDetailTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user webhook detail too many requests response has a 4xx status code
func (o *GetUserWebhookDetailTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user webhook detail too many requests response has a 5xx status code
func (o *GetUserWebhookDetailTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user webhook detail too many requests response a status code equal to that given
func (o *GetUserWebhookDetailTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserWebhookDetailTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /webhooks/user/{webhookID}][%d] getUserWebhookDetailTooManyRequests ", 429)
}

func (o *GetUserWebhookDetailTooManyRequests) String() string {
	return fmt.Sprintf("[GET /webhooks/user/{webhookID}][%d] getUserWebhookDetailTooManyRequests ", 429)
}

func (o *GetUserWebhookDetailTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserWebhookDetailInternalServerError creates a GetUserWebhookDetailInternalServerError with default headers values
func NewGetUserWebhookDetailInternalServerError() *GetUserWebhookDetailInternalServerError {
	return &GetUserWebhookDetailInternalServerError{}
}

/*
GetUserWebhookDetailInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetUserWebhookDetailInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user webhook detail internal server error response has a 2xx status code
func (o *GetUserWebhookDetailInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user webhook detail internal server error response has a 3xx status code
func (o *GetUserWebhookDetailInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user webhook detail internal server error response has a 4xx status code
func (o *GetUserWebhookDetailInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user webhook detail internal server error response has a 5xx status code
func (o *GetUserWebhookDetailInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user webhook detail internal server error response a status code equal to that given
func (o *GetUserWebhookDetailInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserWebhookDetailInternalServerError) Error() string {
	return fmt.Sprintf("[GET /webhooks/user/{webhookID}][%d] getUserWebhookDetailInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserWebhookDetailInternalServerError) String() string {
	return fmt.Sprintf("[GET /webhooks/user/{webhookID}][%d] getUserWebhookDetailInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserWebhookDetailInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserWebhookDetailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
