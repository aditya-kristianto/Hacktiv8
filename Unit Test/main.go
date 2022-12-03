package main

import (
	"sesi10/server"
	"sesi10/server/controller"
)

func main() {

	userController := controller.NewUserController()

	router := server.NewGinRotuer(":4444", userController)

	router.Start()
}
