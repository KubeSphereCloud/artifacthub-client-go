// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// GetUserSubscriptionsReader is a Reader for the GetUserSubscriptions structure.
type GetUserSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserSubscriptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserSubscriptionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserSubscriptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserSubscriptionsOK creates a GetUserSubscriptionsOK with default headers values
func NewGetUserSubscriptionsOK() *GetUserSubscriptionsOK {
	return &GetUserSubscriptionsOK{}
}

/*
GetUserSubscriptionsOK describes a response with status code 200, with default header values.

GetUserSubscriptionsOK get user subscriptions o k
*/
type GetUserSubscriptionsOK struct {

	/* Total number of subscriptions
	 */
	PaginationTotalCount string

	Payload []*models.SubscriptionsResponse
}

// IsSuccess returns true when this get user subscriptions o k response has a 2xx status code
func (o *GetUserSubscriptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user subscriptions o k response has a 3xx status code
func (o *GetUserSubscriptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user subscriptions o k response has a 4xx status code
func (o *GetUserSubscriptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user subscriptions o k response has a 5xx status code
func (o *GetUserSubscriptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user subscriptions o k response a status code equal to that given
func (o *GetUserSubscriptionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getUserSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetUserSubscriptionsOK) String() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getUserSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetUserSubscriptionsOK) GetPayload() []*models.SubscriptionsResponse {
	return o.Payload
}

func (o *GetUserSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUserSubscriptionsUnauthorized creates a GetUserSubscriptionsUnauthorized with default headers values
func NewGetUserSubscriptionsUnauthorized() *GetUserSubscriptionsUnauthorized {
	return &GetUserSubscriptionsUnauthorized{}
}

/*
GetUserSubscriptionsUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type GetUserSubscriptionsUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user subscriptions unauthorized response has a 2xx status code
func (o *GetUserSubscriptionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user subscriptions unauthorized response has a 3xx status code
func (o *GetUserSubscriptionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user subscriptions unauthorized response has a 4xx status code
func (o *GetUserSubscriptionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user subscriptions unauthorized response has a 5xx status code
func (o *GetUserSubscriptionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user subscriptions unauthorized response a status code equal to that given
func (o *GetUserSubscriptionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserSubscriptionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getUserSubscriptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserSubscriptionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getUserSubscriptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserSubscriptionsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserSubscriptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSubscriptionsTooManyRequests creates a GetUserSubscriptionsTooManyRequests with default headers values
func NewGetUserSubscriptionsTooManyRequests() *GetUserSubscriptionsTooManyRequests {
	return &GetUserSubscriptionsTooManyRequests{}
}

/*
GetUserSubscriptionsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetUserSubscriptionsTooManyRequests struct {
}

// IsSuccess returns true when this get user subscriptions too many requests response has a 2xx status code
func (o *GetUserSubscriptionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user subscriptions too many requests response has a 3xx status code
func (o *GetUserSubscriptionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user subscriptions too many requests response has a 4xx status code
func (o *GetUserSubscriptionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user subscriptions too many requests response has a 5xx status code
func (o *GetUserSubscriptionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user subscriptions too many requests response a status code equal to that given
func (o *GetUserSubscriptionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserSubscriptionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getUserSubscriptionsTooManyRequests ", 429)
}

func (o *GetUserSubscriptionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getUserSubscriptionsTooManyRequests ", 429)
}

func (o *GetUserSubscriptionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserSubscriptionsInternalServerError creates a GetUserSubscriptionsInternalServerError with default headers values
func NewGetUserSubscriptionsInternalServerError() *GetUserSubscriptionsInternalServerError {
	return &GetUserSubscriptionsInternalServerError{}
}

/*
GetUserSubscriptionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetUserSubscriptionsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user subscriptions internal server error response has a 2xx status code
func (o *GetUserSubscriptionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user subscriptions internal server error response has a 3xx status code
func (o *GetUserSubscriptionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user subscriptions internal server error response has a 4xx status code
func (o *GetUserSubscriptionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user subscriptions internal server error response has a 5xx status code
func (o *GetUserSubscriptionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user subscriptions internal server error response a status code equal to that given
func (o *GetUserSubscriptionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserSubscriptionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getUserSubscriptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserSubscriptionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getUserSubscriptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserSubscriptionsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserSubscriptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
