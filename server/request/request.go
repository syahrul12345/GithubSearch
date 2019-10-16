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
	userInfo := getRepoCount(Username)
	fmt.Println(userInfo["count"])
	page := 3
	//Repositories is the golang struct
	repositories := []models.Repository{}
	//Asynchronous channel to handle writes
	dataChan := make(chan *[]models.Repository)
	for i := 1; i < page+1; i++ {
		pageString := strconv.Itoa(i)
		urlTemp := url + "&page=" + pageString
		//Helpers get the respons from the API
		go helper(urlTemp, dataChan)
		c := <-dataChan
		repositories = append(repositories, *c...)
		fmt.Println(i)
		//Ah is a channel, we have to force iterate it
	}
	//Rebuild the list of repos
	resp["results"] = repositories
	return resp
	//Since the URL can be paginated, we keep calling to the next page asynchronously,
	//This is done by incrementing the the page number and spawning one Go routine each
	// to increase performance
}

//helper function to make calls
func helper(url string, dataChan chan *[]models.Repository) map[string]interface{} {
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
	//If the page returns a 0 array, we dont append it to the channel
	dataChan <- repositories
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

//getRepoCount will get the number of repos owned by the user
func getRepoCount(Username string) map[string]interface{} {
	resp := make(map[string]interface{})
	clientID := fmt.Sprintf("?client_id=%s", os.Getenv("client_id"))
	clientSecret := fmt.Sprintf("&client_secret=%s", os.Getenv("client_secret"))
	userURL := mainAPI + Username + clientID + clientSecret
	response, responseErr := http.Get(userURL)
	fmt.Println(userURL)
	//Handle if it's not a 200 first. The error handling for other endpoints is a failsafe
	if response.StatusCode != 200 {
		resp["error"] = "Failed to get user info"
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

	//parse the return object
	user := &models.User{}
	parsedError := json.Unmarshal(body, &user)
	if parsedError != nil {
		resp["error"] = parsedError
		return resp
	}
	resp["count"] = user.PublicRepos
	return resp

}
