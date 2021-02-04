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

// GetRouteByIDReader is a Reader for the GetRouteByID structure.
type GetRouteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRouteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRouteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRouteByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRouteByIDOK creates a GetRouteByIDOK with default headers values
func NewGetRouteByIDOK() *GetRouteByIDOK {
	return &GetRouteByIDOK{}
}

/* GetRouteByIDOK describes a response with status code 200, with default header values.

A representation of the route.
*/
type GetRouteByIDOK struct {
	Payload *models.Route
}

func (o *GetRouteByIDOK) Error() string {
	return fmt.Sprintf("[GET /routes/{id}][%d] getRouteByIdOK  %+v", 200, o.Payload)
}
func (o *GetRouteByIDOK) GetPayload() *models.Route {
	return o.Payload
}

func (o *GetRouteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Route)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRouteByIDDefault creates a GetRouteByIDDefault with default headers values
func NewGetRouteByIDDefault(code int) *GetRouteByIDDefault {
	return &GetRouteByIDDefault{
		_statusCode: code,
	}
}

/* GetRouteByIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetRouteByIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get route by Id default response
func (o *GetRouteByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRouteByIDDefault) Error() string {
	return fmt.Sprintf("[GET /routes/{id}][%d] getRouteById default  %+v", o._statusCode, o.Payload)
}
func (o *GetRouteByIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetRouteByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
