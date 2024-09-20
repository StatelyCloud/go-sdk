package stately

import (
	"strings"
)

// RegionToEndpoint converts a region to an endpoint.
func RegionToEndpoint(s string) string {
	if s == "" {
		return "https://api.stately.cloud"
	}
	// remove aws- prefix
	s, _ = strings.CutPrefix(s, "aws-")
	return "https://" + s + ".aws.api.stately.cloud"
}
