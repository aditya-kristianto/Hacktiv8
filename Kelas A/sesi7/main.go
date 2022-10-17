package main

import (
	"sesi7/config"
	"sesi7/repository/gorm"
	"sesi7/server"
	"sesi7/server/controller"

	"sesi7/repository/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	dbOrm, err := config.ConnectGorm()
	if err != nil {
		panic(err)
	}

	db, err := dbOrm.DB()
	if err != nil {
		panic(err)
	}

	userRepoGorm := gorm.NewUserRepo(dbOrm)
	userRepo := postgres.NewUserRepo(db)
	userHandler := controller.NewUserController(userRepo, userRepoGorm)

	productRepo := gorm.NewProductRepo(dbOrm)
	productHandler := controller.NewProductController(productRepo)

	router := gin.Default()
	server.NewRouter(router, userHandler, productHandler).Start(":4000")
}
