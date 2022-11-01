package safe

import (
	"fmt"

	"github.com/cloudlibraries/cast"
)

func Call(a any) (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("panic: %w", cast.ToError(v))
		}
	}()

	switch f := a.(type) {
	case func():
		f()
	case func() error:
		err = f()
	default:
		panic(fmt.Errorf("invalid function type: %T", f))
	}

	return
}

func AsyncCall(a any) (err error) {
	errCh := make(chan error, 1)

	go func() {
		errCh <- Call(a)
	}()

	return <-errCh
}
