package handler

import (
	"github.com/gin-contrib/cors"
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

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")
	// Register the middleware
	router.Use(cors.New(corsConfig))

	api := router.Group("/api/v1")
	{
		api.POST("/", h.handleSaveRequest)
		api.GET("/:name", h.handleGetParameterValueRequest)
		api.GET("/", h.handleGetByPairRequest)
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

func (h *Handler) handleGetParameterValueRequest(c *gin.Context) {

	name := c.Param("name")

	res, err := h.services.DataService.GetParameterValue(name)
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"parameter": name,
		"values":    res,
	})
}

func (h *Handler) handleGetByPairRequest(c *gin.Context) {
	param := c.Query("parameter")
	value := c.Query("value")

	req := model.Property{
		Parameter: param,
		Value:     value,
	}

	//if err := c.BindJSON(&req); err != nil { // validation is done by this func
	//	model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}

	res, err := h.services.DataService.GetByPair(req)
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
