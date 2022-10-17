package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sesi10/helper"
	"sesi10/server/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetUserGin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"payload": model.Users,
		"myRole":  ctx.Value("role"),
	})
}

func (u *UserController) Register(ctx *gin.Context) {
	var req model.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Id = len(model.Users) + 1
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Password = string(hash)
	model.Users = append(model.Users, req)

	ctx.JSON(http.StatusCreated, req)
}

func (u *UserController) Login(ctx *gin.Context) {
	var req model.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := model.FindbyEmail(req.Email)
	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user with email " + req.Email + " not found!",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := helper.GenerateToken(user.Email, user.Role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":     "Login success",
		"payload": token,
	})

}

func (u *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get users process ...", r.Context().Value("role"))
	time.Sleep(100 * time.Millisecond)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg": "Hello",
	})
}
