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
	LocalityName string `json:"locality_name" binding:"required"`
	ProvinceName string `json:"province_name" binding:"required"`
	CountryName  string `json:"country_name" binding:"required"`
}

type sqlUpdateRequest struct {
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

func NewLocality(s domain.LocalityService) *LocalityController {
	return &LocalityController{
		service: s,
	}
}

// func (c *LocalityController) GetAll() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		l, err := c.service.GetAll(ctx)
// 		if err != nil {
// 			ctx.JSON(http.StatusNotFound, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, response.NewResponse(l))
// 	}
// }

func (c *LocalityController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Id"})
			return
		}

		l, err := c.service.GetById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(l))
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

func (c *LocalityController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req sqlCreateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if req.LocalityName == "" {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("locality name is required"))
			return
		}
		if req.ProvinceName == "" {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("province name is required"))
			return
		}
		if req.CountryName == "" {
			ctx.JSON(http.StatusUnprocessableEntity,
				response.DecodeError("country name is required"))
			return
		}

		l, err := c.service.Create(ctx, req.LocalityName, req.ProvinceName, req.CountryName)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, response.NewResponse(l))

	}
}

func (s *LocalityController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Id"})
			return
		}

		var req sqlUpdateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		l, err := s.service.Update(ctx, id, req.LocalityName, req.ProvinceName, req.CountryName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, l)
	}
}

func (c *LocalityController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Id"})
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}