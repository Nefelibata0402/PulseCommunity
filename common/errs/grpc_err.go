package errs

import (
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcError(err *ErrCore) error {
	return status.Error(codes.Code(err.ErrCode), err.ErrMsg)
}

func ParseGrpcError(err error) (int, string) {
	fromError, _ := status.FromError(err)
	return int(fromError.Code()), fromError.Message()
}

func ToBError(err error) *ErrCore {
	fromError, _ := status.FromError(err)
	return NewErrCore(int(fromError.Code()), fromError.Message())
}
