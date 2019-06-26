package routes

import (
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
)

func pingRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This site wont go down :D!\n")
}

func fastPingRoute(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "This site wont go down :D!\n")
}

// GetPingRouteGroup Gets the routes describtions of the ping routes
func GetPingRouteGroup() []HTTPRoute {
	return []HTTPRoute{
		HTTPRoute{
			Pattern: "/ping",
			Methods: []string{HTTPGet},
			Handler: pingRoute,
		},
	}
}

// GetFastPingRouteGroup Gets the fasthttp routes describtions of the ping routes
func GetFastPingRouteGroup() []FastHTTPRoute {
	return []FastHTTPRoute{
		FastHTTPRoute{
			Pattern: "/ping",
			Handler: fastPingRoute,
			UsesGet: true,
		},
	}
}
