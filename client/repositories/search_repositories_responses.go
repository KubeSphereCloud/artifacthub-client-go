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

// SearchRepositoriesReader is a Reader for the SearchRepositories structure.
type SearchRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchRepositoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSearchRepositoriesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchRepositoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchRepositoriesOK creates a SearchRepositoriesOK with default headers values
func NewSearchRepositoriesOK() *SearchRepositoriesOK {
	return &SearchRepositoriesOK{}
}

/*
SearchRepositoriesOK describes a response with status code 200, with default header values.

SearchRepositoriesOK search repositories o k
*/
type SearchRepositoriesOK struct {

	/* Total number of repositories
	 */
	PaginationTotalCount string

	Payload []*models.Repository
}

// IsSuccess returns true when this search repositories o k response has a 2xx status code
func (o *SearchRepositoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search repositories o k response has a 3xx status code
func (o *SearchRepositoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search repositories o k response has a 4xx status code
func (o *SearchRepositoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search repositories o k response has a 5xx status code
func (o *SearchRepositoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search repositories o k response a status code equal to that given
func (o *SearchRepositoriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchRepositoriesOK) Error() string {
	return fmt.Sprintf("[GET /repositories/search][%d] searchRepositoriesOK  %+v", 200, o.Payload)
}

func (o *SearchRepositoriesOK) String() string {
	return fmt.Sprintf("[GET /repositories/search][%d] searchRepositoriesOK  %+v", 200, o.Payload)
}

func (o *SearchRepositoriesOK) GetPayload() []*models.Repository {
	return o.Payload
}

func (o *SearchRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Pagination-Total-Count
	hdrPaginationTotalCount := response.GetHeader("Pagination-Total-Count")

	if hdrPaginationTotalCount != "" {
		o.PaginationTotalCount = hdrPaginationTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRepositoriesUnauthorized creates a SearchRepositoriesUnauthorized with default headers values
func NewSearchRepositoriesUnauthorized() *SearchRepositoriesUnauthorized {
	return &SearchRepositoriesUnauthorized{}
}

/*
SearchRepositoriesUnauthorized describes a response with status code 401, with default header values.

Valid authentication credentials not provided
*/
type SearchRepositoriesUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this search repositories unauthorized response has a 2xx status code
func (o *SearchRepositoriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search repositories unauthorized response has a 3xx status code
func (o *SearchRepositoriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search repositories unauthorized response has a 4xx status code
func (o *SearchRepositoriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search repositories unauthorized response has a 5xx status code
func (o *SearchRepositoriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search repositories unauthorized response a status code equal to that given
func (o *SearchRepositoriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchRepositoriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/search][%d] searchRepositoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchRepositoriesUnauthorized) String() string {
	return fmt.Sprintf("[GET /repositories/search][%d] searchRepositoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchRepositoriesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchRepositoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRepositoriesTooManyRequests creates a SearchRepositoriesTooManyRequests with default headers values
func NewSearchRepositoriesTooManyRequests() *SearchRepositoriesTooManyRequests {
	return &SearchRepositoriesTooManyRequests{}
}

/*
SearchRepositoriesTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type SearchRepositoriesTooManyRequests struct {
}

// IsSuccess returns true when this search repositories too many requests response has a 2xx status code
func (o *SearchRepositoriesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search repositories too many requests response has a 3xx status code
func (o *SearchRepositoriesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search repositories too many requests response has a 4xx status code
func (o *SearchRepositoriesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this search repositories too many requests response has a 5xx status code
func (o *SearchRepositoriesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this search repositories too many requests response a status code equal to that given
func (o *SearchRepositoriesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *SearchRepositoriesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /repositories/search][%d] searchRepositoriesTooManyRequests ", 429)
}

func (o *SearchRepositoriesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /repositories/search][%d] searchRepositoriesTooManyRequests ", 429)
}

func (o *SearchRepositoriesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchRepositoriesInternalServerError creates a SearchRepositoriesInternalServerError with default headers values
func NewSearchRepositoriesInternalServerError() *SearchRepositoriesInternalServerError {
	return &SearchRepositoriesInternalServerError{}
}

/*
SearchRepositoriesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type SearchRepositoriesInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this search repositories internal server error response has a 2xx status code
func (o *SearchRepositoriesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search repositories internal server error response has a 3xx status code
func (o *SearchRepositoriesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search repositories internal server error response has a 4xx status code
func (o *SearchRepositoriesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search repositories internal server error response has a 5xx status code
func (o *SearchRepositoriesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search repositories internal server error response a status code equal to that given
func (o *SearchRepositoriesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchRepositoriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /repositories/search][%d] searchRepositoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchRepositoriesInternalServerError) String() string {
	return fmt.Sprintf("[GET /repositories/search][%d] searchRepositoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchRepositoriesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchRepositoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
