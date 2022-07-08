package controller

import (
	"github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type InboudOrdersController struct {
	service domain.Service
}

type requestInboudOrders struct {
	Id             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int    `json:"employee_id"`
	ProductBatchId int    `json:"product_batch_id"`
	WarehouseId    int    `json:"warehouse_id"`
}

func NewInboudOrders(e domain.Service) *InboudOrdersController {
	return &InboudOrdersController{
		service: e,
	}
}

func (ioc *InboudOrdersController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestInboudOrders
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				gin.H{
					"error":  "VALIDATEERR-1",
					"messge": "Invalid inputs. Please check your inputs",
				})
			return
		}
		io, err := ioc.service.Create(req.OrderDate, req.OrderNumber, req.EmployeeId, req.ProductBatchId, req.WarehouseId)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, response.NewResponse(io))
	}

}

func (ioc *InboudOrdersController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employee, _ := ctx.GetQuery("employee_id")
		log.Println(employee)
		employeeId, _ := strconv.Atoi(employee)
		log.Println(employeeId)
		i, err := ioc.service.GetByEmployee(int64(employeeId))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, response.NewResponse(i))
	}
}
