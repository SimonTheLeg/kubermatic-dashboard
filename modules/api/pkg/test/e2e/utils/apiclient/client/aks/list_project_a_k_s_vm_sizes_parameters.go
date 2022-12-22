// Code generated by go-swagger; DO NOT EDIT.

package aks

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
)

// NewListProjectAKSVMSizesParams creates a new ListProjectAKSVMSizesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectAKSVMSizesParams() *ListProjectAKSVMSizesParams {
	return &ListProjectAKSVMSizesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectAKSVMSizesParamsWithTimeout creates a new ListProjectAKSVMSizesParams object
// with the ability to set a timeout on a request.
func NewListProjectAKSVMSizesParamsWithTimeout(timeout time.Duration) *ListProjectAKSVMSizesParams {
	return &ListProjectAKSVMSizesParams{
		timeout: timeout,
	}
}

// NewListProjectAKSVMSizesParamsWithContext creates a new ListProjectAKSVMSizesParams object
// with the ability to set a context for a request.
func NewListProjectAKSVMSizesParamsWithContext(ctx context.Context) *ListProjectAKSVMSizesParams {
	return &ListProjectAKSVMSizesParams{
		Context: ctx,
	}
}

// NewListProjectAKSVMSizesParamsWithHTTPClient creates a new ListProjectAKSVMSizesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectAKSVMSizesParamsWithHTTPClient(client *http.Client) *ListProjectAKSVMSizesParams {
	return &ListProjectAKSVMSizesParams{
		HTTPClient: client,
	}
}

/*
ListProjectAKSVMSizesParams contains all the parameters to send to the API endpoint

	for the list project a k s VM sizes operation.

	Typically these are written to a http.Request.
*/
type ListProjectAKSVMSizesParams struct {

	// ClientID.
	ClientID *string

	// ClientSecret.
	ClientSecret *string

	// Credential.
	Credential *string

	/* Location.

	   Location - Resource location
	*/
	Location *string

	// SubscriptionID.
	SubscriptionID *string

	// TenantID.
	TenantID *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project a k s VM sizes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAKSVMSizesParams) WithDefaults() *ListProjectAKSVMSizesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project a k s VM sizes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAKSVMSizesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithTimeout(timeout time.Duration) *ListProjectAKSVMSizesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithContext(ctx context.Context) *ListProjectAKSVMSizesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithHTTPClient(client *http.Client) *ListProjectAKSVMSizesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithClientID(clientID *string) *ListProjectAKSVMSizesParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithClientSecret(clientSecret *string) *ListProjectAKSVMSizesParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetClientSecret(clientSecret *string) {
	o.ClientSecret = clientSecret
}

// WithCredential adds the credential to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithCredential(credential *string) *ListProjectAKSVMSizesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithLocation adds the location to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithLocation(location *string) *ListProjectAKSVMSizesParams {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetLocation(location *string) {
	o.Location = location
}

// WithSubscriptionID adds the subscriptionID to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithSubscriptionID(subscriptionID *string) *ListProjectAKSVMSizesParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetSubscriptionID(subscriptionID *string) {
	o.SubscriptionID = subscriptionID
}

// WithTenantID adds the tenantID to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithTenantID(tenantID *string) *ListProjectAKSVMSizesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithProjectID adds the projectID to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) WithProjectID(projectID string) *ListProjectAKSVMSizesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project a k s VM sizes params
func (o *ListProjectAKSVMSizesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectAKSVMSizesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// header param ClientID
		if err := r.SetHeaderParam("ClientID", *o.ClientID); err != nil {
			return err
		}
	}

	if o.ClientSecret != nil {

		// header param ClientSecret
		if err := r.SetHeaderParam("ClientSecret", *o.ClientSecret); err != nil {
			return err
		}
	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.Location != nil {

		// header param Location
		if err := r.SetHeaderParam("Location", *o.Location); err != nil {
			return err
		}
	}

	if o.SubscriptionID != nil {

		// header param SubscriptionID
		if err := r.SetHeaderParam("SubscriptionID", *o.SubscriptionID); err != nil {
			return err
		}
	}

	if o.TenantID != nil {

		// header param TenantID
		if err := r.SetHeaderParam("TenantID", *o.TenantID); err != nil {
			return err
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}