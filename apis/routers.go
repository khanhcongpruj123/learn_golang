package main

import (
	"encoding/json"
	"example/core"
	"example/core/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create task")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get all task")
}

func login(w http.ResponseWriter, r *http.Request) {
	var u UserDTO
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.Random()

	core.Login(u.username, u.password)

	json.NewEncoder(w).Encode(u)
}

func AuthRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/api/v1/auth").Subrouter()
	r.HandleFunc("/login", login).Methods("POST")
	return r
}

func TaskRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/api/v1/tasks").Subrouter()
	r.HandleFunc("/", createTask).Methods("POST")
	r.HandleFunc("/", getTasks).Methods("GET")
	return r
}
