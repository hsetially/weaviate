//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// GetRolesForUserDeprecatedReader is a Reader for the GetRolesForUserDeprecated structure.
type GetRolesForUserDeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesForUserDeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRolesForUserDeprecatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRolesForUserDeprecatedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRolesForUserDeprecatedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRolesForUserDeprecatedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRolesForUserDeprecatedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetRolesForUserDeprecatedUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRolesForUserDeprecatedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRolesForUserDeprecatedOK creates a GetRolesForUserDeprecatedOK with default headers values
func NewGetRolesForUserDeprecatedOK() *GetRolesForUserDeprecatedOK {
	return &GetRolesForUserDeprecatedOK{}
}

/*
GetRolesForUserDeprecatedOK describes a response with status code 200, with default header values.

Role assigned users
*/
type GetRolesForUserDeprecatedOK struct {
	Payload models.RolesListResponse
}

// IsSuccess returns true when this get roles for user deprecated o k response has a 2xx status code
func (o *GetRolesForUserDeprecatedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get roles for user deprecated o k response has a 3xx status code
func (o *GetRolesForUserDeprecatedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user deprecated o k response has a 4xx status code
func (o *GetRolesForUserDeprecatedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles for user deprecated o k response has a 5xx status code
func (o *GetRolesForUserDeprecatedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user deprecated o k response a status code equal to that given
func (o *GetRolesForUserDeprecatedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get roles for user deprecated o k response
func (o *GetRolesForUserDeprecatedOK) Code() int {
	return 200
}

func (o *GetRolesForUserDeprecatedOK) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedOK  %+v", 200, o.Payload)
}

func (o *GetRolesForUserDeprecatedOK) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedOK  %+v", 200, o.Payload)
}

func (o *GetRolesForUserDeprecatedOK) GetPayload() models.RolesListResponse {
	return o.Payload
}

func (o *GetRolesForUserDeprecatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForUserDeprecatedBadRequest creates a GetRolesForUserDeprecatedBadRequest with default headers values
func NewGetRolesForUserDeprecatedBadRequest() *GetRolesForUserDeprecatedBadRequest {
	return &GetRolesForUserDeprecatedBadRequest{}
}

/*
GetRolesForUserDeprecatedBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetRolesForUserDeprecatedBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles for user deprecated bad request response has a 2xx status code
func (o *GetRolesForUserDeprecatedBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user deprecated bad request response has a 3xx status code
func (o *GetRolesForUserDeprecatedBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user deprecated bad request response has a 4xx status code
func (o *GetRolesForUserDeprecatedBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles for user deprecated bad request response has a 5xx status code
func (o *GetRolesForUserDeprecatedBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user deprecated bad request response a status code equal to that given
func (o *GetRolesForUserDeprecatedBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get roles for user deprecated bad request response
func (o *GetRolesForUserDeprecatedBadRequest) Code() int {
	return 400
}

func (o *GetRolesForUserDeprecatedBadRequest) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesForUserDeprecatedBadRequest) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesForUserDeprecatedBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesForUserDeprecatedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForUserDeprecatedUnauthorized creates a GetRolesForUserDeprecatedUnauthorized with default headers values
func NewGetRolesForUserDeprecatedUnauthorized() *GetRolesForUserDeprecatedUnauthorized {
	return &GetRolesForUserDeprecatedUnauthorized{}
}

/*
GetRolesForUserDeprecatedUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type GetRolesForUserDeprecatedUnauthorized struct {
}

// IsSuccess returns true when this get roles for user deprecated unauthorized response has a 2xx status code
func (o *GetRolesForUserDeprecatedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user deprecated unauthorized response has a 3xx status code
func (o *GetRolesForUserDeprecatedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user deprecated unauthorized response has a 4xx status code
func (o *GetRolesForUserDeprecatedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles for user deprecated unauthorized response has a 5xx status code
func (o *GetRolesForUserDeprecatedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user deprecated unauthorized response a status code equal to that given
func (o *GetRolesForUserDeprecatedUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get roles for user deprecated unauthorized response
func (o *GetRolesForUserDeprecatedUnauthorized) Code() int {
	return 401
}

func (o *GetRolesForUserDeprecatedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedUnauthorized ", 401)
}

func (o *GetRolesForUserDeprecatedUnauthorized) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedUnauthorized ", 401)
}

func (o *GetRolesForUserDeprecatedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesForUserDeprecatedForbidden creates a GetRolesForUserDeprecatedForbidden with default headers values
func NewGetRolesForUserDeprecatedForbidden() *GetRolesForUserDeprecatedForbidden {
	return &GetRolesForUserDeprecatedForbidden{}
}

/*
GetRolesForUserDeprecatedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRolesForUserDeprecatedForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles for user deprecated forbidden response has a 2xx status code
func (o *GetRolesForUserDeprecatedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user deprecated forbidden response has a 3xx status code
func (o *GetRolesForUserDeprecatedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user deprecated forbidden response has a 4xx status code
func (o *GetRolesForUserDeprecatedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles for user deprecated forbidden response has a 5xx status code
func (o *GetRolesForUserDeprecatedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user deprecated forbidden response a status code equal to that given
func (o *GetRolesForUserDeprecatedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get roles for user deprecated forbidden response
func (o *GetRolesForUserDeprecatedForbidden) Code() int {
	return 403
}

func (o *GetRolesForUserDeprecatedForbidden) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedForbidden  %+v", 403, o.Payload)
}

func (o *GetRolesForUserDeprecatedForbidden) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedForbidden  %+v", 403, o.Payload)
}

func (o *GetRolesForUserDeprecatedForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesForUserDeprecatedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForUserDeprecatedNotFound creates a GetRolesForUserDeprecatedNotFound with default headers values
func NewGetRolesForUserDeprecatedNotFound() *GetRolesForUserDeprecatedNotFound {
	return &GetRolesForUserDeprecatedNotFound{}
}

/*
GetRolesForUserDeprecatedNotFound describes a response with status code 404, with default header values.

no role found for user
*/
type GetRolesForUserDeprecatedNotFound struct {
}

// IsSuccess returns true when this get roles for user deprecated not found response has a 2xx status code
func (o *GetRolesForUserDeprecatedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user deprecated not found response has a 3xx status code
func (o *GetRolesForUserDeprecatedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user deprecated not found response has a 4xx status code
func (o *GetRolesForUserDeprecatedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles for user deprecated not found response has a 5xx status code
func (o *GetRolesForUserDeprecatedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user deprecated not found response a status code equal to that given
func (o *GetRolesForUserDeprecatedNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get roles for user deprecated not found response
func (o *GetRolesForUserDeprecatedNotFound) Code() int {
	return 404
}

func (o *GetRolesForUserDeprecatedNotFound) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedNotFound ", 404)
}

func (o *GetRolesForUserDeprecatedNotFound) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedNotFound ", 404)
}

func (o *GetRolesForUserDeprecatedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesForUserDeprecatedUnprocessableEntity creates a GetRolesForUserDeprecatedUnprocessableEntity with default headers values
func NewGetRolesForUserDeprecatedUnprocessableEntity() *GetRolesForUserDeprecatedUnprocessableEntity {
	return &GetRolesForUserDeprecatedUnprocessableEntity{}
}

/*
GetRolesForUserDeprecatedUnprocessableEntity describes a response with status code 422, with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type GetRolesForUserDeprecatedUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles for user deprecated unprocessable entity response has a 2xx status code
func (o *GetRolesForUserDeprecatedUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user deprecated unprocessable entity response has a 3xx status code
func (o *GetRolesForUserDeprecatedUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user deprecated unprocessable entity response has a 4xx status code
func (o *GetRolesForUserDeprecatedUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles for user deprecated unprocessable entity response has a 5xx status code
func (o *GetRolesForUserDeprecatedUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user deprecated unprocessable entity response a status code equal to that given
func (o *GetRolesForUserDeprecatedUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get roles for user deprecated unprocessable entity response
func (o *GetRolesForUserDeprecatedUnprocessableEntity) Code() int {
	return 422
}

func (o *GetRolesForUserDeprecatedUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetRolesForUserDeprecatedUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetRolesForUserDeprecatedUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesForUserDeprecatedUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForUserDeprecatedInternalServerError creates a GetRolesForUserDeprecatedInternalServerError with default headers values
func NewGetRolesForUserDeprecatedInternalServerError() *GetRolesForUserDeprecatedInternalServerError {
	return &GetRolesForUserDeprecatedInternalServerError{}
}

/*
GetRolesForUserDeprecatedInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type GetRolesForUserDeprecatedInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles for user deprecated internal server error response has a 2xx status code
func (o *GetRolesForUserDeprecatedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user deprecated internal server error response has a 3xx status code
func (o *GetRolesForUserDeprecatedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user deprecated internal server error response has a 4xx status code
func (o *GetRolesForUserDeprecatedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles for user deprecated internal server error response has a 5xx status code
func (o *GetRolesForUserDeprecatedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get roles for user deprecated internal server error response a status code equal to that given
func (o *GetRolesForUserDeprecatedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get roles for user deprecated internal server error response
func (o *GetRolesForUserDeprecatedInternalServerError) Code() int {
	return 500
}

func (o *GetRolesForUserDeprecatedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesForUserDeprecatedInternalServerError) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserDeprecatedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesForUserDeprecatedInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesForUserDeprecatedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
