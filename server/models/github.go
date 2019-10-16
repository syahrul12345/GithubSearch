package models

//RepoRequestPayload is the incoming request made to the server to get all repo
//It contains only the username and will be provided by the caller
type RepoRequestPayload struct {
	Username string
}

//ReadmeRequestPayload is the incoming request made to the server to get data of a particular repo
type ReadmeRequestPayload struct {
	Username   string
	Repository string
}

//Repository represents one repository
type Repository struct {
	Name string
}

//User represents the user information
type User struct {
	PublicRepos int `"json:public_repos"`
}
