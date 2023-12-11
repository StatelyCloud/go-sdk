package data

import (
	"fmt"
	"strings"
)

// Item response.
type Item struct {
	KeyPath
	JSON  string
	Proto []byte
}

// KeyPath is the string representation of the path of an item in the database.
// See https://docs.google.com/document/d/1TV-b-49elWvQdc2zDrj_2Bi7qDfnL3_tv_zscypW_JM/edit#heading=h.ji4wddkhv1a1
type KeyPath string

// NewKeyPath formulates a valid stately key path from the given segments.
// Input is a list of strings which should alternate between types and IDs.
// which will be built into the complete key path.
// Output will be of the form /<type1>-<id1>/<type2>-<id2>.
// The terminating element can be either a type or ID.
func NewKeyPath(segments ...string) KeyPath {
	keyPath := ""
	for i, segment := range segments {
		if i%2 == 0 {
			keyPath = keyPath + "/" + segment
		} else {
			keyPath = keyPath + "-" + segment
		}
	}
	return KeyPath(keyPath)
}

// String returns the KeyPath as a string.
func (k KeyPath) String() string {
	return string(k)
}

// ExtractID extracts the item ID from the given item.
// This assumes that the supplied item has a key path of the form /<type>-<id>.
// If the path does not contain an ID or the ID is an empty string
// then this API will return undefined.
func (k KeyPath) ExtractID() (string, error) {
	pairs := strings.Split(k.String(), "/")
	if len(pairs) < 1 {
		return "", fmt.Errorf("KeyPath '%s' is empty", k.String())
	}

	lastPair := pairs[len(pairs)-1]
	splitPair := strings.Split(lastPair, "-")
	if len(splitPair) < 2 {
		return "", fmt.Errorf("KeyPath '%s' does not contain an ID segment", k.String())
	}

	return splitPair[1], nil
}
