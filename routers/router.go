package routers

import (
	"be-library-go/handlers"

	"github.com/gorilla/mux"
)

func setUpRouter() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/categories", handlers.UpdateCategories).Methods("GET")
	router.HandleFunc("/categories/{id}",handlers.CreateCategories ).Methods("POST")
		router.HandleFunc("/categories/{id}",handlers.UpdateCategories ).Methods("PUT")
			router.HandleFunc("/categories/{id}",handlers.DeleteCategories ).Methods("DELETE")
			
return router
}