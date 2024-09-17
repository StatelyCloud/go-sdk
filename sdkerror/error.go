package sdkerror

import (
	"errors"
	"slices"

	"connectrpc.com/connect"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"

	pberrors "github.com/StatelyCloud/go-sdk/pb/errors"
)

// Error is the error struct that all SDK code vends.
type Error struct {
	// ConnectCode is the top-level RPC code for this error.
	ConnectCode connect.Code

	// The StatelyCode property is the specific error code
	// For more information about a specific code read:
	// https://docs.stately.cloud/reference/error_codes/#{STATELY_CODE}
	StatelyCode StatelyErrorCode

	// Message is a human-readable description of the error itself.
	Message string

	// If present this is the underlying error that caused this error.
	CauseErr error

	attrs map[string]string
}

// AddAttr adds a structured attribute to the error which will be included
// under the primary message.
func (e *Error) AddAttr(key, value string) {
	if e.attrs == nil {
		e.attrs = make(map[string]string)
	}
	e.attrs[key] = value
}

// ConnectError converts this to a Connect Error.
func (e *Error) ConnectError() *connect.Error {
	ce := connect.NewError(e.ConnectCode, errors.New(e.baseMessage()))
	wireDetail := &pberrors.StatelyErrorDetails{
		StatelyCode:   string(e.StatelyCode),
		Message:       e.Message,
		UpstreamCause: strerr(e.CauseErr),
	}

	detail, err := connect.NewErrorDetail(wireDetail)
	if err != nil {
		// Not much we can do, we'll just send the whole thing over:
		return connect.NewError(e.ConnectCode, e)
	}

	ce.AddDetail(detail)
	return ce
}

// Error implements the "error" interface and returns a message of the form:
// (RPCCode/StatelyCode) Message
// \t{ AttrK1: AttrV1 }
// Caused by: <message>.
func (e *Error) Error() string {
	result := e.baseMessage()
	if e.Message != "" {
		result += " " + e.Message
	}
	const separator = ", "
	hasDetails := false

	keys := maps.Keys(e.attrs)
	slices.Sort(keys)
	for _, k := range keys {
		if hasDetails {
			result += separator
		} else {
			result += "\n\t{ "
			hasDetails = true
		}
		result += k + ": " + e.attrs[k]
	}
	if hasDetails {
		result += " }"
	}

	if e.CauseErr != nil {
		result += "\nCaused by: " + e.CauseErr.Error()
	}
	return result
}

// Unwrap is for unwrapping errors to get to the cause error.
func (e *Error) Unwrap() error {
	if e != nil && e.CauseErr != nil {
		return e.CauseErr
	}
	return nil
}

// Is implements the required errors.Is interface.
func (e *Error) Is(err error) bool {
	if err == nil || e == nil {
		return false
	}
	if unwrapped := errors.Unwrap(err); unwrapped != nil {
		return e.Is(unwrapped)
	}

	return false
}

// As allows `errors.As` to convert this to connect.Error.
func (e *Error) As(target any) bool {
	if target == nil {
		return false
	}
	if ce, ok := target.(**connect.Error); ok {
		*ce = e.ConnectError()
		return true
	}
	if e.CauseErr != nil {
		return errors.As(e.CauseErr, target)
	}
	return false
}

// builds a formatted message of the form:
// (Code/StatelyCode).
func (e *Error) baseMessage() string {
	return "(" + strcase.ToCamel(e.ConnectCode.String()) + "/" + string(e.StatelyCode) + ")"
}

func strerr(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
