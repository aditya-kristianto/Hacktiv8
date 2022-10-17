package controller

import (
	"net/http"
	"sesi6/webserver"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEmployeesHandler(ctx *gin.Context) {
	employess := webserver.GetEmployees()
	if len(employess) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "data not found !",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": employess,
		"total":   len(employess),
	})
}

func CreateEmployee(ctx *gin.Context) {
	var req webserver.Employee
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.ID = len(webserver.GetEmployees()) + 1
	resp := webserver.CreateEmployee(&req)
	ctx.JSON(http.StatusOK, gin.H{
		"payload": resp,
	})
}

func GetEmployeeByID(ctx *gin.Context) {
	emps := webserver.GetEmployees()
	id, bool := ctx.Params.Get("id")
	if !bool {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "params id should be included",
		})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	emp := emps[idInt-1]

	ctx.JSON(http.StatusOK, gin.H{
		"payload": emp,
	})
}
