package router

import (
	"sesi7/gin/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	employee *controller.EmployeeHandler
}

func NewRouter(employee *controller.EmployeeHandler) *Router {
	return &Router{
		employee: employee,
	}
}

func (r *Router) CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/employees", r.employee.GetEmployeesHandler)
	router.GET("/employees/:id", r.employee.GetEmployeeByID)
	router.POST("/employees", r.employee.CreateEmployee)

	return router
}
