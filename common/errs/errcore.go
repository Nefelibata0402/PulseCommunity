package errs

import (
	"errors"
	"fmt"
)

type ErrCore struct {
	ErrCode int
	ErrMsg  string
}

func (e ErrCore) Error() string {
	return fmt.Sprintf("ErrCode = %d, ErrMsg = %s", e.ErrCode, e.ErrMsg)
}

func NewErrCore(code int, msg string) *ErrCore {
	return &ErrCore{code, msg}
}

func (e ErrCore) WithMessage(msg string) ErrCore {
	e.ErrMsg = msg
	return e
}

func ConvertErr(err error) *ErrCore {
	core := ErrCore{}
	if errors.As(err, &core) {
		return &core
	}

	s := InternalError
	s.ErrMsg = err.Error()
	return s
}
