package controller

import (
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type LocalityController struct {
	service domain.LocalityService
}

type sqlCreateRequest struct {
	ZipCode      string `json:"zip_code" binding:"required"`
	LocalityName string `json:"locality_name" binding:"required"`
	ProvinceName string `json:"province_name" binding:"required"`
	CountryName  string `json:"country_name" binding:"required"`
}

func NewLocality(s domain.LocalityService) *LocalityController {
	return &LocalityController{
		service: s,
	}
}

func (c *LocalityController) GetBySellers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		localityId, byId := ctx.GetQuery("id")
		id, err := strconv.Atoi(localityId)
		if byId == true {
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Id"})
				return
			}
		}

		l, err := c.service.GetBySellers(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(l))
	}
}

func (c *LocalityController) GetByCarriers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		localityId, byId := ctx.GetQuery("id")
		id, err := strconv.Atoi(localityId)
		if byId == true {
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Id"})
				return
			}
		}

		l, err := c.service.GetByCarriers(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(l))
	}
}

func (c *LocalityController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req sqlCreateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		l, err := c.service.Create(ctx, req.ZipCode, req.LocalityName, req.ProvinceName, req.CountryName)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, response.NewResponse(l))

	}
}
