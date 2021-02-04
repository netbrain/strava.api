// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RunningRace running race
//
// swagger:model runningRace
type RunningRace struct {

	// The name of the city in which the race is taking place.
	City string `json:"city,omitempty"`

	// The name of the country in which the race is taking place.
	Country string `json:"country,omitempty"`

	// The race's distance, in meters.
	Distance float32 `json:"distance,omitempty"`

	// The unique identifier of this race.
	ID int64 `json:"id,omitempty"`

	// The unit system in which the race should be displayed.
	// Enum: [feet meters]
	MeasurementPreference string `json:"measurement_preference,omitempty"`

	// The name of this race.
	Name string `json:"name,omitempty"`

	// The set of routes that cover this race's course.
	RouteIds []int64 `json:"route_ids"`

	// The type of this race.
	RunningRaceType int64 `json:"running_race_type,omitempty"`

	// The time at which the race begins started in the local timezone.
	// Format: date-time
	StartDateLocal strfmt.DateTime `json:"start_date_local,omitempty"`

	// The name of the state or geographical region in which the race is taking place.
	State string `json:"state,omitempty"`

	// The vanity URL of this race on Strava.
	URL string `json:"url,omitempty"`

	// The URL of this race's website.
	WebsiteURL string `json:"website_url,omitempty"`
}

// Validate validates this running race
func (m *RunningRace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeasurementPreference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateLocal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var runningRaceTypeMeasurementPreferencePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["feet","meters"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runningRaceTypeMeasurementPreferencePropEnum = append(runningRaceTypeMeasurementPreferencePropEnum, v)
	}
}

const (

	// RunningRaceMeasurementPreferenceFeet captures enum value "feet"
	RunningRaceMeasurementPreferenceFeet string = "feet"

	// RunningRaceMeasurementPreferenceMeters captures enum value "meters"
	RunningRaceMeasurementPreferenceMeters string = "meters"
)

// prop value enum
func (m *RunningRace) validateMeasurementPreferenceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, runningRaceTypeMeasurementPreferencePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RunningRace) validateMeasurementPreference(formats strfmt.Registry) error {
	if swag.IsZero(m.MeasurementPreference) { // not required
		return nil
	}

	// value enum
	if err := m.validateMeasurementPreferenceEnum("measurement_preference", "body", m.MeasurementPreference); err != nil {
		return err
	}

	return nil
}

func (m *RunningRace) validateStartDateLocal(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDateLocal) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date_local", "body", "date-time", m.StartDateLocal.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this running race based on context it is used
func (m *RunningRace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunningRace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunningRace) UnmarshalBinary(b []byte) error {
	var res RunningRace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
