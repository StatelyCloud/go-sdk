package internal

import (
	"errors"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"

	"github.com/StatelyCloud/go-sdk/pb/common"
)

// MapProtoError maps any common error modeled by Stately to a Connect error but passes it back as
// a regular golang error so we don't lose it's nil-ness in the interface conversion.
func MapProtoError(err *common.OperationError) error {
	if err == nil || err.GrpcCode == uint32(codes.OK) {
		return nil
	}
	return connect.NewError(connect.Code(err.GrpcCode), errors.New(err.GetDescription()))
}
