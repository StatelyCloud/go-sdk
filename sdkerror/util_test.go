package sdkerror_test

import (
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"

	"github.com/StatelyCloud/go-sdk/sdkerror"
)

func TestIs(t *testing.T) {
	err := &sdkerror.Error{
		ConnectCode: connect.CodeInvalidArgument,
		StatelyCode: "IllegalFieldImpersonation",
		Message:     "don't impersonate fields",
	}
	assert.True(t, sdkerror.Is(err, "IllegalFieldImpersonation"))
	var wrappedErr error = &wrappedTestError{err: err}
	assert.True(t, sdkerror.Is(wrappedErr, "IllegalFieldImpersonation"))
}

type wrappedTestError struct {
	err error
}

func (e *wrappedTestError) Error() string {
	return "wrapped" + e.err.Error()
}

func (e *wrappedTestError) Unwrap() error {
	if e != nil {
		return e.err
	}
	return e
}
