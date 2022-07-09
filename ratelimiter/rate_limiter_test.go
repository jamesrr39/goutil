package ratelimiter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RateLimiter(t *testing.T) {
	// test takes part within a fake "minute". We start from 0 and set the timeNowSeconds as the test goes on
	durationsList := []time.Duration{time.Second, time.Second * 3}
	var timeNowSeconds int
	nowFunc := func() time.Time {
		return time.Date(2000, 1, 1, 1, 1, timeNowSeconds, 0, time.UTC)
	}

	key1 := "test1"

	rateLimiter, err := NewRateLimiter(durationsList, nowFunc)
	require.NoError(t, err)

	t.Run("test first available", func(t *testing.T) {
		avail, err := rateLimiter.IsAvailable(key1)
		require.NoError(t, err)
		assert.True(t, avail)
	})

	t.Run("register entry and check limiter is not available", func(t *testing.T) {
		err = rateLimiter.RegisterEntry(key1)
		require.NoError(t, err)

		status := rateLimiter.GetStatus(key1)
		assert.Equal(t, int64(1), status.CountEntries)

		avail, err := rateLimiter.IsAvailable(key1)
		require.NoError(t, err)
		assert.False(t, avail)

		assert.Equal(t, int64(1), status.CountEntries)
	})

	t.Run("update time now to '2' seconds, rate limiter should be ready again", func(t *testing.T) {
		timeNowSeconds = 2
		avail, err := rateLimiter.IsAvailable(key1)
		require.NoError(t, err)
		assert.True(t, avail)

		// register 2 entries.
		err = rateLimiter.RegisterEntry(key1)
		require.NoError(t, err)

		status := rateLimiter.GetStatus(key1)

		assert.Equal(t, int64(1), status.CountEntries)

		err = rateLimiter.RegisterEntry(key1)
		require.NoError(t, err)

		assert.Equal(t, int64(1), status.CountEntries)

		// Since time now is at "2" seconds when the second one is set, it should be ready at "5 seconds"
		timeNowSeconds = 4
		avail, err = rateLimiter.IsAvailable(key1)
		require.NoError(t, err)
		assert.False(t, avail)

		timeNowSeconds = 5
		avail, err = rateLimiter.IsAvailable(key1)
		require.NoError(t, err)
		assert.True(t, avail)
	})
}

func Test_uses_last_value_after_length_of_list_has_been_reached(t *testing.T) {
	const lastTimeSeconds = 10
	durationsList := []time.Duration{time.Second, time.Second * lastTimeSeconds}
	timeNowSeconds := 0
	nowFunc := func() time.Time {
		return time.Date(2000, 1, 1, 1, 1, timeNowSeconds, 0, time.UTC)
	}

	key1 := "test1"

	rateLimiter, err := NewRateLimiter(durationsList, nowFunc)
	require.NoError(t, err)

	require.Len(t, durationsList, 2, "if durations list has been expanded, the amount of calls to RegisterEntry should be increased")

	err = rateLimiter.RegisterEntry(key1)
	require.NoError(t, err)

	err = rateLimiter.RegisterEntry(key1)
	require.NoError(t, err)

	err = rateLimiter.RegisterEntry(key1)
	require.NoError(t, err)

	err = rateLimiter.RegisterEntry(key1)
	require.NoError(t, err)

	status := rateLimiter.GetStatus(key1)
	assert.Equal(t, time.Date(2000, 1, 1, 1, 1, lastTimeSeconds, 0, time.UTC), status.NextAvailableTime)
}
