package webservice

import (
	"exoplanet-microservice/domain"
	"exoplanet-microservice/models"
	"github.com/gorilla/mux"
	"net/http"
)

type ExoplanetService struct {
	Exoplanets map[string]models.Exoplanet
	domain     domain.Service
}

func RegisterRoutes(router *mux.Router, service *ExoplanetService) {
	router.HandleFunc("/add/exoplanets", service.AddExoplanet).Methods(http.MethodPost)
	router.HandleFunc("/list/exoplanets", service.ListExoplanets).Methods(http.MethodGet)
	router.HandleFunc("/get/exoplanet/{id}", service.GetExoplanetById).Methods(http.MethodGet)
	router.HandleFunc("/update/exoplanet/{id}", service.UpdateExoplanetById).Methods(http.MethodPost)
	router.HandleFunc("/delete/exoplanet/{id}", service.DeleteExoplanetById).Methods(http.MethodDelete)
	router.HandleFunc("/fuel/estimation/{id}", service.FuelEstimation).Methods(http.MethodGet)
}

func NewAPIHandler(
	d domain.Service) *ExoplanetService {
	return &ExoplanetService{
		domain: d,
	}
}
