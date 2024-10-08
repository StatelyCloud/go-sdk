package stately

import (
	"strconv"
	"strings"
)

// AllowStale - If true, you're okay with getting a slightly stale item - that
// is, if you had just changed an item and then call get or list on it, you
// might get the old version of the item. This can result in improved
// performance, availability, and cost.
type AllowStale bool

// AppendIDAssignment specifies the ID assignment strategy when Appending an Item.
type AppendIDAssignment int

func (a AppendIDAssignment) String() string {
	switch a {
	case AppendIDAssignmentSequence:
		return "sequence"
	case AppendIDAssignmentUUID:
		return "uuid"
	case AppendIDAssignmentRand53:
		return "rand53"
	default:
		return "invalid (" + strconv.FormatInt(int64(a), 10) + ")"
	}
}

const (
	// AppendIDAssignmentSequence will assign the item a monotonically
	// increasing, contiguous ID that is unique *within the parent path and
	// item type*. This is only valid for non-root items.
	AppendIDAssignmentSequence AppendIDAssignment = iota + 1

	// AppendIDAssignmentUUID will assign the item a globally unique 128-bit
	// UUID. This will be encoded in the item key path as a binary ID. This
	// is usable anywhere, in any store config.
	AppendIDAssignmentUUID

	// AppendIDAssignmentRand53 will assign the item a random 53-bit numeric ID that
	// is unique *within the parent path and item type*, but is not globally
	// unique. This is usable anywhere, in any store config. We use 53 bits
	// instead of 64 because 53 bits is still a lot of bits, and it's the largest
	// integer that can be represented exactly in JavaScript.
	AppendIDAssignmentRand53
)

// NewAppendIDAssignment converts from a name to a AppendIDAssignment type.
// An invalid value will return 0.
func NewAppendIDAssignment(s string) AppendIDAssignment {
	switch strings.ToLower(s) {
	case "sequence":
		return AppendIDAssignmentSequence
	case "uuid":
		return AppendIDAssignmentUUID
	case "rand53":
		return AppendIDAssignmentRand53
	default:
		return 0
	}
}
