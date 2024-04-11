package data

import (
	"strings"
)

// ItemKey is a complete item key consisting of one (the root) and any more key components. The root component is special
// as you MUST always specify the `ID` associated with `type` of on the root, while other components only require type.
// Example: /conversation-foo/message and /conversation-foo/message-1 are valid ItemKey(s), while /conversation is not
// because it lacks the ID `-foo` of the segment. The key path must be at least 1 element long.
type ItemKey []*ItemKeyComponent

// KeyPath converts an ItemKey structure to a customer-facing string path,
// e.g. "/users-123/posts-a-fun-recipe". It is assumed that the ItemKey was
// validated when it was first converted from a string, so no validation is done
// in this direction.
func (k ItemKey) KeyPath() string {
	path := strings.Builder{}

	for _, keyComponent := range k {
		path.WriteString(PathDelim)
		path.WriteString(keyComponent.String())
	}

	return path.String()
}
