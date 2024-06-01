package main

import (
	"exoplanet-microservice/domain/standard"
	"exoplanet-microservice/webservice"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	stdDomain := standard.NewExoplanetDomain()
	handlerInstance := webservice.NewAPIHandler(stdDomain)
	webservice.RegisterRoutes(router, handlerInstance)
	log.Fatal(http.ListenAndServe(":8080", router))
}
