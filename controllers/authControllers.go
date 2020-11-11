package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var request models.WebRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	accessTokenString, refreshTokenString, err := models.CreateTokenPair(request.GUID)
	if err != nil {
		utils.Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	request.AccessToken = accessTokenString
	request.RefreshToken = refreshTokenString

	var response models.WebResponse
	response.Payload["request"] = request

	utils.Respond(w, response)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	var request models.WebRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	refreshTokenValid, message := models.IsValidToken(request.RefreshToken)
	if !refreshTokenValid {
		utils.Respond(w, models.WebResponse{Message: message})
		return
	}

	accessTokenString, refreshTokenString, err := models.CreateTokenPair(request.GUID)
	if err != nil {
		utils.Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	request.AccessToken = accessTokenString
	request.RefreshToken = refreshTokenString

	var response models.WebResponse
	response.Payload["request"] = request

	utils.Respond(w, response)
}

func DeleteOneRefreshToken(w http.ResponseWriter, r *http.Request) {
	var request models.WebRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	accessTokenValid, message := models.IsValidToken(request.AccessToken)
	if !accessTokenValid {
		utils.Respond(w, models.WebResponse{Message: message})
		return
	}

	utils.Respond(w, models.WebResponse{Message: "Ok"})
}

func DeleteAllRefreshToken(w http.ResponseWriter, r *http.Request) {
	var request models.WebRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Respond(w, models.WebResponse{Message: err.Error()})
		return
	}

	accessTokenValid, message := models.IsValidToken(request.AccessToken)
	if !accessTokenValid {
		utils.Respond(w, models.WebResponse{Message: message})
		return
	}

	utils.Respond(w, models.WebResponse{Message: "Ok"})
}
