package auth

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSleepWithContext_ErrorSource(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()
	err := sleepWithContext(ctx, time.Minute)
	assert.Equal(t, "context canceled", err.Error())
}
