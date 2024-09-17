package stately

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

var encoding = base64.RawURLEncoding.Strict()

// ToKeyID converts an ID to a key ID. If the ID is a byte slice, it will be
// base64-encoded. If the ID is a string, it will be used as-is.
func ToKeyID[T ~[16]byte | string | ~uint64 | ~int64 | []byte](id T) string {
	var keyID string

	switch v := any(id).(type) {
	case [16]byte:
		keyID = encoding.EncodeToString(v[:])
	case []byte:
		keyID = encoding.EncodeToString(v)
	case uuid.UUID:
		keyID = encoding.EncodeToString(v[:])
	case string:
		keyID = v
	case uint64:
		keyID = strconv.FormatUint(v, 10)
	case int64:
		keyID = strconv.FormatInt(v, 10)
	default:
		panic(fmt.Sprintf("unmapped type: %T", v))
	}

	return keyID
}
