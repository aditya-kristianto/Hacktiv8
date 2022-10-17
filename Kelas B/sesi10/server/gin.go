package server

import (
	"fmt"
	"log"
	"net/http"
	"sesi10/helper"
	"sesi10/server/controller"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	port string
	user *controller.UserController
}

func NewGinRotuer(port string, user *controller.UserController) *GinRouter {
	return &GinRouter{
		port: port,
		user: user,
	}
}

func (g *GinRouter) Start() {
	router := gin.Default()

	router.Use(Trace)
	router.POST("/users", g.user.CreateUser)
	router.POST("/users/login", g.user.Login)

	router.Use(Auth)
	router.GET("/users", g.user.GetAll)

	router.Run(g.port)
}

func Trace(ctx *gin.Context) {
	now := time.Now()

	ctx.Next()
	end := time.Since(now).Milliseconds()

	log.Println("response time :", end, "ms")
}

func Auth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "need signin",
		})
		return
	}

	bearer := strings.Split(auth, "Bearer ")

	if len(bearer) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "need signin",
		})
		return
	}

	tokStr := bearer[1]

	tok, err := helper.ValidateToken(tokStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("%+v\n", tok)
	ctx.Set("role", tok.Role)
	ctx.Next()
}
