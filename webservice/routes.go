package webservice

import (
	"encoding/json"
	"exoplanet-microservice/models"
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
