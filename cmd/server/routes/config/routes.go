package config

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"os"
)

func ConfigurationRoutes(router *gin.Engine) *gin.Engine {

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseUrl := router.Group("/api/v1/")
	{
		routes.BuyersRoutes(baseUrl)
		routes.EmployeesRoutes(baseUrl)
		routes.ProductsRoutes(baseUrl)
		routes.SectionsRoutes(baseUrl)
		routes.SellersRoutes(baseUrl)
		routes.WarehousesRoutes(baseUrl)
		routes.LocalitiesRoutes(baseUrl)
	}

	return router
}
