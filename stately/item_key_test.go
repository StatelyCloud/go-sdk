package stately_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/StatelyCloud/go-sdk/stately"
)

func TestToKeyID(t *testing.T) {
	type UUID [16]byte
	id := UUID{5, 4, 2, 1, 3, 4, 7, 3, 1, 2, 3, 4, 6, 1, 3, 5}

	type testCase struct {
		name     string
		given    string
		expected string
	}
	tests := []testCase{
		{
			name: "ByteSlice16",
			given: stately.ToKeyID([16]byte{
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x50,
			}),
			expected: "AAAAAAAAAAAAAAAAAAAAUA",
		},
		{
			name:     "typed uuid",
			given:    stately.ToKeyID([16]byte(id)),
			expected: "BQQCAQMEBwMBAgMEBgEDBQ",
		},
		{
			name:     "String",
			given:    stately.ToKeyID("string"),
			expected: "string",
		},
		{
			name:     "ByteSlice",
			given:    stately.ToKeyID([]byte{0x00, 0x01, 0x02, 0x03}),
			expected: "AAECAw",
		},
		{
			name:     "Uint64",
			given:    stately.ToKeyID(uint64(1234)),
			expected: "1234",
		},
		{
			name:     "Uint32",
			given:    stately.ToKeyID(uint32(1234)),
			expected: "1234",
		},
		{
			name:     "int64",
			given:    stately.ToKeyID(int64(-1234)),
			expected: "-1234",
		},
		{
			name:     "int32",
			given:    stately.ToKeyID(int32(-1234)),
			expected: "-1234",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, tt.given, "ToKeyID(%v)", tt.given)
		})
	}
}
