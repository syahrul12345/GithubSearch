package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/controller"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	//Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Create a new router
	router := mux.NewRouter()

	//API CALLS

	//Get all repos of a user
	router.HandleFunc("/api/v1/getUser", controller.GetUser).Methods("POST")
	//Get the readme of one repo from one user
	router.HandleFunc("/api/v1/getRepo", controller.GetReadme).Methods("POST")

	port := os.Getenv("PORT")
	fmt.Println(port)
	//For Testing, disabled cors for 8080
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	//Listen to calls to port 9999
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Println(err)
	}

}
