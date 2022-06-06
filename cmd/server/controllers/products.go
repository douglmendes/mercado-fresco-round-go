package controllers

import (
	"net/http"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
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
		products, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(products))
	}
}

type productsRequest struct {
	ProductCode                    string  `json:"product_code" binding:"required"`
	Description                    string  `json:"description" binding:"required"`
	Width                          float64 `json:"width" binding:"required"`
	Height                         float64 `json:"height" binding:"required"`
	Length                         float64 `json:"length" binding:"required"`
	NetWeight                      float64 `json:"net_weight" binding:"required"`
	ExpirationRate                 float64 `json:"expiration_rate" binding:"required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	FreezingRate                   float64 `json:"freezing_rate" binding:"required"`
	ProductTypeId                  int     `json:"product_type_id" binding:"required"`
	SellerId                       int     `json:"seller_id" binding:"required"`
}

func (c *ProductController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req productsRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := c.service.Store(req.ProductCode, req.Description, req.Width, req.Height, req.Length, req.NetWeight, req.ExpirationRate, req.RecommendedFreezingTemperature, req.FreezingRate, req.ProductTypeId, req.SellerId)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, response.NewResponse(product))
	}
}
