package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func placeholderHandle(operation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"operation": operation})
	}
}

func main() {
	router := gin.Default()

	sectionsRoutes := router.Group("/api/v1/sections")
	{
		sectionsRoutes.GET("/", placeholderHandle("get all sections"))
		sectionsRoutes.GET("/:id", placeholderHandle("get section by id"))
		sectionsRoutes.POST("/", placeholderHandle("create section"))
		sectionsRoutes.PATCH("/:id", placeholderHandle("update section"))
		sectionsRoutes.DELETE("/:id", placeholderHandle("delete section"))
	}

	router.Run()
}
