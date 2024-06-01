package webservice

import (
	"encoding/json"
	"exoplanet-microservice/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *ExoplanetService) AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.domain.AddExoplanet(exoplanet)

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(exoplanet)
	if err != nil {
		return
	}
}

func (s *ExoplanetService) ListExoplanets(w http.ResponseWriter, r *http.Request) {
	exoplanets := s.domain.ListExoplanets()
	err := json.NewEncoder(w).Encode(exoplanets)
	if err != nil {
		return
	}
}

func (s *ExoplanetService) GetExoplanetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	exoplanet, exists := s.domain.GetExoplanetById(id)
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(exoplanet)
	if err != nil {
		return
	}
}

func (s *ExoplanetService) UpdateExoplanetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var exoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !s.domain.UpdateExoplanetById(id, exoplanet) {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(exoplanet)
	if err != nil {
		return
	}
}

func (s *ExoplanetService) DeleteExoplanetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if !s.domain.DeleteExoplanet(id) {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *ExoplanetService) FuelEstimation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	exoplanet, exists := s.domain.GetExoplanetById(id)
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	crewCapacity, err := strconv.Atoi(r.URL.Query().Get("crew"))
	if err != nil {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}

	fuel, typeErr := s.domain.GetFuelEstimation(exoplanet, int64(crewCapacity))
	if typeErr != nil {
		http.Error(w, typeErr.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]float64{"Fuel Estimation ": fuel})
	if err != nil {
		return
	}
}
