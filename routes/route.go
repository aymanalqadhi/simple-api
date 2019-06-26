package routes

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

const (
	// HTTPGet holds the Http GET method string
	HTTPGet = "GET"
	// HTTPPost holds the Http POST method string
	HTTPPost = "POST"
	// HTTPDelete holds the Http DELETE method string
	HTTPDelete = "DELETE"
	// HTTPPut holds the Http PUT method string
	HTTPPut = "PUT"
)

// HTTPRoute is a struct to represent an HTTP route
type HTTPRoute struct {
	Pattern   string
	Methods   []string
	Handler   func(http.ResponseWriter, *http.Request)
	NeedsAuth bool
}

// FastHTTPRoute is a route type to work with fasthttp framework
type FastHTTPRoute struct {
	Pattern   string
	Handler   func(*fasthttp.RequestCtx)
	NeedsAuth bool

	UsesGet    bool
	UsesPost   bool
	UsesDelete bool
	UsesPut    bool
}
