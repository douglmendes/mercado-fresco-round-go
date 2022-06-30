package routes

import (
	sellersController "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/controller"
	sellersRepository "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/repository"
	sellerService "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/service"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func SellersRoutes(group *gin.RouterGroup) {

	sellerRouterGroup := group.Group("/sellers")
	{
		sellersDb := store.New(store.FileType, store.PathBuilder("/sellers.json"))
		sellersRepo := sellersRepository.NewRepository(sellersDb)
		sellersService := sellerService.NewService(sellersRepo)
		s := sellersController.NewSeller(sellersService)

		sellerRouterGroup.POST("/", s.Create())
		sellerRouterGroup.GET("/", s.GetAll())
		sellerRouterGroup.GET("/:id", s.GetById())
		sellerRouterGroup.PATCH("/:id", s.Update())
		sellerRouterGroup.DELETE("/:id", s.Delete())

	}
}
