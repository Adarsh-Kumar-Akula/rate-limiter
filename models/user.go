package models

import (
	"sync"
	"time"
)

type UserStats struct {
	UserID        string
	TotalRequests int
	Allowed       int
	Blocked       int
}

type UserData struct {
	Timestamps []time.Time
	Stats      UserStats
	Mu         sync.Mutex
}
