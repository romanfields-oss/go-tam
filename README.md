go-tam
======

The `tam/client` and `netbox/models` packages are a Go client for the REST API of [RWS Tridion Access Management](https://docs.rws.com/en-US/tridion-sites-10-1-main-documentation-1174622/tridion-access-management-768664) service.

Versioning
==========

To be determined. Meanwhile, look at brances and tags.

Using the client
================
The client is a go-swagger client: build a transport for your Tridion Access Management host, put the API token on it,
and every operation can then be called with `nil` for `authInfo`. For example:

```go
package main

import (
    "log"
    "os"

    "github.com/romanfields-oss/go-tam/tam/client"
    "github.com/romanfields-oss/go-tam/tam/client/system"
    httptransport "github.com/go-openapi/runtime/client"
)

func main() {
    token := os.GetEnv("TAM_TOKEN")
    if token == "" {
        log.Fatalf("Please provide Tridion Access Management API token via env var TAM_TOKEN")
    }

    tamHost := os.GetEnv("TAM_HOST")
    if tamHost == "" {
        log.Fatalf("Please provide tam host via env var TAM_HOST")
    }

    transport := httptransport.New(tamHost, client.DefaultBasePath, []string{"https"})
    transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Bearer "+token)

    c := client.New(transport, nil)

    req = system.NewSystemGetLoginStatusRetrieveParams()
    res, err := c.System.SystemGetLoginStatusRetrieve(req, nil)
    if err != nil {
        log.Fatalf("Cannot get login status: %v", err)
    }
    log.Printf("res: %v", res)
}
```

Go Module support
=================

Go 1.26.8+

`go get github.com/romanfields-oss/go-tam`

More complex configuration
==========================
The [godocs for the go-openapi/runtime/client module](https://godoc.org/github.com/go-openapi/runtime/client) explain the client options in detail, including different authentication and debugging options. Worth knowing: setting the `DEBUG` environment variable dumps every request to standard out.