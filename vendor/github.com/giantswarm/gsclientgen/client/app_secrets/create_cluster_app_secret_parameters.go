// Code generated by go-swagger; DO NOT EDIT.

package app_secrets

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

// NewCreateClusterAppSecretParams creates a new CreateClusterAppSecretParams object
// with the default values initialized.
func NewCreateClusterAppSecretParams() *CreateClusterAppSecretParams {
	var ()
	return &CreateClusterAppSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterAppSecretParamsWithTimeout creates a new CreateClusterAppSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterAppSecretParamsWithTimeout(timeout time.Duration) *CreateClusterAppSecretParams {
	var ()
	return &CreateClusterAppSecretParams{

		timeout: timeout,
	}
}

// NewCreateClusterAppSecretParamsWithContext creates a new CreateClusterAppSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClusterAppSecretParamsWithContext(ctx context.Context) *CreateClusterAppSecretParams {
	var ()
	return &CreateClusterAppSecretParams{

		Context: ctx,
	}
}

// NewCreateClusterAppSecretParamsWithHTTPClient creates a new CreateClusterAppSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClusterAppSecretParamsWithHTTPClient(client *http.Client) *CreateClusterAppSecretParams {
	var ()
	return &CreateClusterAppSecretParams{
		HTTPClient: client,
	}
}

/*CreateClusterAppSecretParams contains all the parameters to send to the API endpoint
for the create cluster app secret operation typically these are written to a http.Request
*/
type CreateClusterAppSecretParams struct {

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
	/*AppName
	  App Name

	*/
	AppName string
	/*Body*/
	Body models.V4CreateClusterAppSecretRequest
	/*ClusterID
	  Cluster ID

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cluster app secret params
func (o *CreateClusterAppSecretParams) WithTimeout(timeout time.Duration) *CreateClusterAppSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster app secret params
func (o *CreateClusterAppSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster app secret params
func (o *CreateClusterAppSecretParams) WithContext(ctx context.Context) *CreateClusterAppSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster app secret params
func (o *CreateClusterAppSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster app secret params
func (o *CreateClusterAppSecretParams) WithHTTPClient(client *http.Client) *CreateClusterAppSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster app secret params
func (o *CreateClusterAppSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the create cluster app secret params
func (o *CreateClusterAppSecretParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *CreateClusterAppSecretParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the create cluster app secret params
func (o *CreateClusterAppSecretParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the create cluster app secret params
func (o *CreateClusterAppSecretParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *CreateClusterAppSecretParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the create cluster app secret params
func (o *CreateClusterAppSecretParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the create cluster app secret params
func (o *CreateClusterAppSecretParams) WithXRequestID(xRequestID *string) *CreateClusterAppSecretParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create cluster app secret params
func (o *CreateClusterAppSecretParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAppName adds the appName to the create cluster app secret params
func (o *CreateClusterAppSecretParams) WithAppName(appName string) *CreateClusterAppSecretParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the create cluster app secret params
func (o *CreateClusterAppSecretParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBody adds the body to the create cluster app secret params
func (o *CreateClusterAppSecretParams) WithBody(body models.V4CreateClusterAppSecretRequest) *CreateClusterAppSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster app secret params
func (o *CreateClusterAppSecretParams) SetBody(body models.V4CreateClusterAppSecretRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create cluster app secret params
func (o *CreateClusterAppSecretParams) WithClusterID(clusterID string) *CreateClusterAppSecretParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create cluster app secret params
func (o *CreateClusterAppSecretParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterAppSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
