package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/employees"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type EmployeesController struct {
	service employees.Service
}

type requestEmployee struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

func NewEmployees(e employees.Service) *EmployeesController {
	return &EmployeesController{
		service: e,
	}
}

// ListEmployees godoc
// @Summary List all employees
// @Tags Employees
// @Description get all employees
// @Produce  json
// @Success 200 {object} request
// @Failure 404 {object} string
// @Router /api/v1/employees [get]
func (c *EmployeesController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response.NewResponse(e))
	}
}

// ListEmployees godoc
// @Summary List employee
// @Tags Employees
// @Description get employees
// @Accept  json
// @Produce  json
// @Param id   path int true "Employee ID"
// @Success 200 {object} request
// @Failure 404 {object} string
// @Router /api/v1/employees/{id} [get]

func (c *EmployeesController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		e, err := c.service.GetById(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(e))
	}
}

// Create godoc
// @Summary Create employees
// @Tags Employees
// @Description create employees
// @Accept  json
// @Produce  json
// @Param product body request true "Employee to create"
// @Success 201 {object} string
// @Failure 422 {object} string
// @Failure 409 {object} string
// @Router /api/v1/employees [post]

func (c *EmployeesController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestEmployee
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":  "VALIDATEERR-1",
					"messge": "Invalid inputs. Please check your inputs",
				})
			return
		}
		e, err := c.service.Create(req.Id, req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(e))
	}

}

// ListEmployees godoc
// @Summary Update employee
// @Tags Employees
// @Description update employee
// @Accept  json
// @Produce  json
// @Param product body request true "Employee to create"
// @Param id   path int true "Employee ID"
// @Success 200 {object} request
// @Router /api/v1/employees/{id} [patch]
func (c *EmployeesController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		var req requestEmployee
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error})
			return
		}
		e, err := c.service.Update(int(id), req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, e)

	}

}

// ListEmployees godoc
// @Summary Delete employee
// @Tags Employees
// @Description delete employee
// @Param id   path int true "Employee ID"
// @Success 204 {object} request
// @Router /api/v1/employees/{id} [delete]
func (c *EmployeesController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("employee %d was removed", id)})
	}
}
