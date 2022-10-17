package controller

import (
	"database/sql"
	"net/http"
	"sesi7/gin/repository"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	repo repository.UserRepo
}

func NewEmployeeHandler(repo repository.UserRepo) *EmployeeHandler {
	return &EmployeeHandler{
		repo: repo,
	}
}

func (e *EmployeeHandler) GetEmployeesHandler(ctx *gin.Context) {
	employess, err := e.repo.GetUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	if len(*employess) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "data not found !",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": employess,
		"total":   len(*employess),
	})
}

func (e *EmployeeHandler) CreateEmployee(ctx *gin.Context) {
	var req repository.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = e.repo.CreateUser(&req)
	if err != nil {
		if err == sql.ErrConnDone {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "DB ERROR " + err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user created!",
	})
}

func (e *EmployeeHandler) GetEmployeeByID(ctx *gin.Context) {
	// emps := webserver.GetEmployees()
	// id, bool := ctx.Params.Get("id")
	// if !bool {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error": "params id should be included",
	// 	})
	// 	return
	// }
	// idInt, err := strconv.Atoi(id)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// emp := emps[idInt-1]

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"payload": emp,
	// })
}
