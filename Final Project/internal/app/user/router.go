package user

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type (
	Router struct {
		controller *Controller
	}
)

func NewRouter(db *gorm.DB) *Router {
	return &Router{
		controller: NewController(db),
	}
}

func (r *Router) Init(e *echo.Echo, config *middleware.JWTConfig) {
	g := e.Group("/users")
	g.POST("/register", r.controller.Register)
	g.POST("/login", r.controller.Login)
	g.PUT("", r.controller.UpdateUser, middleware.JWTWithConfig(*config))
	g.DELETE("", r.controller.DeleteUser, middleware.JWTWithConfig(*config))
}
