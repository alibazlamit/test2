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

// UpdateIPSecConnectionReader is a Reader for the UpdateIPSecConnection structure.
type UpdateIPSecConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIPSecConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateIPSecConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateIPSecConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateIPSecConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateIPSecConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateIPSecConnectionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateIPSecConnectionPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateIPSecConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateIPSecConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateIPSecConnectionOK creates a UpdateIPSecConnectionOK with default headers values
func NewUpdateIPSecConnectionOK() *UpdateIPSecConnectionOK {
	return &UpdateIPSecConnectionOK{}
}

/*UpdateIPSecConnectionOK handles this case with default header values.

The IPSec connection was updated.
*/
type UpdateIPSecConnectionOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.IPSecConnection
}

func (o *UpdateIPSecConnectionOK) Error() string {
	return fmt.Sprintf("[PUT /ipsecConnections/{ipscId}][%d] updateIpSecConnectionOK  %+v", 200, o.Payload)
}

func (o *UpdateIPSecConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.IPSecConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPSecConnectionBadRequest creates a UpdateIPSecConnectionBadRequest with default headers values
func NewUpdateIPSecConnectionBadRequest() *UpdateIPSecConnectionBadRequest {
	return &UpdateIPSecConnectionBadRequest{}
}

/*UpdateIPSecConnectionBadRequest handles this case with default header values.

Bad Request
*/
type UpdateIPSecConnectionBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateIPSecConnectionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipsecConnections/{ipscId}][%d] updateIpSecConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateIPSecConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPSecConnectionUnauthorized creates a UpdateIPSecConnectionUnauthorized with default headers values
func NewUpdateIPSecConnectionUnauthorized() *UpdateIPSecConnectionUnauthorized {
	return &UpdateIPSecConnectionUnauthorized{}
}

/*UpdateIPSecConnectionUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateIPSecConnectionUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateIPSecConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /ipsecConnections/{ipscId}][%d] updateIpSecConnectionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateIPSecConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPSecConnectionNotFound creates a UpdateIPSecConnectionNotFound with default headers values
func NewUpdateIPSecConnectionNotFound() *UpdateIPSecConnectionNotFound {
	return &UpdateIPSecConnectionNotFound{}
}

/*UpdateIPSecConnectionNotFound handles this case with default header values.

Not Found
*/
type UpdateIPSecConnectionNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateIPSecConnectionNotFound) Error() string {
	return fmt.Sprintf("[PUT /ipsecConnections/{ipscId}][%d] updateIpSecConnectionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateIPSecConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPSecConnectionConflict creates a UpdateIPSecConnectionConflict with default headers values
func NewUpdateIPSecConnectionConflict() *UpdateIPSecConnectionConflict {
	return &UpdateIPSecConnectionConflict{}
}

/*UpdateIPSecConnectionConflict handles this case with default header values.

Conflict
*/
type UpdateIPSecConnectionConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateIPSecConnectionConflict) Error() string {
	return fmt.Sprintf("[PUT /ipsecConnections/{ipscId}][%d] updateIpSecConnectionConflict  %+v", 409, o.Payload)
}

func (o *UpdateIPSecConnectionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPSecConnectionPreconditionFailed creates a UpdateIPSecConnectionPreconditionFailed with default headers values
func NewUpdateIPSecConnectionPreconditionFailed() *UpdateIPSecConnectionPreconditionFailed {
	return &UpdateIPSecConnectionPreconditionFailed{}
}

/*UpdateIPSecConnectionPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateIPSecConnectionPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateIPSecConnectionPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /ipsecConnections/{ipscId}][%d] updateIpSecConnectionPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateIPSecConnectionPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPSecConnectionInternalServerError creates a UpdateIPSecConnectionInternalServerError with default headers values
func NewUpdateIPSecConnectionInternalServerError() *UpdateIPSecConnectionInternalServerError {
	return &UpdateIPSecConnectionInternalServerError{}
}

/*UpdateIPSecConnectionInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateIPSecConnectionInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateIPSecConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /ipsecConnections/{ipscId}][%d] updateIpSecConnectionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateIPSecConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPSecConnectionDefault creates a UpdateIPSecConnectionDefault with default headers values
func NewUpdateIPSecConnectionDefault(code int) *UpdateIPSecConnectionDefault {
	return &UpdateIPSecConnectionDefault{
		_statusCode: code,
	}
}

/*UpdateIPSecConnectionDefault handles this case with default header values.

An error has occurred.
*/
type UpdateIPSecConnectionDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update IP sec connection default response
func (o *UpdateIPSecConnectionDefault) Code() int {
	return o._statusCode
}

func (o *UpdateIPSecConnectionDefault) Error() string {
	return fmt.Sprintf("[PUT /ipsecConnections/{ipscId}][%d] UpdateIPSecConnection default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateIPSecConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
