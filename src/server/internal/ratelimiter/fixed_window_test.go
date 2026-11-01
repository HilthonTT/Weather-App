package ratelimiter

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestFixedWindowLimiter_AllowsUpToLimit(t *testing.T) {
	rl := NewFixedWindowLimiter(3, time.Minute)

	for i := range 3 {
		if ok, _ := rl.Allow("a"); !ok {
			t.Fatalf("request %d should be allowed", i)
		}
	}

	ok, retry := rl.Allow("a")
	if ok {
		t.Fatal("4th request should be rejected")
	}
	if retry != time.Minute {
		t.Errorf("retry = %v, want %v", retry, time.Minute)
	}

	if ok, _ := rl.Allow("b"); !ok {
		t.Error("other clients must have their own bucket")
	}
}

func TestFixedWindowLimiter_ResetsAfterWindow(t *testing.T) {
	rl := NewFixedWindowLimiter(1, 20*time.Millisecond)

	rl.Allow("a")
	if ok, _ := rl.Allow("a"); ok {
		t.Fatal("second request should be rejected")
	}

	time.Sleep(50 * time.Millisecond)

	if ok, _ := rl.Allow("a"); !ok {
		t.Error("request after the window should be allowed")
	}
}

func TestFixedWindowLimiter_Concurrent(t *testing.T) {
	const limit = 10
	rl := NewFixedWindowLimiter(limit, time.Minute)

	var allowed atomic.Int32
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if ok, _ := rl.Allow("a"); ok {
				allowed.Add(1)
			}
		}()
	}
	wg.Wait()

	if got := allowed.Load(); got != limit {
		t.Errorf("allowed %d concurrent requests, want exactly %d", got, limit)
	}
}
