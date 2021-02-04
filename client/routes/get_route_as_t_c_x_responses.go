// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbrain/strava.api/models"
)

// GetRouteAsTCXReader is a Reader for the GetRouteAsTCX structure.
type GetRouteAsTCXReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRouteAsTCXReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRouteAsTCXOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRouteAsTCXDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRouteAsTCXOK creates a GetRouteAsTCXOK with default headers values
func NewGetRouteAsTCXOK() *GetRouteAsTCXOK {
	return &GetRouteAsTCXOK{}
}

/* GetRouteAsTCXOK describes a response with status code 200, with default header values.

A TCX file with the route.
*/
type GetRouteAsTCXOK struct {
}

func (o *GetRouteAsTCXOK) Error() string {
	return fmt.Sprintf("[GET /routes/{id}/export_tcx][%d] getRouteAsTCXOK ", 200)
}

func (o *GetRouteAsTCXOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRouteAsTCXDefault creates a GetRouteAsTCXDefault with default headers values
func NewGetRouteAsTCXDefault(code int) *GetRouteAsTCXDefault {
	return &GetRouteAsTCXDefault{
		_statusCode: code,
	}
}

/* GetRouteAsTCXDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetRouteAsTCXDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get route as t c x default response
func (o *GetRouteAsTCXDefault) Code() int {
	return o._statusCode
}

func (o *GetRouteAsTCXDefault) Error() string {
	return fmt.Sprintf("[GET /routes/{id}/export_tcx][%d] getRouteAsTCX default  %+v", o._statusCode, o.Payload)
}
func (o *GetRouteAsTCXDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetRouteAsTCXDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
