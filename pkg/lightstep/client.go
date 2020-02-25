package lightstep

import (
	"github.com/go-openapi/runtime"
	apiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	v01 "github.com/ledor473/lightstep-api-go/pkg/v0.1/client"
	v02 "github.com/ledor473/lightstep-api-go/pkg/v0.2/client"
)

// NewClientV01 returns a v0.1 client
func NewClientV01() *v01.Lightstep {
	return v01.New(customClientTransport(v01.DefaultHost, v01.DefaultBasePath, v01.DefaultSchemes), strfmt.Default)
}

// NewClientV02 returns a v0.2 client
func NewClientV02() *v02.Lightstep {
	return v02.New(customClientTransport(v02.DefaultHost, v02.DefaultBasePath, v02.DefaultSchemes), strfmt.Default)
}

func customClientTransport(host, basePath string, schemes []string) *apiclient.Runtime {
	rt := apiclient.New(host, basePath, schemes)
	rt.Consumers["application/vnd.api+json"] = runtime.JSONConsumer()

	return rt
}
