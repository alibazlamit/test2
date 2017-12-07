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

// GetCrossConnectStatusReader is a Reader for the GetCrossConnectStatus structure.
type GetCrossConnectStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCrossConnectStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCrossConnectStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetCrossConnectStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCrossConnectStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCrossConnectStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCrossConnectStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCrossConnectStatusOK creates a GetCrossConnectStatusOK with default headers values
func NewGetCrossConnectStatusOK() *GetCrossConnectStatusOK {
	return &GetCrossConnectStatusOK{}
}

/*GetCrossConnectStatusOK handles this case with default header values.

The cross-connect's status was retrieved.
*/
type GetCrossConnectStatusOK struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.CrossConnectStatus
}

func (o *GetCrossConnectStatusOK) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/status][%d] getCrossConnectStatusOK  %+v", 200, o.Payload)
}

func (o *GetCrossConnectStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.CrossConnectStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrossConnectStatusUnauthorized creates a GetCrossConnectStatusUnauthorized with default headers values
func NewGetCrossConnectStatusUnauthorized() *GetCrossConnectStatusUnauthorized {
	return &GetCrossConnectStatusUnauthorized{}
}

/*GetCrossConnectStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCrossConnectStatusUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCrossConnectStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/status][%d] getCrossConnectStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCrossConnectStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrossConnectStatusNotFound creates a GetCrossConnectStatusNotFound with default headers values
func NewGetCrossConnectStatusNotFound() *GetCrossConnectStatusNotFound {
	return &GetCrossConnectStatusNotFound{}
}

/*GetCrossConnectStatusNotFound handles this case with default header values.

Not Found
*/
type GetCrossConnectStatusNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCrossConnectStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/status][%d] getCrossConnectStatusNotFound  %+v", 404, o.Payload)
}

func (o *GetCrossConnectStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrossConnectStatusInternalServerError creates a GetCrossConnectStatusInternalServerError with default headers values
func NewGetCrossConnectStatusInternalServerError() *GetCrossConnectStatusInternalServerError {
	return &GetCrossConnectStatusInternalServerError{}
}

/*GetCrossConnectStatusInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetCrossConnectStatusInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetCrossConnectStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/status][%d] getCrossConnectStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCrossConnectStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrossConnectStatusDefault creates a GetCrossConnectStatusDefault with default headers values
func NewGetCrossConnectStatusDefault(code int) *GetCrossConnectStatusDefault {
	return &GetCrossConnectStatusDefault{
		_statusCode: code,
	}
}

/*GetCrossConnectStatusDefault handles this case with default header values.

An error has occurred.
*/
type GetCrossConnectStatusDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get cross connect status default response
func (o *GetCrossConnectStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetCrossConnectStatusDefault) Error() string {
	return fmt.Sprintf("[GET /crossConnects/{crossConnectId}/status][%d] GetCrossConnectStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetCrossConnectStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
