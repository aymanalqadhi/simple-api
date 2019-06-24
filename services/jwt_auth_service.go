package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

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

// ValidateToken validates a Jwt Token
func (auth *JwtAuthService) ValidateToken(string) error {
	return nil
}

// AuthorizedHandler returns a version that requires authorization of simple handler
func (auth *JwtAuthService) AuthorizedHandler(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Header["Authorization"] != nil {
			token, err := jwt.Parse(req.Header["Authorization"][0], func(*jwt.Token) (interface{}, error) {
				return []byte(shared.AuthPassword), nil
			})

			if err != nil || !token.Valid {
				res.WriteHeader(http.StatusUnauthorized)
			} else {
				h(res, req)
			}

		} else {
			res.WriteHeader(http.StatusUnauthorized)
		}
	})
}
