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
		writer.WriteHeader(http.StatusInternalServerError)
		utils.Respond(writer, utils.Message(false, "Request payload deformed, should contain a STRING in the username field of JSON"))
		return
	}
	//Handle an empty username provided
	if requestPayload.Username == "" {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.Respond(writer, utils.Message(false, "Request payload does not contain username string"))
		return
	}
	resp := request.GetRepos(requestPayload.Username)

	//The request module will have an error object if it was unable to fetch the data from github API
	if resp["error"] != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.Respond(writer, utils.Message(false, "Failed to get userdata, it might not exist"))
		return
	}
	utils.Respond(writer, request.GetRepos(requestPayload.Username))
}

//GetReadme is called to return one repository from the user
var GetReadme = func(writer http.ResponseWriter, req *http.Request) {
	requestPayload := &models.ReadmeRequestPayload{}
	err := json.NewDecoder(req.Body).Decode(requestPayload)
	//Handle malformed requests
	//For ever error set the status code
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.Respond(writer, utils.Message(false, "Request payload deformed, should contain a STRING in the username & repo field of JSON"))
		return
	}
	//Handle an empty username provided
	if requestPayload.Username == "" {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.Respond(writer, utils.Message(false, "Request payload to get repo does not contain username string"))
		return
	}
	//Handle emoty repo name provider
	if requestPayload.Repository == "" {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.Respond(writer, utils.Message(false, "Request payload to get repo does not contain repository string"))
		return
	}
	resp := request.GetReadme(requestPayload.Username, requestPayload.Repository)
	if resp["error"] != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		utils.Respond(writer, utils.Message(false, "Incorrect username or repository provided"))
		return
	}
	utils.Respond(writer, request.GetReadme(requestPayload.Username, requestPayload.Repository))
}
