package domain

import "exoplanet-microservice/models"

type exoplanet interface {
	AddExoplanet(exoplanet models.Exoplanet)
}
