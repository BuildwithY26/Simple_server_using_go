package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to Marshal response")
		w.WriteHeader(500)
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}