package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sesi10/helper"
	"sesi10/server/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		// before
		fmt.Println("Middleware 1 : Before ...")

		// send data to context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "role", "admin")
		r = r.WithContext(ctx)
		next(w, r)

		// after
		end := time.Since(now).Milliseconds()
		log.Println("request success with response time :", end, "ms")
		// fmt.Println("")
	}
}

func Middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		fmt.Println("Role nya adalah :", ctx.Value("role"))
		fmt.Println("Middleware 2 : Before ...")
		next(w, r)
		fmt.Println("Middleware 2 : After ...")
	}
}

func Middleware1And2(next http.HandlerFunc) http.HandlerFunc {
	return Middleware1(Middleware2(func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}))
}

func Middleware1Gin(ctx *gin.Context) {
	fmt.Println("Middleware 1 Gin : Before ...")
	ctx.Next()
	fmt.Println("Middleware 1 Gin : After ...")
}

func CheckAuth(ctx *gin.Context) {
	tokenHeader := ctx.Request.Header.Get("Authorization")
	tokenArr := strings.Split(tokenHeader, "Bearer ")
	if len(tokenArr) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "need login",
		})
		return
	}

	tokenStr := tokenArr[1]

	payload, err := helper.ValidateToken(tokenStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Set("email", payload["email"])
	ctx.Set("role", payload["role"])
	ctx.Next()
}

func AdminRole(ctx *gin.Context) {
	user := model.FindbyEmail(ctx.GetString("email"))
	isCanAccess := checkRole(user.Role, []string{"admin"})
	if !isCanAccess {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "forbidden",
		})
		return
	}
	ctx.Next()
}

func checkRole(userRole string, needs []string) bool {
	for _, role := range needs {
		if userRole == role {
			return true
		}
	}
	return false
}
