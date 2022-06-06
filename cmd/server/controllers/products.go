package controllers

import (
	"net/http"
	"strconv"
	"strings"

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
			ctx.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(products))
	}
}

func (c *ProductController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
			return
		}

		products, err := c.service.GetById(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
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
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.DecodeError(err.Error()))
			return
		}

		product, err := c.service.Store(req.ProductCode, req.Description, req.Width, req.Height, req.Length, req.NetWeight, req.ExpirationRate, req.RecommendedFreezingTemperature, req.FreezingRate, req.ProductTypeId, req.SellerId)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, response.DecodeError(err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, response.NewResponse(product))
	}
}

type optionalProductsRequest struct {
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ExpirationRate                 float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductTypeId                  int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}

func (c *ProductController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
			return
		}

		var req optionalProductsRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.DecodeError(err.Error()))
			return
		}

		product, err := c.service.Update(int(id), req.ProductCode, req.Description, req.Width, req.Height, req.Length, req.NetWeight, req.ExpirationRate, req.RecommendedFreezingTemperature, req.FreezingRate, req.ProductTypeId, req.SellerId)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				ctx.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
				return
			}

			ctx.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(product))
	}
}

func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
			return
		}

		ctx.AbortWithStatus(http.StatusNoContent)
	}
}
