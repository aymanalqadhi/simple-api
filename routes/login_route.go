package routes

import (
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"

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

func fastLoginRoute(ctx *fasthttp.RequestCtx) {
	if !ctx.PostArgs().Has("username") || !ctx.PostArgs().Has("password") {
		ctx.Error("Username and password are required", fasthttp.StatusBadRequest)
		return
	}

	username := string(ctx.PostArgs().Peek("username"))
	password := string(ctx.PostArgs().Peek("password"))

	if username == "" || password == "" {
		ctx.Error("Username and password are required", fasthttp.StatusBadRequest)
		return
	}

	token, err := services.Services.AuthService.Authenticate(username, password)
	if err != nil {
		ctx.Error("Invalid username or password", fasthttp.StatusUnauthorized)
		return
	}

	fmt.Fprint(ctx, token)
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

func fastRegisterRoute(ctx *fasthttp.RequestCtx) {
	if !ctx.PostArgs().Has("username") || !ctx.PostArgs().Has("password") {
		ctx.Error("Username and password are required", fasthttp.StatusBadRequest)
		return
	}

	username := string(ctx.PostArgs().Peek("username"))
	password := string(ctx.PostArgs().Peek("password"))
	email := string(ctx.PostArgs().Peek("email"))

	if username == "" || email == "" || password == "" {
		ctx.Error("Empty values are not allowed", fasthttp.StatusBadRequest)
		return
	}

	if err := services.Services.UsersRepo.AddUser(username, password, email, 2); err != nil {
		ctx.Error("Could not register user", fasthttp.StatusInternalServerError)
	} else {
		ctx.Response.SetStatusCode(fasthttp.StatusCreated)
	}
}

// GetLoginRouteGroup Gets the routes descriptions of the login routes
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

// GetFastLoginRouteGroup Gets the fasthttp routes descriptions of the login routes
func GetFastLoginRouteGroup() []FastHTTPRoute {
	return []FastHTTPRoute{
		FastHTTPRoute{
			Pattern:  "/login",
			UsesPost: true,
			Handler:  fastLoginRoute,
		},
		FastHTTPRoute{
			Pattern:  "/register",
			UsesPost: true,
			Handler:  fastRegisterRoute,
		},
	}
}
