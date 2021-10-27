package retry

import (
	"fmt"
	"time"
)

func Do(maxRetries int, f func() error) error {
	retries := 0
	for {
		if err := f(); err != nil {
			retries++
			if retries >= maxRetries {
				return fmt.Errorf("max retries [%v] reached with error: %v", retries, err.Error())
			}
			continue
		}
		return nil
	}
}

func WithBackOff(maxRetries int, initialBackOff int, f func() error) error {
	retries := 0
	for {
		if retries > 0 {
			time.Sleep(time.Duration(int64(initialBackOff) * int64(retries) * int64(time.Millisecond)))
		}
		if err := f(); err != nil {
			retries++
			if retries >= maxRetries {
				return fmt.Errorf("max retries with back off [%v] reached with error: %v", retries, err.Error())
			}
			continue
		}
		return nil
	}
}