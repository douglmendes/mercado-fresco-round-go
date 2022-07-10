package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	carriersController "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/controller"
	carriersRepository "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/repository"
	carriersService "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/service"
	localityRepo "github.com/douglmendes/mercado-fresco-round-go/internal/localities/repository"

	"github.com/gin-gonic/gin"
)

func CarriersRoutes(group *gin.RouterGroup) {

	carriersRouterGroup := group.Group("/carriers")
	{
		connection := connections.NewConnection()

		localitiesRepo := localityRepo.NewRepository(connection)
		carriersRepo := carriersRepository.NewRepository(connection)
		carrierService := carriersService.NewService(carriersRepo, localitiesRepo)
		controller := carriersController.NewCarries(carrierService)

		carriersRouterGroup.POST("/", controller.Create())

	}
}
