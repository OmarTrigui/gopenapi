package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/OmarTrigui/gopenapi/pkg/swagger/server/restapi"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/OmarTrigui/gopenapi/pkg/swagger/server/restapi/operations"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGopenapiAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	// Implement the CheckHealth handler
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)

	// Implement the GetPasswordsGenerateRandomHandler handler
	api.GetPasswordsGenerateRandomHandler = operations.GetPasswordsGenerateRandomHandlerFunc(GetRandomPassword)

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

//Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

//GetRandomPassword returns a random password with a specific length
func GetRandomPassword(randomString operations.GetPasswordsGenerateRandomParams) middleware.Responder {

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, randomString.Length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	payload := operations.GetPasswordsGenerateRandomOKBody{Password: string(b)}

	return operations.NewGetPasswordsGenerateRandomOK().WithPayload(&payload)
}
