package ratelimit

import (
	"sync"

	"github.com/google/uuid"
)

// EventRateLimiter allows preventing multiple events from events.
type EventRateLimiter struct {
	mu          sync.Mutex
	rateLimited map[uuid.UUID]bool
}

// NewEventRateLimiter is a constructor for NewRateLimiterEvent.
func NewEventRateLimiter() *EventRateLimiter {
	return &EventRateLimiter{
		rateLimited: make(map[uuid.UUID]bool),
	}
}

// IsAllowed indicates if event is allowed to happen.
func (rateLimiter *EventRateLimiter) IsAllowed(key uuid.UUID) bool {
	allow, exists := rateLimiter.rateLimited[key]
	if exists {
		return allow
	}

	return true
}

// SetLimit sets limit from the list of rate limited entities.
func (rateLimiter *EventRateLimiter) SetLimit(userID uuid.UUID) error {
	rateLimiter.mu.Lock()
	defer rateLimiter.mu.Unlock()

	rateLimiter.rateLimited[userID] = false
	return nil
}

// AllowFormEvent allowed event.
func (rateLimiter *EventRateLimiter) AllowFormEvent(key uuid.UUID) {
	rateLimiter.mu.Lock()
	defer rateLimiter.mu.Unlock()

	_, exists := rateLimiter.rateLimited[key]
	if exists {
		rateLimiter.rateLimited[key] = true
	}
}
