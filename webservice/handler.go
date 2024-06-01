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
	router.HandleFunc("/get/exoplanet/{id}", service.GetExoplanetById).Methods("GET")
	router.HandleFunc("/update/exoplanet/{id}", service.UpdateExoplanetById).Methods("POST")
	router.HandleFunc("/delete/exoplanet/{id}", service.DeleteExoplanetById).Methods("DELETE")
}

func NewAPIHandler(
	d domain.Service) *ExoplanetService {
	return &ExoplanetService{
		domain: d,
	}
}
