package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sesi10/helper"
	"sesi10/server/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping...")
	ctx := r.Context()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "OKE",
		"role":    ctx.Value("role"),
	})
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var req models.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Id = len(models.Users) + 1

	hash, err := helper.GeneratePassword([]byte(req.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Password = hash

	models.Users = append(models.Users, req)

	ctx.JSON(http.StatusCreated, gin.H{
		"payload": req,
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	var req models.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.FindByEmail(req.Email)
	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "no data",
		})
		return
	}

	err = helper.ValidatePassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	payload := helper.Token{
		Email: user.Email,
		Role:  user.Role,
	}

	token, err := helper.GenerateToken(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"payload": token,
	})

}

func (u *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"payload": models.Users,
		"myRole":  ctx.GetString("role"),
	})
}
