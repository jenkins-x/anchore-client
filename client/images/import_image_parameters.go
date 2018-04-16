// Code generated by go-swagger; DO NOT EDIT.

package images

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

	"github.com/jenkins-x/anchore-client/models"
)

// NewImportImageParams creates a new ImportImageParams object
// with the default values initialized.
func NewImportImageParams() *ImportImageParams {
	var ()
	return &ImportImageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImportImageParamsWithTimeout creates a new ImportImageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImportImageParamsWithTimeout(timeout time.Duration) *ImportImageParams {
	var ()
	return &ImportImageParams{

		timeout: timeout,
	}
}

// NewImportImageParamsWithContext creates a new ImportImageParams object
// with the default values initialized, and the ability to set a context for a request
func NewImportImageParamsWithContext(ctx context.Context) *ImportImageParams {
	var ()
	return &ImportImageParams{

		Context: ctx,
	}
}

// NewImportImageParamsWithHTTPClient creates a new ImportImageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImportImageParamsWithHTTPClient(client *http.Client) *ImportImageParams {
	var ()
	return &ImportImageParams{
		HTTPClient: client,
	}
}

/*ImportImageParams contains all the parameters to send to the API endpoint
for the import image operation typically these are written to a http.Request
*/
type ImportImageParams struct {

	/*AnalysisReport*/
	AnalysisReport models.ImageAnalysisReport

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the import image params
func (o *ImportImageParams) WithTimeout(timeout time.Duration) *ImportImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import image params
func (o *ImportImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import image params
func (o *ImportImageParams) WithContext(ctx context.Context) *ImportImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import image params
func (o *ImportImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import image params
func (o *ImportImageParams) WithHTTPClient(client *http.Client) *ImportImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import image params
func (o *ImportImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnalysisReport adds the analysisReport to the import image params
func (o *ImportImageParams) WithAnalysisReport(analysisReport models.ImageAnalysisReport) *ImportImageParams {
	o.SetAnalysisReport(analysisReport)
	return o
}

// SetAnalysisReport adds the analysisReport to the import image params
func (o *ImportImageParams) SetAnalysisReport(analysisReport models.ImageAnalysisReport) {
	o.AnalysisReport = analysisReport
}

// WriteToRequest writes these params to a swagger request
func (o *ImportImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
