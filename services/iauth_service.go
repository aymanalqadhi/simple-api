package services

import "net/http"

// IAuthService is an interface represening an authentication service
type IAuthService interface {
	Authenticate(string, string) (string, error)
	ValidateToken(string) error
	AuthorizedHandler(func(http.ResponseWriter, *http.Request)) http.Handler
}
