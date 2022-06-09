package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type SellerController struct {
	service sellers.Service
}

type request struct {
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

func NewSeller(s sellers.Service) *SellerController {
	return &SellerController{
		service: s,
	}

}

// ListSellers godoc
// @Summary List sellers
// @Tags Sellers
// @Description get sellers
// @Produce  json
// @Success 200 {object} request
// @Failure 404 {object} string
// @Router /api/v1/sellers [get]
func (c *SellerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(s))
	}
}

// ListSeller godoc
// @Summary List seller
// @Tags Sellers
// @Description get seller
// @Accept  json
// @Produce  json
// @Param id   path int true "Seller ID"
// @Success 200 {object} request
// @Failure 404 {object} string
// @Router /api/v1/sellers/{id} [get]
func (c *SellerController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		s, err := c.service.GetById(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(s))
	}
}

// Create godoc
// @Summary Create sellers
// @Tags Sellers
// @Description create sellers
// @Accept  json
// @Produce  json
// @Param product body request true "Seller to create"
// @Success 201 {object} string
// @Failure 422 {object} string
// @Failure 409 {object} string
// @Router /api/v1/sellers [post]
func (c *SellerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if req.Cid == 0 {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("cid is required"))
			return
		}
		if req.CompanyName == "" {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("company name is required"))
			return
		}
		if req.Address == "" {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("address is required"))
			return
		}
		if req.Telephone == "" {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("telephone is required"))
			return
		}

		s, err := c.service.Create(req.Cid, req.CompanyName, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, response.NewResponse(s))
	}

}

// ListSellers godoc
// @Summary Update seller
// @Tags Sellers
// @Description update seller
// @Accept  json
// @Produce  json
// @Param product body request true "Seller to create"
// @Param id   path int true "Seller ID"
// @Success 200 {object} request
// @Router /api/v1/sellers/{id} [patch]
func (s *SellerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		s, err := s.service.Update(int(id), req.Cid, req.CompanyName, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, s)
	}

}

// ListSellers godoc
// @Summary Delete seller
// @Tags Sellers
// @Description delete seller
// @Param id   path int true "Seller ID"
// @Success 204 {object} request
// @Router /api/v1/sellers/{id} [delete]
func (c *SellerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("seller %d was removed", id)})
	}

}
