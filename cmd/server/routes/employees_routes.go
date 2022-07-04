package routes

import (
	"github.com/douglmendes/mercado-fresco-round-go/connections"
	employeesController "github.com/douglmendes/mercado-fresco-round-go/internal/employees/controller"
	employeesRepository "github.com/douglmendes/mercado-fresco-round-go/internal/employees/repository"
	employeesService "github.com/douglmendes/mercado-fresco-round-go/internal/employees/service"
	"github.com/gin-gonic/gin"
)

func EmployeesRoutes(group *gin.RouterGroup) {

	employeeRouterGroup := group.Group("/employees")
	{
		employeesRepo := employeesRepository.NewRepository(connections.NewConnection())
		employeeService := employeesService.NewService(employeesRepo)
		e := employeesController.NewEmployees(employeeService)

		employeeRouterGroup.POST("/", e.Create())
		employeeRouterGroup.GET("/", e.GetAll())
		employeeRouterGroup.GET("/:id", e.GetById())
		employeeRouterGroup.PATCH("/:id", e.Update())
		employeeRouterGroup.DELETE("/:id", e.Delete())

	}
}
