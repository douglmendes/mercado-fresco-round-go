package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/controller"
	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/repository"
	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/service"
	"github.com/gin-gonic/gin"
)

func LocalitiesRoutes(group *gin.RouterGroup)  {

	localityRouterGroup := group.Group("/localities")
	{
		localitiesDb := connections.NewConnection()
		localitiesRepo := repository.NewRepository(localitiesDb)
		localitiesService := service.NewService(localitiesRepo)
		l := controller.NewLocality(localitiesService)

		localityRouterGroup.POST("/", l.Create())
		localityRouterGroup.GET("/", l.GetAll())
		localityRouterGroup.GET("/:id", l.GetById())
		localityRouterGroup.PATCH("/:id", l.Update())
		localityRouterGroup.DELETE("/:id", l.Delete())
	}
}