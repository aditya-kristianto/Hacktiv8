package server

import (
	"sesi7/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	user    *controller.UserController
	product *controller.ProductController
}

func NewRouter(router *gin.Engine, user *controller.UserController, product *controller.ProductController) *Router {
	return &Router{
		router:  router,
		user:    user,
		product: product,
	}
}

func (r *Router) Start(port string) {

	r.router.GET("/v1/users", r.user.GetUsers)
	r.router.POST("/v1/users", r.user.CreateUser)

	r.router.GET("/v1/products", r.product.GetProducts)
	r.router.POST("/v1/products", r.product.CreateProduct)
	r.router.Run(port)
}
