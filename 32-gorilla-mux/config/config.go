package config

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	c "gorilla-mux/controllers"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", c.HomePage).Methods("GET")
	myRouter.HandleFunc("/all", c.ReturnAll).Methods("GET")
	myRouter.HandleFunc("/article/{id}", c.ReturnSingleArticle)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
