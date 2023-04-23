// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// AddOrganizationReader is a Reader for the AddOrganization structure.
type AddOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddOrganizationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddOrganizationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddOrganizationCreated creates a AddOrganizationCreated with default headers values
func NewAddOrganizationCreated() *AddOrganizationCreated {
	return &AddOrganizationCreated{}
}

/*
AddOrganizationCreated describes a response with status code 201, with default header values.

The request has succeeded and has led to the creation of a resource
*/
type AddOrganizationCreated struct {
}

// IsSuccess returns true when this add organization created response has a 2xx status code
func (o *AddOrganizationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add organization created response has a 3xx status code
func (o *AddOrganizationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization created response has a 4xx status code
func (o *AddOrganizationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add organization created response has a 5xx status code
func (o *AddOrganizationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add organization created response a status code equal to that given
func (o *AddOrganizationCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddOrganizationCreated) Error() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationCreated ", 201)
}

func (o *AddOrganizationCreated) String() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationCreated ", 201)
}

func (o *AddOrganizationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddOrganizationBadRequest creates a AddOrganizationBadRequest with default headers values
func NewAddOrganizationBadRequest() *AddOrganizationBadRequest {
	return &AddOrganizationBadRequest{}
}

/*
AddOrganizationBadRequest describes a response with status code 400, with default header values.

The request sent was not valid
*/
type AddOrganizationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this add organization bad request response has a 2xx status code
func (o *AddOrganizationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add organization bad request response has a 3xx status code
func (o *AddOrganizationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization bad request response has a 4xx status code
func (o *AddOrganizationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add organization bad request response has a 5xx status code
func (o *AddOrganizationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add organization bad request response a status code equal to that given
func (o *AddOrganizationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *AddOrganizationBadRequest) String() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *AddOrganizationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrganizationUnauthorized creates a AddOrganizationUnauthorized with default headers values
func NewAddOrganizationUnauthorized() *AddOrganizationUnauthorized {
	return &AddOrganizationUnauthorized{}
}

/*
AddOrganizationUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type AddOrganizationUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this add organization unauthorized response has a 2xx status code
func (o *AddOrganizationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add organization unauthorized response has a 3xx status code
func (o *AddOrganizationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization unauthorized response has a 4xx status code
func (o *AddOrganizationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add organization unauthorized response has a 5xx status code
func (o *AddOrganizationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add organization unauthorized response a status code equal to that given
func (o *AddOrganizationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AddOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *AddOrganizationUnauthorized) String() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *AddOrganizationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrganizationTooManyRequests creates a AddOrganizationTooManyRequests with default headers values
func NewAddOrganizationTooManyRequests() *AddOrganizationTooManyRequests {
	return &AddOrganizationTooManyRequests{}
}

/*
AddOrganizationTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type AddOrganizationTooManyRequests struct {
}

// IsSuccess returns true when this add organization too many requests response has a 2xx status code
func (o *AddOrganizationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add organization too many requests response has a 3xx status code
func (o *AddOrganizationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization too many requests response has a 4xx status code
func (o *AddOrganizationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add organization too many requests response has a 5xx status code
func (o *AddOrganizationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add organization too many requests response a status code equal to that given
func (o *AddOrganizationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AddOrganizationTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationTooManyRequests ", 429)
}

func (o *AddOrganizationTooManyRequests) String() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationTooManyRequests ", 429)
}

func (o *AddOrganizationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddOrganizationInternalServerError creates a AddOrganizationInternalServerError with default headers values
func NewAddOrganizationInternalServerError() *AddOrganizationInternalServerError {
	return &AddOrganizationInternalServerError{}
}

/*
AddOrganizationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type AddOrganizationInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this add organization internal server error response has a 2xx status code
func (o *AddOrganizationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add organization internal server error response has a 3xx status code
func (o *AddOrganizationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add organization internal server error response has a 4xx status code
func (o *AddOrganizationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add organization internal server error response has a 5xx status code
func (o *AddOrganizationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add organization internal server error response a status code equal to that given
func (o *AddOrganizationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *AddOrganizationInternalServerError) String() string {
	return fmt.Sprintf("[POST /orgs][%d] addOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *AddOrganizationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}