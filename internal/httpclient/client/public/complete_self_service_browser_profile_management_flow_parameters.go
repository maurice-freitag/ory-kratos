// Code generated by go-swagger; DO NOT EDIT.

package public

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

	"github.com/ory/kratos/internal/httpclient/models"
)

// NewCompleteSelfServiceBrowserProfileManagementFlowParams creates a new CompleteSelfServiceBrowserProfileManagementFlowParams object
// with the default values initialized.
func NewCompleteSelfServiceBrowserProfileManagementFlowParams() *CompleteSelfServiceBrowserProfileManagementFlowParams {
	var ()
	return &CompleteSelfServiceBrowserProfileManagementFlowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteSelfServiceBrowserProfileManagementFlowParamsWithTimeout creates a new CompleteSelfServiceBrowserProfileManagementFlowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompleteSelfServiceBrowserProfileManagementFlowParamsWithTimeout(timeout time.Duration) *CompleteSelfServiceBrowserProfileManagementFlowParams {
	var ()
	return &CompleteSelfServiceBrowserProfileManagementFlowParams{

		timeout: timeout,
	}
}

// NewCompleteSelfServiceBrowserProfileManagementFlowParamsWithContext creates a new CompleteSelfServiceBrowserProfileManagementFlowParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompleteSelfServiceBrowserProfileManagementFlowParamsWithContext(ctx context.Context) *CompleteSelfServiceBrowserProfileManagementFlowParams {
	var ()
	return &CompleteSelfServiceBrowserProfileManagementFlowParams{

		Context: ctx,
	}
}

// NewCompleteSelfServiceBrowserProfileManagementFlowParamsWithHTTPClient creates a new CompleteSelfServiceBrowserProfileManagementFlowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompleteSelfServiceBrowserProfileManagementFlowParamsWithHTTPClient(client *http.Client) *CompleteSelfServiceBrowserProfileManagementFlowParams {
	var ()
	return &CompleteSelfServiceBrowserProfileManagementFlowParams{
		HTTPClient: client,
	}
}

/*CompleteSelfServiceBrowserProfileManagementFlowParams contains all the parameters to send to the API endpoint
for the complete self service browser profile management flow operation typically these are written to a http.Request
*/
type CompleteSelfServiceBrowserProfileManagementFlowParams struct {

	/*Body*/
	Body *models.CompleteSelfServiceBrowserProfileManagementFlowPayload
	/*Request
	  Request is the request ID.

	*/
	Request string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) WithTimeout(timeout time.Duration) *CompleteSelfServiceBrowserProfileManagementFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) WithContext(ctx context.Context) *CompleteSelfServiceBrowserProfileManagementFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) WithHTTPClient(client *http.Client) *CompleteSelfServiceBrowserProfileManagementFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) WithBody(body *models.CompleteSelfServiceBrowserProfileManagementFlowPayload) *CompleteSelfServiceBrowserProfileManagementFlowParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) SetBody(body *models.CompleteSelfServiceBrowserProfileManagementFlowPayload) {
	o.Body = body
}

// WithRequest adds the request to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) WithRequest(request string) *CompleteSelfServiceBrowserProfileManagementFlowParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the complete self service browser profile management flow params
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) SetRequest(request string) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteSelfServiceBrowserProfileManagementFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param request
	qrRequest := o.Request
	qRequest := qrRequest
	if qRequest != "" {
		if err := r.SetQueryParam("request", qRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
