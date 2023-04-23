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

// GetBackstagePluginsVersionDetailsReader is a Reader for the GetBackstagePluginsVersionDetails structure.
type GetBackstagePluginsVersionDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackstagePluginsVersionDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackstagePluginsVersionDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBackstagePluginsVersionDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBackstagePluginsVersionDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBackstagePluginsVersionDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackstagePluginsVersionDetailsOK creates a GetBackstagePluginsVersionDetailsOK with default headers values
func NewGetBackstagePluginsVersionDetailsOK() *GetBackstagePluginsVersionDetailsOK {
	return &GetBackstagePluginsVersionDetailsOK{}
}

/*
GetBackstagePluginsVersionDetailsOK describes a response with status code 200, with default header values.

GetBackstagePluginsVersionDetailsOK get backstage plugins version details o k
*/
type GetBackstagePluginsVersionDetailsOK struct {
	Payload *models.BackstagePlugin
}

// IsSuccess returns true when this get backstage plugins version details o k response has a 2xx status code
func (o *GetBackstagePluginsVersionDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get backstage plugins version details o k response has a 3xx status code
func (o *GetBackstagePluginsVersionDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backstage plugins version details o k response has a 4xx status code
func (o *GetBackstagePluginsVersionDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backstage plugins version details o k response has a 5xx status code
func (o *GetBackstagePluginsVersionDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get backstage plugins version details o k response a status code equal to that given
func (o *GetBackstagePluginsVersionDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBackstagePluginsVersionDetailsOK) Error() string {
	return fmt.Sprintf("[GET /packages/backstage/{repoName}/{packageName}/{version}][%d] getBackstagePluginsVersionDetailsOK  %+v", 200, o.Payload)
}

func (o *GetBackstagePluginsVersionDetailsOK) String() string {
	return fmt.Sprintf("[GET /packages/backstage/{repoName}/{packageName}/{version}][%d] getBackstagePluginsVersionDetailsOK  %+v", 200, o.Payload)
}

func (o *GetBackstagePluginsVersionDetailsOK) GetPayload() *models.BackstagePlugin {
	return o.Payload
}

func (o *GetBackstagePluginsVersionDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackstagePlugin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackstagePluginsVersionDetailsNotFound creates a GetBackstagePluginsVersionDetailsNotFound with default headers values
func NewGetBackstagePluginsVersionDetailsNotFound() *GetBackstagePluginsVersionDetailsNotFound {
	return &GetBackstagePluginsVersionDetailsNotFound{}
}

/*
GetBackstagePluginsVersionDetailsNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GetBackstagePluginsVersionDetailsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get backstage plugins version details not found response has a 2xx status code
func (o *GetBackstagePluginsVersionDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backstage plugins version details not found response has a 3xx status code
func (o *GetBackstagePluginsVersionDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backstage plugins version details not found response has a 4xx status code
func (o *GetBackstagePluginsVersionDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backstage plugins version details not found response has a 5xx status code
func (o *GetBackstagePluginsVersionDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get backstage plugins version details not found response a status code equal to that given
func (o *GetBackstagePluginsVersionDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetBackstagePluginsVersionDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/backstage/{repoName}/{packageName}/{version}][%d] getBackstagePluginsVersionDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetBackstagePluginsVersionDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /packages/backstage/{repoName}/{packageName}/{version}][%d] getBackstagePluginsVersionDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetBackstagePluginsVersionDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackstagePluginsVersionDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackstagePluginsVersionDetailsTooManyRequests creates a GetBackstagePluginsVersionDetailsTooManyRequests with default headers values
func NewGetBackstagePluginsVersionDetailsTooManyRequests() *GetBackstagePluginsVersionDetailsTooManyRequests {
	return &GetBackstagePluginsVersionDetailsTooManyRequests{}
}

/*
GetBackstagePluginsVersionDetailsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetBackstagePluginsVersionDetailsTooManyRequests struct {
}

// IsSuccess returns true when this get backstage plugins version details too many requests response has a 2xx status code
func (o *GetBackstagePluginsVersionDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backstage plugins version details too many requests response has a 3xx status code
func (o *GetBackstagePluginsVersionDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backstage plugins version details too many requests response has a 4xx status code
func (o *GetBackstagePluginsVersionDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backstage plugins version details too many requests response has a 5xx status code
func (o *GetBackstagePluginsVersionDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get backstage plugins version details too many requests response a status code equal to that given
func (o *GetBackstagePluginsVersionDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetBackstagePluginsVersionDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/backstage/{repoName}/{packageName}/{version}][%d] getBackstagePluginsVersionDetailsTooManyRequests ", 429)
}

func (o *GetBackstagePluginsVersionDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/backstage/{repoName}/{packageName}/{version}][%d] getBackstagePluginsVersionDetailsTooManyRequests ", 429)
}

func (o *GetBackstagePluginsVersionDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBackstagePluginsVersionDetailsInternalServerError creates a GetBackstagePluginsVersionDetailsInternalServerError with default headers values
func NewGetBackstagePluginsVersionDetailsInternalServerError() *GetBackstagePluginsVersionDetailsInternalServerError {
	return &GetBackstagePluginsVersionDetailsInternalServerError{}
}

/*
GetBackstagePluginsVersionDetailsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetBackstagePluginsVersionDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get backstage plugins version details internal server error response has a 2xx status code
func (o *GetBackstagePluginsVersionDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backstage plugins version details internal server error response has a 3xx status code
func (o *GetBackstagePluginsVersionDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backstage plugins version details internal server error response has a 4xx status code
func (o *GetBackstagePluginsVersionDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backstage plugins version details internal server error response has a 5xx status code
func (o *GetBackstagePluginsVersionDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get backstage plugins version details internal server error response a status code equal to that given
func (o *GetBackstagePluginsVersionDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetBackstagePluginsVersionDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/backstage/{repoName}/{packageName}/{version}][%d] getBackstagePluginsVersionDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBackstagePluginsVersionDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/backstage/{repoName}/{packageName}/{version}][%d] getBackstagePluginsVersionDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBackstagePluginsVersionDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackstagePluginsVersionDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
