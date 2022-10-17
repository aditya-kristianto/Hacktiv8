package app

import (
	"final-project/internal/app/comment"
	"final-project/internal/app/photo"
	"final-project/internal/app/socialmedia"
	"final-project/internal/app/user"
	"final-project/internal/pkg/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Router struct {
	user        *user.Router
	photo       *photo.Router
	comment     *comment.Router
	socialmedia *socialmedia.Router
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{
		user:        user.NewRouter(db),
		photo:       photo.NewRouter(db),
		comment:     comment.NewRouter(db),
		socialmedia: socialmedia.NewRouter(db),
	}
}

func (r *Router) Init(e *echo.Echo) {
	config := middleware.JWTConfig{
		Claims:     &helper.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.user.Init(e, &config)
	r.photo.Init(e, &config)
	r.comment.Init(e, &config)
	r.socialmedia.Init(e, &config)
}
