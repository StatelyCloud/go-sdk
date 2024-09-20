package stately_test

import (
	"testing"

	"github.com/StatelyCloud/go-sdk/stately"
)

func TestEnvironmentString(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "Default Region",
			env:  "",
			want: "https://api.stately.cloud",
		},
		{
			name: "AWS US East 1",
			env:  "aws-us-east-1",
			want: "https://us-east-1.aws.api.stately.cloud",
		},
		{
			name: "AWS US West 2",
			env:  "us-west-2",
			want: "https://us-west-2.aws.api.stately.cloud",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stately.RegionToEndpoint(tt.env); got != tt.want {
				t.Errorf("Region.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
