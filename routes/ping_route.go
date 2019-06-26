package routes

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func pingRoute(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "This site wont go down :D!\n")
}

// GetPingRouteGroup Gets the routes describtions of the ping routes
func GetPingRouteGroup() []HTTPRoute {
	return []HTTPRoute{
		HTTPRoute{
			Pattern: "/ping",
			Handler: pingRoute,
			UsesGet: true,
		},
	}
}
