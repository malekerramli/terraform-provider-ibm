// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteVpcsVpcIDAddressPrefixesIDParams creates a new DeleteVpcsVpcIDAddressPrefixesIDParams object
// with the default values initialized.
func NewDeleteVpcsVpcIDAddressPrefixesIDParams() *DeleteVpcsVpcIDAddressPrefixesIDParams {
	var ()
	return &DeleteVpcsVpcIDAddressPrefixesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVpcsVpcIDAddressPrefixesIDParamsWithTimeout creates a new DeleteVpcsVpcIDAddressPrefixesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVpcsVpcIDAddressPrefixesIDParamsWithTimeout(timeout time.Duration) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	var ()
	return &DeleteVpcsVpcIDAddressPrefixesIDParams{

		timeout: timeout,
	}
}

// NewDeleteVpcsVpcIDAddressPrefixesIDParamsWithContext creates a new DeleteVpcsVpcIDAddressPrefixesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVpcsVpcIDAddressPrefixesIDParamsWithContext(ctx context.Context) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	var ()
	return &DeleteVpcsVpcIDAddressPrefixesIDParams{

		Context: ctx,
	}
}

// NewDeleteVpcsVpcIDAddressPrefixesIDParamsWithHTTPClient creates a new DeleteVpcsVpcIDAddressPrefixesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVpcsVpcIDAddressPrefixesIDParamsWithHTTPClient(client *http.Client) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	var ()
	return &DeleteVpcsVpcIDAddressPrefixesIDParams{
		HTTPClient: client,
	}
}

/*DeleteVpcsVpcIDAddressPrefixesIDParams contains all the parameters to send to the API endpoint
for the delete vpcs vpc ID address prefixes ID operation typically these are written to a http.Request
*/
type DeleteVpcsVpcIDAddressPrefixesIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The prefix identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcID
	  The VPC identifier

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) WithTimeout(timeout time.Duration) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) WithContext(ctx context.Context) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) WithHTTPClient(client *http.Client) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) WithGeneration(generation int64) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) WithID(id string) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) WithVersion(version string) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcID adds the vpcID to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) WithVpcID(vpcID string) *DeleteVpcsVpcIDAddressPrefixesIDParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the delete vpcs vpc ID address prefixes ID params
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVpcsVpcIDAddressPrefixesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	// path param vpc_id
	if err := r.SetPathParam("vpc_id", o.VpcID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
