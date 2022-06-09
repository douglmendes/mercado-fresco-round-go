package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type WareHouseController struct {
	service warehouses.Service
}

func (w *WareHouseController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var whRequest whRequest

		if err := ctx.Bind(&whRequest); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		warehouse, err := w.service.Create(
			whRequest.Address,
			whRequest.Telephone,
			whRequest.WarehouseCode,
			whRequest.MinimunCapacity,
			whRequest.MinimunTemperature,
		)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, warehouse)
	}
}

func (w *WareHouseController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		warehousesList, err := w.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, response.NewResponse(warehousesList))
	}
}

func (w *WareHouseController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.DecodeError("id is not valid"))
			return
		}

		warehouse, err := w.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, response.NewResponse(warehouse))
	}
}

func (w *WareHouseController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid"})
			return
		}

		var whRequest whRequest
		if err := ctx.ShouldBindJSON(&whRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		warehouse, err := w.service.Update(
			id,
			whRequest.Address,
			whRequest.Telephone,
			whRequest.WarehouseCode,
			whRequest.MinimunCapacity,
			whRequest.MinimunTemperature,
		)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, warehouse)
	}
}

func (w *WareHouseController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid"})
			return
		}

		err = w.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("warehouse with id %d hs been removed", id)})
	}
}

func NewWareHouse(w warehouses.Service) *WareHouseController {
	return &WareHouseController{
		service: w,
	}
}

type whRequest struct {
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}
