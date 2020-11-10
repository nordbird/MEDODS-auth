package utils

import (
	"../models"
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data models.WebResponse) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
