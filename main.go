package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	// Type go get -u github.com/gorilla/mux to install
	"github.com/gorilla/mux" // Unused packages will create compilation error
)

type Item struct {
	Name string `json:"Title"`
	Desc string `json:"Desc"`
	Price float64 `json:"Price"`
}

type Inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	inventory := Inventory{
		Item{Name: "Cheese", Desc: "A fine block of cheese.", Price: 4.99},
	}

	fmt.Println("Endpoint Called: getInventory()")
	json.NewEncoder(w).Encode(inventory)
}

func handleRequests(){
	// := is the short variable declaration operator
	// Automatically determines type for variable
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}