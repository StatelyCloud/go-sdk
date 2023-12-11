package time

import (
	"crypto/rand"
	"math/big"
	"time"
)

// Jitter creates a cryptographically random duration between the given bounds.
func Jitter(min, max time.Duration) (time.Duration, error) {
	jitterNanos, err := rand.Int(rand.Reader, big.NewInt(max.Nanoseconds()-min.Nanoseconds()))
	if err != nil {
		return 0, err
	}
	return min + time.Duration(jitterNanos.Int64()), nil
}
