package main

import (
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/buaazp/fasthttprouter"

	"github.com/xSHAD0Wx/simple-api/services"
	"github.com/xSHAD0Wx/simple-api/shared"
)

func main() {
	// Defer cleanup
	defer services.Services.DB.Close()

	log.Printf("Started server on: %s\n", time.Now())

	// Create a new router
	mainRouter := fasthttprouter.New()

	// Configure app routes
	log.Printf("Configuring router...")
	if !ConfigureFastRouter(mainRouter) {
		return
	}

	log.Printf("Configuring routes...")
	if !ConfigureFastRoutes(mainRouter) {
		return
	}

	// Start the server
	log.Printf("Starting server on: localhost:%d", shared.ListenPort)
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%d", shared.ListenPort), mainRouter.Handler))
}
