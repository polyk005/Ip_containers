package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetContainers(c *gin.Context) {
	containers, err := h.services.Container.GetContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, containers)
}

func (h *Handler) AddContainer(c *gin.Context) {
	var input struct {
		IPAddress          string `json:"ip_address"`
		PingTime           string `json:"ping_time"`
		LastSuccessfulPing string `json:"last_successful_ping"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.services.Container.AddContainer(input.IPAddress, input.PingTime, input.LastSuccessfulPing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
