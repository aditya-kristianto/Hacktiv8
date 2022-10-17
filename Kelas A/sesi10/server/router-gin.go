package server

import (
	"sesi10/server/controller"

	"github.com/gin-gonic/gin"
)

type RouterGin struct {
	user *controller.UserController
}

func NewRouterGin(user *controller.UserController) *RouterGin {
	return &RouterGin{
		user: user,
	}
}

func (r *RouterGin) Start(port string) {
	router := gin.Default()

	router.GET("/gin/users", CheckAuth, AdminRole, r.user.GetUserGin)
	router.POST("/gin/users", r.user.Register)
	router.POST("/gin/users/login", r.user.Login)

	router.Run(port)
}
