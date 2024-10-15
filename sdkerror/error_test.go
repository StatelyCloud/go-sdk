package sdkerror_test

import (
	"errors"
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"

	"github.com/StatelyCloud/go-sdk/sdkerror"
)

func TestError_Error(t *testing.T) {
	t.Run("error with message", func(t *testing.T) {
		err := &sdkerror.Error{
			Code:        connect.CodeInternal,
			StatelyCode: "EnthalpyLimitExceeded",
			Message:     "just a message",
		}
		assert.Equal(t, "(Internal/EnthalpyLimitExceeded) just a message", err.Error())
	})

	t.Run("error with attribute and message", func(t *testing.T) {
		err := &sdkerror.Error{
			Code:        connect.CodeInternal,
			StatelyCode: "EnthalpyLimitExceeded",
			Message:     "just a message",
		}
		err.AddAttr("key", "value")
		assert.Equal(t, "(Internal/EnthalpyLimitExceeded) just a message\n\t{ key: value }", err.Error())
	})

	t.Run("error with attribute and message and cause", func(t *testing.T) {
		err := &sdkerror.Error{
			Code:        connect.CodeInternal,
			StatelyCode: "EnthalpyLimitExceeded",
			Message:     "just a message",
			CauseErr:    errors.New("some cause"),
		}
		err.AddAttr("key", "value")
		assert.Equal(t,
			"(Internal/EnthalpyLimitExceeded) just a message\n\t{ key: value }\nCaused by: some cause",
			err.Error())
	})
}
