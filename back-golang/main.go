package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/silviorosas/crud-golang-svelte/commons"
	"github.com/silviorosas/crud-golang-svelte/routes"
	// Asegúrate de importar tus paquetes necesarios
)

func main() {
	// Realiza las migraciones necesarias
	commons.Migrate()

	// Configura tus rutas con mux
	router := mux.NewRouter()
	routes.SetupPersonaRoutes(router)

	// Configura CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Cambia esta URL por la URL de tu aplicación Svelte
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// Aplica el manejador CORS a todas las rutas
	handler := corsHandler.Handler(router)

	// Crea un servidor HTTP
	server := &http.Server{
		Addr:    ":9000",
		Handler: handler, // Utiliza el manejador CORS para todas las rutas
	}

	// Inicia el servidor y muestra cualquier error que ocurra
	log.Println("Servidor ejecutándose en el puerto 9000")
	log.Println(server.ListenAndServe())
}
