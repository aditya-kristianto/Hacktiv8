package webserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port = ":4000"

func Start() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employees/create", createEmployee)

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	employees := GetEmployees()
	userAgent := r.Header.Get("x-client")
	if userAgent == "API" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"payload": employees,
			"total":   len(employees),
		})
	} else {
		tpl, err := template.ParseFiles("./webserver/static/index.html", "./webserver/static/layout.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		tpl.Execute(w, map[string]interface{}{
			"Payload":  employees,
			"Title":    "Halaman Employee",
			"listMenu": []string{},
		})
	}
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": fmt.Sprintf("method %s not allowed!", method),
		})
		return
	}
	var req Employee
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	req.ID = len(emps) + 1
	employees := CreateEmployee(&req)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": employees,
		"total":   len(employees),
	})
}
