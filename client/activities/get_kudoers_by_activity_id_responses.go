// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbrain/strava.api/models"
)

// GetKudoersByActivityIDReader is a Reader for the GetKudoersByActivityID structure.
type GetKudoersByActivityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKudoersByActivityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKudoersByActivityIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetKudoersByActivityIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKudoersByActivityIDOK creates a GetKudoersByActivityIDOK with default headers values
func NewGetKudoersByActivityIDOK() *GetKudoersByActivityIDOK {
	return &GetKudoersByActivityIDOK{}
}

/* GetKudoersByActivityIDOK describes a response with status code 200, with default header values.

Comments.
*/
type GetKudoersByActivityIDOK struct {
	Payload []*models.SummaryAthlete
}

func (o *GetKudoersByActivityIDOK) Error() string {
	return fmt.Sprintf("[GET /activities/{id}/kudos][%d] getKudoersByActivityIdOK  %+v", 200, o.Payload)
}
func (o *GetKudoersByActivityIDOK) GetPayload() []*models.SummaryAthlete {
	return o.Payload
}

func (o *GetKudoersByActivityIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKudoersByActivityIDDefault creates a GetKudoersByActivityIDDefault with default headers values
func NewGetKudoersByActivityIDDefault(code int) *GetKudoersByActivityIDDefault {
	return &GetKudoersByActivityIDDefault{
		_statusCode: code,
	}
}

/* GetKudoersByActivityIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetKudoersByActivityIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get kudoers by activity Id default response
func (o *GetKudoersByActivityIDDefault) Code() int {
	return o._statusCode
}

func (o *GetKudoersByActivityIDDefault) Error() string {
	return fmt.Sprintf("[GET /activities/{id}/kudos][%d] getKudoersByActivityId default  %+v", o._statusCode, o.Payload)
}
func (o *GetKudoersByActivityIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetKudoersByActivityIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
