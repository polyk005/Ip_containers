package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/polyk005/Ip_containers/back/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/containers", h.GetContainers)
		api.POST("/containers", h.AddContainer)
	}

	return router
}
