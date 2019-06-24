package main

import (
	"log"

	"github.com/xSHAD0Wx/simple-api/services"

	"github.com/gorilla/mux"
	"github.com/xSHAD0Wx/simple-api/routes"
)

// ConfigureRouter configures a mux router
func ConfigureRouter(router *mux.Router) bool {
	router.StrictSlash(true)
	return true
}

// ConfigureRoutes configures all public app routes
func ConfigureRoutes(router *mux.Router) bool {
	// Get app routes
	routeGroups := [][]routes.HTTPRoute{
		routes.GetPingRouteGroup(),
		routes.GetClientsRouteGroup(),
		routes.GetLoginRouteGroup(),
	}

	// Map the routes
	for _, routeGroup := range routeGroups {
		for _, route := range routeGroup {
			log.Printf("Adding route %q...", route.Pattern)
			if route.NeedsAuth {
				router.Handle(route.Pattern, services.Services.AuthService.AuthorizedHandler(route.Handler)).Methods(route.Methods...)
			} else {
				router.HandleFunc(route.Pattern, route.Handler).Methods(route.Methods...)
			}
		}
	}

	return true
}
