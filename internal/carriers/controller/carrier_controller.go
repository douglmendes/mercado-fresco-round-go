package controller

import (
	"github.com/douglmendes/mercado-fresco-round-go/internal/carriers/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/logger"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CarrierController struct {
	service domain.CarrierService
}

func NewCarries(c domain.CarrierService) *CarrierController {
	return &CarrierController{
		service: c,
	}
}

// Create godoc
// @Summary Create carriers
// @Tags Warehouses
// @Description create one carrie
// @Accept  json
// @Produce  json
// @Param warehouses body carriesCreateRequest true "Carrier to create"
// @Success 201 {object} domain.Carrier
// @Failure 422 {object} response.Response
// @Router /api/v1/carriers [post]
func (c *CarrierController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request carriesCreateRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		carries, err := c.service.CreateCarrier(
			ctx,
			request.Cid,
			request.CompanyName,
			request.Address,
			request.Telephone,
			request.LocalityId,
		)
		if err != nil {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, carries)
	}
}

type carriesCreateRequest struct {
	Cid         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityId  int    `json:"locality_id" binding:"required"`
}
