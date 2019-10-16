package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server/controller"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sevlyar/go-daemon"
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()
	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")
	serve()
}

func serve() {
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
	router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("web page loaded")
		const staticPath = "../dist"
		const indexPath = "index.html"
		fileServer := http.FileServer(http.Dir(staticPath))
		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		path = filepath.Join(staticPath, path)
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			// file does not exist, serve index.html
			http.ServeFile(w, r, filepath.Join(staticPath, indexPath))
			return
		} else if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileServer.ServeHTTP(w, r)
	}))
	port := os.Getenv("PORT")
	fmt.Println("Listening on: http://localhost:", port)
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
