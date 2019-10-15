package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/models"
)

var (
	mainAPI   string = "https://api.github.com/users/"
	readmeAPI string = "https://raw.githubusercontent.com/"
)

//GetRepos will accept a string which is the username to be searched
func GetRepos(Username string) map[string]interface{} {
	//Resp is what we will send back to the frontend to parse
	resp := make(map[string]interface{})
	url := mainAPI + Username + "/repos"
	response, responseErr := http.Get(url)

	//Handle Wrong URL
	if responseErr != nil {
		resp["error"] = responseErr
		return resp
	}
	//Close body to prevent leakages
	defer response.Body.Close()

	//Read the response
	body, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		resp["error"] = bodyErr
		return resp
	}
	//Body returns an Array
	//Parse the response
	repositories := &[]models.Repository{}
	parsedErr := json.Unmarshal(body, &repositories)
	if parsedErr != nil {
		resp["error"] = parsedErr
		return resp
	}
	resp["repositories"] = repositories
	return resp
}

//GetReadme returns the readme of one repo belonging to a user
func GetReadme(username string, repo string) map[string]interface{} {
	// ReadmeAPI should have this appended: syahrul12345/Battlecards/master/README.md
	resp := make(map[string]interface{})
	readmeURL := readmeAPI + username + "/" + repo + "/" + "/master/README.md"
	response, responseErr := http.Get(readmeURL)

	//Handle if it's not a 200 first. The error handling for other endpoints is a failsafe
	if response.StatusCode != 200 {
		resp["error"] = "Failed to get readme"
		return resp
	}
	//Handle Wrong URL
	if responseErr != nil {
		resp["error"] = responseErr
		return resp
	}
	//Close body to prevent leakages
	defer response.Body.Close()

	//Read the response
	body, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		resp["error"] = responseErr
		return resp
	}
	//Body returns a string
	resp["readme"] = string(body)
	return resp
}
