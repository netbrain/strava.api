// Code generated by go-swagger; DO NOT EDIT.

package clubs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetClubAdminsByIDParams creates a new GetClubAdminsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClubAdminsByIDParams() *GetClubAdminsByIDParams {
	return &GetClubAdminsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClubAdminsByIDParamsWithTimeout creates a new GetClubAdminsByIDParams object
// with the ability to set a timeout on a request.
func NewGetClubAdminsByIDParamsWithTimeout(timeout time.Duration) *GetClubAdminsByIDParams {
	return &GetClubAdminsByIDParams{
		timeout: timeout,
	}
}

// NewGetClubAdminsByIDParamsWithContext creates a new GetClubAdminsByIDParams object
// with the ability to set a context for a request.
func NewGetClubAdminsByIDParamsWithContext(ctx context.Context) *GetClubAdminsByIDParams {
	return &GetClubAdminsByIDParams{
		Context: ctx,
	}
}

// NewGetClubAdminsByIDParamsWithHTTPClient creates a new GetClubAdminsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClubAdminsByIDParamsWithHTTPClient(client *http.Client) *GetClubAdminsByIDParams {
	return &GetClubAdminsByIDParams{
		HTTPClient: client,
	}
}

/* GetClubAdminsByIDParams contains all the parameters to send to the API endpoint
   for the get club admins by Id operation.

   Typically these are written to a http.Request.
*/
type GetClubAdminsByIDParams struct {

	/* ID.

	   The identifier of the club.

	   Format: int64
	*/
	ID int64

	/* Page.

	   Page number. Defaults to 1.
	*/
	Page *int64

	/* PerPage.

	   Number of items per page. Defaults to 30.

	   Default: 30
	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get club admins by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClubAdminsByIDParams) WithDefaults() *GetClubAdminsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get club admins by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClubAdminsByIDParams) SetDefaults() {
	var (
		perPageDefault = int64(30)
	)

	val := GetClubAdminsByIDParams{
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get club admins by Id params
func (o *GetClubAdminsByIDParams) WithTimeout(timeout time.Duration) *GetClubAdminsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get club admins by Id params
func (o *GetClubAdminsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get club admins by Id params
func (o *GetClubAdminsByIDParams) WithContext(ctx context.Context) *GetClubAdminsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get club admins by Id params
func (o *GetClubAdminsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get club admins by Id params
func (o *GetClubAdminsByIDParams) WithHTTPClient(client *http.Client) *GetClubAdminsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get club admins by Id params
func (o *GetClubAdminsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get club admins by Id params
func (o *GetClubAdminsByIDParams) WithID(id int64) *GetClubAdminsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get club admins by Id params
func (o *GetClubAdminsByIDParams) SetID(id int64) {
	o.ID = id
}

// WithPage adds the page to the get club admins by Id params
func (o *GetClubAdminsByIDParams) WithPage(page *int64) *GetClubAdminsByIDParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get club admins by Id params
func (o *GetClubAdminsByIDParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get club admins by Id params
func (o *GetClubAdminsByIDParams) WithPerPage(perPage *int64) *GetClubAdminsByIDParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get club admins by Id params
func (o *GetClubAdminsByIDParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetClubAdminsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
