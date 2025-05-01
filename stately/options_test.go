package stately_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/StatelyCloud/go-sdk/stately"
)

func TestOptions_Endpoint(t *testing.T) {
	tests := []struct {
		name        string
		Region      string
		Endpoint    string
		want        string
		shouldPanic bool
	}{
		{
			name:     "Supply nothing",
			Region:   "",
			Endpoint: "",
			want:     "https://api.stately.cloud",
		},
		{
			name:     "Supply endpoint, no environment",
			Region:   "",
			Endpoint: "https://example.com",
			want:     "https://example.com",
		},
		{
			name:     "Supply environment, no endpoint",
			Region:   "us-east-1",
			Endpoint: "",
			want:     "https://us-east-1.aws.api.stately.cloud",
		},
		{
			name:        "Supply both, but different, expect error",
			Region:      "us-west-2",
			Endpoint:    "https://example.com",
			want:        "",
			shouldPanic: true,
		},
		{
			name:     "Supply both, but same, expect no error",
			Region:   "us-east-1",
			Endpoint: "https://us-east-1.aws.api.stately.cloud",
			want:     "https://us-east-1.aws.api.stately.cloud",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &stately.Options{
				Region:            tt.Region,
				Endpoint:          tt.Endpoint,
				AuthTokenProvider: func(_ context.Context, _ bool) (string, error) { return t.Name(), nil },
			}
			if tt.shouldPanic {
				assert.Panics(t, func() {
					o.ApplyDefaults(context.TODO())
				})
			} else {
				got := o.ApplyDefaults(context.TODO())
				assert.Equal(t, tt.want, got.Endpoint)
			}
		})
	}
}

func TestOptions_Merge(t *testing.T) {
	tests := []struct {
		name     string
		initial  stately.Options
		merge    stately.Options
		expected stately.Options
	}{
		{
			name:    "Empty initial",
			initial: stately.Options{},
			merge: stately.Options{
				AccessKey:          "access-key",
				Region:             "us-east-1",
				Endpoint:           "https://us-east-1.aws.api.stately.cloud",
				JSONResponseFormat: true,
				// AuthTokenProvider isn't tested because EqualExportedValues doesn't work with interfaces.
			},
			expected: stately.Options{
				AccessKey:          "access-key",
				Region:             "us-east-1",
				Endpoint:           "https://us-east-1.aws.api.stately.cloud",
				JSONResponseFormat: true,
			},
		},
		{
			name: "Overwrite initial",
			initial: stately.Options{
				AccessKey:          "access-key-old",
				Region:             "us-east-1-old",
				Endpoint:           "https://us-east-1.aws.api.stately.cloud-old",
				JSONResponseFormat: false,
				// AuthTokenProvider isn't tested because EqualExportedValues doesn't work with interfaces.
			},
			merge: stately.Options{
				AccessKey:          "access-key",
				Region:             "us-east-1",
				Endpoint:           "https://us-east-1.aws.api.stately.cloud",
				JSONResponseFormat: true,
				// AuthTokenProvider isn't tested because EqualExportedValues doesn't work with interfaces.
			},
			expected: stately.Options{
				AccessKey:          "access-key",
				Region:             "us-east-1",
				Endpoint:           "https://us-east-1.aws.api.stately.cloud",
				JSONResponseFormat: true,
			},
		},
		{
			name: "Don't overwrite if not set",
			initial: stately.Options{
				AccessKey:          "access-key-old",
				Region:             "us-east-1-old",
				Endpoint:           "https://us-east-1.aws.api.stately.cloud-old",
				JSONResponseFormat: false,
				// AuthTokenProvider isn't tested because EqualExportedValues doesn't work with interfaces.
			},
			merge: stately.Options{},
			expected: stately.Options{
				AccessKey:          "access-key-old",
				Region:             "us-east-1-old",
				Endpoint:           "https://us-east-1.aws.api.stately.cloud-old",
				JSONResponseFormat: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualExportedValues(t, tt.expected, *tt.initial.Merge(&tt.merge))
		})
	}
}
