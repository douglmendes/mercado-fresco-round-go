package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func SectionsRoutes(group *gin.RouterGroup) {

	sectionRouterGroup := group.Group("/sections")
	{
		sectionsRepository := sections.NewRepository(store.New(store.FileType, store.PathBuilder("/sections.json")))
		sectionsService := sections.NewService(sectionsRepository)
		sectionsController := controllers.NewSectionsController(sectionsService)

		sectionRouterGroup.POST("/", sectionsController.Create)
		sectionRouterGroup.GET("/", sectionsController.GetAll)
		sectionRouterGroup.GET("/:id", sectionsController.GetById)
		sectionRouterGroup.PATCH("/:id", sectionsController.Update)
		sectionRouterGroup.DELETE("/:id", sectionsController.Delete)

	}
}
