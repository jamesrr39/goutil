package ratelimiter

import (
	"errors"
	"sync"
	"time"
)

type limiterEntryType struct {
	CountEntries      int64
	NextAvailableTime time.Time
}

type RateLimiter struct {
	delaySequence []time.Duration
	// map[key]sequenceEntry
	m       map[string]limiterEntryType
	mu      *sync.RWMutex
	nowFunc func() time.Time // for tests. Use time.Now() for normal use
}

// NewRateLimiter creates a new rate limiter
// delaySequence should be a list of delay durations (and after the last one has been reached, the limiter will keep using the last one)
// nowFunc is a function that returns the current time, and is provided for convience of testing. Normally, you should pass `time.Now` in here.
func NewRateLimiter(delaySequence []time.Duration, nowFunc func() time.Time) (*RateLimiter, error) {
	if len(delaySequence) == 0 {
		return nil, errors.New("delay sequence must have at least one entry")
	}

	return &RateLimiter{
		delaySequence: delaySequence,
		m:             make(map[string]limiterEntryType),
		mu:            new(sync.RWMutex),
		nowFunc:       nowFunc,
	}, nil
}

func (s *RateLimiter) RegisterEntry(key string) error {
	if key == "" {
		return errors.New("must have a non-zero length key")
	}

	now := s.nowFunc()
	s.mu.Lock()
	defer s.mu.Unlock()

	existingSequenceEntry := s.m[key]
	var amountOfTimeToAdd time.Duration
	if existingSequenceEntry.CountEntries < int64(len(s.delaySequence)) {
		// has not reached the max yet
		amountOfTimeToAdd = s.delaySequence[existingSequenceEntry.CountEntries]
	} else {
		// has reached the max, so keep using the max
		amountOfTimeToAdd = s.delaySequence[len(s.delaySequence)-1]
	}

	s.m[key] = limiterEntryType{
		CountEntries:      existingSequenceEntry.CountEntries + 1,
		NextAvailableTime: now.Add(amountOfTimeToAdd),
	}

	return nil
}

func (s *RateLimiter) IsAvailable(key string) (bool, error) {
	if key == "" {
		return false, errors.New("must have a non-zero length key")
	}

	now := s.nowFunc()
	s.mu.RLock()
	defer s.mu.RUnlock()

	existingEntry, ok := s.m[key]
	if !ok {
		// no rate limiter set for this key
		return true, nil
	}

	if existingEntry.NextAvailableTime.After(now) {
		return false, nil
	}

	// rate limiter was set, but it has expired. We can delete it.
	delete(s.m, key)
	return true, nil
}

func (s *RateLimiter) GetStatus(key string) limiterEntryType {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.m[key]
}
