package webservice

import (
	"encoding/json"
	"exoplanet-microservice/models"
	"github.com/gorilla/mux"
	"net/http"
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

func (s *ExoplanetService) GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
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
