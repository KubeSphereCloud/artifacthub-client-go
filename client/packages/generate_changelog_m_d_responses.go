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

// GenerateChangelogMDReader is a Reader for the GenerateChangelogMD structure.
type GenerateChangelogMDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateChangelogMDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateChangelogMDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGenerateChangelogMDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGenerateChangelogMDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenerateChangelogMDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateChangelogMDOK creates a GenerateChangelogMDOK with default headers values
func NewGenerateChangelogMDOK() *GenerateChangelogMDOK {
	return &GenerateChangelogMDOK{}
}

/*
GenerateChangelogMDOK describes a response with status code 200, with default header values.

GenerateChangelogMDOK generate changelog m d o k
*/
type GenerateChangelogMDOK struct {
	Payload string
}

// IsSuccess returns true when this generate changelog m d o k response has a 2xx status code
func (o *GenerateChangelogMDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate changelog m d o k response has a 3xx status code
func (o *GenerateChangelogMDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate changelog m d o k response has a 4xx status code
func (o *GenerateChangelogMDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate changelog m d o k response has a 5xx status code
func (o *GenerateChangelogMDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate changelog m d o k response a status code equal to that given
func (o *GenerateChangelogMDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GenerateChangelogMDOK) Error() string {
	return fmt.Sprintf("[GET /packages/{repoKindParam}/{repoName}/{packageName}/changelog.md][%d] generateChangelogMDOK  %+v", 200, o.Payload)
}

func (o *GenerateChangelogMDOK) String() string {
	return fmt.Sprintf("[GET /packages/{repoKindParam}/{repoName}/{packageName}/changelog.md][%d] generateChangelogMDOK  %+v", 200, o.Payload)
}

func (o *GenerateChangelogMDOK) GetPayload() string {
	return o.Payload
}

func (o *GenerateChangelogMDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateChangelogMDNotFound creates a GenerateChangelogMDNotFound with default headers values
func NewGenerateChangelogMDNotFound() *GenerateChangelogMDNotFound {
	return &GenerateChangelogMDNotFound{}
}

/*
GenerateChangelogMDNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type GenerateChangelogMDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this generate changelog m d not found response has a 2xx status code
func (o *GenerateChangelogMDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate changelog m d not found response has a 3xx status code
func (o *GenerateChangelogMDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate changelog m d not found response has a 4xx status code
func (o *GenerateChangelogMDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate changelog m d not found response has a 5xx status code
func (o *GenerateChangelogMDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this generate changelog m d not found response a status code equal to that given
func (o *GenerateChangelogMDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GenerateChangelogMDNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/{repoKindParam}/{repoName}/{packageName}/changelog.md][%d] generateChangelogMDNotFound  %+v", 404, o.Payload)
}

func (o *GenerateChangelogMDNotFound) String() string {
	return fmt.Sprintf("[GET /packages/{repoKindParam}/{repoName}/{packageName}/changelog.md][%d] generateChangelogMDNotFound  %+v", 404, o.Payload)
}

func (o *GenerateChangelogMDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GenerateChangelogMDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateChangelogMDTooManyRequests creates a GenerateChangelogMDTooManyRequests with default headers values
func NewGenerateChangelogMDTooManyRequests() *GenerateChangelogMDTooManyRequests {
	return &GenerateChangelogMDTooManyRequests{}
}

/*
GenerateChangelogMDTooManyRequests describes a response with status code 429, with default header values.

The user has sent too many requests in a given amount of time
*/
type GenerateChangelogMDTooManyRequests struct {
}

// IsSuccess returns true when this generate changelog m d too many requests response has a 2xx status code
func (o *GenerateChangelogMDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate changelog m d too many requests response has a 3xx status code
func (o *GenerateChangelogMDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate changelog m d too many requests response has a 4xx status code
func (o *GenerateChangelogMDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate changelog m d too many requests response has a 5xx status code
func (o *GenerateChangelogMDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this generate changelog m d too many requests response a status code equal to that given
func (o *GenerateChangelogMDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GenerateChangelogMDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /packages/{repoKindParam}/{repoName}/{packageName}/changelog.md][%d] generateChangelogMDTooManyRequests ", 429)
}

func (o *GenerateChangelogMDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /packages/{repoKindParam}/{repoName}/{packageName}/changelog.md][%d] generateChangelogMDTooManyRequests ", 429)
}

func (o *GenerateChangelogMDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenerateChangelogMDInternalServerError creates a GenerateChangelogMDInternalServerError with default headers values
func NewGenerateChangelogMDInternalServerError() *GenerateChangelogMDInternalServerError {
	return &GenerateChangelogMDInternalServerError{}
}

/*
GenerateChangelogMDInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition that prevented it from fulfilling the request
*/
type GenerateChangelogMDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this generate changelog m d internal server error response has a 2xx status code
func (o *GenerateChangelogMDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate changelog m d internal server error response has a 3xx status code
func (o *GenerateChangelogMDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate changelog m d internal server error response has a 4xx status code
func (o *GenerateChangelogMDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate changelog m d internal server error response has a 5xx status code
func (o *GenerateChangelogMDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this generate changelog m d internal server error response a status code equal to that given
func (o *GenerateChangelogMDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GenerateChangelogMDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /packages/{repoKindParam}/{repoName}/{packageName}/changelog.md][%d] generateChangelogMDInternalServerError  %+v", 500, o.Payload)
}

func (o *GenerateChangelogMDInternalServerError) String() string {
	return fmt.Sprintf("[GET /packages/{repoKindParam}/{repoName}/{packageName}/changelog.md][%d] generateChangelogMDInternalServerError  %+v", 500, o.Payload)
}

func (o *GenerateChangelogMDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GenerateChangelogMDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
