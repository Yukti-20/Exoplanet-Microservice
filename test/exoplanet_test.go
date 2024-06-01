package test

import (
	"exoplanet-microservice/domain/standard"
	"exoplanet-microservice/models"
	"exoplanet-microservice/webservice"
	"fmt"
	"github.com/gorilla/mux"
	"testing"
)

func TestAddExoplanet(t *testing.T) {
	domain := standard.NewExoplanetDomain()
	service := webservice.NewAPIHandler(domain)
	router := mux.NewRouter()
	webservice.RegisterRoutes(router, service)

	exoplanet := models.Exoplanet{
		Name:              "Test Planet A",
		Description:       "This exoplanet is a fascinating celestial body",
		DistanceFromEarth: 500,
		Radius:            1.5,
		Mass:              1.0,
		Type:              "GasGiant",
	}
	resp := domain.AddExoplanet(exoplanet)
	fmt.Println("The response on adding exoplanet is: ", resp)
}

func TestListExoplanets(t *testing.T) {
	domain := standard.NewExoplanetDomain()
	service := webservice.NewAPIHandler(domain)
	router := mux.NewRouter()
	webservice.RegisterRoutes(router, service)

	resp := domain.ListExoplanets()
	fmt.Println("The list of exoplanets is: ", resp)
}

func TestGetExoplanetByID(t *testing.T) {
	domain := standard.NewExoplanetDomain()
	service := webservice.NewAPIHandler(domain)
	router := mux.NewRouter()
	webservice.RegisterRoutes(router, service)

	resp, ok := domain.GetExoplanetById("5d1fbecc-d650-4bec-8435-e2bac94d2c7f")
	if ok {
		fmt.Println("The response by Id is: ", resp)
	} else {
		fmt.Println("There is no exoplanet with the given Id")
	}
}

func TestUpdateExoplanet(t *testing.T) {
	domain := standard.NewExoplanetDomain()
	service := webservice.NewAPIHandler(domain)
	router := mux.NewRouter()
	webservice.RegisterRoutes(router, service)

	exoplanet := models.Exoplanet{
		Name:              "Test Planet",
		Description:       "A mysterious planet.",
		DistanceFromEarth: 100,
		Radius:            1.5,
		Mass:              1.0,
		Type:              "Terrestrial",
	}
	ok := domain.UpdateExoplanetById("5d1fbecc-d650-4bec-8435-e2bac94d2c7f", exoplanet)
	if ok {
		fmt.Println("The exoplanet updated successfully")
	}

}

func TestDeleteExoplanetById(t *testing.T) {
	domain := standard.NewExoplanetDomain()
	service := webservice.NewAPIHandler(domain)
	router := mux.NewRouter()
	webservice.RegisterRoutes(router, service)

	ok := domain.DeleteExoplanet("9ca6bed7-0c03-4dc5-9259-3e04685fabd5")
	if ok {
		fmt.Println("The exoplanet deleted successfully")
	}
}

func TestFuelEstimation(t *testing.T) {
	domain := standard.NewExoplanetDomain()
	service := webservice.NewAPIHandler(domain)
	router := mux.NewRouter()
	webservice.RegisterRoutes(router, service)

	exoplanet := models.Exoplanet{
		ID:                "5d1fbecc-d650-4bec-8435-e2bac94d2c7f",
		Name:              "Test Planet A",
		Description:       "This exoplanet is a fascinating celestial body",
		DistanceFromEarth: 500,
		Radius:            1.5,
		Mass:              1.0,
		Type:              "GasGiant",
	}

	fuel, err := domain.GetFuelEstimation(exoplanet, 12)
	if err != nil {
		fmt.Println("The is unknown exoplanet type")
	}
	fmt.Println("The fuel estimation is: ", fuel)
}
