package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/controller"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/repository"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/service"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func SectionsRoutes(group *gin.RouterGroup) {

	sectionRouterGroup := group.Group("/sections")
	{
		sectionsRepository := repository.NewRepository(store.New(store.FileType, store.PathBuilder("/sections.json")))
		sectionsService := service.NewService(sectionsRepository)
		sectionsController := controller.NewSectionsController(sectionsService)

		sectionRouterGroup.POST("/", sectionsController.Create)
		sectionRouterGroup.GET("/", sectionsController.GetAll)
		sectionRouterGroup.GET("/:id", sectionsController.GetById)
		sectionRouterGroup.PATCH("/:id", sectionsController.Update)
		sectionRouterGroup.DELETE("/:id", sectionsController.Delete)

	}
}
