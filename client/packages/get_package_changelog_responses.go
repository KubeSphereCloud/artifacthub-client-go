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

// GetPackageChangelogReader is a Reader for the GetPackageChangelog structure.
type GetPackageChangelogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageChangelogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageChangelogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPackageChangelogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPackageChangelogTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPackageChangelogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageChangelogOK creates a GetPackageChangelogOK with default headers values
func NewGetPackageChangelogOK() *GetPackageChangelogOK {
	return &GetPackageChangelogOK{}
}

/*
GetPackageChangelogOK describes a response with status code 200, with default header values.

GetPackageChangelogOK get package changelog o k
*/
type GetPackageChangelogOK struct {
	Payload []*models.PackagesChangelogResponse
}

// IsSuccess returns true when this get package changelog o k response has a 2xx status code
func (o *GetPackageChangelogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get package changelog o k response has a 3xx status code
func (o *GetPackageChangelogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package changelog o k response has a 4xx status code
func (o *GetPackageChangelogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package changelog o k response has a 5xx status code
func (o *GetPackageChangelogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get package changelog o k response a status code equal to that given
func (o *GetPackageChangelogOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPackageChangelogOK) Error() string {
	return fmt.Sprintf("[GET /packages/{packageID}/changelog][%d] getPackageChangelogOK  %+v", 200, o.Payload)
}

func (o *GetPackageChangelogOK) String() string {
	return fmt.Sprintf("[GET /packages/{packageID}/changelog][%d] getPackageChangelogOK  %+v", 200, o.Payload)
}

func (o *GetPackageChangelogOK) GetPayload() []*models.PackagesChangelogResponse {
	return o.Payload
}

func (o *GetPackageChangelogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageChangelogNotFound creates a GetPackageChangelogNotFound with default headers values
func NewGetPackageChangelogNotFound() *GetPackageChangelogNotFound {
	return &GetPackageChangelogNotFound{}
}

/*
GetPackageChangelogNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GetPackageChangelogNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get package changelog not found response has a 2xx status code
func (o *GetPackageChangelogNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package changelog not found response has a 3xx status code
func (o *GetPackageChangelogNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package changelog not found response has a 4xx status code
func (o *GetPackageChangelogNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package changelog not found response has a 5xx status code
func (o *GetPackageChangelogNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get package changelog not found response a status code equal to that given
func (o *GetPackageChangelogNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPackageChangelogNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/{packageID}/changelog][%d] getPackageChangelogNotFound  %+v", 404, o.Payload)
}

func (o *GetPackageChangelogNotFound) String() string {
	return fmt.Sprintf("[GET /packages/{packageID}/changelog][%d] getPackageChangelogNotFound  %+v", 404, o.Payload)
}

func (o *GetPackageChangelogNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPackageChangelogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageChangelogTooManyRequests creates a GetPackageChangelogTooManyRequests with default headers values
func NewGetPackageChangelogTooManyRequests() *GetPackageChangelogTooManyRequests {
	return &GetPackageChangelogTooManyRequests{}
}

/*
GetPackageChangelogTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetPackageChangelogTooManyRequests struct {
}

// IsSuccess returns true when this get package changelog too many requests response has a 2xx status code
func (o *GetPackageChangelogTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package changelog too many requests response has a 3xx status code
func (o *GetPackageChangelogTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package changelog too many requests response has a 4xx status code
func (o *GetPackageChangelogTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package changelog too many requests response has a 5xx status code
func (o *GetPackageChangelogTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get package changelog too many requests response a status code equal to that given
func (o *GetPackageChangelogTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPackageChangelogTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/{packageID}/changelog][%d] getPackageChangelogTooManyRequests ", 429)
}

func (o *GetPackageChangelogTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/{packageID}/changelog][%d] getPackageChangelogTooManyRequests ", 429)
}

func (o *GetPackageChangelogTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPackageChangelogInternalServerError creates a GetPackageChangelogInternalServerError with default headers values
func NewGetPackageChangelogInternalServerError() *GetPackageChangelogInternalServerError {
	return &GetPackageChangelogInternalServerError{}
}

/*
GetPackageChangelogInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetPackageChangelogInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get package changelog internal server error response has a 2xx status code
func (o *GetPackageChangelogInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package changelog internal server error response has a 3xx status code
func (o *GetPackageChangelogInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package changelog internal server error response has a 4xx status code
func (o *GetPackageChangelogInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package changelog internal server error response has a 5xx status code
func (o *GetPackageChangelogInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get package changelog internal server error response a status code equal to that given
func (o *GetPackageChangelogInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPackageChangelogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/{packageID}/changelog][%d] getPackageChangelogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackageChangelogInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/{packageID}/changelog][%d] getPackageChangelogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackageChangelogInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPackageChangelogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
