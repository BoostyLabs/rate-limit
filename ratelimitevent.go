package ratelimit

import (
	"sync"

	"github.com/google/uuid"
)

// RateLimiterEvent allows to prevent multiple events in fixed period of time.
type RateLimiterEvent struct {
	mu          sync.Mutex
	rateLimited map[uuid.UUID]bool
}

// NewRateLimiterEvent is a constructor for NewRateLimiter.
func NewRateLimiterEvent() *RateLimiterEvent {
	return &RateLimiterEvent{
		rateLimited: make(map[uuid.UUID]bool),
	}
}

// IsAllowed indicates if event is allowed to happen.
func (rateLimiter *RateLimiterEvent) IsAllowed(key uuid.UUID) bool {
	allow, exists := rateLimiter.rateLimited[key]
	if exists {
		return allow
	}

	return true
}

// SetLimit sets limit from the list of rate limited entities.
func (rateLimiter *RateLimiterEvent) SetLimit(userID uuid.UUID) error {
	rateLimiter.mu.Lock()
	defer rateLimiter.mu.Unlock()

	rateLimiter.rateLimited[userID] = false
	return nil
}

// AllowFormEvent allowed event.
func (rateLimiter *RateLimiterEvent) AllowFormEvent(key uuid.UUID) {
	rateLimiter.mu.Lock()
	defer rateLimiter.mu.Unlock()

	_, exists := rateLimiter.rateLimited[key]
	if exists {
		rateLimiter.rateLimited[key] = true
	}
}
