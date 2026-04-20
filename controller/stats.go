package controller

import "github.com/gin-gonic/gin"

func (h *Handler) GetStats(c *gin.Context) {
	userID := c.Query("user_id")

	stats := h.rateLimiter.GetStats(userID)

	c.JSON(200, stats)
}
