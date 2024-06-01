package standard

import (
	"exoplanet-microservice/models"
	"github.com/google/uuid"
)

func (d *Domain) AddExoplanet(exoplanet models.Exoplanet) {
	exoplanet.ID = uuid.New().String() //Generating unique Id
	d.Exoplanets[exoplanet.ID] = exoplanet
}
