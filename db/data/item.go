package data

import (
	"encoding/json"
	"time"
)

// RawItem represents an Item in StatelyDB.
type RawItem struct {
	Key           ItemKey
	ItemType      string
	ParentKeyPath ItemKey
	JSONData      json.RawMessage
	ID            KeyValue
	Metadata      ItemMetadata
}

// Item represents a parsed RawItem.
type Item[T any] struct {
	*RawItem
	Data T
}

// ItemMetadata describes autopopulated fields about every item in your Stately store. This removes the need to
// add very common attributes like item creation timestamp or last modified timestamp.
type ItemMetadata struct {
	CreatedAt             time.Time
	LastModifiedAt        time.Time
	CreatedAtVersion      uint64
	LastModifiedAtVersion uint64
}

// KeyPath is a convenience function.
func (i *RawItem) KeyPath() string {
	return i.Key.KeyPath()
}
