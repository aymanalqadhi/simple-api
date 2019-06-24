package routes

import (
	"fmt"
	"net/http"

	"github.com/xSHAD0Wx/simple-api/services"
)

func loginRoute(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	username := req.Form.Get("username")
	password := req.Form.Get("password")

	fmt.Println(username)
	fmt.Println(password)
	if username == "" || password == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := services.Services.AuthService.Authenticate(username, password)
	if err != nil {
		http.Error(res, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	fmt.Fprint(res, token)
}

func registerRoute(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	username := req.Form.Get("username")
	email := req.Form.Get("email")
	password := req.Form.Get("password")

	if username == "" || email == "" || password == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := services.Services.UsersRepo.AddUser(username, password, email, 2); err != nil {
		http.Error(res, "Could not register user", http.StatusInternalServerError)
	} else {
		res.WriteHeader(http.StatusCreated)
	}
}

// GetLoginRouteGroup Gets the routes describtions of the login routes
func GetLoginRouteGroup() []HTTPRoute {
	return []HTTPRoute{
		HTTPRoute{
			Pattern: "/login",
			Methods: []string{HTTPPost},
			Handler: loginRoute,
		},
		HTTPRoute{
			Pattern: "/register",
			Methods: []string{HTTPPost},
			Handler: registerRoute,
		},
	}
}
