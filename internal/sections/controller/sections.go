package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type SectionsController struct {
	service domain.Service
}

// ListSections godoc
// @Summary      List all sections
// @Description  List all sections currently in the system
// @Tags         sections
// @Accept       json
// @Produce      json
// @Success      200  {array}  sections.Section
// @Success      204  "Empty database"
// @Failure      500  {object}  response.Response
// @Router       /api/v1/sections [get]
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

// GetSection godoc
// @Summary      Get a section by id
// @Description  Get a section from the system searching by id
// @Tags         sections
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Section id"
// @Success      200  {object}  sections.Section
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /api/v1/sections/{id} [get]
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

// CreateSection godoc
// @Summary      Create a new section
// @Description  Create a new section in the system
// @Tags         sections
// @Accept       json
// @Produce      json
// @Param        section  body      sectionsRequest  true  "Section to be created"
// @Success      201      {object}  sections.Section
// @Failure      409      {object}  response.Response
// @Failure      422      {object}  response.Response
// @Router       /api/v1/sections [post]
func (s *SectionsController) Create(c *gin.Context) {
	var req sectionsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.DecodeError(err.Error()))
		return
	}

	section, err := s.service.Create(
		req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature,
		req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity,
		req.WarehouseId, req.ProductTypeId,
	)
	if err != nil {
		c.JSON(http.StatusConflict, response.DecodeError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.NewResponse(section))
}

// UpdateSection godoc
// @Summary      Update a section
// @Description  Update a section in the system, selecting by id
// @Tags         sections
// @Accept       json
// @Produce      json
// @Param        id       path      int              true   "Section id"
// @Param        section  body      sectionsRequest  false  "Section to be updated (all fields are optional)"
// @Success      200      {object}  sections.Section
// @Failure      400      {object}  response.Response
// @Failure      404      {object}  response.Response
// @Failure      409      {object}  response.Response
// @Failure      422      {object}  response.Response
// @Failure      500      {object}  response.Response
// @Router       /api/v1/sections/{id} [patch]
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
			errNF := &domain.ErrorNotFound{}
			if errors.As(err, &errNF) {
				return http.StatusNotFound
			}

			errCF := &domain.ErrorConflict{}
			if errors.As(err, &errCF) {
				return http.StatusConflict
			}

			return http.StatusInternalServerError
		}(), response.DecodeError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(section))
}

// DeleteSection godoc
// @Summary      Delete a section
// @Description  Delete a section from the system, selecting by id
// @Tags         sections
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Section id"
// @Success      204  "Successfully deleted"
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /api/v1/sections/{id} [delete]
func (s *SectionsController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.DecodeError(err.Error()))
		return
	}

	_, err = s.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.DecodeError(err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, response.NewResponse(nil))
}

func NewSectionsController(s domain.Service) *SectionsController {
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
