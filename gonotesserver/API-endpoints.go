package main

func setAPIEndpoints() {
	router.HandleFunc("/api/v1/getCategories", getCategoriesHandler).Methods("GET", "OPTIONS")
}
