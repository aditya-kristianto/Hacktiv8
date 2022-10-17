package main

import (
	"net/http"
	"swaggo/server/controller"

	_ "swaggo/docs"

	_ "github.com/alecthomas/template"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// @title Orders API
// @description Ini adalah simple API Spec untuk Orders
// @version v1.0
// @termsOfService https://swagger.io/terms
// @contact.name Hacktiv8
// @contact.email admin@hacktiv8.com
// @host localhost:4000
// @BasePath /v1
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/orders", controller.GetOrders).Methods("GET")
	router.HandleFunc("/v1/orders/{id}", controller.CreateOrder).Methods("POST")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":4000", router)
}
