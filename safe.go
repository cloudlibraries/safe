package safe

import (
	"context"
	"fmt"
	"time"

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

func DoWithContext(ctx context.Context, a any) (err error) {
	errCh := make(chan error, 1)

	go func() {
		errCh <- Do(a)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err = <-errCh:
		return err
	}
}

func DoWithTimeout(a any, d time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	return DoWithContext(ctx, a)
}
