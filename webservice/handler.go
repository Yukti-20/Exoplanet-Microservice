package webservice

import (
	"exoplanet-microservice/domain"
	"exoplanet-microservice/models"
	"github.com/gorilla/mux"
)

type ExoplanetService struct {
	Exoplanets map[string]models.Exoplanet
	domain     domain.Service
}

func RegisterRoutes(router *mux.Router, service *ExoplanetService) {
	router.HandleFunc("/add/exoplanets", service.AddExoplanet).Methods("POST")
	router.HandleFunc("/list/exoplanets", service.ListExoplanets).Methods("GET")
}

func NewAPIHandler(
	d domain.Service) *ExoplanetService {
	return &ExoplanetService{
		domain: d,
	}
}
