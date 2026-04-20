package controller

import "github.com/gin-gonic/gin"

type RequestBody struct {
	UserID  string `json:"user_id"`
	Payload string `json:"payload"`
}

func (h *Handler) HandleRequest(c *gin.Context) {
	var req RequestBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	allowed := h.rateLimiter.Allow(req.UserID)

	if !allowed {
		c.JSON(429, gin.H{"error": "rate limit exceeded"})
		return
	}

	c.JSON(200, gin.H{"message": "request accepted"})
}
