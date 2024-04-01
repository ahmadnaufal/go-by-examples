package main

import (
	"fmt"
	"time"
)

const (
	MaxBucketSize = 3
	RefillRate    = 1
	// in seconds
	RefillDuration = 2
)

type TokenBucket struct {
	currentBucketSize int
	// should be in nanoseconds
	lastRefillTimestamp time.Time
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func (tb *TokenBucket) refill() {
	nowTime := time.Now()

	dur := nowTime.Sub(tb.lastRefillTimestamp)

	tokensToAdd := (int(dur.Seconds()) / RefillDuration) * RefillRate
	tb.currentBucketSize = min(tb.currentBucketSize+tokensToAdd, MaxBucketSize)
	tb.lastRefillTimestamp = nowTime
}

func (tb *TokenBucket) GrantRequest() bool {
	tb.refill()

	if tb.currentBucketSize < 1 {
		return false
	}

	tb.currentBucketSize--
	return true
}

func main() {
	tb := TokenBucket{
		currentBucketSize: 3,
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("Request %d: %v\n", i+1, tb.GrantRequest())
		time.Sleep(3 * time.Second)
	}
}
