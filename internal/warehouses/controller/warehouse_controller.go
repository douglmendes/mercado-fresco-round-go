package controller

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/logger"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type WarehousesController struct {
	service domain.WarehouseService
}

// Create godoc
// @Summary Create warehouses
// @Tags Warehouses
// @Description create one warehouse
// @Accept  json
// @Produce  json
// @Param warehouses body whCreateRequest true "Warehouse to create"
// @Success 201 {object} domain.Warehouse
// @Failure 422 {object} response.Response
// @Router /api/v1/warehouses [post]
func (w *WarehousesController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var whRequest whCreateRequest

		if err := ctx.ShouldBindJSON(&whRequest); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		warehouse, err := w.service.Create(
			ctx,
			whRequest.Address,
			whRequest.Telephone,
			whRequest.WarehouseCode,
			whRequest.LocalityId,
		)
		if err != nil {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, warehouse)
	}
}

// GetAll godoc
// @Summary List warehouses
// @Tags Warehouses
// @Description List all available warehouses
// @Produce  json
// @Success 200 {array} response.Response{data=warehouses.Warehouse} "desc"
// @Failure 404 {object} response.Response
// @Router /api/v1/warehouses [get]
func (w *WarehousesController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		warehousesList, err := w.service.GetAll(ctx)
		if err != nil {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			ctx.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, response.NewResponse(warehousesList))
	}
}

// GetById Warehouse godoc
// @Summary Warehouse
// @Tags Warehouses
// @Description Read one warehouse
// @Accept  json
// @Produce  json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} domain.Warehouse
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /api/v1/warehouses/{id} [get]
func (w *WarehousesController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.DecodeError("id is not valid"))
			return
		}

		warehouse, err := w.service.GetById(ctx, id)
		if err != nil {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			ctx.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(warehouse))
	}
}

// Update godoc
// @Summary Update warehouse
// @Tags Warehouses
// @Description Update a warehouse by ID
// @Accept  json
// @Produce  json
// @Param warehouse body whUpdateRequest true "Warehouse to update"
// @Param id path int true "Warehouse ID"
// @Success 200 {object} domain.Warehouse
// @Failure 404 {object} response.Response
// @Router /api/v1/warehouses/{id} [patch]
func (w *WarehousesController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid"})
			return
		}

		var whRequest whUpdateRequest
		if err := ctx.ShouldBindJSON(&whRequest); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		warehouse, err := w.service.Update(
			ctx,
			id,
			whRequest.Address,
			whRequest.Telephone,
			whRequest.WarehouseCode,
			whRequest.LocalityId,
		)
		if err != nil {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, warehouse)
	}
}

// Delete Warehouse godoc
// @Summary Delete warehouse
// @Tags Warehouses
// @Description Delete a warehouse by ID
// @Param id path int true "Warehouse ID"
// @Success 204 {object} string
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /api/v1/warehouses/{id} [delete]
func (w *WarehousesController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid"})
			return
		}

		err = w.service.Delete(ctx, id)
		if err != nil {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("warehouse with id %d has been removed", id)})
	}
}

func NewWarehouse(w domain.WarehouseService) *WarehousesController {
	return &WarehousesController{
		service: w,
	}
}

type whCreateRequest struct {
	Address       string `json:"address" binding:"required"`
	Telephone     string `json:"telephone" binding:"required"`
	WarehouseCode string `json:"warehouse_code" binding:"required"`
	LocalityId    int    `json:"locality_id" binding:"required"`
}

type whUpdateRequest struct {
	Address       string `json:"address"`
	Telephone     string `json:"telephone"`
	WarehouseCode string `json:"warehouse_code"`
	LocalityId    int    `json:"locality_id"`
}
