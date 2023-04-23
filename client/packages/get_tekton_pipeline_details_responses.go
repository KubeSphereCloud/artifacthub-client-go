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

// GetTektonPipelineDetailsReader is a Reader for the GetTektonPipelineDetails structure.
type GetTektonPipelineDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTektonPipelineDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTektonPipelineDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTektonPipelineDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTektonPipelineDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTektonPipelineDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTektonPipelineDetailsOK creates a GetTektonPipelineDetailsOK with default headers values
func NewGetTektonPipelineDetailsOK() *GetTektonPipelineDetailsOK {
	return &GetTektonPipelineDetailsOK{}
}

/*
GetTektonPipelineDetailsOK describes a response with status code 200, with default header values.

GetTektonPipelineDetailsOK get tekton pipeline details o k
*/
type GetTektonPipelineDetailsOK struct {
	Payload *models.TektonPipelinePackage
}

// IsSuccess returns true when this get tekton pipeline details o k response has a 2xx status code
func (o *GetTektonPipelineDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tekton pipeline details o k response has a 3xx status code
func (o *GetTektonPipelineDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline details o k response has a 4xx status code
func (o *GetTektonPipelineDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tekton pipeline details o k response has a 5xx status code
func (o *GetTektonPipelineDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline details o k response a status code equal to that given
func (o *GetTektonPipelineDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTektonPipelineDetailsOK) Error() string {
	return fmt.Sprintf("[GET /packages/tekton-pipeline/{repoName}/{packageName}][%d] getTektonPipelineDetailsOK  %+v", 200, o.Payload)
}

func (o *GetTektonPipelineDetailsOK) String() string {
	return fmt.Sprintf("[GET /packages/tekton-pipeline/{repoName}/{packageName}][%d] getTektonPipelineDetailsOK  %+v", 200, o.Payload)
}

func (o *GetTektonPipelineDetailsOK) GetPayload() *models.TektonPipelinePackage {
	return o.Payload
}

func (o *GetTektonPipelineDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TektonPipelinePackage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTektonPipelineDetailsNotFound creates a GetTektonPipelineDetailsNotFound with default headers values
func NewGetTektonPipelineDetailsNotFound() *GetTektonPipelineDetailsNotFound {
	return &GetTektonPipelineDetailsNotFound{}
}

/*
GetTektonPipelineDetailsNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GetTektonPipelineDetailsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tekton pipeline details not found response has a 2xx status code
func (o *GetTektonPipelineDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tekton pipeline details not found response has a 3xx status code
func (o *GetTektonPipelineDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline details not found response has a 4xx status code
func (o *GetTektonPipelineDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tekton pipeline details not found response has a 5xx status code
func (o *GetTektonPipelineDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline details not found response a status code equal to that given
func (o *GetTektonPipelineDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTektonPipelineDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/tekton-pipeline/{repoName}/{packageName}][%d] getTektonPipelineDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetTektonPipelineDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /packages/tekton-pipeline/{repoName}/{packageName}][%d] getTektonPipelineDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetTektonPipelineDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTektonPipelineDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTektonPipelineDetailsTooManyRequests creates a GetTektonPipelineDetailsTooManyRequests with default headers values
func NewGetTektonPipelineDetailsTooManyRequests() *GetTektonPipelineDetailsTooManyRequests {
	return &GetTektonPipelineDetailsTooManyRequests{}
}

/*
GetTektonPipelineDetailsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetTektonPipelineDetailsTooManyRequests struct {
}

// IsSuccess returns true when this get tekton pipeline details too many requests response has a 2xx status code
func (o *GetTektonPipelineDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tekton pipeline details too many requests response has a 3xx status code
func (o *GetTektonPipelineDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline details too many requests response has a 4xx status code
func (o *GetTektonPipelineDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tekton pipeline details too many requests response has a 5xx status code
func (o *GetTektonPipelineDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline details too many requests response a status code equal to that given
func (o *GetTektonPipelineDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTektonPipelineDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/tekton-pipeline/{repoName}/{packageName}][%d] getTektonPipelineDetailsTooManyRequests ", 429)
}

func (o *GetTektonPipelineDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/tekton-pipeline/{repoName}/{packageName}][%d] getTektonPipelineDetailsTooManyRequests ", 429)
}

func (o *GetTektonPipelineDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTektonPipelineDetailsInternalServerError creates a GetTektonPipelineDetailsInternalServerError with default headers values
func NewGetTektonPipelineDetailsInternalServerError() *GetTektonPipelineDetailsInternalServerError {
	return &GetTektonPipelineDetailsInternalServerError{}
}

/*
GetTektonPipelineDetailsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetTektonPipelineDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tekton pipeline details internal server error response has a 2xx status code
func (o *GetTektonPipelineDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tekton pipeline details internal server error response has a 3xx status code
func (o *GetTektonPipelineDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline details internal server error response has a 4xx status code
func (o *GetTektonPipelineDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tekton pipeline details internal server error response has a 5xx status code
func (o *GetTektonPipelineDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tekton pipeline details internal server error response a status code equal to that given
func (o *GetTektonPipelineDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTektonPipelineDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/tekton-pipeline/{repoName}/{packageName}][%d] getTektonPipelineDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTektonPipelineDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/tekton-pipeline/{repoName}/{packageName}][%d] getTektonPipelineDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTektonPipelineDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTektonPipelineDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
