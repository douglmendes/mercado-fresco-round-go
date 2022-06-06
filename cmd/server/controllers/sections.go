package controllers

import (
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sections"
	"github.com/gin-gonic/gin"
)

type SectionsController struct {
	service sections.Service
}

func (s *SectionsController) GetAll(c *gin.Context) {
	sections, err := s.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Não foi possível obter as seções",
		})
		return
	}

	c.JSON(func() int {
		if len(sections) == 0 {
			return http.StatusNoContent
		}
		return http.StatusAccepted
	}(), gin.H{"data": sections})
}

func (s *SectionsController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "O ID informado não é válido",
		})
		return
	}

	section, err := s.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Não foi possível obter a seção procurada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": section})
}

func (s SectionsController) Create(c *gin.Context) {
	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   err.Error(),
			"message": "Não foi possível processar os dados da requisição",
		})
		return
	}

	section, err := s.service.Create(
		req.SectionNumber, req.CurrentCapacity, req.MinimumCapacity,
		req.MaximumCapacity, req.WarehouseId, req.ProductTypeId,
		req.CurrentTemperature, req.MinimumTemperature,
	)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error":   err.Error(),
			"message": "Não é possível criar duas seções com o mesmo número",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": section})
}

func NewSectionsController(s sections.Service) *SectionsController {
	return &SectionsController{s}
}

type request struct {
	SectionNumber      int     `json:"section_number" binding:"required"`
	CurrentTemperature float64 `json:"current_temperature" binding:"required"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int     `json:"current_capacity" binding:"required"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int     `json:"maximum_capacity" binding:"required"`
	WarehouseId        int     `json:"warehouse_id" binding:"required"`
	ProductTypeId      int     `json:"product_type_id" binding:"required"`
}
