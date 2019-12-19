// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

	models "github.com/giantswarm/gsclientgen/models"
)

// NewAddOrganizationParams creates a new AddOrganizationParams object
// with the default values initialized.
func NewAddOrganizationParams() *AddOrganizationParams {
	var ()
	return &AddOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddOrganizationParamsWithTimeout creates a new AddOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddOrganizationParamsWithTimeout(timeout time.Duration) *AddOrganizationParams {
	var ()
	return &AddOrganizationParams{

		timeout: timeout,
	}
}

// NewAddOrganizationParamsWithContext creates a new AddOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddOrganizationParamsWithContext(ctx context.Context) *AddOrganizationParams {
	var ()
	return &AddOrganizationParams{

		Context: ctx,
	}
}

// NewAddOrganizationParamsWithHTTPClient creates a new AddOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddOrganizationParamsWithHTTPClient(client *http.Client) *AddOrganizationParams {
	var ()
	return &AddOrganizationParams{
		HTTPClient: client,
	}
}

/*AddOrganizationParams contains all the parameters to send to the API endpoint
for the add organization operation typically these are written to a http.Request
*/
type AddOrganizationParams struct {

	/*XGiantSwarmActivity
	  Name of an activity to track, like "list-clusters". This allows to
	analyze several API requests sent in context and gives an idea on
	the purpose.


	*/
	XGiantSwarmActivity *string
	/*XGiantSwarmCmdLine
	  If activity has been issued by a CLI, this header can contain the
	command line


	*/
	XGiantSwarmCmdLine *string
	/*XRequestID
	  A randomly generated key that can be used to track a request throughout
	services of Giant Swarm.


	*/
	XRequestID *string
	/*Body*/
	Body *models.V4Organization
	/*OrganizationID
	  An ID for the organization.
	This ID must be unique and match this regular
	expression: ^[a-z0-9_]{4,30}$


	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add organization params
func (o *AddOrganizationParams) WithTimeout(timeout time.Duration) *AddOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add organization params
func (o *AddOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add organization params
func (o *AddOrganizationParams) WithContext(ctx context.Context) *AddOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add organization params
func (o *AddOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add organization params
func (o *AddOrganizationParams) WithHTTPClient(client *http.Client) *AddOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add organization params
func (o *AddOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the add organization params
func (o *AddOrganizationParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *AddOrganizationParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the add organization params
func (o *AddOrganizationParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the add organization params
func (o *AddOrganizationParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *AddOrganizationParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the add organization params
func (o *AddOrganizationParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the add organization params
func (o *AddOrganizationParams) WithXRequestID(xRequestID *string) *AddOrganizationParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the add organization params
func (o *AddOrganizationParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the add organization params
func (o *AddOrganizationParams) WithBody(body *models.V4Organization) *AddOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add organization params
func (o *AddOrganizationParams) SetBody(body *models.V4Organization) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the add organization params
func (o *AddOrganizationParams) WithOrganizationID(organizationID string) *AddOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the add organization params
func (o *AddOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *AddOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XGiantSwarmActivity != nil {

		// header param X-Giant-Swarm-Activity
		if err := r.SetHeaderParam("X-Giant-Swarm-Activity", *o.XGiantSwarmActivity); err != nil {
			return err
		}

	}

	if o.XGiantSwarmCmdLine != nil {

		// header param X-Giant-Swarm-CmdLine
		if err := r.SetHeaderParam("X-Giant-Swarm-CmdLine", *o.XGiantSwarmCmdLine); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-ID
		if err := r.SetHeaderParam("X-Request-ID", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
