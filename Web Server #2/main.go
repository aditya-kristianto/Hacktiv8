package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var PORT = ":8080"
var user []User
var userService = NewUserSvc()

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/register", register)
	http.HandleFunc("/users", getUsers)

	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Println(w, msg)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var data registerRequest

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	userService.registerUser(data)

	responseJSON := response{
		Status: "success",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseJSON)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var users = userService.getUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

type registerRequest struct {
	Name string `json:"name"`
}

type response struct {
	Status string `json:"status"`
}

func (u *service) registerUser(user registerRequest) {
	u.users = append(u.users, User{
		Name: user.Name,
	})

	fmt.Println(user.Name + " berhasil didaftarkan")
}

type userSvc interface {
	registerUser(user registerRequest)
	getUsers() []User
}

func NewUserSvc() userSvc {
	return &service{users: user}
}

type User struct {
	Name string `json:"name"`
}

type service struct {
	users []User `json:"users"`
}

func (u *service) getUsers() []User {
	return u.users
}
