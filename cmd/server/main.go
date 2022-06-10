package main

import (
	"log"
	"os"

	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/docs"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Mercado Fresco
// @version         1.0
// @description     This API Handle MELI fresh products
// @termsOfService  https://developers.mercadolivre.com.br/pt_br/termos-e-condicoes

// @contact.name  API Support
// @contact.url   https://developers.mercadolivre.com.br/support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load(store.PathBuilder("/.env"))
	if err != nil {
		log.Fatal("failed to load .env")
	}

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	sellersDb := store.New(store.FileType, store.PathBuilder("/sellers.json"))
	sellersRepo := sellers.NewRepository(sellersDb)
	sellersService := sellers.NewService(sellersRepo)

	s := controllers.NewSeller(sellersService)

	sl := router.Group("/api/v1/sellers")
	{
		sl.GET("/", s.GetAll())
		sl.GET("/:id", s.GetById())
		sl.POST("/", s.Create())
		sl.PATCH("/:id", s.Update())
		sl.DELETE("/:id", s.Delete())
	}

	sectionsRepository := sections.NewRepository(store.New(store.FileType, "../../sections.json"))
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

	warehousesDB := store.New(store.FileType, store.PathBuilder("/warehouses.json"))
	warehousesRepo := warehouses.NewRepository(warehousesDB)
	warehousesService := warehouses.NewService(warehousesRepo)
	whController := controllers.NewWareHouse(warehousesService)

	wh := router.Group("/api/v1/warehouses")
	{
		wh.POST("/", whController.Create())
		wh.GET("/", whController.GetAll())
		wh.GET("/:id", whController.GetById())
		wh.PATCH("/:id", whController.Update())
		wh.DELETE("/:id", whController.Delete())
	}

	productsDb := store.New(store.FileType, "products.json")
	productsRepository := products.NewRepository(productsDb)
	productsService := products.NewService(productsRepository)
	productsController := controllers.NewProductController(productsService)

	productsRoutes := router.Group("/api/v1/products")
	{
		productsRoutes.GET("/", productsController.GetAll())
		productsRoutes.GET("/:id", productsController.GetById())
		productsRoutes.POST("/", productsController.Create())
		productsRoutes.PATCH("/:id", productsController.Update())
		productsRoutes.DELETE("/:id", productsController.Delete())
	}

	employeesDb := store.New(store.FileType, store.PathBuilder("/employees.json"))
	employeesRepo := employees.NewRepository(employeesDb)
	employeesService := employees.NewService(employeesRepo)

	e := controllers.NewEmployees(employeesService)

	emp := router.Group("/api/v1/employees")
	{
		emp.GET("/", e.GetAll())
		emp.GET("/:id", e.GetById())
		emp.POST("/", e.Create())
		emp.PATCH("/:id", e.Update())
		emp.DELETE("/:id", e.Delete())
	}

	buyersDb := store.New(store.FileType, "buyers.json")
	buyersDbRepo := buyers.NewRepository(buyersDb)
	buyersService := buyers.NewService(buyersDbRepo)

	b := controllers.NewBuyer(buyersService)

	buy := router.Group("/api/v1/buyers")
	{
		buy.GET("/", b.GetAll())
		buy.GET("/:id", b.GetById())
		buy.POST("/", b.Create())
		buy.PATCH("/:id", b.Update())
		buy.DELETE("/:id", b.Delete())
	}

	router.Run()
}
