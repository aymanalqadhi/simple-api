package services

import (
	"github.com/valyala/fasthttp"
)

// IAuthService is an interface represening an authentication service
type IAuthService interface {
	Authenticate(string, string) (string, error)
	AuthorizedFastHandler(fasthttp.RequestHandler) fasthttp.RequestHandler
}
