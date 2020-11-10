package controllers

import (
	"../models"
	"../utils"
	"encoding/base64"
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
	request.RefreshToken = base64.StdEncoding.EncodeToString([]byte(refreshTokenString))

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

	accessTokenValid, message := models.IsValidAccessToken(request.AccessToken)
	if !accessTokenValid {
		utils.Respond(w, models.WebResponse{Message: message})
		return
	}

	tokenData, err := base64.StdEncoding.DecodeString(request.RefreshToken)
	if err != nil {
		utils.Respond(w, models.WebResponse{Message: err.Error()})
	}

	refreshTokenValid, message := models.IsValidRefreshToken(string(tokenData), request.AccessToken)
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
	request.RefreshToken = base64.StdEncoding.EncodeToString([]byte(refreshTokenString))

	var response models.WebResponse
	response.Payload["request"] = request

	utils.Respond(w, response)
}
