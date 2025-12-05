package retry

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	originalErr := errors.New("foo")
	err := Do(5, func() error {
		return originalErr
	})
	assert.Error(t, err)
	assert.ErrorIs(t, err, originalErr)
	assert.Contains(t, err.Error(), "max retries [5] reached")

	err = Do(5, func() error {
		return nil
	})
	assert.NoError(t, err)
}

func TestWithBackOff(t *testing.T) {
	originalErr := errors.New("foo")
	err := WithBackOff(5, 50, func() error {
		return originalErr
	})
	assert.Error(t, err)
	assert.ErrorIs(t, err, originalErr)
	assert.Contains(t, err.Error(), "max retries [5] reached")
}
