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
	UID string `json:"UID"`
	Name string `json:"Name"`
	Desc string `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Function Called: homePage()")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getInventory()")

	json.NewEncoder(w).Encode(inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item) // Obtain item from request JSON

	inventory = append(inventory, item)	// Add item to inventory

	json.NewEncoder(w).Encode(item) // Show item in response JSON for verification
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtUid(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			// Delete item from Slice
			inventory = append(inventory[:index], inventory[index+1:]...)
			break 
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item) // Obtain item from request JSON
	
	params := mux.Vars(r)
	
	_deleteItemAtUid(params["uid"]) // Delete item
	inventory = append(inventory, item) // Create it again with data from request

	json.NewEncoder(w).Encode(inventory)
}

func handleRequests(){
	// := is the short variable declaration operator
	// Automatically determines type for variable
	router := mux.NewRouter().StrictSlash(true)	

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory", createItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	// Data store
	inventory = append(inventory, Item{
		UID: "0", 
		Name: "Cheese", 
		Desc: "A fine block of cheese.",
		Price: 4.99,
	}) 

	handleRequests()
}