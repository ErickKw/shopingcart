package main

import (
	"gofibershop/controllers/authcontrollers"
	"gofibershop/controllers/productcontrollers"
	"gofibershop/middleware"
	"gofibershop/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontrollers.Login).Methods("POST")
	r.HandleFunc("/register", authcontrollers.Register).Methods("POST")
	r.HandleFunc("/logout", authcontrollers.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontrollers.Index).Methods("GET")
	api.Use(middleware.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
