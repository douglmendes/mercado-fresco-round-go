package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	pbController "github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/controller"
	pbRepository "github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/repository"
	pbService "github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/service"
	productsRepository "github.com/douglmendes/mercado-fresco-round-go/internal/products/repository/mariadb"
	sectionsRepository "github.com/douglmendes/mercado-fresco-round-go/internal/sections/repository"

	"github.com/gin-gonic/gin"
)

func ProductBatchesRoutes(group *gin.RouterGroup) {

	productBatchesRouterGroup := group.Group("/productBatches")
	{
		connection := connections.NewConnection()

		pbRepo := pbRepository.NewRepository(connection)
		productRepo := productsRepository.NewRepository(connection)
		sectionRepo := sectionsRepository.NewRepository(connection)
		service := pbService.NewService(pbRepo, productRepo, sectionRepo)
		controller := pbController.NewController(service)

		productBatchesRouterGroup.POST("/", controller.Create())
	}
}
