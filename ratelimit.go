// Copyright (C) 2021-2022 Amuzed GmbH finn@amuzed.io.
// This file is part of the project AMUZED.
// AMUZED can not be copied and/or distributed without the express.
// permission of Amuzed GmbH.

package ratelimit

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// RateLimiter allows to prevent multiple events in fixed period of time.
type RateLimiter struct {
	mu          sync.Mutex
	rateLimited map[uuid.UUID]time.Time
	duration    time.Duration
}

// Config defines values needed to start rate limiter.
type Config struct {
	LimitForBet time.Duration `json:"limitForBet"`
}

// NewRateLimiter is a constructor for NewRateLimiter.
func NewRateLimiter(config Config) *RateLimiter {
	return &RateLimiter{
		rateLimited: make(map[uuid.UUID]time.Time),
		duration:    config.LimitForBet,
	}
}

// IsAllowed indicates if event is allowed to happen.
func (rateLimiter *RateLimiter) IsAllowed(key uuid.UUID, now time.Time) bool {
	occursAt, exists := rateLimiter.rateLimited[key]
	if exists {

		return occursAt.Before(now)
	}

	return true
}

// SetLimit sets limit from the list of rate limited entities.
func (rateLimiter *RateLimiter) SetLimit(userID uuid.UUID) error {
	rateLimiter.mu.Lock()
	defer rateLimiter.mu.Unlock()

	occursAt := time.Now().UTC().Add(rateLimiter.duration)
	rateLimiter.rateLimited[userID] = occursAt
	return nil
}

// GetDuration gets limit duration.
func (rateLimiter *RateLimiter) GetDuration() time.Duration {
	return rateLimiter.duration
}
