package stately

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

var encoding = base64.RawURLEncoding.Strict()

// ToKeyID converts an ID to a key ID. If the ID is a byte slice, it will be
// base64-encoded. If the ID is a string, it will be used as-is.
func ToKeyID[T [16]byte | string | uint64 | uint32 | int64 | int32 | []byte](id T) string {
	var keyID string

	switch v := any(id).(type) {
	case [16]byte:
		keyID = encoding.EncodeToString(v[:])
	case []byte:
		keyID = encoding.EncodeToString(v)
	case string:
		keyID = v
	case uint64:
		keyID = strconv.FormatUint(v, 10)
	case uint32:
		keyID = strconv.FormatUint(uint64(v), 10)
	case int32:
		keyID = strconv.FormatInt(int64(v), 10)
	case int64:
		keyID = strconv.FormatInt(v, 10)
	default:
		panic(fmt.Sprintf("unmapped type: %T", v))
	}

	return keyID
}
