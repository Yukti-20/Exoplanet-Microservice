package domain

import "exoplanet-microservice/models"

type exoplanet interface {
	AddExoplanet(exoplanet models.Exoplanet)
	ListExoplanets() []models.Exoplanet
	GetExoplanetById(id string) (models.Exoplanet, bool)
	UpdateExoplanetById(id string, exoplanet models.Exoplanet) bool
}
