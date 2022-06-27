package main

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/docs"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers"
	employeesController "github.com/douglmendes/mercado-fresco-round-go/internal/employees/controller"
	employeesRepository "github.com/douglmendes/mercado-fresco-round-go/internal/employees/repository"
	employeesService "github.com/douglmendes/mercado-fresco-round-go/internal/employees/service"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections"
	sellersController "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/controller"
	sellersRepository "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/repository"
	sellersService "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/service"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func Start() {
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	sellersDb := store.New(store.FileType, store.PathBuilder("/sellers.json"))
	sellersRepo := sellersRepository.NewRepository(sellersDb)
	sellersService := sellersService.NewService(sellersRepo)

	s := sellersController.NewSeller(sellersService)

	sl := router.Group("/api/v1/sellers")
	{
		sl.GET("/", s.GetAll())
		sl.GET("/:id", s.GetById())
		sl.POST("/", s.Create())
		sl.PATCH("/:id", s.Update())
		sl.DELETE("/:id", s.Delete())
	}

	sectionsRepository := sections.NewRepository(store.New(store.FileType, store.PathBuilder("/sections.json")))
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
	whController := controllers.NewWarehouse(warehousesService)

	wh := router.Group("/api/v1/warehouses")
	{
		wh.POST("/", whController.Create())
		wh.GET("/", whController.GetAll())
		wh.GET("/:id", whController.GetById())
		wh.PATCH("/:id", whController.Update())
		wh.DELETE("/:id", whController.Delete())
	}

	productsDb := store.New(store.FileType, store.PathBuilder("/products.json"))
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
	employeesRepo := employeesRepository.NewRepository(employeesDb)
	employeesService := employeesService.NewService(employeesRepo)

	e := employeesController.NewEmployees(employeesService)

	emp := router.Group("/api/v1/employees")
	{
		emp.GET("/", e.GetAll())
		emp.GET("/:id", e.GetById())
		emp.POST("/", e.Create())
		emp.PATCH("/:id", e.Update())
		emp.DELETE("/:id", e.Delete())
	}

	buyersDb := store.New(store.FileType, store.PathBuilder("/buyers.json"))
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
