package main

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	employeesDb := store.New(store.FileType, "employees.json")
	employeesRepo := employees.NewRepository(employeesDb)
	employeesService := employees.NewService(empl)
	router := gin.Default()

	e := controllers.NewEmployees(employeesService)

	emp := router.Group("/api/v1/employees")
	{
		emp.GET("/", e.GetAll())
		emp.GET("/:id", e.GetById())
		emp.POST("/", e.Store())
		emp.PUT("/:id", e.Update())
		emp.DELETE("/:id", e.Delete())
	}
}
