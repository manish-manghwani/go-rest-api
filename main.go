package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func getTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, " Getting data...")

	jsonFile, err := os.Open("tasks.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened tasks.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tasks Tasks

	json.Unmarshal(byteValue, &tasks)

	// return tasks
	for i := 0; i < len(tasks.Tasks); i++ {
		fmt.Fprint(w, "Tasks title: "+tasks.Tasks[i].Title+"\n")
		fmt.Fprint(w, "Tasks description: "+tasks.Tasks[i].Description+"\n")
		fmt.Fprint(w, "Tasks created at: "+tasks.Tasks[i].CreatedAt+"\n")
		fmt.Fprint(w, "Facebook updated at: "+tasks.Tasks[i].UpdatedAt+"\n")
	}
}

func addTask(w http.ResponseWriter, r *http.Request) {

	newTask := Task{"Title", "Adding new task", "2019-12-10 12:22:07", "2019-12-22 12:22:07"}

	d, err := json.Marshal(newTask)

	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	ioutil.WriteFile("tasks.json", d, 0755)

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
