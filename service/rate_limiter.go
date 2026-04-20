package service

import (
	"rate-limiter/models"
	"rate-limiter/store"
	"time"
)

type RateLimiter struct {
	store  *store.MemoryStore
	limit  int
	window time.Duration
}

func NewRateLimiter(store *store.MemoryStore, limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		store:  store,
		limit:  limit,
		window: window,
	}
}

func (r *RateLimiter) Allow(userID string) bool {
	user := r.store.GetUser(userID)

	user.Mu.Lock()
	defer user.Mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-r.window)

	// 1. Remove old timestamps
	var valid []time.Time
	for _, t := range user.Timestamps {
		if t.After(cutoff) {
			valid = append(valid, t)
		}
	}
	user.Timestamps = valid

	user.Stats.TotalRequests++

	// 2. Check limit
	if len(user.Timestamps) >= r.limit {
		user.Stats.Blocked++
		return false
	}

	// 3. Allow request
	user.Timestamps = append(user.Timestamps, now)
	user.Stats.Allowed++

	return true
}

func (r *RateLimiter) GetStats(userID string) models.UserStats {
	user := r.store.GetUser(userID)

	user.Mu.Lock()
	defer user.Mu.Unlock()

	return user.Stats
}
