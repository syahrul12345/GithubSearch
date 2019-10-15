package controller

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/request"
	"server/utils"
)

//GetUser is called to get the repositories of the user.
var GetUser = func(writer http.ResponseWriter, req *http.Request) {
	requestPayload := &models.RepoRequestPayload{}
	err := json.NewDecoder(req.Body).Decode(requestPayload)
	//Handle malformed requests
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Request payload deformed, should contain a STRING in the username field of JSON"))
		return
	}
	//Handle an empty username provided
	if requestPayload.Username == "" {
		utils.Respond(writer, utils.Message(false, "Request payload does not contain username string"))
		return
	}
	utils.Respond(writer, request.GetRepos(requestPayload.Username))
}

//GetReadme is called to return one repository from the user
var GetReadme = func(writer http.ResponseWriter, req *http.Request) {
	requestPayload := &models.ReadmeRequestPayload{}
	err := json.NewDecoder(req.Body).Decode(requestPayload)
	//Handle malformed requests
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Request payload deformed, should contain a STRING in the username & repo field of JSON"))
		return
	}
	//Handle an empty username provided
	if requestPayload.Username == "" {
		utils.Respond(writer, utils.Message(false, "Request payload does not contain username string"))
		return
	}
	//Handle emoty repo name provider
	if requestPayload.Repository == "" {
		utils.Respond(writer, utils.Message(false, "Request payload does not contain repository string"))
		return
	}

	utils.Respond(writer, request.GetReadme(requestPayload.Username, requestPayload.Repository))
}
