package webservice

import (
	"exoplanet-microservice/models"
)

type ExoplanetService struct {
	exoplanets map[string]models.Exoplanet
}

func NewExoplanetService() *ExoplanetService {
	return &ExoplanetService{
		exoplanets: make(map[string]models.Exoplanet),
	}
}
