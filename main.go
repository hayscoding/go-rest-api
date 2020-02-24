package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux" // Unused packages will create compilation error
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(){
	// := is the short variable declaration operator
	// Automatically determines type for variable
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}