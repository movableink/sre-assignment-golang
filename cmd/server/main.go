package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/movableink/sre-assignment-golang/internal/config"
	"github.com/movableink/sre-assignment-golang/internal/geoip"
)

type errorResponse struct {
	Error string `json:"error"`
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	geoipService := geoip.New(cfg)
	router := mux.NewRouter()

	router.HandleFunc("/lookup/{ip}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ip := vars["ip"]

		result, err := geoipService.LookupIP(ip)
		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(result)
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Server starting on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
