package server

import (
	"sesi8-latihan/server/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	people *controllers.PersonController
}

func NewRouter(people *controllers.PersonController) *Router {
	return &Router{people: people}
}

func (r *Router) Start(port string) {
	router := gin.Default()

	router.GET("/people", r.people.GetPeople)
	router.Run(port)
}
