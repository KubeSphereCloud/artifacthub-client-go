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

// GetUserWebhooksReader is a Reader for the GetUserWebhooks structure.
type GetUserWebhooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserWebhooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserWebhooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserWebhooksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserWebhooksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserWebhooksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserWebhooksOK creates a GetUserWebhooksOK with default headers values
func NewGetUserWebhooksOK() *GetUserWebhooksOK {
	return &GetUserWebhooksOK{}
}

/*
GetUserWebhooksOK describes a response with status code 200, with default header values.

GetUserWebhooksOK get user webhooks o k
*/
type GetUserWebhooksOK struct {

	/* Total number of user's webhooks
	 */
	PaginationTotalCount string

	Payload []*models.Webhook
}

// IsSuccess returns true when this get user webhooks o k response has a 2xx status code
func (o *GetUserWebhooksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user webhooks o k response has a 3xx status code
func (o *GetUserWebhooksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user webhooks o k response has a 4xx status code
func (o *GetUserWebhooksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user webhooks o k response has a 5xx status code
func (o *GetUserWebhooksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user webhooks o k response a status code equal to that given
func (o *GetUserWebhooksOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserWebhooksOK) Error() string {
	return fmt.Sprintf("[GET /webhooks/user][%d] getUserWebhooksOK  %+v", 200, o.Payload)
}

func (o *GetUserWebhooksOK) String() string {
	return fmt.Sprintf("[GET /webhooks/user][%d] getUserWebhooksOK  %+v", 200, o.Payload)
}

func (o *GetUserWebhooksOK) GetPayload() []*models.Webhook {
	return o.Payload
}

func (o *GetUserWebhooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Pagination-Total-Count
	hdrPaginationTotalCount := response.GetHeader("Pagination-Total-Count")

	if hdrPaginationTotalCount != "" {
		o.PaginationTotalCount = hdrPaginationTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserWebhooksUnauthorized creates a GetUserWebhooksUnauthorized with default headers values
func NewGetUserWebhooksUnauthorized() *GetUserWebhooksUnauthorized {
	return &GetUserWebhooksUnauthorized{}
}

/*
GetUserWebhooksUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type GetUserWebhooksUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user webhooks unauthorized response has a 2xx status code
func (o *GetUserWebhooksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user webhooks unauthorized response has a 3xx status code
func (o *GetUserWebhooksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user webhooks unauthorized response has a 4xx status code
func (o *GetUserWebhooksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user webhooks unauthorized response has a 5xx status code
func (o *GetUserWebhooksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user webhooks unauthorized response a status code equal to that given
func (o *GetUserWebhooksUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserWebhooksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /webhooks/user][%d] getUserWebhooksUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserWebhooksUnauthorized) String() string {
	return fmt.Sprintf("[GET /webhooks/user][%d] getUserWebhooksUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserWebhooksUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserWebhooksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserWebhooksTooManyRequests creates a GetUserWebhooksTooManyRequests with default headers values
func NewGetUserWebhooksTooManyRequests() *GetUserWebhooksTooManyRequests {
	return &GetUserWebhooksTooManyRequests{}
}

/*
GetUserWebhooksTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetUserWebhooksTooManyRequests struct {
}

// IsSuccess returns true when this get user webhooks too many requests response has a 2xx status code
func (o *GetUserWebhooksTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user webhooks too many requests response has a 3xx status code
func (o *GetUserWebhooksTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user webhooks too many requests response has a 4xx status code
func (o *GetUserWebhooksTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user webhooks too many requests response has a 5xx status code
func (o *GetUserWebhooksTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user webhooks too many requests response a status code equal to that given
func (o *GetUserWebhooksTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserWebhooksTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /webhooks/user][%d] getUserWebhooksTooManyRequests ", 429)
}

func (o *GetUserWebhooksTooManyRequests) String() string {
	return fmt.Sprintf("[GET /webhooks/user][%d] getUserWebhooksTooManyRequests ", 429)
}

func (o *GetUserWebhooksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserWebhooksInternalServerError creates a GetUserWebhooksInternalServerError with default headers values
func NewGetUserWebhooksInternalServerError() *GetUserWebhooksInternalServerError {
	return &GetUserWebhooksInternalServerError{}
}

/*
GetUserWebhooksInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetUserWebhooksInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user webhooks internal server error response has a 2xx status code
func (o *GetUserWebhooksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user webhooks internal server error response has a 3xx status code
func (o *GetUserWebhooksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user webhooks internal server error response has a 4xx status code
func (o *GetUserWebhooksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user webhooks internal server error response has a 5xx status code
func (o *GetUserWebhooksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user webhooks internal server error response a status code equal to that given
func (o *GetUserWebhooksInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserWebhooksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /webhooks/user][%d] getUserWebhooksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserWebhooksInternalServerError) String() string {
	return fmt.Sprintf("[GET /webhooks/user][%d] getUserWebhooksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserWebhooksInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserWebhooksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}