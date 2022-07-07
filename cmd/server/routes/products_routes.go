package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func ProductsRoutes(group *gin.RouterGroup) {

	productRouterGroup := group.Group("/products")
	{
		productsDb := store.New(store.FileType, store.PathBuilder("/products.json"))
		productsRepository := products.NewRepository(productsDb)
		productsService := products.NewService(productsRepository)
		productsController := controllers.NewProductController(productsService)

		productRouterGroup.POST("/", productsController.Create())
		productRouterGroup.GET("/", productsController.GetAll())
		productRouterGroup.GET("/:id", productsController.GetById())
		productRouterGroup.PATCH("/:id", productsController.Update())
		productRouterGroup.DELETE("/:id", productsController.Delete())

	}
}
