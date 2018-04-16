// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewPostSystemPruneCandidatesParams creates a new PostSystemPruneCandidatesParams object
// with the default values initialized.
func NewPostSystemPruneCandidatesParams() *PostSystemPruneCandidatesParams {
	var ()
	return &PostSystemPruneCandidatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSystemPruneCandidatesParamsWithTimeout creates a new PostSystemPruneCandidatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSystemPruneCandidatesParamsWithTimeout(timeout time.Duration) *PostSystemPruneCandidatesParams {
	var ()
	return &PostSystemPruneCandidatesParams{

		timeout: timeout,
	}
}

// NewPostSystemPruneCandidatesParamsWithContext creates a new PostSystemPruneCandidatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSystemPruneCandidatesParamsWithContext(ctx context.Context) *PostSystemPruneCandidatesParams {
	var ()
	return &PostSystemPruneCandidatesParams{

		Context: ctx,
	}
}

// NewPostSystemPruneCandidatesParamsWithHTTPClient creates a new PostSystemPruneCandidatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSystemPruneCandidatesParamsWithHTTPClient(client *http.Client) *PostSystemPruneCandidatesParams {
	var ()
	return &PostSystemPruneCandidatesParams{
		HTTPClient: client,
	}
}

/*PostSystemPruneCandidatesParams contains all the parameters to send to the API endpoint
for the post system prune candidates operation typically these are written to a http.Request
*/
type PostSystemPruneCandidatesParams struct {

	/*Bodycontent
	  resource objects to prune

	*/
	Bodycontent models.PruneCandidate
	/*Resourcetype
	  resource type

	*/
	Resourcetype string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) WithTimeout(timeout time.Duration) *PostSystemPruneCandidatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) WithContext(ctx context.Context) *PostSystemPruneCandidatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) WithHTTPClient(client *http.Client) *PostSystemPruneCandidatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBodycontent adds the bodycontent to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) WithBodycontent(bodycontent models.PruneCandidate) *PostSystemPruneCandidatesParams {
	o.SetBodycontent(bodycontent)
	return o
}

// SetBodycontent adds the bodycontent to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) SetBodycontent(bodycontent models.PruneCandidate) {
	o.Bodycontent = bodycontent
}

// WithResourcetype adds the resourcetype to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) WithResourcetype(resourcetype string) *PostSystemPruneCandidatesParams {
	o.SetResourcetype(resourcetype)
	return o
}

// SetResourcetype adds the resourcetype to the post system prune candidates params
func (o *PostSystemPruneCandidatesParams) SetResourcetype(resourcetype string) {
	o.Resourcetype = resourcetype
}

// WriteToRequest writes these params to a swagger request
func (o *PostSystemPruneCandidatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourcetype
	if err := r.SetPathParam("resourcetype", o.Resourcetype); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}