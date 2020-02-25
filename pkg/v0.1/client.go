package lightstep

import (
	"github.com/go-openapi/runtime"
	openapi "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/ledor473/lightstep-api-go/pkg/v0.1/client"
)

// NewClient returns a v0.1 client
func NewClient() *client.Lightstep {
	return client.New(customClientTransport(), strfmt.Default)
}

// NewClientAuthInfoWriter returns a ClientAuthInfoWriter to be used with the client
func NewClientAuthInfoWriter(apikey string) runtime.ClientAuthInfoWriter {
	return openapi.APIKeyAuth("Authorization", "header", apikey)
}

func customClientTransport() *openapi.Runtime {
	rt := openapi.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
	rt.Consumers["application/vnd.api+json"] = runtime.JSONConsumer()

	return rt
}
