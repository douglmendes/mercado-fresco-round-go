package controller

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain"
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type EmployeesController struct {
	service domain.Service
}

type requestEmployee struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

func NewEmployees(e domain.Service) *EmployeesController {
	return &EmployeesController{
		service: e,
	}
}

// ListEmployees godoc
// @Summary      List employees
// @Tags         employees
// @Description  get employees
// @Produce      json
// @Success      200  {object}  request
// @Failure      404  {object}  string
// @Router       /api/v1/employees [get]
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

// GetEmployee godoc
// @Summary      Get a employee by id
// @Description  Get a employee from the system searching by id
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Employee id"
// @Success      200  {object}  employees.Employee
// @Failure      404  {object}  string
// @Router       /api/v1/employees/{id} [get]
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

// CreateEmployee godoc
// @Summary      Create a new employee
// @Description  Create a new employee in the system
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        employee  body      requestEmployee  true  "Employee to be created"
// @Success      201       {object}  employees.Employee
// @Failure      409       {object}  string
// @Failure      422       {object}  string
// @Router       /api/v1/employees [post]
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
		e, err := c.service.Create(req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, response.NewResponse(e))
	}

}

// ListEmployees godoc
// @Summary      Update employee
// @Tags         employees
// @Description  update employee
// @Accept       json
// @Produce      json
// @Param        product  body      requestEmployee  true  "Employee to create"
// @Param        id       path      int      true  "Employee ID"
// @Success      200      {object}  employees.Employee
// @Failure      400       {object}  string
// @Failure      404       {object}  string
// @Router       /api/v1/employees/{id} [patch]
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

// DeleteEmployee godoc
// @Summary      Delete a employee
// @Description  Delete a employee from the system, selecting by id
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Employee id"
// @Success      204  "Successfully deleted"
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Router       /api/v1/employees/{id} [delete]
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
