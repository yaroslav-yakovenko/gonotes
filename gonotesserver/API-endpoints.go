package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setAPIEndpoints(router *mux.Router) {

	router.HandleFunc("/api/v1/getCategories", getCategoriesHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/addCategory", addCategoryHandler).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/v1/getTags", getTagsHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/addTag", addTagHandler).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/v1/getNotes", getNotesHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/getNotesByCategory", getNotesByCategoryHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/getNotesByTag", getNotesByTagHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/addNote", addNoteHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/updateNote", updateNoteHandler).Methods("POST", "OPTIONS")

	// serving VueJS SPA
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./dist"))))

}
