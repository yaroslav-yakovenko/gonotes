package main

import (
	"gonotes/gonotesserver/pkg/model"
	"gonotes/gonotesserver/pkg/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	db       *storage.MongoDB
	sessions []model.Session
)

func main() {

	// Init Database connection
	var err error
	db, err = storage.NewMongoDBClient()
	if err != nil {
		log.Fatal(err)
	}

	// Init web Router
	router := mux.NewRouter()

	// Init API endpoints
	setAPIEndpoints(router)

	// Starting web server with Gorilla Mux Router
	log.Fatal(http.ListenAndServe(":80", router))

}
