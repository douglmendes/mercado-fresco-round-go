package main

import (
	"net/http"

	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections"
	"github.com/gin-gonic/gin"
)

func placeholderHandle(operation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"operation": operation})
	}
}

func main() {
	router := gin.Default()

	sectionsRepository := sections.NewRepository()
	sectionsService := sections.NewService(sectionsRepository)
	sectionsController := controllers.NewSectionsController(sectionsService)

	sectionsRoutes := router.Group("/api/v1/sections")
	{
		sectionsRoutes.GET("/", sectionsController.GetAll)
		sectionsRoutes.GET("/:id", placeholderHandle("get section by id"))
		sectionsRoutes.POST("/", placeholderHandle("create section"))
		sectionsRoutes.PATCH("/:id", placeholderHandle("update section"))
		sectionsRoutes.DELETE("/:id", placeholderHandle("delete section"))
	}

	router.Run()
}
