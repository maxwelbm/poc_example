package controller

import (
	"encoding/json"
	"net/http"
)

func respondJSON(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func handleError(w http.ResponseWriter, status int, message string) {
	body := &ResponseBodyProduct{
		Message: message,
		Data:    nil,
		Error:   true,
	}
	respondJSON(w, status, body)
}
