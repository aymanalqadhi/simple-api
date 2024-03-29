package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/xSHAD0Wx/simple-api/models"
	"github.com/xSHAD0Wx/simple-api/services"
)

func getClientsRoute(ctx *fasthttp.RequestCtx) {
	clients, err := services.Services.ClientsRepo.GetAll()

	if err != nil {
		ctx.Error(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Could not fetch clients for: %s", ctx.RemoteAddr())
		return
	}

	ctx.Response.Header.Add("Content-Type", "application/json")
	json.NewEncoder(ctx).Encode(clients)
}

func addClientRoute(ctx *fasthttp.RequestCtx) {
	var client models.Client
	err := json.Unmarshal(ctx.PostBody(), &client)

	if err != nil {
		ctx.Error(http.StatusText(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		log.Printf("Bad request from %s", ctx.RemoteAddr())
		return
	}

	id, err := services.Services.ClientsRepo.Add(client)

	if err != nil {
		ctx.Error(http.StatusText(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		log.Printf("Could not add a client with name: %q, for: %s", client.Name, ctx.RemoteAddr())
	} else {
		log.Printf("A client was added, result id: %d", id)
		ctx.Response.SetStatusCode(fasthttp.StatusCreated)
		fmt.Fprint(ctx, id)
	}
}

func deleteClientRoute(ctx *fasthttp.RequestCtx) {

	id, err := strconv.ParseUint(ctx.UserValue("id").(string), 10, 0)
	if err != nil {
		ctx.Error(http.StatusText(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		log.Printf("Bad request from %s", ctx.RemoteAddr())
		return
	}

	err = services.Services.ClientsRepo.Remove(uint(id))
	if err != nil {
		ctx.Error(http.StatusText(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		log.Printf("Could not delete client with id: %d, for: %s", id, ctx.RemoteAddr())
	} else {
		log.Printf("A client was deleted, result id: %d", id)
	}
}

// GetClientsRouteGroup Gets the routes descriptions of the clients routes
func GetClientsRouteGroup() []HTTPRoute {
	return []HTTPRoute{
		HTTPRoute{
			Pattern:   "/clients",
			Handler:   getClientsRoute,
			UsesGet:   true,
			NeedsAuth: true,
		},
		HTTPRoute{
			Pattern:   "/clients",
			Handler:   addClientRoute,
			UsesPost:  true,
			NeedsAuth: true,
		},
		HTTPRoute{
			Pattern:    "/clients/:id",
			Handler:    deleteClientRoute,
			UsesDelete: true,
			NeedsAuth:  true,
		},
	}
}
