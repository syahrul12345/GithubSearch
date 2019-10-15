package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"server/models"
	"strconv"
)

var (
	mainAPI   string = "https://api.github.com/users/"
	readmeAPI string = "https://raw.githubusercontent.com/"
)

//GetRepos will accept a string which is the username to be searched
func GetRepos(Username string) map[string]interface{} {
	resp := make(map[string]interface{})
	//Create the Oauth authenticated request
	clientID := fmt.Sprintf("?client_id=%s", os.Getenv("client_id"))
	clientSecret := fmt.Sprintf("&client_secret=%s", os.Getenv("client_secret"))
	url := mainAPI + Username + "/repos" + clientID + clientSecret
	page := 0

	//Asynchronous channel to handle writes
	dataChan := make(chan models.Repository)
	for page < 100 {
		pageString := strconv.Itoa(page)
		urlTemp := url + "&page=" + pageString
		go helper(urlTemp, dataChan)
		page += 1
	}

	//Rebuild the list of repos
	repositories := []models.Repository{}
	for elem := range dataChan {
		repositories = append(repositories, elem)
	}
	fmt.Println("END")
	fmt.Println(repositories)
	resp["results"] = repositories
	return resp
	//Since the URL can be paginated, we keep calling to the next page asynchronously,
	//This is done by incrementing the the page number and spawning one Go routine each
	// to increase performance

}

//helper function to make calls
func helper(url string, dataChan chan models.Repository) map[string]interface{} {
	//Resp is what we will send back to the frontend to parse
	resp := make(map[string]interface{})
	//Make the get response
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
	//If the page returns a 0 array, we dont append it
	if len(*repositories) > 0 {
		repo := *repositories
		for i := range repo {
			dataChan <- repo[i]
		}
	}
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
