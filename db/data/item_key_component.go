package data

import (
	"fmt"
	"strconv"
)

// ItemKeyComponent identifies a segment of an ItemKey, representing a tuple of
// an item type and an optional ID.
type ItemKeyComponent struct {
	// ItemType identifies the type of this item, which:
	//  1. Associates it with a type in schema (TBD)
	//  2. Determines which version counter / ID sequence to use for these items,
	//     for both ID assignment and change tracking.
	//
	// TODO: In the future we could represent this as an ordinal into a schema,
	// for more compact storage.
	ItemType string

	// ItemID is an optional, unique value that identifies this item among other
	// items of the same item type within its containing item type. The ID may be
	// omitted in cases where the item is a singleton. This ID may be
	// user-provided, or assigned automatically by append.
	//
	// Types that are assignable to ItemId:
	//
	//	*KeyValueUint
	//	*KeyValueString
	//	*KeyValueBinary
	ItemID KeyValue
}

func (c ItemKeyComponent) String() string {
	if c.ItemID == nil {
		return c.ItemType
	}
	return c.ItemType + TypeIDDelim + c.ItemID.String()
}

// KeyValue represents a key in an item path. It can be a uint64, string, or bytes.
type KeyValue interface {
	fmt.Stringer
}

// KeyValueUint is a uint64 key value.
type KeyValueUint struct {
	// uint_value is the common case of an unsigned 64-bit integer as an ID.
	// This is a great default for most IDs and is what AppendItem automatically
	// uses when assigning IDs.
	Value uint64
}

func (kv *KeyValueUint) String() string {
	return strconv.FormatUint(kv.Value, 10)
}

// KeyValueInt is a int64 key value. This is currently only used for LSIs; it is not
// a part of primary keys.
type KeyValueInt struct {
	// Value is the int64 value of the key.
	Value int64
}

func (kv *KeyValueInt) String() string {
	return strconv.FormatInt(kv.Value, 10)
}

// KeyValueString is a string key value.
type KeyValueString struct {
	// string_value allows users to supply arbitrary keys as a string.
	Value string
}

func (kv *KeyValueString) String() string {
	// The string value is unprefixed
	return kv.Value
}

// KeyValueBinary is a byte slice key value.
type KeyValueBinary struct {
	// binary_value allows users to supply arbitrary binary value keys, which is
	// good for things like UUIDs.
	Value []byte
}

func (kv *KeyValueBinary) String() string {
	// URLEncoding won't use the / character, so it's safe to use in paths.
	// It is prefixed so we can tell the difference between it and a string.
	return byteSigil + encoding.EncodeToString(kv.Value)
}
