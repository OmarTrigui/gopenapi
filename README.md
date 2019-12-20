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
$ curl http://localhost:8080/api/v2/healthz
OK

$ curl http://localhost:8080/api/v2/passwords/generateRandom?length=50
{
    password: "ciHWcHmvllC0u3igjneNEIDkQ8TyzOj6pQfF5TbsAX5dJIcaAE"
}
```

## Generate swagger files

After editing `pkg/swagger/swagger.yml` file you need to generate swagger files again:

`$ make gen`

## Test swagger file validity

`$ make validate`