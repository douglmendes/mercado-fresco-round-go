package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	employeesRepository "github.com/douglmendes/mercado-fresco-round-go/internal/employees/repository"
	employeesController "github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/controller"
	inboudOrdersRepository "github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/repository"
	inboudOrdersService "github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/service"
	"github.com/gin-gonic/gin"
)

func InboudOrdersRoutes(group *gin.RouterGroup) {
	connection := connections.NewConnection()
	inboudOrdersRouterGroup := group.Group("/inboud-orders")
	{
		inboudOrdersRepo := inboudOrdersRepository.NewRepository(connection)
		employeesRepo := employeesRepository.NewRepository(connection)

		inboudOrderService := inboudOrdersService.NewService(inboudOrdersRepo, employeesRepo)
		io := employeesController.NewInboudOrders(inboudOrderService)

		inboudOrdersRouterGroup.POST("/", io.Create())
		inboudOrdersRouterGroup.GET("/report-inboud-orders", io.GetById())

	}
}
