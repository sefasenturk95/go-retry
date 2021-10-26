package retry

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDo(t *testing.T) {
	err := Do(5, func() error {
		return errors.New("foo")
	})
	assert.Error(t, err)
	assert.Equal(t, errors.New("max retries [5] reached with error: foo"), err)

	err = Do(5, func() error {
		return nil
	})
	assert.NoError(t, err)
}

func TestWithBackOff(t *testing.T) {
	err := WithBackOff(5, func() error {
		return errors.New("foo")
	})
	assert.Error(t, err)
	assert.Equal(t, errors.New("max retries with back off [5] reached with error: foo"), err)
}