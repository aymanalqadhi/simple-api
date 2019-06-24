package routes

import (
	"fmt"
	"net/http"
)

func pingRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This site wont go down :D!\n")
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
