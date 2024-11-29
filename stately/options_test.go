package stately_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/StatelyCloud/go-sdk/stately"
)

// Test that the stately package is correctly documented.

func TestOptions_Endpoint(t *testing.T) {
	tests := []struct {
		name     string
		Region   string
		Endpoint string
		want     string
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:     "Supply nothing",
			Region:   "",
			Endpoint: "",
			want:     "https://api.stately.cloud",
			wantErr:  assert.NoError,
		},
		{
			name:     "Supply endpoint, no environment",
			Region:   "",
			Endpoint: "https://example.com",
			want:     "https://example.com",
			wantErr:  assert.NoError,
		},
		{
			name:     "Supply environment, no endpoint",
			Region:   "us-east-1",
			Endpoint: "",
			want:     "https://us-east-1.aws.api.stately.cloud",
			wantErr:  assert.NoError,
		},
		{
			name:     "Supply both, but different, expect error",
			Region:   "us-west-2",
			Endpoint: "https://example.com",
			want:     "",
			wantErr:  assert.Error,
		},
		{
			name:     "Supply both, but same, expect no error",
			Region:   "us-east-1",
			Endpoint: "https://us-east-1.aws.api.stately.cloud",
			want:     "https://us-east-1.aws.api.stately.cloud",
			wantErr:  assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &stately.Options{
				Region:            tt.Region,
				Endpoint:          tt.Endpoint,
				AuthTokenProvider: func(_ context.Context, _ bool) (string, error) { return t.Name(), nil },
			}
			got, err := o.ApplyDefaults(context.TODO())
			if !tt.wantErr(t, err) {
				return
			}

			if got == nil {
				assert.Equal(t, tt.want, "")
			}
		})
	}
}
