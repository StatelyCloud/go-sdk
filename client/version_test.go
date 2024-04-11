package client_test

import (
	"regexp"
	"testing"

	"github.com/StatelyCloud/go-sdk/client"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	t.Parallel()
	// regex for semver from here: https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
	assert.Regexp(
		t,
		regexp.MustCompile(
			`^v(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`,
		),
		client.Version(),
	)
}
