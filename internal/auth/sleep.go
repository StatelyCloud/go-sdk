package auth

import (
	"context"
	"time"
)

// Copied out of gocommon/ctxutils/sleep.go because sdk/go shouldn't depend on gocommon

// SleepWithContext does what it says - it sleeps, but if the context is
// canceled it stops early. It returns an error in the case of context
// cancellation.
func sleepWithContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
		return ctx.Err()
	case <-timer.C:
	}
	return nil
}
