package photo

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
	g := e.Group("/photos")
	g.Use(middleware.JWTWithConfig(*config))
	g.POST("", r.controller.CreatePhotos)
	g.GET("", r.controller.GetPhotos)
	g.PUT("/:photoId", r.controller.UpdatePhotos)
	g.DELETE("/:photoId", r.controller.DeletePhotos)
}
