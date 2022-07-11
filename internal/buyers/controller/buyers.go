package controller

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BuyerController struct {
	service domain.Service
}

type buyerRequest struct {
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

func NewBuyer(s domain.Service) *BuyerController {
	return &BuyerController{
		service: s,
	}

}

// ListSellers godoc
// @Summary List buyers
// @Tags Buyers
// @Description get buyers
// @Produce  json
// @Success 200 {array} buyers.Buyer
// @Failure 404 {object} string
// @Router /api/v1/buyers [get]
func (c *BuyerController) GetAll() gin.HandlerFunc {
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
// @Summary List buyer
// @Tags Buyers
// @Description get buyer
// @Accept  json
// @Produce  json
// @Param id   path int true "Buyer ID"
// @Success 200 {object} buyers.Buyer
// @Failure 404 {object} string
// @Router /api/v1/buyers/{id} [get]
func (c *BuyerController) GetById() gin.HandlerFunc {
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

func (c *BuyerController) GetOrdersByBuyers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buyerId, byId := ctx.GetQuery("id")
		id, err := strconv.Atoi(buyerId)
		if byId == true {
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Id"})
				return
			}
		}

		l, err := c.service.GetOrdersByBuyers(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(l))
	}
}

// Create godoc
// @Summary Create buyers
// @Tags Buyers
// @Description create buyers
// @Accept  json
// @Produce  json
// @Param buyer body buyerRequest true "Buyer to create"
// @Success 201 {object} buyers.Buyer
// @Failure 422 {object} string
// @Failure 409 {object} string
// @Router /api/v1/buyers [post]
func (c *BuyerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req buyerRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if req.CardNumberId == "" {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("card number is required"))
			return
		}
		if req.FirstName == "" {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("first name is required"))
			return
		}
		if req.LastName == "" {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("last name is required"))
			return
		}

		s, err := c.service.Create(req.CardNumberId, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, response.NewResponse(s))
	}

}

// ListSellers godoc
// @Summary Update buyer
// @Tags Buyers
// @Description update buyer
// @Accept  json
// @Produce  json
// @Param buyer body buyerRequest true "Buyer to create"
// @Param id   path int true "Buyer ID"
// @Success 200 {object} buyers.Buyer
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/buyers/{id} [patch]
func (s *BuyerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		var req buyerRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		s, err := s.service.Update(int(id), req.CardNumberId, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, s)
	}

}

// ListSellers godoc
// @Summary Delete buyer
// @Tags Buyers
// @Description delete buyer
// @Param id   path int true "Buyer ID"
// @Success 204 {object} request
// @Router /api/v1/buyers/{id} [delete]
func (c *BuyerController) Delete() gin.HandlerFunc {
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
