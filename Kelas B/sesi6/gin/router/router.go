package router

import (
	"sesi6/gin/controller"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/employees", controller.GetEmployeesHandler)
	router.GET("/employees/:id", controller.GetEmployeeByID)
	router.POST("/employees", controller.CreateEmployee)

	return router
}
