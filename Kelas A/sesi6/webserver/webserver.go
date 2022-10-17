package webserver

import (
	"fmt"
	"log"
	"net/http"
	"sesi6/webserver/controllers"
)

const PORT = ":4000"

func Start() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/users", controllers.GetUsersHandler)
	http.HandleFunc("/users", controllers.CreateUserHandler)
	log.Println("Server running at port", PORT)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}
