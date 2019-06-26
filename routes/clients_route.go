package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/gorilla/mux"

	"github.com/xSHAD0Wx/simple-api/models"
	"github.com/xSHAD0Wx/simple-api/services"
)

func getClientsRoute(w http.ResponseWriter, r *http.Request) {
	clients, err := services.Services.ClientsRepo.GetAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Could not fetch clients for: %s", r.RemoteAddr)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func fastGetClientsRoute(ctx *fasthttp.RequestCtx) {
	clients, err := services.Services.ClientsRepo.GetAll()

	if err != nil {
		ctx.Error(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Could not fetch clients for: %s", ctx.RemoteAddr())
		return
	}

	ctx.Response.Header.Add("Content-Type", "application/json")
	json.NewEncoder(ctx).Encode(clients)
}

func addClientRoute(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Bad request from %s", r.RemoteAddr)
		return
	}

	id, err := services.Services.ClientsRepo.Add(client)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Could not add a client with name: %q, for: %s", client.Name, r.RemoteAddr)
	} else {
		log.Printf("A client was added, result id: %d", id)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, id)
	}
}

func fastAddClientRoute(ctx *fasthttp.RequestCtx) {
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

func deleteClientRoute(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 0)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Bad request from %s", r.RemoteAddr)
		return
	}

	err = services.Services.ClientsRepo.Remove(uint(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Could not delete client with id: %d, for: %s", id, r.RemoteAddr)
	} else {
		log.Printf("A client was deleted, result id: %d", id)
	}
}

func fastDeleteClientRoute(ctx *fasthttp.RequestCtx) {

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
			Methods:   []string{HTTPGet},
			Handler:   getClientsRoute,
			NeedsAuth: true,
		},
		HTTPRoute{
			Pattern:   "/clients",
			Methods:   []string{HTTPPost},
			Handler:   addClientRoute,
			NeedsAuth: true,
		},
		HTTPRoute{
			Pattern:   "/clients/{id}",
			Methods:   []string{HTTPDelete},
			Handler:   deleteClientRoute,
			NeedsAuth: true,
		},
	}
}

// GetFastClientsRouteGroup Gets the fasthttp routes descriptions of the clients routes
func GetFastClientsRouteGroup() []FastHTTPRoute {
	return []FastHTTPRoute{
		FastHTTPRoute{
			Pattern:   "/clients",
			Handler:   fastGetClientsRoute,
			UsesGet:   true,
			NeedsAuth: true,
		},
		FastHTTPRoute{
			Pattern:   "/clients",
			Handler:   fastAddClientRoute,
			UsesPost:  true,
			NeedsAuth: true,
		},
		FastHTTPRoute{
			Pattern:    "/clients/:id",
			Handler:    fastDeleteClientRoute,
			UsesDelete: true,
			NeedsAuth:  true,
		},
	}
}
