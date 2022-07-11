package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	buyersController "github.com/douglmendes/mercado-fresco-round-go/internal/buyers/controller"
	buyersRepository "github.com/douglmendes/mercado-fresco-round-go/internal/buyers/repository"
	buyersService "github.com/douglmendes/mercado-fresco-round-go/internal/buyers/service"
	"github.com/gin-gonic/gin"
)

func BuyersRoutes(group *gin.RouterGroup) {

	buyerRouterGroup := group.Group("/buyers")
	{
		buyersDbRepo := buyersRepository.NewRepository(connections.NewConnection())
		buyersService := buyersService.NewService(buyersDbRepo)
		b := buyersController.NewBuyer(buyersService)

		buyerRouterGroup.POST("/", b.Create())
		buyerRouterGroup.GET("/", b.GetAll())
		buyerRouterGroup.GET("/:id", b.GetById())
		buyerRouterGroup.GET("/reportPurchaseOrders", b.GetOrdersByBuyers())
		buyerRouterGroup.PATCH("/:id", b.Update())
		buyerRouterGroup.DELETE("/:id", b.Delete())
	}
}
