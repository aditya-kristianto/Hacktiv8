package server

import (
	"net/http"
	"sesi10/server/controller"
)

type Router struct {
	user *controller.UserController
}

func NewRouter(u *controller.UserController) *Router {
	return &Router{
		user: u,
	}
}

func (r *Router) Start(port string) {
	router := http.NewServeMux()

	// endpoint := http.HandlerFunc(controller.GetUsers)

	router.HandleFunc("/users", Middleware1And2(r.user.GetUsers))
	router.HandleFunc("/users", Middleware1(Middleware2(r.user.GetUsers)))
	router.HandleFunc("/users", Middleware2(r.user.GetUsers))

	http.ListenAndServe(port, router)
}
