package retry

import "time"

func Do(maxRetries int, f func() error) error {
	retries := 0

	for {
		err := f()

		if err != nil {
			retries++

			if retries >= maxRetries {
				return err
			}

			continue
		}

		return nil
	}
}

func WithBackOff(maxRetries int, f func() error) error {
	retries := 0
	backOffMillis := 500

	for {
		if retries > 0 {
			time.Sleep(time.Duration(int64(backOffMillis) * int64(retries) * int64(time.Millisecond)))
		}

		err := f()

		if err != nil {
			retries++

			if retries >= maxRetries {
				return err
			}

			continue
		}

		return nil
	}
}