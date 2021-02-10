// Code generated by go-swagger; DO NOT EDIT.

package admin

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

	"github.com/ory/kratos-client-go/models"
)

// NewUpdateIdentityParams creates a new UpdateIdentityParams object
// with the default values initialized.
func NewUpdateIdentityParams() *UpdateIdentityParams {
	var ()
	return &UpdateIdentityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIdentityParamsWithTimeout creates a new UpdateIdentityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateIdentityParamsWithTimeout(timeout time.Duration) *UpdateIdentityParams {
	var ()
	return &UpdateIdentityParams{

		timeout: timeout,
	}
}

// NewUpdateIdentityParamsWithContext creates a new UpdateIdentityParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateIdentityParamsWithContext(ctx context.Context) *UpdateIdentityParams {
	var ()
	return &UpdateIdentityParams{

		Context: ctx,
	}
}

// NewUpdateIdentityParamsWithHTTPClient creates a new UpdateIdentityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateIdentityParamsWithHTTPClient(client *http.Client) *UpdateIdentityParams {
	var ()
	return &UpdateIdentityParams{
		HTTPClient: client,
	}
}

/*UpdateIdentityParams contains all the parameters to send to the API endpoint
for the update identity operation typically these are written to a http.Request
*/
type UpdateIdentityParams struct {

	/*Body*/
	Body *models.UpdateIdentity
	/*ID
	  ID must be set to the ID of identity you want to update

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update identity params
func (o *UpdateIdentityParams) WithTimeout(timeout time.Duration) *UpdateIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update identity params
func (o *UpdateIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update identity params
func (o *UpdateIdentityParams) WithContext(ctx context.Context) *UpdateIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update identity params
func (o *UpdateIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update identity params
func (o *UpdateIdentityParams) WithHTTPClient(client *http.Client) *UpdateIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update identity params
func (o *UpdateIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update identity params
func (o *UpdateIdentityParams) WithBody(body *models.UpdateIdentityRequestPayload) *UpdateIdentityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update identity params
func (o *UpdateIdentityParams) SetBody(body *models.UpdateIdentityRequestPayload) {
	o.Body = body
}

// WithID adds the id to the update identity params
func (o *UpdateIdentityParams) WithID(id string) *UpdateIdentityParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update identity params
func (o *UpdateIdentityParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
