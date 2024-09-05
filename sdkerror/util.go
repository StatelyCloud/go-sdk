package sdkerror

import (
	"errors"
)

// Is checks if the error is a Stately error with the given code.
// It can be used like this:
//
//	if sdkerror.Is(err, sdkerror.StoreRequestLimitExceeded) {
//	   // handle
//	}
func Is(err error, targetCode StatelyErrorCode) bool {
	if err == nil {
		return false
	}
	for err != nil {
		if e, ok := err.(*Error); ok && e.StatelyCode == targetCode {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}
