package controller

import (
	"github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PurchaseOrder struct {
	service domain.Service
}

type requestPurchaseOrders struct {
	OrderNumber     string `json:"order_number"`
	OrderDate       string `json:"order_date"`
	TrackingCode    string `json:"tracking_code"`
	BuyerId         int    `json:"buyer_id"`
	CarrierId       int    `jsons:"carrier_id"`
	ProductRecordId int    `json:"product_record_id"`
	OrderStatusId   int    `json:"order_status_id"`
}

func NewPurchaseOrders(s domain.Service) *PurchaseOrder {
	return &PurchaseOrder{
		service: s,
	}
}

func (por *PurchaseOrder) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestPurchaseOrders
		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.Println(req)
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				gin.H{
					"error":  "VALIDATEERR-1",
					"messge": "Invalid inputs. Please check your inputs",
				})
			return
		}
		po, err := por.service.Create(ctx, req.OrderNumber, req.OrderDate, req.TrackingCode, req.BuyerId, req.CarrierId, req.ProductRecordId, req.OrderStatusId)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, response.NewResponse(po))
	}
}
