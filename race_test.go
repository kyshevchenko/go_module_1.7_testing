package stringutil

import "testing"

// TestConcurrentCounter — проверка ConcurrentCounter под нагрузкой.
// Долгий тест: пропускается в short-режиме (go test -short).
func TestConcurrentCounter(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping race-sensitive test in short mode")
	}

	const (
		goroutines = 100
		attempts   = 10
	)

	for i := 0; i < attempts; i++ {
		got := ConcurrentCounter(goroutines)
		if got != goroutines {
			t.Fatalf("attempt %d: ConcurrentCounter(%d) = %d, want %d",
				i, goroutines, got, goroutines)
		}
	}
}
