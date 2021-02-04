// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/netbrain/strava.api/client/activities"
	"github.com/netbrain/strava.api/client/athletes"
	"github.com/netbrain/strava.api/client/clubs"
	"github.com/netbrain/strava.api/client/gears"
	"github.com/netbrain/strava.api/client/routes"
	"github.com/netbrain/strava.api/client/running_races"
	"github.com/netbrain/strava.api/client/segment_efforts"
	"github.com/netbrain/strava.api/client/segments"
	"github.com/netbrain/strava.api/client/streams"
	"github.com/netbrain/strava.api/client/uploads"
)

// Default strava API v3 HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "www.strava.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v3"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new strava API v3 HTTP client.
func NewHTTPClient(formats strfmt.Registry) *StravaAPIV3 {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new strava API v3 HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *StravaAPIV3 {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new strava API v3 client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *StravaAPIV3 {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(StravaAPIV3)
	cli.Transport = transport
	cli.Activities = activities.New(transport, formats)
	cli.Athletes = athletes.New(transport, formats)
	cli.Clubs = clubs.New(transport, formats)
	cli.Gears = gears.New(transport, formats)
	cli.Routes = routes.New(transport, formats)
	cli.RunningRaces = running_races.New(transport, formats)
	cli.SegmentEfforts = segment_efforts.New(transport, formats)
	cli.Segments = segments.New(transport, formats)
	cli.Streams = streams.New(transport, formats)
	cli.Uploads = uploads.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// StravaAPIV3 is a client for strava API v3
type StravaAPIV3 struct {
	Activities activities.ClientService

	Athletes athletes.ClientService

	Clubs clubs.ClientService

	Gears gears.ClientService

	Routes routes.ClientService

	RunningRaces running_races.ClientService

	SegmentEfforts segment_efforts.ClientService

	Segments segments.ClientService

	Streams streams.ClientService

	Uploads uploads.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *StravaAPIV3) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Activities.SetTransport(transport)
	c.Athletes.SetTransport(transport)
	c.Clubs.SetTransport(transport)
	c.Gears.SetTransport(transport)
	c.Routes.SetTransport(transport)
	c.RunningRaces.SetTransport(transport)
	c.SegmentEfforts.SetTransport(transport)
	c.Segments.SetTransport(transport)
	c.Streams.SetTransport(transport)
	c.Uploads.SetTransport(transport)
}
