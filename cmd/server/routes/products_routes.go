package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	productRecordController "github.com/douglmendes/mercado-fresco-round-go/internal/product_record/controller"
	productRecordMariaDB "github.com/douglmendes/mercado-fresco-round-go/internal/product_record/repository/mariadb"
	productRecordService "github.com/douglmendes/mercado-fresco-round-go/internal/product_record/service"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/controller"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/repository/mariadb"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/service"
	"github.com/gin-gonic/gin"
)

func ProductsRoutes(group *gin.RouterGroup) {
	productRouterGroup := group.Group("/products")
	{
		connection := connections.NewConnection()

		productsRepository := mariadb.NewRepository(connection)
		productsService := service.NewService(productsRepository)
		productsController := controller.NewProductController(productsService)

		productRecordsRepository := productRecordMariaDB.NewRepository(connection)
		productRecordsService := productRecordService.NewProductRecordService(productRecordsRepository, productsRepository)
		productRecordController := productRecordController.NewProductRecordController(productRecordsService)

		productRouterGroup.POST("/", productsController.Create())
		productRouterGroup.GET("/", productsController.GetAll())
		productRouterGroup.GET("/:id", productsController.GetById())
		productRouterGroup.PATCH("/:id", productsController.Update())
		productRouterGroup.DELETE("/:id", productsController.Delete())

		productRouterGroup.GET("/reportRecords", productRecordController.GetByProductId())
	}
}
