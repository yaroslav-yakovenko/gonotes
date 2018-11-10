package main

import (
	"gonotes/gonotesserver/pkg/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	err    error
	db     *storage.MongoDB
	router *mux.Router
)

func main() {

	// Init Database connection
	db, err = storage.NewMongoDBClient()
	if err != nil {
		log.Fatal(err)
	}

	// Init web Router
	router = mux.NewRouter()

	// Init API endpoints
	setAPIEndpoints()

	// Starting web server with Gorilla Mux Router
	log.Fatal(http.ListenAndServe("localhost:10000", router))

	/*doc := model.Category{
		Name:        "Data Structures",
		Description: "Well-known Data Structures and Patterns",
	}
	_ = doc
	err = db.AddCategory(doc)
	if err != nil {
		fmt.Println(err)
	}*/

	/*categories, err := db.GetCategories()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(categories)*/

}
