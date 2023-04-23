// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// TogglePackageStarReader is a Reader for the TogglePackageStar structure.
type TogglePackageStarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TogglePackageStarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTogglePackageStarNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTogglePackageStarUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewTogglePackageStarTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTogglePackageStarInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTogglePackageStarNoContent creates a TogglePackageStarNoContent with default headers values
func NewTogglePackageStarNoContent() *TogglePackageStarNoContent {
	return &TogglePackageStarNoContent{}
}

/*
TogglePackageStarNoContent describes a response with status code 204, with default header values.

The request has succeeded, no content returned
*/
type TogglePackageStarNoContent struct {
}

// IsSuccess returns true when this toggle package star no content response has a 2xx status code
func (o *TogglePackageStarNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this toggle package star no content response has a 3xx status code
func (o *TogglePackageStarNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this toggle package star no content response has a 4xx status code
func (o *TogglePackageStarNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this toggle package star no content response has a 5xx status code
func (o *TogglePackageStarNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this toggle package star no content response a status code equal to that given
func (o *TogglePackageStarNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *TogglePackageStarNoContent) Error() string {
	return fmt.Sprintf("[PUT /packages/{packageID}/stars][%d] togglePackageStarNoContent ", 204)
}

func (o *TogglePackageStarNoContent) String() string {
	return fmt.Sprintf("[PUT /packages/{packageID}/stars][%d] togglePackageStarNoContent ", 204)
}

func (o *TogglePackageStarNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTogglePackageStarUnauthorized creates a TogglePackageStarUnauthorized with default headers values
func NewTogglePackageStarUnauthorized() *TogglePackageStarUnauthorized {
	return &TogglePackageStarUnauthorized{}
}

/*
TogglePackageStarUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type TogglePackageStarUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this toggle package star unauthorized response has a 2xx status code
func (o *TogglePackageStarUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this toggle package star unauthorized response has a 3xx status code
func (o *TogglePackageStarUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this toggle package star unauthorized response has a 4xx status code
func (o *TogglePackageStarUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this toggle package star unauthorized response has a 5xx status code
func (o *TogglePackageStarUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this toggle package star unauthorized response a status code equal to that given
func (o *TogglePackageStarUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TogglePackageStarUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /packages/{packageID}/stars][%d] togglePackageStarUnauthorized  %+v", 401, o.Payload)
}

func (o *TogglePackageStarUnauthorized) String() string {
	return fmt.Sprintf("[PUT /packages/{packageID}/stars][%d] togglePackageStarUnauthorized  %+v", 401, o.Payload)
}

func (o *TogglePackageStarUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *TogglePackageStarUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTogglePackageStarTooManyRequests creates a TogglePackageStarTooManyRequests with default headers values
func NewTogglePackageStarTooManyRequests() *TogglePackageStarTooManyRequests {
	return &TogglePackageStarTooManyRequests{}
}

/*
TogglePackageStarTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type TogglePackageStarTooManyRequests struct {
}

// IsSuccess returns true when this toggle package star too many requests response has a 2xx status code
func (o *TogglePackageStarTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this toggle package star too many requests response has a 3xx status code
func (o *TogglePackageStarTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this toggle package star too many requests response has a 4xx status code
func (o *TogglePackageStarTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this toggle package star too many requests response has a 5xx status code
func (o *TogglePackageStarTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this toggle package star too many requests response a status code equal to that given
func (o *TogglePackageStarTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *TogglePackageStarTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /packages/{packageID}/stars][%d] togglePackageStarTooManyRequests ", 429)
}

func (o *TogglePackageStarTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /packages/{packageID}/stars][%d] togglePackageStarTooManyRequests ", 429)
}

func (o *TogglePackageStarTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTogglePackageStarInternalServerError creates a TogglePackageStarInternalServerError with default headers values
func NewTogglePackageStarInternalServerError() *TogglePackageStarInternalServerError {
	return &TogglePackageStarInternalServerError{}
}

/*
TogglePackageStarInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type TogglePackageStarInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this toggle package star internal server error response has a 2xx status code
func (o *TogglePackageStarInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this toggle package star internal server error response has a 3xx status code
func (o *TogglePackageStarInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this toggle package star internal server error response has a 4xx status code
func (o *TogglePackageStarInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this toggle package star internal server error response has a 5xx status code
func (o *TogglePackageStarInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this toggle package star internal server error response a status code equal to that given
func (o *TogglePackageStarInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TogglePackageStarInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /packages/{packageID}/stars][%d] togglePackageStarInternalServerError  %+v", 500, o.Payload)
}

func (o *TogglePackageStarInternalServerError) String() string {
	return fmt.Sprintf("[PUT /packages/{packageID}/stars][%d] togglePackageStarInternalServerError  %+v", 500, o.Payload)
}

func (o *TogglePackageStarInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *TogglePackageStarInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
