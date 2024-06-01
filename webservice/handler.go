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
	router.HandleFunc("/exoplanets", service.AddExoplanet).Methods("POST")
}

func NewAPIHandler(
	d domain.Service) *ExoplanetService {
	return &ExoplanetService{
		domain: d,
	}
}
