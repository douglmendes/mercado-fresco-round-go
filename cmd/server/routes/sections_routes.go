package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	pbController "github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/controller"
	pbRepo "github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/repository"
	pbService "github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/service"
	productsRepository "github.com/douglmendes/mercado-fresco-round-go/internal/products/repository/mariadb"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/controller"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/repository"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/service"
	"github.com/gin-gonic/gin"
)

func SectionsRoutes(group *gin.RouterGroup) {

	sectionRouterGroup := group.Group("/sections")
	{
		connection := connections.NewConnection()

		sectionsRepository := repository.NewRepository(connection)
		sectionsService := service.NewService(sectionsRepository)
		sectionsController := controller.NewSectionsController(sectionsService)

		productBatchesRepository := pbRepo.NewRepository(connection)
		productRepo := productsRepository.NewRepository(connection)
		productBatchesService := pbService.NewService(productBatchesRepository, productRepo, sectionsRepository)
		productBatchesController := pbController.NewController(productBatchesService)

		sectionRouterGroup.POST("/", sectionsController.Create)
		sectionRouterGroup.GET("/", sectionsController.GetAll)
		sectionRouterGroup.GET("/:id", sectionsController.GetById)
		sectionRouterGroup.PATCH("/:id", sectionsController.Update)
		sectionRouterGroup.DELETE("/:id", sectionsController.Delete)

		sectionRouterGroup.GET("/reportProducts", productBatchesController.GetBySectionId())
	}
}
