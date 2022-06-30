package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func BuyersRoutes(group *gin.RouterGroup) {

	buyerRouterGroup := group.Group("/buyers")
	{
		buyersDb := store.New(store.FileType, store.PathBuilder("/buyers.json"))
		buyersDbRepo := buyers.NewRepository(buyersDb)
		buyersService := buyers.NewService(buyersDbRepo)
		b := controllers.NewBuyer(buyersService)

		buyerRouterGroup.POST("/", b.Create())
		buyerRouterGroup.GET("/", b.GetAll())
		buyerRouterGroup.GET("/:id", b.GetById())
		buyerRouterGroup.PATCH("/:id", b.Update())
		buyerRouterGroup.DELETE("/:id", b.Delete())

	}
}
