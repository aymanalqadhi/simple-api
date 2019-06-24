package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/xSHAD0Wx/simple-api/services"
	"github.com/xSHAD0Wx/simple-api/shared"

	"github.com/gorilla/mux"
)

func main() {
	// Defer cleanup
	defer services.Services.DB.Close()

	log.Printf("Started server on: %s\n", time.Now())

	// Create a new router
	mainRouter := mux.NewRouter()

	// Configure app routes
	log.Printf("Configuring router...")
	if !ConfigureRouter(mainRouter) {
		return
	}

	log.Printf("Configuring routes...")
	if !ConfigureRoutes(mainRouter) {
		return
	}

	// Start the server
	log.Printf("Starting server on: localhost:%d", shared.ListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", shared.ListenPort), mainRouter))
}
