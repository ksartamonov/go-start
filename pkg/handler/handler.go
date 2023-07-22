package handler

import (
	"github.com/gin-gonic/gin"
	"go-start/pkg/model"
	"go-start/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) RouteHandlers() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		api.POST("/", h.handleSaveRequest)
		api.GET("/:name", h.handleGetByNameRequest)
	}
	return router
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) handleSaveRequest(c *gin.Context) {
	var req model.SaveRequest

	if err := c.BindJSON(&req); err != nil { // validation is done by this func
		model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.DataService.WriteData(req)
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) handleGetByNameRequest(c *gin.Context) {

	name := c.Param("name")

	res, err := h.services.DataService.GetPropertyByName(name)
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
