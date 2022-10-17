package socialmedia

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
	g := e.Group("/socialmedias")
	g.Use(middleware.JWTWithConfig(*config))
	g.POST("", r.controller.CreateSocialmedias)
	g.GET("", r.controller.GetSocialmedias)
	g.PUT("/:socialMediaId", r.controller.UpdateSocialmedias)
	g.DELETE("/:socialMediaId", r.controller.DeleteSocialmedias)
}
