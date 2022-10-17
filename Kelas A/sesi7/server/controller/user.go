package controller

import (
	"net/http"
	"sesi7/model"
	"sesi7/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepoOrm repository.UserRepository
	userRepo    repository.UserRepository
}

func NewUserController(userRepo, orm repository.UserRepository) *UserController {
	return &UserController{
		userRepo:    userRepo,
		userRepoOrm: orm,
	}
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	users, err := u.userRepo.GetUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": users,
	})
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = u.userRepoOrm.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "created user success",
	})
}
