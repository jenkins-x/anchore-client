// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetImageContentByTypeImageIDParams creates a new GetImageContentByTypeImageIDParams object
// with the default values initialized.
func NewGetImageContentByTypeImageIDParams() *GetImageContentByTypeImageIDParams {
	var ()
	return &GetImageContentByTypeImageIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageContentByTypeImageIDParamsWithTimeout creates a new GetImageContentByTypeImageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageContentByTypeImageIDParamsWithTimeout(timeout time.Duration) *GetImageContentByTypeImageIDParams {
	var ()
	return &GetImageContentByTypeImageIDParams{

		timeout: timeout,
	}
}

// NewGetImageContentByTypeImageIDParamsWithContext creates a new GetImageContentByTypeImageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageContentByTypeImageIDParamsWithContext(ctx context.Context) *GetImageContentByTypeImageIDParams {
	var ()
	return &GetImageContentByTypeImageIDParams{

		Context: ctx,
	}
}

// NewGetImageContentByTypeImageIDParamsWithHTTPClient creates a new GetImageContentByTypeImageIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageContentByTypeImageIDParamsWithHTTPClient(client *http.Client) *GetImageContentByTypeImageIDParams {
	var ()
	return &GetImageContentByTypeImageIDParams{
		HTTPClient: client,
	}
}

/*GetImageContentByTypeImageIDParams contains all the parameters to send to the API endpoint
for the get image content by type image Id operation typically these are written to a http.Request
*/
type GetImageContentByTypeImageIDParams struct {

	/*Ctype*/
	Ctype string
	/*ImageID*/
	ImageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) WithTimeout(timeout time.Duration) *GetImageContentByTypeImageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) WithContext(ctx context.Context) *GetImageContentByTypeImageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) WithHTTPClient(client *http.Client) *GetImageContentByTypeImageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCtype adds the ctype to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) WithCtype(ctype string) *GetImageContentByTypeImageIDParams {
	o.SetCtype(ctype)
	return o
}

// SetCtype adds the ctype to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) SetCtype(ctype string) {
	o.Ctype = ctype
}

// WithImageID adds the imageID to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) WithImageID(imageID string) *GetImageContentByTypeImageIDParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the get image content by type image Id params
func (o *GetImageContentByTypeImageIDParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageContentByTypeImageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ctype
	if err := r.SetPathParam("ctype", o.Ctype); err != nil {
		return err
	}

	// path param imageId
	if err := r.SetPathParam("imageId", o.ImageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
