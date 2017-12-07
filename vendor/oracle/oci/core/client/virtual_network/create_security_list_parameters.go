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

	"oracle/oci/core/models"
)

// NewCreateSecurityListParams creates a new CreateSecurityListParams object
// with the default values initialized.
func NewCreateSecurityListParams() *CreateSecurityListParams {
	var ()
	return &CreateSecurityListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSecurityListParamsWithTimeout creates a new CreateSecurityListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSecurityListParamsWithTimeout(timeout time.Duration) *CreateSecurityListParams {
	var ()
	return &CreateSecurityListParams{

		timeout: timeout,
	}
}

// NewCreateSecurityListParamsWithContext creates a new CreateSecurityListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSecurityListParamsWithContext(ctx context.Context) *CreateSecurityListParams {
	var ()
	return &CreateSecurityListParams{

		Context: ctx,
	}
}

// NewCreateSecurityListParamsWithHTTPClient creates a new CreateSecurityListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSecurityListParamsWithHTTPClient(client *http.Client) *CreateSecurityListParams {
	var ()
	return &CreateSecurityListParams{
		HTTPClient: client,
	}
}

/*CreateSecurityListParams contains all the parameters to send to the API endpoint
for the create security list operation typically these are written to a http.Request
*/
type CreateSecurityListParams struct {

	/*CreateSecurityListDetails
	  Details regarding the security list to create.

	*/
	CreateSecurityListDetails *models.CreateSecurityListDetails
	/*OpcRetryToken
	  A token that uniquely identifies a request so it can be retried in case of a timeout or
	server error without risk of executing that same action again. Retry tokens expire after 24
	hours, but can be invalidated before then due to conflicting operations (e.g., if a resource
	has been deleted and purged from the system, then a retry of the original creation request
	may be rejected).


	*/
	OpcRetryToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create security list params
func (o *CreateSecurityListParams) WithTimeout(timeout time.Duration) *CreateSecurityListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create security list params
func (o *CreateSecurityListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create security list params
func (o *CreateSecurityListParams) WithContext(ctx context.Context) *CreateSecurityListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create security list params
func (o *CreateSecurityListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create security list params
func (o *CreateSecurityListParams) WithHTTPClient(client *http.Client) *CreateSecurityListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create security list params
func (o *CreateSecurityListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateSecurityListDetails adds the createSecurityListDetails to the create security list params
func (o *CreateSecurityListParams) WithCreateSecurityListDetails(createSecurityListDetails *models.CreateSecurityListDetails) *CreateSecurityListParams {
	o.SetCreateSecurityListDetails(createSecurityListDetails)
	return o
}

// SetCreateSecurityListDetails adds the createSecurityListDetails to the create security list params
func (o *CreateSecurityListParams) SetCreateSecurityListDetails(createSecurityListDetails *models.CreateSecurityListDetails) {
	o.CreateSecurityListDetails = createSecurityListDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create security list params
func (o *CreateSecurityListParams) WithOpcRetryToken(opcRetryToken *string) *CreateSecurityListParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create security list params
func (o *CreateSecurityListParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSecurityListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateSecurityListDetails == nil {
		o.CreateSecurityListDetails = new(models.CreateSecurityListDetails)
	}

	if err := r.SetBodyParam(o.CreateSecurityListDetails); err != nil {
		return err
	}

	if o.OpcRetryToken != nil {

		// header param opc-retry-token
		if err := r.SetHeaderParam("opc-retry-token", *o.OpcRetryToken); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
