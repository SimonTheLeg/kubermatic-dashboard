// Code generated by go-swagger; DO NOT EDIT.

package networkdefaults

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

// NewGetNetworkDefaultsParams creates a new GetNetworkDefaultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkDefaultsParams() *GetNetworkDefaultsParams {
	return &GetNetworkDefaultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkDefaultsParamsWithTimeout creates a new GetNetworkDefaultsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkDefaultsParamsWithTimeout(timeout time.Duration) *GetNetworkDefaultsParams {
	return &GetNetworkDefaultsParams{
		timeout: timeout,
	}
}

// NewGetNetworkDefaultsParamsWithContext creates a new GetNetworkDefaultsParams object
// with the ability to set a context for a request.
func NewGetNetworkDefaultsParamsWithContext(ctx context.Context) *GetNetworkDefaultsParams {
	return &GetNetworkDefaultsParams{
		Context: ctx,
	}
}

// NewGetNetworkDefaultsParamsWithHTTPClient creates a new GetNetworkDefaultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkDefaultsParamsWithHTTPClient(client *http.Client) *GetNetworkDefaultsParams {
	return &GetNetworkDefaultsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkDefaultsParams contains all the parameters to send to the API endpoint

	for the get network defaults operation.

	Typically these are written to a http.Request.
*/
type GetNetworkDefaultsParams struct {

	// Dc.
	DC string

	// ProviderName.
	ProviderName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkDefaultsParams) WithDefaults() *GetNetworkDefaultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkDefaultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network defaults params
func (o *GetNetworkDefaultsParams) WithTimeout(timeout time.Duration) *GetNetworkDefaultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network defaults params
func (o *GetNetworkDefaultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network defaults params
func (o *GetNetworkDefaultsParams) WithContext(ctx context.Context) *GetNetworkDefaultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network defaults params
func (o *GetNetworkDefaultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network defaults params
func (o *GetNetworkDefaultsParams) WithHTTPClient(client *http.Client) *GetNetworkDefaultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network defaults params
func (o *GetNetworkDefaultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDC adds the dc to the get network defaults params
func (o *GetNetworkDefaultsParams) WithDC(dc string) *GetNetworkDefaultsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the get network defaults params
func (o *GetNetworkDefaultsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProviderName adds the providerName to the get network defaults params
func (o *GetNetworkDefaultsParams) WithProviderName(providerName string) *GetNetworkDefaultsParams {
	o.SetProviderName(providerName)
	return o
}

// SetProviderName adds the providerName to the get network defaults params
func (o *GetNetworkDefaultsParams) SetProviderName(providerName string) {
	o.ProviderName = providerName
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkDefaultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param provider_name
	if err := r.SetPathParam("provider_name", o.ProviderName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}