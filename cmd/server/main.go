package main

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	sectionsRepository := sections.NewRepository(store.New("../../sections.json"))
	sectionsService := sections.NewService(sectionsRepository)
	sectionsController := controllers.NewSectionsController(sectionsService)

	sectionsRoutes := router.Group("/api/v1/sections")
	{
		sectionsRoutes.GET("/", sectionsController.GetAll)
		sectionsRoutes.GET("/:id", sectionsController.GetById)
		sectionsRoutes.POST("/", sectionsController.Create)
		sectionsRoutes.PATCH("/:id", sectionsController.Update)
		sectionsRoutes.DELETE("/:id", sectionsController.Delete)
	}

	router.Run()
}
