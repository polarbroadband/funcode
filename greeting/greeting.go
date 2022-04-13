package main

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	log.Infof("GET request")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(map[string]string{"from": "T&T", "status": "ALL GOOD"}); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
