package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// UpdateSecurityListReader is a Reader for the UpdateSecurityList structure.
type UpdateSecurityListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSecurityListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSecurityListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSecurityListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateSecurityListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSecurityListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateSecurityListConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateSecurityListPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateSecurityListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateSecurityListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSecurityListOK creates a UpdateSecurityListOK with default headers values
func NewUpdateSecurityListOK() *UpdateSecurityListOK {
	return &UpdateSecurityListOK{}
}

/*UpdateSecurityListOK handles this case with default header values.

The security list was updated.
*/
type UpdateSecurityListOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.SecurityList
}

func (o *UpdateSecurityListOK) Error() string {
	return fmt.Sprintf("[PUT /securityLists/{securityListId}][%d] updateSecurityListOK  %+v", 200, o.Payload)
}

func (o *UpdateSecurityListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.SecurityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityListBadRequest creates a UpdateSecurityListBadRequest with default headers values
func NewUpdateSecurityListBadRequest() *UpdateSecurityListBadRequest {
	return &UpdateSecurityListBadRequest{}
}

/*UpdateSecurityListBadRequest handles this case with default header values.

Bad Request
*/
type UpdateSecurityListBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSecurityListBadRequest) Error() string {
	return fmt.Sprintf("[PUT /securityLists/{securityListId}][%d] updateSecurityListBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSecurityListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityListUnauthorized creates a UpdateSecurityListUnauthorized with default headers values
func NewUpdateSecurityListUnauthorized() *UpdateSecurityListUnauthorized {
	return &UpdateSecurityListUnauthorized{}
}

/*UpdateSecurityListUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateSecurityListUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSecurityListUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /securityLists/{securityListId}][%d] updateSecurityListUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateSecurityListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityListNotFound creates a UpdateSecurityListNotFound with default headers values
func NewUpdateSecurityListNotFound() *UpdateSecurityListNotFound {
	return &UpdateSecurityListNotFound{}
}

/*UpdateSecurityListNotFound handles this case with default header values.

Not Found
*/
type UpdateSecurityListNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSecurityListNotFound) Error() string {
	return fmt.Sprintf("[PUT /securityLists/{securityListId}][%d] updateSecurityListNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSecurityListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityListConflict creates a UpdateSecurityListConflict with default headers values
func NewUpdateSecurityListConflict() *UpdateSecurityListConflict {
	return &UpdateSecurityListConflict{}
}

/*UpdateSecurityListConflict handles this case with default header values.

Conflict
*/
type UpdateSecurityListConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSecurityListConflict) Error() string {
	return fmt.Sprintf("[PUT /securityLists/{securityListId}][%d] updateSecurityListConflict  %+v", 409, o.Payload)
}

func (o *UpdateSecurityListConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityListPreconditionFailed creates a UpdateSecurityListPreconditionFailed with default headers values
func NewUpdateSecurityListPreconditionFailed() *UpdateSecurityListPreconditionFailed {
	return &UpdateSecurityListPreconditionFailed{}
}

/*UpdateSecurityListPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateSecurityListPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSecurityListPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /securityLists/{securityListId}][%d] updateSecurityListPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateSecurityListPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityListInternalServerError creates a UpdateSecurityListInternalServerError with default headers values
func NewUpdateSecurityListInternalServerError() *UpdateSecurityListInternalServerError {
	return &UpdateSecurityListInternalServerError{}
}

/*UpdateSecurityListInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateSecurityListInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateSecurityListInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /securityLists/{securityListId}][%d] updateSecurityListInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSecurityListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityListDefault creates a UpdateSecurityListDefault with default headers values
func NewUpdateSecurityListDefault(code int) *UpdateSecurityListDefault {
	return &UpdateSecurityListDefault{
		_statusCode: code,
	}
}

/*UpdateSecurityListDefault handles this case with default header values.

An error has occurred.
*/
type UpdateSecurityListDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update security list default response
func (o *UpdateSecurityListDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSecurityListDefault) Error() string {
	return fmt.Sprintf("[PUT /securityLists/{securityListId}][%d] UpdateSecurityList default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSecurityListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
