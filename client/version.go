package client

import _ "embed"

//go:embed version.txt
var version string

// Version returns the current version of the SDK.
func Version() string {
	return version
}
