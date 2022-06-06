package controllers

import (
	"errors"
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
		return http.StatusOK
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

func (s *SectionsController) Create(c *gin.Context) {
	var req sectionsRequest

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

func (s *SectionsController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "O ID informado não é válido",
		})
		return
	}

	var args map[string]int
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   err.Error(),
			"message": "Não foi possível processar os dados da requisição",
		})
		return
	}

	section, err := s.service.Update(id, args)
	if err != nil {
		errNF := &sections.ErrorNotFound{}
		errCF := &sections.ErrorConflict{}
		if errors.As(err, &errNF) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "A seção que deseja atualizar não existe",
			})
			return
		}
		if errors.As(err, &errCF) {
			c.JSON(http.StatusConflict, gin.H{
				"error":   err.Error(),
				"message": "Não é possível criar duas seções com o mesmo número",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Não foi possível atualizar a seção procurada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": section})
}

func (s *SectionsController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "O ID informado não é válido",
		})
		return
	}

	err = s.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Não foi possível remover a seção procurada",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func NewSectionsController(s sections.Service) *SectionsController {
	return &SectionsController{s}
}

type sectionsRequest struct {
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required"`
	WarehouseId        int `json:"warehouse_id" binding:"required"`
	ProductTypeId      int `json:"product_type_id" binding:"required"`
}
