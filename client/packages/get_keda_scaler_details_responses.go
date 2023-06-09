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

// GetKedaScalerDetailsReader is a Reader for the GetKedaScalerDetails structure.
type GetKedaScalerDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKedaScalerDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKedaScalerDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetKedaScalerDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKedaScalerDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKedaScalerDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKedaScalerDetailsOK creates a GetKedaScalerDetailsOK with default headers values
func NewGetKedaScalerDetailsOK() *GetKedaScalerDetailsOK {
	return &GetKedaScalerDetailsOK{}
}

/*
GetKedaScalerDetailsOK describes a response with status code 200, with default header values.

GetKedaScalerDetailsOK get keda scaler details o k
*/
type GetKedaScalerDetailsOK struct {
	Payload *models.KedaScalerPackage
}

// IsSuccess returns true when this get keda scaler details o k response has a 2xx status code
func (o *GetKedaScalerDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get keda scaler details o k response has a 3xx status code
func (o *GetKedaScalerDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get keda scaler details o k response has a 4xx status code
func (o *GetKedaScalerDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get keda scaler details o k response has a 5xx status code
func (o *GetKedaScalerDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get keda scaler details o k response a status code equal to that given
func (o *GetKedaScalerDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKedaScalerDetailsOK) Error() string {
	return fmt.Sprintf("[GET /packages/keda-scaler/{repoName}/{packageName}][%d] getKedaScalerDetailsOK  %+v", 200, o.Payload)
}

func (o *GetKedaScalerDetailsOK) String() string {
	return fmt.Sprintf("[GET /packages/keda-scaler/{repoName}/{packageName}][%d] getKedaScalerDetailsOK  %+v", 200, o.Payload)
}

func (o *GetKedaScalerDetailsOK) GetPayload() *models.KedaScalerPackage {
	return o.Payload
}

func (o *GetKedaScalerDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KedaScalerPackage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKedaScalerDetailsNotFound creates a GetKedaScalerDetailsNotFound with default headers values
func NewGetKedaScalerDetailsNotFound() *GetKedaScalerDetailsNotFound {
	return &GetKedaScalerDetailsNotFound{}
}

/*
GetKedaScalerDetailsNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GetKedaScalerDetailsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get keda scaler details not found response has a 2xx status code
func (o *GetKedaScalerDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get keda scaler details not found response has a 3xx status code
func (o *GetKedaScalerDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get keda scaler details not found response has a 4xx status code
func (o *GetKedaScalerDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get keda scaler details not found response has a 5xx status code
func (o *GetKedaScalerDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get keda scaler details not found response a status code equal to that given
func (o *GetKedaScalerDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKedaScalerDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/keda-scaler/{repoName}/{packageName}][%d] getKedaScalerDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetKedaScalerDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /packages/keda-scaler/{repoName}/{packageName}][%d] getKedaScalerDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetKedaScalerDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetKedaScalerDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKedaScalerDetailsTooManyRequests creates a GetKedaScalerDetailsTooManyRequests with default headers values
func NewGetKedaScalerDetailsTooManyRequests() *GetKedaScalerDetailsTooManyRequests {
	return &GetKedaScalerDetailsTooManyRequests{}
}

/*
GetKedaScalerDetailsTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GetKedaScalerDetailsTooManyRequests struct {
}

// IsSuccess returns true when this get keda scaler details too many requests response has a 2xx status code
func (o *GetKedaScalerDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get keda scaler details too many requests response has a 3xx status code
func (o *GetKedaScalerDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get keda scaler details too many requests response has a 4xx status code
func (o *GetKedaScalerDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get keda scaler details too many requests response has a 5xx status code
func (o *GetKedaScalerDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get keda scaler details too many requests response a status code equal to that given
func (o *GetKedaScalerDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKedaScalerDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/keda-scaler/{repoName}/{packageName}][%d] getKedaScalerDetailsTooManyRequests ", 429)
}

func (o *GetKedaScalerDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/keda-scaler/{repoName}/{packageName}][%d] getKedaScalerDetailsTooManyRequests ", 429)
}

func (o *GetKedaScalerDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKedaScalerDetailsInternalServerError creates a GetKedaScalerDetailsInternalServerError with default headers values
func NewGetKedaScalerDetailsInternalServerError() *GetKedaScalerDetailsInternalServerError {
	return &GetKedaScalerDetailsInternalServerError{}
}

/*
GetKedaScalerDetailsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GetKedaScalerDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get keda scaler details internal server error response has a 2xx status code
func (o *GetKedaScalerDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get keda scaler details internal server error response has a 3xx status code
func (o *GetKedaScalerDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get keda scaler details internal server error response has a 4xx status code
func (o *GetKedaScalerDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get keda scaler details internal server error response has a 5xx status code
func (o *GetKedaScalerDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get keda scaler details internal server error response a status code equal to that given
func (o *GetKedaScalerDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKedaScalerDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/keda-scaler/{repoName}/{packageName}][%d] getKedaScalerDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKedaScalerDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/keda-scaler/{repoName}/{packageName}][%d] getKedaScalerDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKedaScalerDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetKedaScalerDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
