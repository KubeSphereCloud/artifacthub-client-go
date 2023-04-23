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

// GetHelmPackageVersionDetailsReader is a Reader for the GetHelmPackageVersionDetails structure.
type GetHelmPackageVersionDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHelmPackageVersionDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHelmPackageVersionDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHelmPackageVersionDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetHelmPackageVersionDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHelmPackageVersionDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHelmPackageVersionDetailsOK creates a GetHelmPackageVersionDetailsOK with default headers values
func NewGetHelmPackageVersionDetailsOK() *GetHelmPackageVersionDetailsOK {
	return &GetHelmPackageVersionDetailsOK{}
}

/*
GetHelmPackageVersionDetailsOK describes a response with status code 200, with default header values.

GetHelmPackageVersionDetailsOK get helm package version details o k
*/
type GetHelmPackageVersionDetailsOK struct {
	Payload *models.HelmPackage
}

// IsSuccess returns true when this get helm package version details o k response has a 2xx status code
func (o *GetHelmPackageVersionDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get helm package version details o k response has a 3xx status code
func (o *GetHelmPackageVersionDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm package version details o k response has a 4xx status code
func (o *GetHelmPackageVersionDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get helm package version details o k response has a 5xx status code
func (o *GetHelmPackageVersionDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm package version details o k response a status code equal to that given
func (o *GetHelmPackageVersionDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHelmPackageVersionDetailsOK) Error() string {
	return fmt.Sprintf("[GET /packages/helm/{repoName}/{packageName}/{version}][%d] getHelmPackageVersionDetailsOK  %+v", 200, o.Payload)
}

func (o *GetHelmPackageVersionDetailsOK) String() string {
	return fmt.Sprintf("[GET /packages/helm/{repoName}/{packageName}/{version}][%d] getHelmPackageVersionDetailsOK  %+v", 200, o.Payload)
}

func (o *GetHelmPackageVersionDetailsOK) GetPayload() *models.HelmPackage {
	return o.Payload
}

func (o *GetHelmPackageVersionDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HelmPackage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHelmPackageVersionDetailsNotFound creates a GetHelmPackageVersionDetailsNotFound with default headers values
func NewGetHelmPackageVersionDetailsNotFound() *GetHelmPackageVersionDetailsNotFound {
	return &GetHelmPackageVersionDetailsNotFound{}
}

/*
GetHelmPackageVersionDetailsNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GetHelmPackageVersionDetailsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get helm package version details not found response has a 2xx status code
func (o *GetHelmPackageVersionDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm package version details not found response has a 3xx status code
func (o *GetHelmPackageVersionDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm package version details not found response has a 4xx status code
func (o *GetHelmPackageVersionDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get helm package version details not found response has a 5xx status code
func (o *GetHelmPackageVersionDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm package version details not found response a status code equal to that given
func (o *GetHelmPackageVersionDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetHelmPackageVersionDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/helm/{repoName}/{packageName}/{version}][%d] getHelmPackageVersionDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetHelmPackageVersionDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /packages/helm/{repoName}/{packageName}/{version}][%d] getHelmPackageVersionDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetHelmPackageVersionDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHelmPackageVersionDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHelmPackageVersionDetailsTooManyRequests creates a GetHelmPackageVersionDetailsTooManyRequests with default headers values
func NewGetHelmPackageVersionDetailsTooManyRequests() *GetHelmPackageVersionDetailsTooManyRequests {
	return &GetHelmPackageVersionDetailsTooManyRequests{}
}

/*
GetHelmPackageVersionDetailsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetHelmPackageVersionDetailsTooManyRequests struct {
}

// IsSuccess returns true when this get helm package version details too many requests response has a 2xx status code
func (o *GetHelmPackageVersionDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm package version details too many requests response has a 3xx status code
func (o *GetHelmPackageVersionDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm package version details too many requests response has a 4xx status code
func (o *GetHelmPackageVersionDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get helm package version details too many requests response has a 5xx status code
func (o *GetHelmPackageVersionDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm package version details too many requests response a status code equal to that given
func (o *GetHelmPackageVersionDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetHelmPackageVersionDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/helm/{repoName}/{packageName}/{version}][%d] getHelmPackageVersionDetailsTooManyRequests ", 429)
}

func (o *GetHelmPackageVersionDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/helm/{repoName}/{packageName}/{version}][%d] getHelmPackageVersionDetailsTooManyRequests ", 429)
}

func (o *GetHelmPackageVersionDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHelmPackageVersionDetailsInternalServerError creates a GetHelmPackageVersionDetailsInternalServerError with default headers values
func NewGetHelmPackageVersionDetailsInternalServerError() *GetHelmPackageVersionDetailsInternalServerError {
	return &GetHelmPackageVersionDetailsInternalServerError{}
}

/*
GetHelmPackageVersionDetailsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetHelmPackageVersionDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get helm package version details internal server error response has a 2xx status code
func (o *GetHelmPackageVersionDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm package version details internal server error response has a 3xx status code
func (o *GetHelmPackageVersionDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm package version details internal server error response has a 4xx status code
func (o *GetHelmPackageVersionDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get helm package version details internal server error response has a 5xx status code
func (o *GetHelmPackageVersionDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get helm package version details internal server error response a status code equal to that given
func (o *GetHelmPackageVersionDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetHelmPackageVersionDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/helm/{repoName}/{packageName}/{version}][%d] getHelmPackageVersionDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHelmPackageVersionDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/helm/{repoName}/{packageName}/{version}][%d] getHelmPackageVersionDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHelmPackageVersionDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHelmPackageVersionDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
