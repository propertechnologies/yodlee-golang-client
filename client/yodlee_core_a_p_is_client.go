// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"yodlee-golang-client/client/accounts"
	"yodlee-golang-client/client/auth"
	"yodlee-golang-client/client/cobrand"
	"yodlee-golang-client/client/data_extracts"
	"yodlee-golang-client/client/derived"
	"yodlee-golang-client/client/documents"
	"yodlee-golang-client/client/holdings"
	"yodlee-golang-client/client/provider_accounts"
	"yodlee-golang-client/client/providers"
	"yodlee-golang-client/client/statements"
	"yodlee-golang-client/client/transactions"
	"yodlee-golang-client/client/user"
	"yodlee-golang-client/client/verification"
)

// Default yodlee core a p is HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new yodlee core a p is HTTP client.
func NewHTTPClient(formats strfmt.Registry) *YodleeCoreAPIs {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new yodlee core a p is HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *YodleeCoreAPIs {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new yodlee core a p is client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *YodleeCoreAPIs {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(YodleeCoreAPIs)
	cli.Transport = transport

	cli.Accounts = accounts.New(transport, formats)

	cli.Auth = auth.New(transport, formats)

	cli.Cobrand = cobrand.New(transport, formats)

	cli.DataExtracts = data_extracts.New(transport, formats)

	cli.Derived = derived.New(transport, formats)

	cli.Documents = documents.New(transport, formats)

	cli.Holdings = holdings.New(transport, formats)

	cli.ProviderAccounts = provider_accounts.New(transport, formats)

	cli.Providers = providers.New(transport, formats)

	cli.Statements = statements.New(transport, formats)

	cli.Transactions = transactions.New(transport, formats)

	cli.User = user.New(transport, formats)

	cli.Verification = verification.New(transport, formats)

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
	Host     string
	BasePath string
	Schemes  []string
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

// YodleeCoreAPIs is a client for yodlee core a p is
type YodleeCoreAPIs struct {
	Accounts *accounts.Client

	Auth *auth.Client

	Cobrand *cobrand.Client

	DataExtracts *data_extracts.Client

	Derived *derived.Client

	Documents *documents.Client

	Holdings *holdings.Client

	ProviderAccounts *provider_accounts.Client

	Providers *providers.Client

	Statements *statements.Client

	Transactions *transactions.Client

	User *user.Client

	Verification *verification.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *YodleeCoreAPIs) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Accounts.SetTransport(transport)

	c.Auth.SetTransport(transport)

	c.Cobrand.SetTransport(transport)

	c.DataExtracts.SetTransport(transport)

	c.Derived.SetTransport(transport)

	c.Documents.SetTransport(transport)

	c.Holdings.SetTransport(transport)

	c.ProviderAccounts.SetTransport(transport)

	c.Providers.SetTransport(transport)

	c.Statements.SetTransport(transport)

	c.Transactions.SetTransport(transport)

	c.User.SetTransport(transport)

	c.Verification.SetTransport(transport)

}
