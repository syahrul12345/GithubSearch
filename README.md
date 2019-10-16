# GithubSearch
Search github repositories by username and obtain the read me

## Getting Started
These instructions will help you get started to build your own local version of github serach

## Prerequisites
You'll need to install these dependencies first:
```
Golang (Go v1.11 and up ONLY)
Node & NPM
```
## Installing
First clone the repo:
```
git clone https://github.com/syahrul12345/githubsearch
```
## Build the frontend:
This project is built using vuejs. We build the frontend files, and it will be served by the golang server.
```
cd githubsearch
cd frontend
npm install
npm run build
```
## Building the backend:
First, you'll need to set up the environment files. To prevent rate limiting, the backend needs to be authenticated as a github OAuth app. You can get your client_id and client_secret
[here](https://github.com/settings/developers)

Enter the server folder and create a new `.env` file by modifying the sample.env file.
```
cd server
mv sample.env .env
```
You can now modify the `.env` in the following format. Make sure to change all the environment variables!
```
PORT= CHANGE_TO_YOUR_PORT
client_id = YOUR_CLIENT_ID
client_secret = YOUR_BACKEND_ID
```
Now you can build the server application. Please ensure that you have Go (v1.11) and above!

```
go build
```
While in the `server` folder, you can execute the server application by typing:
```
./server
```
This will serve the html files built using VueJS at `http://127.0.0.1/CHANGE_TO_YOUR_PORTT`.
## Testing
All tests live in the `tests` folder.
Tests are split into backend and frontend accordingly