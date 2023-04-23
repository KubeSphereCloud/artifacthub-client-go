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

// GetHelmPluginVersionDetailsReader is a Reader for the GetHelmPluginVersionDetails structure.
type GetHelmPluginVersionDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHelmPluginVersionDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHelmPluginVersionDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHelmPluginVersionDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetHelmPluginVersionDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHelmPluginVersionDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHelmPluginVersionDetailsOK creates a GetHelmPluginVersionDetailsOK with default headers values
func NewGetHelmPluginVersionDetailsOK() *GetHelmPluginVersionDetailsOK {
	return &GetHelmPluginVersionDetailsOK{}
}

/*
GetHelmPluginVersionDetailsOK describes a response with status code 200, with default header values.

GetHelmPluginVersionDetailsOK get helm plugin version details o k
*/
type GetHelmPluginVersionDetailsOK struct {
	Payload *models.HelmPluginPackage
}

// IsSuccess returns true when this get helm plugin version details o k response has a 2xx status code
func (o *GetHelmPluginVersionDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get helm plugin version details o k response has a 3xx status code
func (o *GetHelmPluginVersionDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm plugin version details o k response has a 4xx status code
func (o *GetHelmPluginVersionDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get helm plugin version details o k response has a 5xx status code
func (o *GetHelmPluginVersionDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm plugin version details o k response a status code equal to that given
func (o *GetHelmPluginVersionDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHelmPluginVersionDetailsOK) Error() string {
	return fmt.Sprintf("[GET /packages/helm-plugin/{repoName}/{packageName}/{version}][%d] getHelmPluginVersionDetailsOK  %+v", 200, o.Payload)
}

func (o *GetHelmPluginVersionDetailsOK) String() string {
	return fmt.Sprintf("[GET /packages/helm-plugin/{repoName}/{packageName}/{version}][%d] getHelmPluginVersionDetailsOK  %+v", 200, o.Payload)
}

func (o *GetHelmPluginVersionDetailsOK) GetPayload() *models.HelmPluginPackage {
	return o.Payload
}

func (o *GetHelmPluginVersionDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HelmPluginPackage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHelmPluginVersionDetailsNotFound creates a GetHelmPluginVersionDetailsNotFound with default headers values
func NewGetHelmPluginVersionDetailsNotFound() *GetHelmPluginVersionDetailsNotFound {
	return &GetHelmPluginVersionDetailsNotFound{}
}

/*
GetHelmPluginVersionDetailsNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GetHelmPluginVersionDetailsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get helm plugin version details not found response has a 2xx status code
func (o *GetHelmPluginVersionDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm plugin version details not found response has a 3xx status code
func (o *GetHelmPluginVersionDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm plugin version details not found response has a 4xx status code
func (o *GetHelmPluginVersionDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get helm plugin version details not found response has a 5xx status code
func (o *GetHelmPluginVersionDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm plugin version details not found response a status code equal to that given
func (o *GetHelmPluginVersionDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetHelmPluginVersionDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/helm-plugin/{repoName}/{packageName}/{version}][%d] getHelmPluginVersionDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetHelmPluginVersionDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /packages/helm-plugin/{repoName}/{packageName}/{version}][%d] getHelmPluginVersionDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetHelmPluginVersionDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHelmPluginVersionDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHelmPluginVersionDetailsTooManyRequests creates a GetHelmPluginVersionDetailsTooManyRequests with default headers values
func NewGetHelmPluginVersionDetailsTooManyRequests() *GetHelmPluginVersionDetailsTooManyRequests {
	return &GetHelmPluginVersionDetailsTooManyRequests{}
}

/*
GetHelmPluginVersionDetailsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetHelmPluginVersionDetailsTooManyRequests struct {
}

// IsSuccess returns true when this get helm plugin version details too many requests response has a 2xx status code
func (o *GetHelmPluginVersionDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm plugin version details too many requests response has a 3xx status code
func (o *GetHelmPluginVersionDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm plugin version details too many requests response has a 4xx status code
func (o *GetHelmPluginVersionDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get helm plugin version details too many requests response has a 5xx status code
func (o *GetHelmPluginVersionDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm plugin version details too many requests response a status code equal to that given
func (o *GetHelmPluginVersionDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetHelmPluginVersionDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/helm-plugin/{repoName}/{packageName}/{version}][%d] getHelmPluginVersionDetailsTooManyRequests ", 429)
}

func (o *GetHelmPluginVersionDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/helm-plugin/{repoName}/{packageName}/{version}][%d] getHelmPluginVersionDetailsTooManyRequests ", 429)
}

func (o *GetHelmPluginVersionDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHelmPluginVersionDetailsInternalServerError creates a GetHelmPluginVersionDetailsInternalServerError with default headers values
func NewGetHelmPluginVersionDetailsInternalServerError() *GetHelmPluginVersionDetailsInternalServerError {
	return &GetHelmPluginVersionDetailsInternalServerError{}
}

/*
GetHelmPluginVersionDetailsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetHelmPluginVersionDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get helm plugin version details internal server error response has a 2xx status code
func (o *GetHelmPluginVersionDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm plugin version details internal server error response has a 3xx status code
func (o *GetHelmPluginVersionDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm plugin version details internal server error response has a 4xx status code
func (o *GetHelmPluginVersionDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get helm plugin version details internal server error response has a 5xx status code
func (o *GetHelmPluginVersionDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get helm plugin version details internal server error response a status code equal to that given
func (o *GetHelmPluginVersionDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetHelmPluginVersionDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/helm-plugin/{repoName}/{packageName}/{version}][%d] getHelmPluginVersionDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHelmPluginVersionDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/helm-plugin/{repoName}/{packageName}/{version}][%d] getHelmPluginVersionDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHelmPluginVersionDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHelmPluginVersionDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
