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

// GetPackageValuesSchemaReader is a Reader for the GetPackageValuesSchema structure.
type GetPackageValuesSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageValuesSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageValuesSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPackageValuesSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPackageValuesSchemaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPackageValuesSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageValuesSchemaOK creates a GetPackageValuesSchemaOK with default headers values
func NewGetPackageValuesSchemaOK() *GetPackageValuesSchemaOK {
	return &GetPackageValuesSchemaOK{}
}

/*
GetPackageValuesSchemaOK describes a response with status code 200, with default header values.

GetPackageValuesSchemaOK get package values schema o k
*/
type GetPackageValuesSchemaOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get package values schema o k response has a 2xx status code
func (o *GetPackageValuesSchemaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get package values schema o k response has a 3xx status code
func (o *GetPackageValuesSchemaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package values schema o k response has a 4xx status code
func (o *GetPackageValuesSchemaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package values schema o k response has a 5xx status code
func (o *GetPackageValuesSchemaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get package values schema o k response a status code equal to that given
func (o *GetPackageValuesSchemaOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPackageValuesSchemaOK) Error() string {
	return fmt.Sprintf("[GET /packages/{packageID}/{version}/values-schema][%d] getPackageValuesSchemaOK  %+v", 200, o.Payload)
}

func (o *GetPackageValuesSchemaOK) String() string {
	return fmt.Sprintf("[GET /packages/{packageID}/{version}/values-schema][%d] getPackageValuesSchemaOK  %+v", 200, o.Payload)
}

func (o *GetPackageValuesSchemaOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetPackageValuesSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageValuesSchemaNotFound creates a GetPackageValuesSchemaNotFound with default headers values
func NewGetPackageValuesSchemaNotFound() *GetPackageValuesSchemaNotFound {
	return &GetPackageValuesSchemaNotFound{}
}

/*
GetPackageValuesSchemaNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GetPackageValuesSchemaNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get package values schema not found response has a 2xx status code
func (o *GetPackageValuesSchemaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package values schema not found response has a 3xx status code
func (o *GetPackageValuesSchemaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package values schema not found response has a 4xx status code
func (o *GetPackageValuesSchemaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package values schema not found response has a 5xx status code
func (o *GetPackageValuesSchemaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get package values schema not found response a status code equal to that given
func (o *GetPackageValuesSchemaNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPackageValuesSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/{packageID}/{version}/values-schema][%d] getPackageValuesSchemaNotFound  %+v", 404, o.Payload)
}

func (o *GetPackageValuesSchemaNotFound) String() string {
	return fmt.Sprintf("[GET /packages/{packageID}/{version}/values-schema][%d] getPackageValuesSchemaNotFound  %+v", 404, o.Payload)
}

func (o *GetPackageValuesSchemaNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPackageValuesSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageValuesSchemaTooManyRequests creates a GetPackageValuesSchemaTooManyRequests with default headers values
func NewGetPackageValuesSchemaTooManyRequests() *GetPackageValuesSchemaTooManyRequests {
	return &GetPackageValuesSchemaTooManyRequests{}
}

/*
GetPackageValuesSchemaTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetPackageValuesSchemaTooManyRequests struct {
}

// IsSuccess returns true when this get package values schema too many requests response has a 2xx status code
func (o *GetPackageValuesSchemaTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package values schema too many requests response has a 3xx status code
func (o *GetPackageValuesSchemaTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package values schema too many requests response has a 4xx status code
func (o *GetPackageValuesSchemaTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package values schema too many requests response has a 5xx status code
func (o *GetPackageValuesSchemaTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get package values schema too many requests response a status code equal to that given
func (o *GetPackageValuesSchemaTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPackageValuesSchemaTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/{packageID}/{version}/values-schema][%d] getPackageValuesSchemaTooManyRequests ", 429)
}

func (o *GetPackageValuesSchemaTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/{packageID}/{version}/values-schema][%d] getPackageValuesSchemaTooManyRequests ", 429)
}

func (o *GetPackageValuesSchemaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPackageValuesSchemaInternalServerError creates a GetPackageValuesSchemaInternalServerError with default headers values
func NewGetPackageValuesSchemaInternalServerError() *GetPackageValuesSchemaInternalServerError {
	return &GetPackageValuesSchemaInternalServerError{}
}

/*
GetPackageValuesSchemaInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetPackageValuesSchemaInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get package values schema internal server error response has a 2xx status code
func (o *GetPackageValuesSchemaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package values schema internal server error response has a 3xx status code
func (o *GetPackageValuesSchemaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package values schema internal server error response has a 4xx status code
func (o *GetPackageValuesSchemaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package values schema internal server error response has a 5xx status code
func (o *GetPackageValuesSchemaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get package values schema internal server error response a status code equal to that given
func (o *GetPackageValuesSchemaInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPackageValuesSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/{packageID}/{version}/values-schema][%d] getPackageValuesSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackageValuesSchemaInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/{packageID}/{version}/values-schema][%d] getPackageValuesSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackageValuesSchemaInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPackageValuesSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
