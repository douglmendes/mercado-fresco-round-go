package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	employeesController "github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/controller"
	inboudOrdersRepository "github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/repository"
	inboudOrdersService "github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/service"
	"github.com/gin-gonic/gin"
)

func InboudOrdersRoutes(group *gin.RouterGroup) {
	inboudOrdersRouterGroup := group.Group("/inboud-orders")
	{
		inboudOrdersRepo := inboudOrdersRepository.NewRepository(connections.NewConnection())
		inboudOrderService := inboudOrdersService.NewService(inboudOrdersRepo)
		io := employeesController.NewInboudOrders(inboudOrderService)

		inboudOrdersRouterGroup.POST("/", io.Create())
		inboudOrdersRouterGroup.GET("/", io.GetById())

	}
}
