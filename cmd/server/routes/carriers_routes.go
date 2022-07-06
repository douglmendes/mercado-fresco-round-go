package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	carriersController "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/controller"
	carriersRepository "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/repository"
	carriersService "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/service"

	"github.com/gin-gonic/gin"
)

func CarriersRoutes(group *gin.RouterGroup) {

	carriersRouterGroup := group.Group("/carriers")
	{
		carriersRepo := carriersRepository.NewRepository(connections.NewConnection())
		carrierService := carriersService.NewService(carriersRepo)
		controller := carriersController.NewCarries(carrierService)

		carriersRouterGroup.POST("/", controller.Create())

	}
}
