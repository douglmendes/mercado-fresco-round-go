package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/controller"
	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/repository/mariadb"
	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/service"
	productsMariaDB "github.com/douglmendes/mercado-fresco-round-go/internal/products/repository/mariadb"
	"github.com/gin-gonic/gin"
)

func ProductRecordsRoutes(group *gin.RouterGroup) {
	productRecordsRouterGroup := group.Group("/productRecords")
	{
		connection := connections.NewConnection()

		productsRepository := productsMariaDB.NewRepository(connection)

		productRecordsRepository := mariadb.NewRepository(connection)
		productRecordsService := service.NewProductRecordService(productRecordsRepository, productsRepository)
		productRecordController := controller.NewProductRecordController(productRecordsService)

		productRecordsRouterGroup.POST("/", productRecordController.Create())
	}
}
