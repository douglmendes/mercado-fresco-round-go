package controller

import (
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
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
				ctx.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
				return
			}
		}

		productRecords, err := c.service.GetByProductId(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(productRecords))
	}
}
