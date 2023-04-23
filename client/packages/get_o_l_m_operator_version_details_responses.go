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

// GetOLMOperatorVersionDetailsReader is a Reader for the GetOLMOperatorVersionDetails structure.
type GetOLMOperatorVersionDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOLMOperatorVersionDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOLMOperatorVersionDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetOLMOperatorVersionDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOLMOperatorVersionDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOLMOperatorVersionDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOLMOperatorVersionDetailsOK creates a GetOLMOperatorVersionDetailsOK with default headers values
func NewGetOLMOperatorVersionDetailsOK() *GetOLMOperatorVersionDetailsOK {
	return &GetOLMOperatorVersionDetailsOK{}
}

/*
GetOLMOperatorVersionDetailsOK describes a response with status code 200, with default header values.

GetOLMOperatorVersionDetailsOK get o l m operator version details o k
*/
type GetOLMOperatorVersionDetailsOK struct {
	Payload *models.OLMPackage
}

// IsSuccess returns true when this get o l m operator version details o k response has a 2xx status code
func (o *GetOLMOperatorVersionDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get o l m operator version details o k response has a 3xx status code
func (o *GetOLMOperatorVersionDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get o l m operator version details o k response has a 4xx status code
func (o *GetOLMOperatorVersionDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get o l m operator version details o k response has a 5xx status code
func (o *GetOLMOperatorVersionDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get o l m operator version details o k response a status code equal to that given
func (o *GetOLMOperatorVersionDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOLMOperatorVersionDetailsOK) Error() string {
	return fmt.Sprintf("[GET /packages/olm/{repoName}/{packageName}/{version}][%d] getOLMOperatorVersionDetailsOK  %+v", 200, o.Payload)
}

func (o *GetOLMOperatorVersionDetailsOK) String() string {
	return fmt.Sprintf("[GET /packages/olm/{repoName}/{packageName}/{version}][%d] getOLMOperatorVersionDetailsOK  %+v", 200, o.Payload)
}

func (o *GetOLMOperatorVersionDetailsOK) GetPayload() *models.OLMPackage {
	return o.Payload
}

func (o *GetOLMOperatorVersionDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OLMPackage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOLMOperatorVersionDetailsNotFound creates a GetOLMOperatorVersionDetailsNotFound with default headers values
func NewGetOLMOperatorVersionDetailsNotFound() *GetOLMOperatorVersionDetailsNotFound {
	return &GetOLMOperatorVersionDetailsNotFound{}
}

/*
GetOLMOperatorVersionDetailsNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GetOLMOperatorVersionDetailsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get o l m operator version details not found response has a 2xx status code
func (o *GetOLMOperatorVersionDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get o l m operator version details not found response has a 3xx status code
func (o *GetOLMOperatorVersionDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get o l m operator version details not found response has a 4xx status code
func (o *GetOLMOperatorVersionDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get o l m operator version details not found response has a 5xx status code
func (o *GetOLMOperatorVersionDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get o l m operator version details not found response a status code equal to that given
func (o *GetOLMOperatorVersionDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOLMOperatorVersionDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/olm/{repoName}/{packageName}/{version}][%d] getOLMOperatorVersionDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetOLMOperatorVersionDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /packages/olm/{repoName}/{packageName}/{version}][%d] getOLMOperatorVersionDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetOLMOperatorVersionDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOLMOperatorVersionDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOLMOperatorVersionDetailsTooManyRequests creates a GetOLMOperatorVersionDetailsTooManyRequests with default headers values
func NewGetOLMOperatorVersionDetailsTooManyRequests() *GetOLMOperatorVersionDetailsTooManyRequests {
	return &GetOLMOperatorVersionDetailsTooManyRequests{}
}

/*
GetOLMOperatorVersionDetailsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetOLMOperatorVersionDetailsTooManyRequests struct {
}

// IsSuccess returns true when this get o l m operator version details too many requests response has a 2xx status code
func (o *GetOLMOperatorVersionDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get o l m operator version details too many requests response has a 3xx status code
func (o *GetOLMOperatorVersionDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get o l m operator version details too many requests response has a 4xx status code
func (o *GetOLMOperatorVersionDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get o l m operator version details too many requests response has a 5xx status code
func (o *GetOLMOperatorVersionDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get o l m operator version details too many requests response a status code equal to that given
func (o *GetOLMOperatorVersionDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOLMOperatorVersionDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/olm/{repoName}/{packageName}/{version}][%d] getOLMOperatorVersionDetailsTooManyRequests ", 429)
}

func (o *GetOLMOperatorVersionDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/olm/{repoName}/{packageName}/{version}][%d] getOLMOperatorVersionDetailsTooManyRequests ", 429)
}

func (o *GetOLMOperatorVersionDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOLMOperatorVersionDetailsInternalServerError creates a GetOLMOperatorVersionDetailsInternalServerError with default headers values
func NewGetOLMOperatorVersionDetailsInternalServerError() *GetOLMOperatorVersionDetailsInternalServerError {
	return &GetOLMOperatorVersionDetailsInternalServerError{}
}

/*
GetOLMOperatorVersionDetailsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetOLMOperatorVersionDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get o l m operator version details internal server error response has a 2xx status code
func (o *GetOLMOperatorVersionDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get o l m operator version details internal server error response has a 3xx status code
func (o *GetOLMOperatorVersionDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get o l m operator version details internal server error response has a 4xx status code
func (o *GetOLMOperatorVersionDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get o l m operator version details internal server error response has a 5xx status code
func (o *GetOLMOperatorVersionDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get o l m operator version details internal server error response a status code equal to that given
func (o *GetOLMOperatorVersionDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOLMOperatorVersionDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/olm/{repoName}/{packageName}/{version}][%d] getOLMOperatorVersionDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOLMOperatorVersionDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/olm/{repoName}/{packageName}/{version}][%d] getOLMOperatorVersionDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOLMOperatorVersionDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOLMOperatorVersionDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
