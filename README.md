# lightstep-api-go
## Unofficial client library for LightStep Public API

Mostly generated using [`go-swagger`](https://goswagger.io/). The details can be found under `internal/`

For documentation about the API, refer to the [LightStep Learning Portal](https://api-docs.lightstep.com/).

## Usage
```
api := client.New(client.Config{AuthInfo: openapi.APIKeyAuth("Authorization", "header", apikey)})
```