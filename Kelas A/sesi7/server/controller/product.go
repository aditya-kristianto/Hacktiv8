package controller

import (
	"net/http"
	"sesi7/model"
	"sesi7/repository"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	repo repository.ProductRepository
}

func NewProductController(repo repository.ProductRepository) *ProductController {
	return &ProductController{
		repo: repo,
	}
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var req model.Product
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = p.repo.CreateProduct(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "created product success",
	})
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.repo.GetProducts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"payload": products,
	})
}
