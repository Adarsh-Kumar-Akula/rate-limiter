package store

import (
	"rate-limiter/models"
	"sync"
)

type MemoryStore struct {
	users map[string]*models.UserData
	mu    sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: make(map[string]*models.UserData),
	}
}

func (s *MemoryStore) GetUser(userID string) *models.UserData {
	s.mu.RLock()
	user, exists := s.users[userID]
	s.mu.RUnlock()

	if exists {
		return user
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// double check (important for concurrency)
	if user, exists = s.users[userID]; exists {
		return user
	}

	user = &models.UserData{}
	user.Stats.UserID = userID
	s.users[userID] = user
	return user
}
