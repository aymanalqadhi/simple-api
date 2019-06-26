package routes

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/xSHAD0Wx/simple-api/services"
)

func loginRoute(ctx *fasthttp.RequestCtx) {
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

func registerRoute(ctx *fasthttp.RequestCtx) {
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

// GetLoginRouteGroup Gets the fasthttp routes descriptions of the login routes
func GetLoginRouteGroup() []FastHTTPRoute {
	return []FastHTTPRoute{
		FastHTTPRoute{
			Pattern:  "/login",
			UsesPost: true,
			Handler:  loginRoute,
		},
		FastHTTPRoute{
			Pattern:  "/register",
			UsesPost: true,
			Handler:  registerRoute,
		},
	}
}
