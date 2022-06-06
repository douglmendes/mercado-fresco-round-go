package controllers

import (
	"net/http"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sections"
	"github.com/gin-gonic/gin"
)

type SectionsController struct {
	service sections.Service
}

func (s SectionsController) GetAll(c *gin.Context) {
	sections, err := s.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   err.Error(),
				"message": "Não foi possível obter as seções",
			},
		)
		return
	}

	c.JSON(func() int {
		if len(sections) == 0 {
			return http.StatusNoContent
		}
		return http.StatusAccepted
	}(), sections)
}

func NewSectionsController(s sections.Service) *SectionsController {
	return &SectionsController{s}
}
