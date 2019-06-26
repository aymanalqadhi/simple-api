package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/xSHAD0Wx/simple-api/shared"

	"github.com/dgrijalva/jwt-go"
)

// JwtAuthService is a service that impelments IAuthService interface
type JwtAuthService struct {
}

// CalimsByLevel is a map of the available claims
var CalimsByLevel = map[uint][]string{
	0: []string{"root", "clients", "ping"},
	1: []string{"clients", "ping"},
	2: []string{"ping"},
}

// Authenticate authenticates a user by username and password
func (auth *JwtAuthService) Authenticate(username string, password string) (string, error) {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	user, err := Services.UsersRepo.GetUser(username, hex.EncodeToString(hasher.Sum(nil)))

	if err != nil {
		return "", errors.New("invalid username or password")
	} else if user.AuthLevel > 2 {
		return "", errors.New("invalid user")
	}

	tokenGen := jwt.New(jwt.SigningMethodHS256)
	claims := tokenGen.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["expires"] = time.Now().Add(time.Hour * time.Duration(24*user.AuthDays)).Unix()

	for _, claim := range CalimsByLevel[user.AuthLevel] {
		claims[claim] = true
	}

	token, err := tokenGen.SignedString([]byte(shared.AuthPassword))
	if err != nil {
		return "", errors.New("could not generate a token")
	}

	return token, nil
}

// AuthorizedFastHandler returns a  fasthttp version that requires authorization of simple handler
func (auth *JwtAuthService) AuthorizedFastHandler(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		if ctx.Request.Header.Peek("Authorization") != nil {
			token, err := jwt.Parse(string(ctx.Request.Header.Peek("Authorization")), func(*jwt.Token) (interface{}, error) {
				return []byte(shared.AuthPassword), nil
			})

			if err != nil || !token.Valid {
				ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
			} else {
				h(ctx)
			}

		} else {
			ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
		}
	})
}
