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

// GetKyvernoPoliciesDetailsReader is a Reader for the GetKyvernoPoliciesDetails structure.
type GetKyvernoPoliciesDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKyvernoPoliciesDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKyvernoPoliciesDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetKyvernoPoliciesDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKyvernoPoliciesDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKyvernoPoliciesDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKyvernoPoliciesDetailsOK creates a GetKyvernoPoliciesDetailsOK with default headers values
func NewGetKyvernoPoliciesDetailsOK() *GetKyvernoPoliciesDetailsOK {
	return &GetKyvernoPoliciesDetailsOK{}
}

/*
GetKyvernoPoliciesDetailsOK describes a response with status code 200, with default header values.

GetKyvernoPoliciesDetailsOK get kyverno policies details o k
*/
type GetKyvernoPoliciesDetailsOK struct {
	Payload *models.KyvernoPolicy
}

// IsSuccess returns true when this get kyverno policies details o k response has a 2xx status code
func (o *GetKyvernoPoliciesDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kyverno policies details o k response has a 3xx status code
func (o *GetKyvernoPoliciesDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kyverno policies details o k response has a 4xx status code
func (o *GetKyvernoPoliciesDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kyverno policies details o k response has a 5xx status code
func (o *GetKyvernoPoliciesDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kyverno policies details o k response a status code equal to that given
func (o *GetKyvernoPoliciesDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKyvernoPoliciesDetailsOK) Error() string {
	return fmt.Sprintf("[GET /packages/kyverno/{repoName}/{packageName}][%d] getKyvernoPoliciesDetailsOK  %+v", 200, o.Payload)
}

func (o *GetKyvernoPoliciesDetailsOK) String() string {
	return fmt.Sprintf("[GET /packages/kyverno/{repoName}/{packageName}][%d] getKyvernoPoliciesDetailsOK  %+v", 200, o.Payload)
}

func (o *GetKyvernoPoliciesDetailsOK) GetPayload() *models.KyvernoPolicy {
	return o.Payload
}

func (o *GetKyvernoPoliciesDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KyvernoPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKyvernoPoliciesDetailsNotFound creates a GetKyvernoPoliciesDetailsNotFound with default headers values
func NewGetKyvernoPoliciesDetailsNotFound() *GetKyvernoPoliciesDetailsNotFound {
	return &GetKyvernoPoliciesDetailsNotFound{}
}

/*
GetKyvernoPoliciesDetailsNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GetKyvernoPoliciesDetailsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get kyverno policies details not found response has a 2xx status code
func (o *GetKyvernoPoliciesDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kyverno policies details not found response has a 3xx status code
func (o *GetKyvernoPoliciesDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kyverno policies details not found response has a 4xx status code
func (o *GetKyvernoPoliciesDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kyverno policies details not found response has a 5xx status code
func (o *GetKyvernoPoliciesDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get kyverno policies details not found response a status code equal to that given
func (o *GetKyvernoPoliciesDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKyvernoPoliciesDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/kyverno/{repoName}/{packageName}][%d] getKyvernoPoliciesDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetKyvernoPoliciesDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /packages/kyverno/{repoName}/{packageName}][%d] getKyvernoPoliciesDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetKyvernoPoliciesDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetKyvernoPoliciesDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKyvernoPoliciesDetailsTooManyRequests creates a GetKyvernoPoliciesDetailsTooManyRequests with default headers values
func NewGetKyvernoPoliciesDetailsTooManyRequests() *GetKyvernoPoliciesDetailsTooManyRequests {
	return &GetKyvernoPoliciesDetailsTooManyRequests{}
}

/*
GetKyvernoPoliciesDetailsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetKyvernoPoliciesDetailsTooManyRequests struct {
}

// IsSuccess returns true when this get kyverno policies details too many requests response has a 2xx status code
func (o *GetKyvernoPoliciesDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kyverno policies details too many requests response has a 3xx status code
func (o *GetKyvernoPoliciesDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kyverno policies details too many requests response has a 4xx status code
func (o *GetKyvernoPoliciesDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kyverno policies details too many requests response has a 5xx status code
func (o *GetKyvernoPoliciesDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get kyverno policies details too many requests response a status code equal to that given
func (o *GetKyvernoPoliciesDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKyvernoPoliciesDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/kyverno/{repoName}/{packageName}][%d] getKyvernoPoliciesDetailsTooManyRequests ", 429)
}

func (o *GetKyvernoPoliciesDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/kyverno/{repoName}/{packageName}][%d] getKyvernoPoliciesDetailsTooManyRequests ", 429)
}

func (o *GetKyvernoPoliciesDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKyvernoPoliciesDetailsInternalServerError creates a GetKyvernoPoliciesDetailsInternalServerError with default headers values
func NewGetKyvernoPoliciesDetailsInternalServerError() *GetKyvernoPoliciesDetailsInternalServerError {
	return &GetKyvernoPoliciesDetailsInternalServerError{}
}

/*
GetKyvernoPoliciesDetailsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetKyvernoPoliciesDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get kyverno policies details internal server error response has a 2xx status code
func (o *GetKyvernoPoliciesDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kyverno policies details internal server error response has a 3xx status code
func (o *GetKyvernoPoliciesDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kyverno policies details internal server error response has a 4xx status code
func (o *GetKyvernoPoliciesDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kyverno policies details internal server error response has a 5xx status code
func (o *GetKyvernoPoliciesDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get kyverno policies details internal server error response a status code equal to that given
func (o *GetKyvernoPoliciesDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKyvernoPoliciesDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/kyverno/{repoName}/{packageName}][%d] getKyvernoPoliciesDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKyvernoPoliciesDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/kyverno/{repoName}/{packageName}][%d] getKyvernoPoliciesDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKyvernoPoliciesDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetKyvernoPoliciesDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
