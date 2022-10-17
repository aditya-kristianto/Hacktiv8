package comment

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
	g := e.Group("/comments")
	g.Use(middleware.JWTWithConfig(*config))
	g.POST("", r.controller.CreateComments)
	g.GET("", r.controller.GetComments)
	g.PUT("/:commentId", r.controller.UpdateComments)
	g.DELETE("/:commentId", r.controller.DeleteComments)
}
