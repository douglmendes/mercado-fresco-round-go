package controller

import (
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductBatchesController struct {
	service domain.ProductBatchesService
}

func NewController(pbs domain.ProductBatchesService) *ProductBatchesController {
	return &ProductBatchesController{
		service: pbs,
	}
}

// Create godoc
// @Summary Create product batches
// @Tags Products
// @Description create a product batch
// @Accept  json
// @Produce  json
// @Param product_batch body createRequest true "ProductBatch to create"
// @Success 201 {object} domain.ProductBatch
// @Failure 422 {object} response.Response
// @Router /api/v1/productBatches [post]
func (c *ProductBatchesController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request createRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		productBatch, err := c.service.Create(
			ctx,
			request.BatchNumber,
			request.CurrentQuantity,
			request.CurrentTemperature,
			request.DueDate,
			request.InitialQuantity,
			request.ManufacturingDate,
			request.ManufacturingHour,
			request.MinimumTemperature,
			request.ProductId,
			request.SectionId,
		)

		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, productBatch)
	}
}

func (c *ProductBatchesController) GetBySectionId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var id int64
		var err error

		if stringId, exists := ctx.GetQuery("id"); exists {
			id, err = strconv.ParseInt(stringId, 10, 64)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
				return
			}
		}

		records, err := c.service.GetBySectionId(ctx, int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(records))
	}
}

type createRequest struct {
	BatchNumber        int    `json:"batch_number" binding:"required"`
	CurrentQuantity    int    `json:"current_quantity" binding:"required"`
	CurrentTemperature int    `json:"current_temperature" binding:"required"`
	DueDate            string `json:"due_date" binding:"required"`
	InitialQuantity    int    `json:"initial_quantity" binding:"required"`
	ManufacturingDate  string `json:"manufacturing_date" binding:"required"`
	ManufacturingHour  int    `json:"manufacturing_hour" binding:"required"`
	MinimumTemperature int    `json:"minimum_temperature" binding:"required"`
	ProductId          int    `json:"product_id" binding:"required"`
	SectionId          int    `json:"section_id" binding:"required"`
}
