package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/controller"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/repository/jsondb"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/service"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func ProductsRoutes(group *gin.RouterGroup) {
	productRouterGroup := group.Group("/products")
	{
		productsDb := store.New(store.FileType, store.PathBuilder("/products.json"))
		productsRepository := jsondb.NewRepository(productsDb)
		productsService := service.NewService(productsRepository)
		productsController := controller.NewProductController(productsService)

		productRouterGroup.POST("/", productsController.Create())
		productRouterGroup.GET("/", productsController.GetAll())
		productRouterGroup.GET("/:id", productsController.GetById())
		productRouterGroup.PATCH("/:id", productsController.Update())
		productRouterGroup.DELETE("/:id", productsController.Delete())

	}
}
