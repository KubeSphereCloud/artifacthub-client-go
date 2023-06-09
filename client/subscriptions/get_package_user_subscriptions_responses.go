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

// GetPackageUserSubscriptionsReader is a Reader for the GetPackageUserSubscriptions structure.
type GetPackageUserSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageUserSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageUserSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPackageUserSubscriptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPackageUserSubscriptionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPackageUserSubscriptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageUserSubscriptionsOK creates a GetPackageUserSubscriptionsOK with default headers values
func NewGetPackageUserSubscriptionsOK() *GetPackageUserSubscriptionsOK {
	return &GetPackageUserSubscriptionsOK{}
}

/*
GetPackageUserSubscriptionsOK describes a response with status code 200, with default header values.

GetPackageUserSubscriptionsOK get package user subscriptions o k
*/
type GetPackageUserSubscriptionsOK struct {
	Payload []*models.SubscriptionsResponse1
}

// IsSuccess returns true when this get package user subscriptions o k response has a 2xx status code
func (o *GetPackageUserSubscriptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get package user subscriptions o k response has a 3xx status code
func (o *GetPackageUserSubscriptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package user subscriptions o k response has a 4xx status code
func (o *GetPackageUserSubscriptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package user subscriptions o k response has a 5xx status code
func (o *GetPackageUserSubscriptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get package user subscriptions o k response a status code equal to that given
func (o *GetPackageUserSubscriptionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPackageUserSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{packageID}][%d] getPackageUserSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetPackageUserSubscriptionsOK) String() string {
	return fmt.Sprintf("[GET /subscriptions/{packageID}][%d] getPackageUserSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetPackageUserSubscriptionsOK) GetPayload() []*models.SubscriptionsResponse1 {
	return o.Payload
}

func (o *GetPackageUserSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageUserSubscriptionsUnauthorized creates a GetPackageUserSubscriptionsUnauthorized with default headers values
func NewGetPackageUserSubscriptionsUnauthorized() *GetPackageUserSubscriptionsUnauthorized {
	return &GetPackageUserSubscriptionsUnauthorized{}
}

/*
GetPackageUserSubscriptionsUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type GetPackageUserSubscriptionsUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get package user subscriptions unauthorized response has a 2xx status code
func (o *GetPackageUserSubscriptionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package user subscriptions unauthorized response has a 3xx status code
func (o *GetPackageUserSubscriptionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package user subscriptions unauthorized response has a 4xx status code
func (o *GetPackageUserSubscriptionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package user subscriptions unauthorized response has a 5xx status code
func (o *GetPackageUserSubscriptionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get package user subscriptions unauthorized response a status code equal to that given
func (o *GetPackageUserSubscriptionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPackageUserSubscriptionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{packageID}][%d] getPackageUserSubscriptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPackageUserSubscriptionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /subscriptions/{packageID}][%d] getPackageUserSubscriptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPackageUserSubscriptionsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPackageUserSubscriptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageUserSubscriptionsTooManyRequests creates a GetPackageUserSubscriptionsTooManyRequests with default headers values
func NewGetPackageUserSubscriptionsTooManyRequests() *GetPackageUserSubscriptionsTooManyRequests {
	return &GetPackageUserSubscriptionsTooManyRequests{}
}

/*
GetPackageUserSubscriptionsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetPackageUserSubscriptionsTooManyRequests struct {
}

// IsSuccess returns true when this get package user subscriptions too many requests response has a 2xx status code
func (o *GetPackageUserSubscriptionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package user subscriptions too many requests response has a 3xx status code
func (o *GetPackageUserSubscriptionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package user subscriptions too many requests response has a 4xx status code
func (o *GetPackageUserSubscriptionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package user subscriptions too many requests response has a 5xx status code
func (o *GetPackageUserSubscriptionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get package user subscriptions too many requests response a status code equal to that given
func (o *GetPackageUserSubscriptionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPackageUserSubscriptionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{packageID}][%d] getPackageUserSubscriptionsTooManyRequests ", 429)
}

func (o *GetPackageUserSubscriptionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /subscriptions/{packageID}][%d] getPackageUserSubscriptionsTooManyRequests ", 429)
}

func (o *GetPackageUserSubscriptionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPackageUserSubscriptionsInternalServerError creates a GetPackageUserSubscriptionsInternalServerError with default headers values
func NewGetPackageUserSubscriptionsInternalServerError() *GetPackageUserSubscriptionsInternalServerError {
	return &GetPackageUserSubscriptionsInternalServerError{}
}

/*
GetPackageUserSubscriptionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetPackageUserSubscriptionsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get package user subscriptions internal server error response has a 2xx status code
func (o *GetPackageUserSubscriptionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package user subscriptions internal server error response has a 3xx status code
func (o *GetPackageUserSubscriptionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package user subscriptions internal server error response has a 4xx status code
func (o *GetPackageUserSubscriptionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package user subscriptions internal server error response has a 5xx status code
func (o *GetPackageUserSubscriptionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get package user subscriptions internal server error response a status code equal to that given
func (o *GetPackageUserSubscriptionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPackageUserSubscriptionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{packageID}][%d] getPackageUserSubscriptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackageUserSubscriptionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /subscriptions/{packageID}][%d] getPackageUserSubscriptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackageUserSubscriptionsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPackageUserSubscriptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
