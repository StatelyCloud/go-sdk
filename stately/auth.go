package stately

import (
	"context"
	"time"

	"github.com/StatelyCloud/go-sdk/internal/auth"
)

// AccessKeyAuth creates an AuthTokenProvider that uses an access key to authenticate.
func AccessKeyAuth(
	ctx context.Context,
	accessKey string,
	endpoint string,
	refreshInterval time.Duration,
) AuthTokenProvider {
	return auth.AccessKeyAuth(ctx, accessKey, endpoint, createTransport(endpoint), refreshInterval)
}
