package identity

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

// NewDeleteAPIKeyParams creates a new DeleteAPIKeyParams object
// with the default values initialized.
func NewDeleteAPIKeyParams() *DeleteAPIKeyParams {
	var ()
	return &DeleteAPIKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIKeyParamsWithTimeout creates a new DeleteAPIKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIKeyParamsWithTimeout(timeout time.Duration) *DeleteAPIKeyParams {
	var ()
	return &DeleteAPIKeyParams{

		timeout: timeout,
	}
}

// NewDeleteAPIKeyParamsWithContext creates a new DeleteAPIKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIKeyParamsWithContext(ctx context.Context) *DeleteAPIKeyParams {
	var ()
	return &DeleteAPIKeyParams{

		Context: ctx,
	}
}

// NewDeleteAPIKeyParamsWithHTTPClient creates a new DeleteAPIKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIKeyParamsWithHTTPClient(client *http.Client) *DeleteAPIKeyParams {
	var ()
	return &DeleteAPIKeyParams{
		HTTPClient: client,
	}
}

/*DeleteAPIKeyParams contains all the parameters to send to the API endpoint
for the delete Api key operation typically these are written to a http.Request
*/
type DeleteAPIKeyParams struct {

	/*Fingerprint
	  The key's fingerprint.

	*/
	Fingerprint string
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*UserID
	  The OCID of the user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete Api key params
func (o *DeleteAPIKeyParams) WithTimeout(timeout time.Duration) *DeleteAPIKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api key params
func (o *DeleteAPIKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api key params
func (o *DeleteAPIKeyParams) WithContext(ctx context.Context) *DeleteAPIKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api key params
func (o *DeleteAPIKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api key params
func (o *DeleteAPIKeyParams) WithHTTPClient(client *http.Client) *DeleteAPIKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api key params
func (o *DeleteAPIKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFingerprint adds the fingerprint to the delete Api key params
func (o *DeleteAPIKeyParams) WithFingerprint(fingerprint string) *DeleteAPIKeyParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the delete Api key params
func (o *DeleteAPIKeyParams) SetFingerprint(fingerprint string) {
	o.Fingerprint = fingerprint
}

// WithIfMatch adds the ifMatch to the delete Api key params
func (o *DeleteAPIKeyParams) WithIfMatch(ifMatch *string) *DeleteAPIKeyParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the delete Api key params
func (o *DeleteAPIKeyParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithUserID adds the userID to the delete Api key params
func (o *DeleteAPIKeyParams) WithUserID(userID string) *DeleteAPIKeyParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete Api key params
func (o *DeleteAPIKeyParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fingerprint
	if err := r.SetPathParam("fingerprint", o.Fingerprint); err != nil {
		return err
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}