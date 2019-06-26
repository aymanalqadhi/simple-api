package routes

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func fastPingRoute(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "This site wont go down :D!\n")
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
