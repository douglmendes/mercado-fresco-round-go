package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	sellersController "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/controller"
	sellersRepository "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/repository"
	localityRepository "github.com/douglmendes/mercado-fresco-round-go/internal/localities/repository"
	sellerService "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/service"
	"github.com/gin-gonic/gin"
)

func SellersRoutes(group *gin.RouterGroup) {

	sellerRouterGroup := group.Group("/sellers")
	{
		sellersDb := connections.NewConnection()
		sellersRepo := sellersRepository.NewRepository(sellersDb)
		localityRepo := localityRepository.NewRepository(sellersDb)
		sellersService := sellerService.NewService(sellersRepo, localityRepo)
		s := sellersController.NewSeller(sellersService)

		sellerRouterGroup.POST("/", s.Create())
		sellerRouterGroup.GET("/", s.GetAll())
		sellerRouterGroup.GET("/:id", s.GetById())
		sellerRouterGroup.PATCH("/:id", s.Update())
		sellerRouterGroup.DELETE("/:id", s.Delete())
	}
}
