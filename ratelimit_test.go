// Copyright (C) 2021-2022 Amuzed GmbH finn@amuzed.io.
// This file is part of the project AMUZED.
// AMUZED can not be copied and/or distributed without the express.
// permission of Amuzed GmbH.

package ratelimit_test

import (
	"testing"
	"time"

	ratelimit "github.com/BoostyLabs/rate-limit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	config := ratelimit.Config{LimitForBet: time.Second * 30}
	rateLimiter := ratelimit.NewRateLimiter(config)

	key := uuid.New()

	_ = rateLimiter.SetLimit(key)

	isAllowed := rateLimiter.IsAllowed(key, time.Now().UTC())
	assert.False(t, isAllowed)

	isAllowed = rateLimiter.IsAllowed(key, time.Now().UTC().Add(rateLimiter.GetDuration()-time.Second))
	assert.False(t, isAllowed)

	isAllowed = rateLimiter.IsAllowed(key, time.Now().UTC().Add(rateLimiter.GetDuration()))
	assert.True(t, isAllowed)
}
