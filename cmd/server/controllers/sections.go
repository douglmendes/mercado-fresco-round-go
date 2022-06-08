package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sections"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type SectionsController struct {
	service sections.Service
}

func (s *SectionsController) GetAll(c *gin.Context) {
	sections, err := s.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.DecodeError(err.Error()))
		return
	}

	c.JSON(func() int {
		if len(sections) == 0 {
			return http.StatusNoContent
		}
		return http.StatusOK
	}(), response.NewResponse(sections))
}

func (s *SectionsController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
		return
	}

	section, err := s.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(section))
}

func (s *SectionsController) Create(c *gin.Context) {
	var req sectionsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.DecodeError(err.Error()))
		return
	}

	section, err := s.service.Create(
		req.SectionNumber, req.CurrentCapacity, req.MinimumCapacity,
		req.MaximumCapacity, req.WarehouseId, req.ProductTypeId,
		req.CurrentTemperature, req.MinimumTemperature,
	)
	if err != nil {
		c.JSON(http.StatusConflict, response.DecodeError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.NewResponse(section))
}

func (s *SectionsController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
		return
	}

	var args map[string]int
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.DecodeError(err.Error()))
		return
	}

	section, err := s.service.Update(id, args)
	if err != nil {
		c.JSON(func() int {
			errNF := &sections.ErrorNotFound{}
			if errors.As(err, &errNF) {
				return http.StatusNotFound
			}

			errCF := &sections.ErrorConflict{}
			if errors.As(err, &errCF) {
				return http.StatusConflict
			}

			return http.StatusInternalServerError
		}(), response.DecodeError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(section))
}

func (s *SectionsController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
		return
	}

	err = s.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, response.NewResponse(nil))
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
