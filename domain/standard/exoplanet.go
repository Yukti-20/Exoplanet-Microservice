package standard

import (
	"exoplanet-microservice/models"
	"github.com/google/uuid"
)

func (d *Domain) AddExoplanet(exoplanet models.Exoplanet) {
	exoplanet.ID = uuid.New().String() //Generating unique Id
	d.Exoplanets[exoplanet.ID] = exoplanet
}

func (d *Domain) ListExoplanets() []models.Exoplanet {
	exoplanets := make([]models.Exoplanet, 0, len(d.Exoplanets))
	for _, exoplanet := range d.Exoplanets {
		exoplanets = append(exoplanets, exoplanet)
	}
	return exoplanets
}
