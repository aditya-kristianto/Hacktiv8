package main

import (
	"net/http"
	"sesi8/server/controllers"

	"github.com/gorilla/mux"

	_ "sesi8/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Orders API
// @description Sample API Spec for Orders
// @version v1.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host localhost:4000
// @contact.name Reyhan
// @contact.email reyhan@gmail.com
func main() {
	port := ":4000"

	router := mux.NewRouter()

	router.HandleFunc("/orders", controllers.GetOrders).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	// router.HandleFunc("/swagger", httpSwagger.WrapHandler).Methods("GET")
	http.ListenAndServe(port, router)
}
