package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/xSHAD0Wx/simple-api/routes"
	"github.com/xSHAD0Wx/simple-api/services"
)

// ConfigureFastRouter configures the fasthttp router
func ConfigureFastRouter(router *fasthttprouter.Router) bool {
	return true
}

// ConfigureFastRoutes configures the fasthttp routes
func ConfigureFastRoutes(router *fasthttprouter.Router) bool {
	// Get app routes
	routeGroups := [][]routes.HTTPRoute{
		routes.GetPingRouteGroup(),
		routes.GetLoginRouteGroup(),
		routes.GetClientsRouteGroup(),
	}

	regFastRoute := func(reg func(string, fasthttp.RequestHandler), route routes.HTTPRoute, value bool) {
		if value {
			reg(route.Pattern, route.Handler)
		}
	}

	// Map the routes
	for _, routeGroup := range routeGroups {
		for _, route := range routeGroup {
			log.Printf("Adding route %q...", route.Pattern)

			if route.NeedsAuth {
				route.Handler = services.Services.AuthService.AuthorizedFastHandler(route.Handler)
			}

			regFastRoute(router.GET, route, route.UsesGet)
			regFastRoute(router.POST, route, route.UsesPost)
			regFastRoute(router.DELETE, route, route.UsesDelete)
			regFastRoute(router.PUT, route, route.UsesPut)
		}
	}

	return true
}
