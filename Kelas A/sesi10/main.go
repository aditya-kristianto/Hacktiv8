package main

import (
	"sesi10/server"
	"sesi10/server/controller"
)

func main() {

	userController := controller.NewUserController()
	router := server.NewRouterGin(userController)

	router.Start(":4444")
}
