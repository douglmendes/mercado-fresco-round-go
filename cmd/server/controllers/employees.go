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

type request struct {
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

func (c *EmployeesController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, e))
	}
}

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

		ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, e))
	}
}

func (c *EmployeesController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":  "VALIDATEERR-1",
					"messge": "Invalid inputs. Please check your inputs",
				})
			return
		}
		e, err := c.service.Store(req.Id, req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, e))
	}

}

func (c *EmployeesController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		var req request
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
