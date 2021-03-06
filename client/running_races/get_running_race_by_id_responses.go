// Code generated by go-swagger; DO NOT EDIT.

package running_races

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbrain/strava.api/models"
)

// GetRunningRaceByIDReader is a Reader for the GetRunningRaceByID structure.
type GetRunningRaceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunningRaceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunningRaceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRunningRaceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunningRaceByIDOK creates a GetRunningRaceByIDOK with default headers values
func NewGetRunningRaceByIDOK() *GetRunningRaceByIDOK {
	return &GetRunningRaceByIDOK{}
}

/* GetRunningRaceByIDOK describes a response with status code 200, with default header values.

Representation of a running race.
*/
type GetRunningRaceByIDOK struct {
	Payload *models.RunningRace
}

func (o *GetRunningRaceByIDOK) Error() string {
	return fmt.Sprintf("[GET /running_races/{id}][%d] getRunningRaceByIdOK  %+v", 200, o.Payload)
}
func (o *GetRunningRaceByIDOK) GetPayload() *models.RunningRace {
	return o.Payload
}

func (o *GetRunningRaceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunningRace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunningRaceByIDDefault creates a GetRunningRaceByIDDefault with default headers values
func NewGetRunningRaceByIDDefault(code int) *GetRunningRaceByIDDefault {
	return &GetRunningRaceByIDDefault{
		_statusCode: code,
	}
}

/* GetRunningRaceByIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetRunningRaceByIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get running race by Id default response
func (o *GetRunningRaceByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRunningRaceByIDDefault) Error() string {
	return fmt.Sprintf("[GET /running_races/{id}][%d] getRunningRaceById default  %+v", o._statusCode, o.Payload)
}
func (o *GetRunningRaceByIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetRunningRaceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
