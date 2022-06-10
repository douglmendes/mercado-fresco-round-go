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

// ListProducts godoc
// @Summary      List all products
// @Description  List all products currently in the system
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}  products.Product
// @Success      204  "Empty database"
// @Failure      500  {object}  response.Response
// @Router       /api/v1/products [get]
func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.DecodeError(err.Error()))
			return
		}

		ctx.JSON(func() int {
			if len(products) == 0 {
				return http.StatusNoContent
			}
			return http.StatusOK
		}(), response.NewResponse(products))
	}
}

// GetProduct godoc
// @Summary      Get a product by id
// @Description  Get a product from the system searching by id
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product id"
// @Success      200  {object}  products.Product
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /api/v1/products/{id} [get]
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

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product in the system
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      productsRequest  true  "Product to be created"
// @Success      201      {object}  products.Product
// @Failure      409      {object}  response.Response
// @Failure      422      {object}  response.Response
// @Router       /api/v1/products [post]
func (c *ProductController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req productsRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.DecodeError(err.Error()))
			return
		}

		product, err := c.service.Create(req.ProductCode, req.Description, req.Width, req.Height, req.Length, req.NetWeight, req.ExpirationRate, req.RecommendedFreezingTemperature, req.FreezingRate, req.ProductTypeId, req.SellerId)
		if err != nil {
			ctx.JSON(http.StatusConflict, response.DecodeError(err.Error()))
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

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update a product in the system, selecting by id
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int              true   "Product id"
// @Param        product  body      productsRequest  false  "Product to be updated (all fields are optional)"
// @Success      200      {object}  products.Product
// @Failure      400      {object}  response.Response
// @Failure      404      {object}  response.Response
// @Failure      422      {object}  response.Response
// @Failure      500      {object}  response.Response
// @Router       /api/v1/products/{id} [patch]
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

			ctx.JSON(http.StatusInternalServerError, response.DecodeError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(product))
	}
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product from the system, selecting by id
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Product id"
// @Success      204  "Successfully deleted"
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /api/v1/products/{id} [delete]
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
