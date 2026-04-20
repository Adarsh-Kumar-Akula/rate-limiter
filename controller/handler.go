package controller

import (
	"rate-limiter/service"
)

type Handler struct {
	rateLimiter *service.RateLimiter
}

func NewHandler(rl *service.RateLimiter) *Handler {
	return &Handler{
		rateLimiter: rl,
	}
}
