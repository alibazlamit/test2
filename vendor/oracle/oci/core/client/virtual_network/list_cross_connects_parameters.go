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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListCrossConnectsParams creates a new ListCrossConnectsParams object
// with the default values initialized.
func NewListCrossConnectsParams() *ListCrossConnectsParams {
	var ()
	return &ListCrossConnectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCrossConnectsParamsWithTimeout creates a new ListCrossConnectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCrossConnectsParamsWithTimeout(timeout time.Duration) *ListCrossConnectsParams {
	var ()
	return &ListCrossConnectsParams{

		timeout: timeout,
	}
}

// NewListCrossConnectsParamsWithContext creates a new ListCrossConnectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCrossConnectsParamsWithContext(ctx context.Context) *ListCrossConnectsParams {
	var ()
	return &ListCrossConnectsParams{

		Context: ctx,
	}
}

// NewListCrossConnectsParamsWithHTTPClient creates a new ListCrossConnectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCrossConnectsParamsWithHTTPClient(client *http.Client) *ListCrossConnectsParams {
	var ()
	return &ListCrossConnectsParams{
		HTTPClient: client,
	}
}

/*ListCrossConnectsParams contains all the parameters to send to the API endpoint
for the list cross connects operation typically these are written to a http.Request
*/
type ListCrossConnectsParams struct {

	/*CompartmentID
	  The OCID of the compartment.

	*/
	CompartmentID string
	/*CrossConnectGroupID
	  The OCID of the cross-connect group.

	*/
	CrossConnectGroupID *string
	/*Limit
	  The maximum number of items to return in a paginated "List" call.

	Example: `500`


	*/
	Limit *int64
	/*Page
	  The value of the `opc-next-page` response header from the previous "List" call.


	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list cross connects params
func (o *ListCrossConnectsParams) WithTimeout(timeout time.Duration) *ListCrossConnectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list cross connects params
func (o *ListCrossConnectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list cross connects params
func (o *ListCrossConnectsParams) WithContext(ctx context.Context) *ListCrossConnectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list cross connects params
func (o *ListCrossConnectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list cross connects params
func (o *ListCrossConnectsParams) WithHTTPClient(client *http.Client) *ListCrossConnectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list cross connects params
func (o *ListCrossConnectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompartmentID adds the compartmentID to the list cross connects params
func (o *ListCrossConnectsParams) WithCompartmentID(compartmentID string) *ListCrossConnectsParams {
	o.SetCompartmentID(compartmentID)
	return o
}

// SetCompartmentID adds the compartmentId to the list cross connects params
func (o *ListCrossConnectsParams) SetCompartmentID(compartmentID string) {
	o.CompartmentID = compartmentID
}

// WithCrossConnectGroupID adds the crossConnectGroupID to the list cross connects params
func (o *ListCrossConnectsParams) WithCrossConnectGroupID(crossConnectGroupID *string) *ListCrossConnectsParams {
	o.SetCrossConnectGroupID(crossConnectGroupID)
	return o
}

// SetCrossConnectGroupID adds the crossConnectGroupId to the list cross connects params
func (o *ListCrossConnectsParams) SetCrossConnectGroupID(crossConnectGroupID *string) {
	o.CrossConnectGroupID = crossConnectGroupID
}

// WithLimit adds the limit to the list cross connects params
func (o *ListCrossConnectsParams) WithLimit(limit *int64) *ListCrossConnectsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list cross connects params
func (o *ListCrossConnectsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the list cross connects params
func (o *ListCrossConnectsParams) WithPage(page *string) *ListCrossConnectsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list cross connects params
func (o *ListCrossConnectsParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *ListCrossConnectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param compartmentId
	qrCompartmentID := o.CompartmentID
	qCompartmentID := qrCompartmentID
	if qCompartmentID != "" {
		if err := r.SetQueryParam("compartmentId", qCompartmentID); err != nil {
			return err
		}
	}

	if o.CrossConnectGroupID != nil {

		// query param crossConnectGroupId
		var qrCrossConnectGroupID string
		if o.CrossConnectGroupID != nil {
			qrCrossConnectGroupID = *o.CrossConnectGroupID
		}
		qCrossConnectGroupID := qrCrossConnectGroupID
		if qCrossConnectGroupID != "" {
			if err := r.SetQueryParam("crossConnectGroupId", qCrossConnectGroupID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}