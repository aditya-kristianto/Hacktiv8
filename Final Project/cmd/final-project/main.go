package main

import (
	"net/http"

	"final-project/internal/app"
	"final-project/internal/pkg/db"
	"final-project/internal/pkg/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	_ "final-project/docs"
)

// @title Swagger Final Project API
// @version 1.0
// @description This is a sample server Final Project server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email kristianto.aditya@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
// @schemes http
func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	postgresDB := db.NewDB()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	helper.NewValidator(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	r := app.NewRouter(postgresDB)
	r.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
