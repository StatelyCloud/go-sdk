package sdkerror

import (
	"context"
	"errors"
	"io"

	"connectrpc.com/connect"

	pberrors "github.com/StatelyCloud/go-sdk/pb/errors"
)

// Below are common attribute keys we use.
const (
	MethodKey = "Method"
)

type clientErrorInterceptor struct{}

// ConnectErrorInterceptor creates interceptors for connect clients.
func ConnectErrorInterceptor() connect.Interceptor {
	return &clientErrorInterceptor{}
}

func (i *clientErrorInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		resp, err := next(ctx, req)
		if err != nil {
			return nil, fromRPC(err, req.Spec().Procedure)
		}
		return resp, nil
	}
}

func (i *clientErrorInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, req connect.Spec) connect.StreamingClientConn {
		resp := next(ctx, req)
		return &errorHandlingClientConn{resp}
	}
}

func (i *clientErrorInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next // noop since this is a client interceptor
}

type errorHandlingClientConn struct {
	connect.StreamingClientConn
}

func (v *errorHandlingClientConn) Receive(m any) error {
	err := v.StreamingClientConn.Receive(m)
	if errors.Is(err, io.EOF) {
		return newEOF("Attempted to receive on a closed stream.", v.StreamingClientConn.Spec().Procedure)
	} else if err != nil {
		return fromRPC(err, v.StreamingClientConn.Spec().Procedure)
	}
	return nil
}

func (v *errorHandlingClientConn) Send(m any) error {
	err := v.StreamingClientConn.Send(m)
	if errors.Is(err, io.EOF) {
		return newEOF("Attempted to send on a closed stream.", v.StreamingClientConn.Spec().Procedure)
	} else if err != nil {
		return fromRPC(err, v.StreamingClientConn.Spec().Procedure)
	}
	return nil
}

// newEOF creates an EOF SDKError.
func newEOF(msg, method string) error {
	return &Error{
		Code:        connect.CodeFailedPrecondition,
		StatelyCode: "StreamClosed",
		Message:     msg,
		attrs: map[string]string{
			MethodKey: method,
		},
		CauseErr: io.EOF,
	}
}

// fromRPC looks for a connect error and attempts to converts it into
// a SDKError with additional expended details.
func fromRPC(err error, method string) error {
	if err == nil {
		return nil
	}
	// Build the default result, we will refine it as we proceed.
	result := &Error{
		Code:        connect.CodeUnknown,
		StatelyCode: "Unknown",
		attrs: map[string]string{
			MethodKey: method,
		},
		Message: err.Error(),
	}

	var ce *connect.Error
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		result.Code = connect.CodeDeadlineExceeded
		result.CauseErr = context.DeadlineExceeded
		result.StatelyCode = "Context"
	case errors.Is(err, context.Canceled):
		result.Code = connect.CodeCanceled
		result.CauseErr = context.Canceled
		result.StatelyCode = "Context"
	case errors.As(err, &ce):
		result.Code = ce.Code()
		if detail := extractStatelyDetails(ce); detail != nil {
			// if we got stately details then this will already have the method in the
			// message so we can remove it from the attributes.
			delete(result.attrs, MethodKey)
			result.StatelyCode = StatelyErrorCode(detail.StatelyCode)
			result.Message = detail.Message
			rid := ce.Meta().Get("st-rid")
			if rid != "" {
				result.Message = result.Message + " (Request ID: " + rid + ")"
			}
			if detail.UpstreamCause != "" {
				result.CauseErr = errors.New(detail.UpstreamCause)
			}
		}
	}
	return result
}

func extractStatelyDetails(ce *connect.Error) *pberrors.StatelyErrorDetails {
	for _, detail := range ce.Details() {
		msg, err := detail.Value()
		if err != nil {
			continue
		}
		sErrDetail, ok := msg.(*pberrors.StatelyErrorDetails)
		if ok {
			return sErrDetail
		}
	}
	return nil
}
