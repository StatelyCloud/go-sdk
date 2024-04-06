package data

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"connectrpc.com/connect"
)

const (
	// PathDelim separates key components.
	PathDelim = "/"
	// TypeIDDelim separates the item type from the item ID.
	TypeIDDelim = "-"
	// byteSigil is the prefix for base64-encoded item IDs in paths. We'll want to
	// prevent users from using this in their own IDs.
	byteSigil = "~"

	maxItemTypeLength = 32
	maxStringLength   = 64
	maxByteLength     = 64
)

var (
	encoding            = base64.RawURLEncoding.Strict()
	maxByteLengthBase64 = base64.RawURLEncoding.EncodedLen(maxByteLength)

	// Ban control characters and other non-printables. This should also include
	// valid characters for numbers and base64.
	stringValueRegex = regexp.MustCompile(`^[^\p{Cc}\p{Co}]+$`)

	// numbers can't start with 0, no negatives.
	numberValueRegex = regexp.MustCompile(
		`^[1-9]\d*$`,
	)
)

// toItemKey converts a customer-facing string path to an ItemKey structure. In
// the absence of a defined schema, all key values are interpreted as strings.
// This works for partial paths, which means that it may return an empty list -
// validate its length if you need to.
func toItemKey(path string) (ItemKey, error) {
	if !strings.HasPrefix(path, PathDelim) {
		return ItemKey{}, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("path must start with %s", PathDelim),
		)
	}
	path = path[1:] // remove leading slash

	// TODO: would be great to do this with a linear scan rather than multiple splits
	parts := strings.Split(path, PathDelim)

	keyPath := make([]*ItemKeyComponent, len(parts))
	for i, part := range parts {
		itemType, itemIDStr, hasID := strings.Cut(part, TypeIDDelim)

		if !hasID {
			if i == len(parts)-1 {
				// The last element is allowed to have no ID.
				keyPath[i] = &ItemKeyComponent{
					ItemType: part,
				}
				continue
			}
			// TODO: Perhaps relax this to allow for hierarchies of singletons, e.g.
			// /settings/chat, /settings/notifications, etc. That makes some item
			// types more like a namespace.
			return ItemKey{}, connect.NewError(
				connect.CodeInvalidArgument,
				fmt.Errorf(
					"all but the last element must contain both an item type and an ID - %q is only an item type (in %q)",
					part,
					path,
				),
			)
		}

		itemID, err := stringToKeyValue(itemIDStr)
		if err != nil {
			return ItemKey{}, err
		}
		keyComponent := ItemKeyComponent{
			ItemType: itemType,
			ItemID:   itemID,
		}
		keyPath[i] = &keyComponent
	}
	return keyPath, nil
}

func stringToKeyValue(str string) (KeyValue, error) {
	if str == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ID path segment cannot be empty"))
	}

	if !stringValueRegex.MatchString(str) {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("ID path segment contains illegal characters"),
		)
	}

	// Binary
	if strings.HasPrefix(str, byteSigil) {
		return parseBinaryIDSegment(str)
	}

	// Positive integer
	if numberValueRegex.MatchString(str) {
		if v, err := strconv.ParseUint(str, 10, 64); err == nil {
			// Numbers are unprefixed, meaning we can't tell the difference between a
			// string that parses as a uint64 and a number. We also don't support signed
			// numbers.
			return &KeyValueUint{Value: v}, nil
		}
	}

	// String
	if len(str) > maxStringLength {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("string ID path segment cannot be more than %b bytes long", maxStringLength),
		)
	}
	return &KeyValueString{Value: str}, nil
}

func parseBinaryIDSegment(str string) (KeyValue, error) {
	str = str[len(byteSigil):] // strip off the sigil
	if len(str) > maxByteLengthBase64 {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("byte ID path segment cannot be more than %d bytes long", maxByteLength),
		)
	}
	rem := len(str) % 4
	if rem == 1 ||
		// Check to see if the remainder string round-trips. If it doesn't, then we reject it as ambiguous.
		rem != 0 && !roundtripBase64(str[len(str)-rem:]) {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("byte ID path segment is not valid base64"),
		)
	}
	b, err := encoding.DecodeString(str)
	if err != nil {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("byte ID path segment is not valid base64"),
		)
	}
	if len(b) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("byte ID path segment cannot empty"))
	}
	return &KeyValueBinary{Value: b}, nil
}

func roundtripBase64(str string) bool {
	b, err := encoding.DecodeString(str)
	if err != nil {
		return false
	}
	r := encoding.EncodeToString(b)
	return str == r
}
