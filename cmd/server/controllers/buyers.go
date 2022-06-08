package controllers

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BuyerController struct {
	service buyers.Service
}

type buyerRequest struct {
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

func NewBuyer(s buyers.Service) *BuyerController {
	return &BuyerController{
		service: s,
	}

}

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

func (s *BuyerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		var req buyerRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
