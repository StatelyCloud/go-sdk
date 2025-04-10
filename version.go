package sdk

import (
	_ "embed"
	"runtime"
	"strings"
)

var (
	//go:embed version.txt
	version string
	// UserAgentString e.g. v33.0.0/go+darwin/amd64.
	UserAgentString = "StatelyDB SDK/" + strings.ReplaceAll(version, "\n", "") + "/go+" + runtime.GOOS + "/" + runtime.
			GOARCH
)
