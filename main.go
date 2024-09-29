package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Todo struct {
	ID        string
	Title     string
	Completed bool
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func updateTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var newTodo Todo
	fmt.Println(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	newTodo.ID = uuid.NewString()
	// add this Todo to our list of todos
	todos = append(todos, newTodo)
	// log.Info("Inserted into todos")

	json.NewEncoder(w).Encode(todos)
}

func handleTodoCrud(w http.ResponseWriter, r *http.Request) {
	// log.Info("Request Received")
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		updateTodos(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func main() {

	fmt.Println("Hello, Restful api")

	//Define routes
	http.HandleFunc("/todos", handleTodoCrud)

	//Start the server
	// log.SetLevel(log.InfoLevel)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
