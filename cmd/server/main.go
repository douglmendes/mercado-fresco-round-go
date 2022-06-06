package main

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main()  {
	
	sellersDb := store.New(store.FileType, "sellers.json")
	sellersRepo := sellers.NewRepository(sellersDb)
	sellersService := sellers.NewService(sellersRepo)

	s := controllers.NewSeller(sellersService)

	r := gin.Default()

	sl := r.Group("/api/v1/sellers")
	{
		sl.GET("/", s.GetAll())
		sl.GET("/:id", s.GetById())
		sl.POST("/", s.Store())
		sl.PATCH("/:id", s.Update())
		sl.DELETE("/:id", s.Delete())
	}

	r.Run()
}