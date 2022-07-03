package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/controller"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/repository/mariadb"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/service"
	"github.com/gin-gonic/gin"
)

func ProductsRoutes(group *gin.RouterGroup) {
	productRouterGroup := group.Group("/products")
	{
		productsRepository := mariadb.NewRepository(connections.NewConnection())
		productsService := service.NewService(productsRepository)
		productsController := controller.NewProductController(productsService)

		productRouterGroup.POST("/", productsController.Create())
		productRouterGroup.GET("/", productsController.GetAll())
		productRouterGroup.GET("/:id", productsController.GetById())
		productRouterGroup.PATCH("/:id", productsController.Update())
		productRouterGroup.DELETE("/:id", productsController.Delete())

	}
}
