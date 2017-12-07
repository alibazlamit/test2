package virtual_network

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

// NewDeleteCpeParams creates a new DeleteCpeParams object
// with the default values initialized.
func NewDeleteCpeParams() *DeleteCpeParams {
	var ()
	return &DeleteCpeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCpeParamsWithTimeout creates a new DeleteCpeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCpeParamsWithTimeout(timeout time.Duration) *DeleteCpeParams {
	var ()
	return &DeleteCpeParams{

		timeout: timeout,
	}
}

// NewDeleteCpeParamsWithContext creates a new DeleteCpeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCpeParamsWithContext(ctx context.Context) *DeleteCpeParams {
	var ()
	return &DeleteCpeParams{

		Context: ctx,
	}
}

// NewDeleteCpeParamsWithHTTPClient creates a new DeleteCpeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCpeParamsWithHTTPClient(client *http.Client) *DeleteCpeParams {
	var ()
	return &DeleteCpeParams{
		HTTPClient: client,
	}
}

/*DeleteCpeParams contains all the parameters to send to the API endpoint
for the delete cpe operation typically these are written to a http.Request
*/
type DeleteCpeParams struct {

	/*CpeID
	  The OCID of the CPE.

	*/
	CpeID string
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cpe params
func (o *DeleteCpeParams) WithTimeout(timeout time.Duration) *DeleteCpeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cpe params
func (o *DeleteCpeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cpe params
func (o *DeleteCpeParams) WithContext(ctx context.Context) *DeleteCpeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cpe params
func (o *DeleteCpeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cpe params
func (o *DeleteCpeParams) WithHTTPClient(client *http.Client) *DeleteCpeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cpe params
func (o *DeleteCpeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCpeID adds the cpeID to the delete cpe params
func (o *DeleteCpeParams) WithCpeID(cpeID string) *DeleteCpeParams {
	o.SetCpeID(cpeID)
	return o
}

// SetCpeID adds the cpeId to the delete cpe params
func (o *DeleteCpeParams) SetCpeID(cpeID string) {
	o.CpeID = cpeID
}

// WithIfMatch adds the ifMatch to the delete cpe params
func (o *DeleteCpeParams) WithIfMatch(ifMatch *string) *DeleteCpeParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the delete cpe params
func (o *DeleteCpeParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCpeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cpeId
	if err := r.SetPathParam("cpeId", o.CpeID); err != nil {
		return err
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
