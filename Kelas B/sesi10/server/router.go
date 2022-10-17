package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sesi10/server/controller"
)

type Router struct {
	port string
	user *controller.UserController
}

func NewRouter(port string, user *controller.UserController) *Router {
	return &Router{
		port: port,
		user: user,
	}
}

func (r *Router) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/ping", middleware1(middleware2(r.user.Ping)))
	log.Println("server running at port", r.port)
	http.ListenAndServe(r.port, mux)
}

func middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// log.Println("Before")
		log.Println("Middleware 1 Started ...")

		ctx := r.Context()
		ctx = context.WithValue(ctx, "role", "admin")
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
		log.Println("Middleware 1 Ended ...")
	}
}

func middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware 2 Started ...")
		if true {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "unauthorized",
			})
			return
		}
		log.Println("access is", r.Context().Value("role"))
		next.ServeHTTP(w, r)
		log.Println("Middleware 2 Ended ...")
	}
}
