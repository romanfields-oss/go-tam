package system

import (
	"context"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new system API client.
func New(transport runtime.ContextualTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new system API client with basic auth credentials.
//
// It takes the following parameters:
// - host: http host (github.com)
// - basePath: any base path for the API client ("/access-management/api", "/api").
// - scheme: http scheme ("http", "https")
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new system API client with a bearer token for authentication.
//
// It takes the following parameters:
// - host: http host (github.com)
// - basePath: any base path for the API client ("/access-management/api", "/api")
// - scheme: http scheme ("http", "https")
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

// Client for the system API
type Client struct {
	transport runtime.ContextualTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods.
type ClientService interface {
	// SystemGetLoginStatusRetrieve Gets the login status of the current user.
	SystemGetLoginStatusRetrieve(params *SystemGetLoginStatusRetrieveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SystemGetLoginStatusRetrieveOK, error)

	// SystemGetLoginStatusRetrieveContext Gets the login status of the current user.
	SystemGetLoginStatusRetrieveContext(ctx context.Context, params *SystemGetLoginStatusRetrieveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SystemGetLoginStatusRetrieveOK, error)

	SetTransport(transport runtime.ContextualTransport)
}

// SystemGetLoginStatusRetrieve Gets the login status of the current user.
//
// This method does not support inject context.
// However, timeout and opentracing contexts are honored whenever enabled.
//
// If you need to pass a specific context, use [Client.SystemGetLoginStatusRetrieveContext] instead.
func (a *Client) SystemGetLoginStatusRetrieve(params *SystemGetLoginStatusRetrieveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SystemGetLoginStatusRetrieveOK, error) {
	var ctx context.Context
	ctx = context.Background()

	return a.SystemGetLoginStatusRetrieveContext(ctx, params, authInfo, opts...)
}

// SystemGetLoginStatusRetrieveContext Gets the login status of the current user.
func (a *Client) SystemGetLoginStatusRetrieveContext(ctx context.Context, params *SystemGetLoginStatusRetrieveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SystemGetLoginStatusRetrieveOK, error) {
	op := &runtime.ClientOperation{
		ID:                 "system_getLoginStatus_retrieve",
		Method:             "GET",
		PathPattern:        "/v1/System/getLoginStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Reader:             &SystemGetLoginStatusRetrieveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Client:             params.HTTPClient,
	}

	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.SubmitContext(ctx, op)
	if err != nil {
		return nil, err
	}

	// only one success response has to be checked
	success, ok := result.(*SystemGetLoginStatusRetrieveOK)
	if ok {
		return success, nil
	}

	// unexpected success response.
	//
	// a default response is provided: fill this and return an error
	unexpectedSuccess := result.(*SystemGetLoginStatusRetrieveDefault)

	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ContextualTransport) {
	a.transport = transport
}

type innerParams struct {
	timeout time.Duration
}
