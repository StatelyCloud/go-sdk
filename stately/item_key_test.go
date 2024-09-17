package stately

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToKeyID(t *testing.T) {
	type testCase struct {
		name     string
		given    string
		expected string
	}
	tests := []testCase{
		{
			name:     "UUID",
			given:    ToKeyID(uuid.MustParse("00000000-0000-0000-0000-000000000005")),
			expected: "AAAAAAAAAAAAAAAAAAAABQ",
		},
		{
			name:     "String",
			given:    ToKeyID("string"),
			expected: "string",
		},
		{
			name: "ByteSlice16",
			given: ToKeyID([16]byte{
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x50,
			}),
			expected: "AAAAAAAAAAAAAAAAAAAAUA",
		},
		{
			name:     "ByteSlice",
			given:    ToKeyID([]byte{0x00, 0x01, 0x02, 0x03}),
			expected: "AAECAw",
		},
		{
			name:     "Uint64",
			given:    ToKeyID(uint64(1234)),
			expected: "1234",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, tt.given, "ToKeyID(%v)", tt.given)
		})
	}
}
