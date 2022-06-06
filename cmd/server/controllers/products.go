package controllers

import (
	"net/http"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service products.Service
}

func NewProductController(service products.Service) *ProductController {
	return &ProductController{service}
}

func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		product, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}
}
