package standard

import (
	"errors"
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

func (d *Domain) GetExoplanetById(id string) (models.Exoplanet, bool) {
	exoplanet, exists := d.Exoplanets[id]
	return exoplanet, exists
}

func (d *Domain) UpdateExoplanetById(id string, exoplanet models.Exoplanet) bool {
	if _, exists := d.Exoplanets[id]; exists {
		exoplanet.ID = id
		d.Exoplanets[id] = exoplanet
		return true
	}
	return false
}

func (d *Domain) DeleteExoplanet(id string) bool {
	if _, exists := d.Exoplanets[id]; exists {
		delete(d.Exoplanets, id)
		return true
	}
	return false
}

func (d *Domain) GetFuelEstimation(exoplanet models.Exoplanet, crewCapacity int64) (float64, error) {
	var gravity, fuel float64
	if exoplanet.Type == "GasGiant" {
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	} else if exoplanet.Type == "Terrestrial" {
		gravity = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	} else {
		return fuel, errors.New("unknown exoplanet type")
	}

	fuel = float64(exoplanet.DistanceFromEarth) / (gravity * gravity) * float64(crewCapacity)
	return fuel, nil
}
