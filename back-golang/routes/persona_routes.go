package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/silviorosas/crud-golang-svelte/controller"
)

func SetupPersonaRoutes(r *mux.Router) {
	r.HandleFunc("/personas", controller.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/personas/{id}", controller.Get).Methods(http.MethodGet)
	r.HandleFunc("/personas", controller.Save).Methods(http.MethodPost)
	r.HandleFunc("/personas/{id}", controller.Delete).Methods(http.MethodDelete)
}
