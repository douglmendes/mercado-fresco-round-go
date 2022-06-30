package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	warehouseController "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/controller"
	warehouseRepository "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/repository"
	warehouseService "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/service"
	"github.com/gin-gonic/gin"
)

func WarehousesRoutes(group *gin.RouterGroup) {

	warehouseRouterGroup := group.Group("/warehouses")
	{
		warehousesRepo := warehouseRepository.NewRepository(connections.NewConnection())
		warehousesService := warehouseService.NewService(warehousesRepo)
		whController := warehouseController.NewWarehouse(warehousesService)

		warehouseRouterGroup.POST("/", whController.Create())
		warehouseRouterGroup.GET("/", whController.GetAll())
		warehouseRouterGroup.GET("/:id", whController.GetById())
		warehouseRouterGroup.PATCH("/:id", whController.Update())
		warehouseRouterGroup.DELETE("/:id", whController.Delete())

	}
}
