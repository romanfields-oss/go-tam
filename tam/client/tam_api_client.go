package client

import (
	"maps"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/romanfields-oss/go-tam/tam/client/system"
)

// Default TAM API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host found in Meta (info) section of the spec file.
	DefaultHost string = "localhost:8081"

	// DefaultBasePath is the default BasePath found in Meta (info) section of the spec file.
	DefaultBasePath string = "/access-management/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of the spec file.
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new TAM API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *TridionAccessManagementAPI {
	return NewHttpClientWithConfig(formats, nil)
}

// NewHttpClientWithConfig creates a new TAM API HTTP client,
// using a customizable transport config.
func NewHttpClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *TridionAccessManagementAPI {
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	maps.Copy(transport.Producers, cfg.Producers)
	maps.Copy(transport.Consumers, cfg.Consumers)

	return New(transport, formats)
}

func New(transport runtime.ContextualTransport, formats strfmt.Registry) *TridionAccessManagementAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(TridionAccessManagementAPI)
	cli.Transport = transport
	cli.System = system.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host      string
	BasePath  string
	Schemes   []string
	Producers map[string]runtime.Producer
	Consumers map[string]runtime.Consumer
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// WithProducers overrides the default producers registered by [httptransport.Runtime].
func (cfg *TransportConfig) WithProducers(producers map[string]runtime.Producer) *TransportConfig {
	cfg.Producers = producers
	return cfg
}

// WithConsumers overrides the default consumers registered by [httptransport.Runtime].
func (cfg *TransportConfig) WithConsumers(consumers map[string]runtime.Consumer) *TransportConfig {
	cfg.Consumers = consumers
	return cfg
}

// TridionAccessManagementAPI is a client for the Tridion Access Management API.
type TridionAccessManagementAPI struct {
	System system.ClientService

	Transport runtime.ContextualTransport
}

// SetTransport changes the tranport on the client and all its subresources.
func (c *TridionAccessManagementAPI) SetTransport(transport runtime.ContextualTransport) {
	c.Transport = transport
	c.System.SetTransport(transport)
}
