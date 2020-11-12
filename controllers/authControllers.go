package controllers

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

func SignIn(w http.ResponseWriter, r *http.Request) {
	var request models.WebRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	accessTokenString, refreshTokenString, err := models.CreateTokenPair(request.GUID)
	if err != nil {
		Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	request.AccessToken = accessTokenString
	request.RefreshToken = refreshTokenString

	var response models.WebResponse
	response.Payload["request"] = request

	Respond(w, response)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	var request models.WebRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	refreshTokenValid, message := models.IsValidToken(request.GUID, request.RefreshToken)
	if !refreshTokenValid {
		Respond(w, models.WebResponse{Message: message})
		return
	}

	accessTokenString, refreshTokenString, err := models.CreateTokenPair(request.GUID)
	if err != nil {
		Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	request.AccessToken = accessTokenString
	request.RefreshToken = refreshTokenString

	var response models.WebResponse
	response.Payload["request"] = request

	Respond(w, response)
}

func DeleteOneRefreshToken(w http.ResponseWriter, r *http.Request) {
	var request models.WebRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	accessTokenValid, message := models.IsValidToken(request.GUID, request.AccessToken)
	if !accessTokenValid {
		Respond(w, models.WebResponse{Message: message})
		return
	}

	Respond(w, models.WebResponse{Message: "Ok"})
}

func DeleteAllRefreshToken(w http.ResponseWriter, r *http.Request) {
	var request models.WebRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	accessTokenValid, message := models.IsValidToken(request.GUID, request.AccessToken)
	if !accessTokenValid {
		Respond(w, models.WebResponse{Message: message})
		return
	}

	Respond(w, models.WebResponse{Message: "Ok"})
}
