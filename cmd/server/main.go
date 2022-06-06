package main

import (
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	employeesDb := store.New(store.FileType, "../../employees.json")
	employeesRepo := employees.NewRepository(employeesDb)
	employeesService := employees.NewService(employeesRepo)

	e := controllers.NewEmployees(employeesService)

	emp := router.Group("/api/v1/employees")
	{
		emp.GET("/", e.GetAll())
		emp.GET("/:id", e.GetById())
		emp.POST("/", e.Store())
		emp.PATCH("/:id", e.Update())
		emp.DELETE("/:id", e.Delete())
	}

	router.Run()
}
