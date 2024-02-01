package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/silviorosas/crud-golang-svelte/commons"
	"github.com/silviorosas/crud-golang-svelte/models"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConexion()
	defer db.Close()

	db.Find(&personas)

	json, _ := json.Marshal(personas)
	commons.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	id := mux.Vars(request)["id"]
	db := commons.GetConexion()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	db := commons.GetConexion()
	defer db.Close()
	error := json.NewDecoder(request.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	// Verificar si ya existe un registro con el mismo correo electrónico
	var existing models.Persona
	if err := db.Where("email = ?", persona.Email).First(&existing).Error; err == nil {
		// Ya existe un registro con este correo electrónico, devolver error
		errorMessage := map[string]string{"error": "Este correo electrónico ya está registrado"}
		jsonError, _ := json.Marshal(errorMessage)
		commons.SendResponse(writer, http.StatusConflict, jsonError)
		return
	}

	// No existe un registro con este correo electrónico, se puede guardar
	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)
	commons.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	id := mux.Vars(request)["id"]
	db := commons.GetConexion()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
