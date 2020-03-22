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

	strfmt "github.com/go-openapi/strfmt"
)

// NewCompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams creates a new CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams object
// with the default values initialized.
func NewCompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams() *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams {

	return &CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParamsWithTimeout creates a new CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParamsWithTimeout(timeout time.Duration) *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams {

	return &CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams{

		timeout: timeout,
	}
}

// NewCompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParamsWithContext creates a new CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParamsWithContext(ctx context.Context) *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams {

	return &CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams{

		Context: ctx,
	}
}

// NewCompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParamsWithHTTPClient creates a new CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParamsWithHTTPClient(client *http.Client) *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams {

	return &CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams{
		HTTPClient: client,
	}
}

/*CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams contains all the parameters to send to the API endpoint
for the complete self service browser profile management password strategy flow operation typically these are written to a http.Request
*/
type CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the complete self service browser profile management password strategy flow params
func (o *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams) WithTimeout(timeout time.Duration) *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete self service browser profile management password strategy flow params
func (o *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete self service browser profile management password strategy flow params
func (o *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams) WithContext(ctx context.Context) *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete self service browser profile management password strategy flow params
func (o *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete self service browser profile management password strategy flow params
func (o *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams) WithHTTPClient(client *http.Client) *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete self service browser profile management password strategy flow params
func (o *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteSelfServiceBrowserProfileManagementPasswordStrategyFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
