package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, " Getting data...")
}

func addTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Adding data...")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Deleting Task...")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Updating Task...")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/feed", getTask)
	router.HandleFunc("/feed/add", addTask)
	router.HandleFunc("/feed/{id}/update", updateTask)
	router.HandleFunc("/feed/delete", deleteTask)
	log.Fatal(http.ListenAndServe(":8080", router))

}
