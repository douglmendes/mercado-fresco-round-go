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
		e := buyersController.NewBuyer(buyersService)

		buyerRouterGroup.POST("/", e.Create())
		buyerRouterGroup.GET("/", e.GetAll())
		buyerRouterGroup.GET("/:id", e.GetById())
		buyerRouterGroup.PATCH("/:id", e.Update())
		buyerRouterGroup.DELETE("/:id", e.Delete())

	}
}
