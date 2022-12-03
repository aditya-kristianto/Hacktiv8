package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
	{ID: 2, Name: "Nanda", Age: 23, Division: "Finance"},
	{ID: 3, Name: "Mailo", Age: 23, Division: "IT"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/employees", getEmployees)

	fmt.Println("service running in port ", PORT)
	http.ListenAndServe(PORT, nil)
	// http.ListenAndServe(PORT, http.FileServer(http.Dir("./doc")))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"

	fmt.Fprint(w, msg)
}
