package safe

import (
	"fmt"

	"github.com/cloudlibraries/cast"
)

func Do(a any) (err error) {
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

func DoWith(a any, f func(err error)) {
	if err := Do(a); err != nil {
		f(err)
	}
}

func Go(a any) (err error) {
	errCh := make(chan error, 1)

	go func() {
		errCh <- Do(a)
	}()

	return <-errCh
}

func GoWith(a any, f func(err error)) {
	go func() {
		f(Go(a))
	}()
}
