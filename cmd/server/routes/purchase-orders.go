package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	purchaOrdersController "github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/controller"
	purchaOrdersRepository "github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/repository"
	purchaOrdersService "github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/service"
	"github.com/gin-gonic/gin"
)

func PurchaseOrdersRoutes(group *gin.RouterGroup) {
	purchaseOrdersRouterGroup := group.Group("/purchase-orders")
	{
		purchaseOrdersRepo := purchaOrdersRepository.NewRepository(connections.NewConnection())
		purchaseOrdersService := purchaOrdersService.NewService(purchaseOrdersRepo)
		po := purchaOrdersController.NewPurchaseOrders(purchaseOrdersService)

		purchaseOrdersRouterGroup.POST("/", po.Create())
	}
}
