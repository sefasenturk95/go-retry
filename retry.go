package retry

import (
	"fmt"
	"time"
)

// Do retries the function f up to maxRetries times until it succeeds.
func Do(maxRetries int, f func() error) error {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		if err := f(); err != nil {
			lastErr = err
			continue
		}
		return nil
	}
	return fmt.Errorf("max retries [%d] reached: %w", maxRetries, lastErr)
}

// WithBackOff retries the function f up to maxRetries times with exponential backoff.
// initialBackOff is the base delay in milliseconds (e.g., 1000 = 1s, 2s, 4s, 8s...).
func WithBackOff(maxRetries int, initialBackOff int, f func() error) error {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff: 1x, 2x, 4x, 8x...
			delay := time.Duration(initialBackOff) * time.Millisecond * (1 << (attempt - 1))
			time.Sleep(delay)
		}
		if err := f(); err != nil {
			lastErr = err
			continue
		}
		return nil
	}
	return fmt.Errorf("max retries [%d] reached: %w", maxRetries, lastErr)
}
