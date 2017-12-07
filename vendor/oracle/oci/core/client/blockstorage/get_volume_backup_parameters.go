package blockstorage

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

// NewGetVolumeBackupParams creates a new GetVolumeBackupParams object
// with the default values initialized.
func NewGetVolumeBackupParams() *GetVolumeBackupParams {
	var ()
	return &GetVolumeBackupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVolumeBackupParamsWithTimeout creates a new GetVolumeBackupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVolumeBackupParamsWithTimeout(timeout time.Duration) *GetVolumeBackupParams {
	var ()
	return &GetVolumeBackupParams{

		timeout: timeout,
	}
}

// NewGetVolumeBackupParamsWithContext creates a new GetVolumeBackupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVolumeBackupParamsWithContext(ctx context.Context) *GetVolumeBackupParams {
	var ()
	return &GetVolumeBackupParams{

		Context: ctx,
	}
}

// NewGetVolumeBackupParamsWithHTTPClient creates a new GetVolumeBackupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVolumeBackupParamsWithHTTPClient(client *http.Client) *GetVolumeBackupParams {
	var ()
	return &GetVolumeBackupParams{
		HTTPClient: client,
	}
}

/*GetVolumeBackupParams contains all the parameters to send to the API endpoint
for the get volume backup operation typically these are written to a http.Request
*/
type GetVolumeBackupParams struct {

	/*VolumeBackupID
	  The OCID of the volume backup.

	*/
	VolumeBackupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get volume backup params
func (o *GetVolumeBackupParams) WithTimeout(timeout time.Duration) *GetVolumeBackupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get volume backup params
func (o *GetVolumeBackupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get volume backup params
func (o *GetVolumeBackupParams) WithContext(ctx context.Context) *GetVolumeBackupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get volume backup params
func (o *GetVolumeBackupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get volume backup params
func (o *GetVolumeBackupParams) WithHTTPClient(client *http.Client) *GetVolumeBackupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get volume backup params
func (o *GetVolumeBackupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVolumeBackupID adds the volumeBackupID to the get volume backup params
func (o *GetVolumeBackupParams) WithVolumeBackupID(volumeBackupID string) *GetVolumeBackupParams {
	o.SetVolumeBackupID(volumeBackupID)
	return o
}

// SetVolumeBackupID adds the volumeBackupId to the get volume backup params
func (o *GetVolumeBackupParams) SetVolumeBackupID(volumeBackupID string) {
	o.VolumeBackupID = volumeBackupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVolumeBackupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param volumeBackupId
	if err := r.SetPathParam("volumeBackupId", o.VolumeBackupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}