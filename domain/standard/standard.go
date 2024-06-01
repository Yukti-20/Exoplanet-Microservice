package standard

import "exoplanet-microservice/models"

type Domain struct {
	Exoplanets map[string]models.Exoplanet
}

func NewExoplanetDomain() *Domain {
	return &Domain{
		Exoplanets: make(map[string]models.Exoplanet),
	}
}
