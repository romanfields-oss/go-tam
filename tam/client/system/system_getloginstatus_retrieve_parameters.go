package system

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewSystemGetLoginStatusRetrieveParams creates a new SystemGetLoginStatusRetrieveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemGetLoginStatusRetrieveParams() *SystemGetLoginStatusRetrieveParams {
	return NewSystemGetLoginStatusRetrieveParamsWithTimeout(cr.DefaultTimeout)
}

// NewSystemGetLoginStatusRetrieveParmasWithTimeout creates a new SystemGetLoginStatusRetrieveParams object
// with the ability to set a timeout on a request.
func NewSystemGetLoginStatusRetrieveParamsWithTimeout(timeout time.Duration) *SystemGetLoginStatusRetrieveParams {
	return &SystemGetLoginStatusRetrieveParams{
		inner: innerParams{
			timeout: timeout,
		},
	}
}

// NewSystemGetLoginStatusRetrieveParamsWithHTTPClient creates a new SystemGetLoginStatusRetrieveParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemGetLoginStatusRetrieveParamsWithHTTPClient(client *http.Client) *SystemGetLoginStatusRetrieveParams {
	return &SystemGetLoginStatusRetrieveParams{
		HTTPClient: client,
	}
}

/*
SystemGetLoginStatusRetrieveParams contains all the parameters to send to the API endpoint

	for the system getLoginStatus operation.

	Typically these are written to a http.Request.
*/
type SystemGetLoginStatusRetrieveParams struct {
	HTTPClient *http.Client

	inner innerParams
}

// SetDefaults hydrates default values in the system getLoginStatus retrieve params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemGetLoginStatusRetrieveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system getLoginStatus retrieve params.
func (o *SystemGetLoginStatusRetrieveParams) WithTimeout(timeout time.Duration) *SystemGetLoginStatusRetrieveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system getLoginStatus retrieve params.
func (o *SystemGetLoginStatusRetrieveParams) SetTimeout(timeout time.Duration) {
	o.inner.timeout = timeout
}

// WithHTTPClient adds the HTTPClient to the system getLoginStatus retrieve params.
func (o *SystemGetLoginStatusRetrieveParams) WithHTTPClient(client *http.Client) *SystemGetLoginStatusRetrieveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system getLoginStatus retrieve params.
func (o *SystemGetLoginStatusRetrieveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a [runtime.ClientRequest].
func (o *SystemGetLoginStatusRetrieveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.inner.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
