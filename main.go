package main

import (
	"exoplanet-microservice/webservice"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	service := webservice.NewExoplanetService()
	registerRoutes(router, service)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerRoutes(router *mux.Router, service *webservice.ExoplanetService) {
}
