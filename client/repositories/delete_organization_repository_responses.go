// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/artifacthub-client-go/models"
)

// DeleteOrganizationRepositoryReader is a Reader for the DeleteOrganizationRepository structure.
type DeleteOrganizationRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationRepositoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteOrganizationRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOrganizationRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOrganizationRepositoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOrganizationRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationRepositoryNoContent creates a DeleteOrganizationRepositoryNoContent with default headers values
func NewDeleteOrganizationRepositoryNoContent() *DeleteOrganizationRepositoryNoContent {
	return &DeleteOrganizationRepositoryNoContent{}
}

/*
DeleteOrganizationRepositoryNoContent describes a response with status code 204, with default header values.

The request has succeeded, no content returned
*/
type DeleteOrganizationRepositoryNoContent struct {
}

// IsSuccess returns true when this delete organization repository no content response has a 2xx status code
func (o *DeleteOrganizationRepositoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization repository no content response has a 3xx status code
func (o *DeleteOrganizationRepositoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization repository no content response has a 4xx status code
func (o *DeleteOrganizationRepositoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization repository no content response has a 5xx status code
func (o *DeleteOrganizationRepositoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization repository no content response a status code equal to that given
func (o *DeleteOrganizationRepositoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteOrganizationRepositoryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryNoContent ", 204)
}

func (o *DeleteOrganizationRepositoryNoContent) String() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryNoContent ", 204)
}

func (o *DeleteOrganizationRepositoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationRepositoryUnauthorized creates a DeleteOrganizationRepositoryUnauthorized with default headers values
func NewDeleteOrganizationRepositoryUnauthorized() *DeleteOrganizationRepositoryUnauthorized {
	return &DeleteOrganizationRepositoryUnauthorized{}
}

/*
DeleteOrganizationRepositoryUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type DeleteOrganizationRepositoryUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete organization repository unauthorized response has a 2xx status code
func (o *DeleteOrganizationRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organization repository unauthorized response has a 3xx status code
func (o *DeleteOrganizationRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization repository unauthorized response has a 4xx status code
func (o *DeleteOrganizationRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organization repository unauthorized response has a 5xx status code
func (o *DeleteOrganizationRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization repository unauthorized response a status code equal to that given
func (o *DeleteOrganizationRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteOrganizationRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrganizationRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrganizationRepositoryUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteOrganizationRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationRepositoryForbidden creates a DeleteOrganizationRepositoryForbidden with default headers values
func NewDeleteOrganizationRepositoryForbidden() *DeleteOrganizationRepositoryForbidden {
	return &DeleteOrganizationRepositoryForbidden{}
}

/*
DeleteOrganizationRepositoryForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the requested operation
*/
type DeleteOrganizationRepositoryForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete organization repository forbidden response has a 2xx status code
func (o *DeleteOrganizationRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organization repository forbidden response has a 3xx status code
func (o *DeleteOrganizationRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization repository forbidden response has a 4xx status code
func (o *DeleteOrganizationRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organization repository forbidden response has a 5xx status code
func (o *DeleteOrganizationRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization repository forbidden response a status code equal to that given
func (o *DeleteOrganizationRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteOrganizationRepositoryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrganizationRepositoryForbidden) String() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrganizationRepositoryForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteOrganizationRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationRepositoryTooManyRequests creates a DeleteOrganizationRepositoryTooManyRequests with default headers values
func NewDeleteOrganizationRepositoryTooManyRequests() *DeleteOrganizationRepositoryTooManyRequests {
	return &DeleteOrganizationRepositoryTooManyRequests{}
}

/*
DeleteOrganizationRepositoryTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type DeleteOrganizationRepositoryTooManyRequests struct {
}

// IsSuccess returns true when this delete organization repository too many requests response has a 2xx status code
func (o *DeleteOrganizationRepositoryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organization repository too many requests response has a 3xx status code
func (o *DeleteOrganizationRepositoryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization repository too many requests response has a 4xx status code
func (o *DeleteOrganizationRepositoryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organization repository too many requests response has a 5xx status code
func (o *DeleteOrganizationRepositoryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization repository too many requests response a status code equal to that given
func (o *DeleteOrganizationRepositoryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteOrganizationRepositoryTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryTooManyRequests ", 429)
}

func (o *DeleteOrganizationRepositoryTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryTooManyRequests ", 429)
}

func (o *DeleteOrganizationRepositoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationRepositoryInternalServerError creates a DeleteOrganizationRepositoryInternalServerError with default headers values
func NewDeleteOrganizationRepositoryInternalServerError() *DeleteOrganizationRepositoryInternalServerError {
	return &DeleteOrganizationRepositoryInternalServerError{}
}

/*
DeleteOrganizationRepositoryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type DeleteOrganizationRepositoryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete organization repository internal server error response has a 2xx status code
func (o *DeleteOrganizationRepositoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organization repository internal server error response has a 3xx status code
func (o *DeleteOrganizationRepositoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization repository internal server error response has a 4xx status code
func (o *DeleteOrganizationRepositoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization repository internal server error response has a 5xx status code
func (o *DeleteOrganizationRepositoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete organization repository internal server error response a status code equal to that given
func (o *DeleteOrganizationRepositoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteOrganizationRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrganizationRepositoryInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /repositories/org/{orgName}/{repoName}][%d] deleteOrganizationRepositoryInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrganizationRepositoryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteOrganizationRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
