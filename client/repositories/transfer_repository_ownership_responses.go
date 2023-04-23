// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// TransferRepositoryOwnershipReader is a Reader for the TransferRepositoryOwnership structure.
type TransferRepositoryOwnershipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferRepositoryOwnershipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTransferRepositoryOwnershipNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTransferRepositoryOwnershipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTransferRepositoryOwnershipForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewTransferRepositoryOwnershipTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTransferRepositoryOwnershipInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTransferRepositoryOwnershipNoContent creates a TransferRepositoryOwnershipNoContent with default headers values
func NewTransferRepositoryOwnershipNoContent() *TransferRepositoryOwnershipNoContent {
	return &TransferRepositoryOwnershipNoContent{}
}

/*
TransferRepositoryOwnershipNoContent describes a response with status code 204, with default header values.

The request has succeeded, no content returned
*/
type TransferRepositoryOwnershipNoContent struct {
}

// IsSuccess returns true when this transfer repository ownership no content response has a 2xx status code
func (o *TransferRepositoryOwnershipNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this transfer repository ownership no content response has a 3xx status code
func (o *TransferRepositoryOwnershipNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer repository ownership no content response has a 4xx status code
func (o *TransferRepositoryOwnershipNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this transfer repository ownership no content response has a 5xx status code
func (o *TransferRepositoryOwnershipNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer repository ownership no content response a status code equal to that given
func (o *TransferRepositoryOwnershipNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *TransferRepositoryOwnershipNoContent) Error() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipNoContent ", 204)
}

func (o *TransferRepositoryOwnershipNoContent) String() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipNoContent ", 204)
}

func (o *TransferRepositoryOwnershipNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferRepositoryOwnershipUnauthorized creates a TransferRepositoryOwnershipUnauthorized with default headers values
func NewTransferRepositoryOwnershipUnauthorized() *TransferRepositoryOwnershipUnauthorized {
	return &TransferRepositoryOwnershipUnauthorized{}
}

/*
TransferRepositoryOwnershipUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type TransferRepositoryOwnershipUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this transfer repository ownership unauthorized response has a 2xx status code
func (o *TransferRepositoryOwnershipUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this transfer repository ownership unauthorized response has a 3xx status code
func (o *TransferRepositoryOwnershipUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer repository ownership unauthorized response has a 4xx status code
func (o *TransferRepositoryOwnershipUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this transfer repository ownership unauthorized response has a 5xx status code
func (o *TransferRepositoryOwnershipUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer repository ownership unauthorized response a status code equal to that given
func (o *TransferRepositoryOwnershipUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TransferRepositoryOwnershipUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipUnauthorized  %+v", 401, o.Payload)
}

func (o *TransferRepositoryOwnershipUnauthorized) String() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipUnauthorized  %+v", 401, o.Payload)
}

func (o *TransferRepositoryOwnershipUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *TransferRepositoryOwnershipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferRepositoryOwnershipForbidden creates a TransferRepositoryOwnershipForbidden with default headers values
func NewTransferRepositoryOwnershipForbidden() *TransferRepositoryOwnershipForbidden {
	return &TransferRepositoryOwnershipForbidden{}
}

/*
TransferRepositoryOwnershipForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the requested operation
*/
type TransferRepositoryOwnershipForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this transfer repository ownership forbidden response has a 2xx status code
func (o *TransferRepositoryOwnershipForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this transfer repository ownership forbidden response has a 3xx status code
func (o *TransferRepositoryOwnershipForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer repository ownership forbidden response has a 4xx status code
func (o *TransferRepositoryOwnershipForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this transfer repository ownership forbidden response has a 5xx status code
func (o *TransferRepositoryOwnershipForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer repository ownership forbidden response a status code equal to that given
func (o *TransferRepositoryOwnershipForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TransferRepositoryOwnershipForbidden) Error() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipForbidden  %+v", 403, o.Payload)
}

func (o *TransferRepositoryOwnershipForbidden) String() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipForbidden  %+v", 403, o.Payload)
}

func (o *TransferRepositoryOwnershipForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *TransferRepositoryOwnershipForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferRepositoryOwnershipTooManyRequests creates a TransferRepositoryOwnershipTooManyRequests with default headers values
func NewTransferRepositoryOwnershipTooManyRequests() *TransferRepositoryOwnershipTooManyRequests {
	return &TransferRepositoryOwnershipTooManyRequests{}
}

/*
TransferRepositoryOwnershipTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type TransferRepositoryOwnershipTooManyRequests struct {
}

// IsSuccess returns true when this transfer repository ownership too many requests response has a 2xx status code
func (o *TransferRepositoryOwnershipTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this transfer repository ownership too many requests response has a 3xx status code
func (o *TransferRepositoryOwnershipTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer repository ownership too many requests response has a 4xx status code
func (o *TransferRepositoryOwnershipTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this transfer repository ownership too many requests response has a 5xx status code
func (o *TransferRepositoryOwnershipTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer repository ownership too many requests response a status code equal to that given
func (o *TransferRepositoryOwnershipTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *TransferRepositoryOwnershipTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipTooManyRequests ", 429)
}

func (o *TransferRepositoryOwnershipTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipTooManyRequests ", 429)
}

func (o *TransferRepositoryOwnershipTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferRepositoryOwnershipInternalServerError creates a TransferRepositoryOwnershipInternalServerError with default headers values
func NewTransferRepositoryOwnershipInternalServerError() *TransferRepositoryOwnershipInternalServerError {
	return &TransferRepositoryOwnershipInternalServerError{}
}

/*
TransferRepositoryOwnershipInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type TransferRepositoryOwnershipInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this transfer repository ownership internal server error response has a 2xx status code
func (o *TransferRepositoryOwnershipInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this transfer repository ownership internal server error response has a 3xx status code
func (o *TransferRepositoryOwnershipInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer repository ownership internal server error response has a 4xx status code
func (o *TransferRepositoryOwnershipInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this transfer repository ownership internal server error response has a 5xx status code
func (o *TransferRepositoryOwnershipInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this transfer repository ownership internal server error response a status code equal to that given
func (o *TransferRepositoryOwnershipInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TransferRepositoryOwnershipInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipInternalServerError  %+v", 500, o.Payload)
}

func (o *TransferRepositoryOwnershipInternalServerError) String() string {
	return fmt.Sprintf("[PUT /repositories/org/{orgName}/{repoName}/transfer][%d] transferRepositoryOwnershipInternalServerError  %+v", 500, o.Payload)
}

func (o *TransferRepositoryOwnershipInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *TransferRepositoryOwnershipInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
