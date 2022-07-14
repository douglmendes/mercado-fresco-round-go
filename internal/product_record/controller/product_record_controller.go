package controller

import (
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/logger"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

type ProductRecordController struct {
	service domain.ProductRecordService
}

func NewProductRecordController(service domain.ProductRecordService) *ProductRecordController {
	return &ProductRecordController{service}
}

// ListProductRecords godoc
// @Summary      List all product records
// @Description  List all product records currently in the system
// @Tags         productRecords
// @Accept       json
// @Produce      json
// @Success      200  {array}  domain.ProductRecordCount
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /api/v1/products/reportRecords [get]
// @Param id query int false "int valid" minimum(1)
func (c *ProductRecordController) GetByProductId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var id int64
		var err error

		if stringId, exists := ctx.GetQuery("id"); exists {
			id, err = strconv.ParseInt(stringId, 10, 64)
			if err != nil {
				message := err.Error()
				logger.Error(ctx, store.GetPathWithLine(), message)

				ctx.JSON(http.StatusBadRequest, response.DecodeError(message))
				return
			}
		}

		productRecords, err := c.service.GetByProductId(ctx, int(id))
		if err != nil {
			message := err.Error()
			logger.Error(ctx, store.GetPathWithLine(), message)

			ctx.JSON(http.StatusNotFound, response.DecodeError(message))
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(productRecords))
	}
}

type productRecordsRequest struct {
	LastUpdateDate string  `json:"last_update_date" binding:"required"`
	PurchasePrice  float64 `json:"purchase_price" binding:"required,min=0"`
	SalePrice      float64 `json:"sale_price" binding:"required,min=0"`
	ProductId      int     `json:"product_id" binding:"required,min=1"`
}

// CreateProductRecord godoc
// @Summary      Create a new product record
// @Description  Create a new product record in the system
// @Tags         productRecords
// @Accept       json
// @Produce      json
// @Param        productRecord  body  productRecordsRequest  true  "Product record to be created"
// @Success      201      {object}  domain.ProductRecord
// @Failure      409      {object}  response.Response
// @Failure      422      {object}  response.Response
// @Router       /api/v1/productRecords [post]
func (c *ProductRecordController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req productRecordsRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			message := err.Error()
			logger.Error(ctx, store.GetPathWithLine(), message)

			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.DecodeError(message))
			return
		}

		arg := domain.ProductRecord{
			LastUpdateDate: req.LastUpdateDate,
			PurchasePrice:  req.PurchasePrice,
			SalePrice:      req.SalePrice,
			ProductId:      req.ProductId,
		}

		product, err := c.service.Create(ctx, arg)
		if err != nil {
			message := err.Error()
			logger.Error(ctx, store.GetPathWithLine(), message)

			ctx.JSON(http.StatusConflict, response.DecodeError(message))
			return
		}

		ctx.JSON(http.StatusCreated, response.NewResponse(product))
	}
}
