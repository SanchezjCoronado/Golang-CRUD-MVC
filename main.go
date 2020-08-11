package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SanchezjCoronado/Golang-CRUD-MVC/connection"
	"github.com/SanchezjCoronado/Golang-CRUD-MVC/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var prefixPath = "/api/compra"

//Encuentra una compra por su id
func FindShoppingByIDController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	shopping, err := connection.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, shopping)
}

func FindShoppingByUserController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idUser, err := strconv.Atoi(params["id_user"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	shoppings, err := connection.FindByUser(idUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error")
		return
	}
	respondWithJSON(w, http.StatusOK, shoppings)
}

//Crear una compra
func CreateShoppingController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shopping model.Shopping
	if err := json.NewDecoder(r.Body).Decode(&shopping); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	shopping.ID = bson.NewObjectId()
	if err := connection.Insert; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, shopping)
}

//Actualiza una compra
func UpdateShoppingController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shopping model.Shopping
	if err := json.NewDecoder(r.Body).Decode(&shopping); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.Update(shopping); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// Borrar una compra
func DeleteShoppingController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shoppingID model.ShoppingID
	if err := json.NewDecoder(r.Body).Decode(&shoppingID); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := connection.Delete(shoppingID.ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc(prefixPath, CreateShoppingController).Methods("POST")
	r.HandleFunc(prefixPath, UpdateShoppingController).Methods("PUT")
	r.HandleFunc(prefixPath, DeleteShoppingController).Methods("DELETE")
	r.HandleFunc(prefixPath+"/{id}", FindShoppingByIDController).Methods("GET")
	r.HandleFunc(prefixPath+"/by-user/{id_user}", FindShoppingByUserController).Methods("GET")

	log.Printf("Listening...")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
