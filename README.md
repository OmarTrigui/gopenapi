# GopenApi

This repo contains a simple/basic HTTP server in Go, with a basic code organization.
We use:
* net/http package to start and serve HTTP server
* Swagger in lorder to serve a REST API compliant with OpenAPI specs

## Pre-requisits

Install Go in 1.13 version minimum.

## Build the app

`$ go build -o bin/gopenapi internal/main.go`

or

`$ make build`

## Run the app

`$ ./bin/gopenapi`

## Test the app

```
$ curl http://localhost:8080/healthz
OK

$ curl http://localhost:8080/random/string?length=50
"oyItTSUaQkYcWAi75XIZhtyIiWhHY7lajXSc5alIx4NfxixU32"
```

## Generate swagger files

After editing `pkg/swagger/swagger.yml` file you need to generate swagger files again:

`$ make gen`

## Test swagger file validity

`$ make validate`