package handler

import (
	"constester-go/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitEndpoints() *gin.Engine {

	r := gin.New()

	//API ENDPOINTS
	api := r.Group("/api")
	{
		api.GET("/test", h.TestEndpoint)
		api.GET("/code", h.TestingCodeEndpoint)
		api.POST("/create", h.AddTaskEndpoint)
	}

	return r
}