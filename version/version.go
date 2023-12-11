package version

import _ "embed"

//go:embed version.txt
var version string

// GetVersion returns the current version of the SDK.
func GetVersion() string {
	return version
}
