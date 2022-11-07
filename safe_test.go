package safe_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/frankban/quicktest"
	"github.com/golibraries/safe"
)

func TestDo(t *testing.T) {
	c := quicktest.New(t)

	c.Run("Do", func(c *quicktest.C) {
		c.Run("nil", func(c *quicktest.C) {
			err := safe.Do(nil)
			c.Assert(err, quicktest.ErrorMatches, "panic: invalid function type: <nil>")
		})

		c.Run("func()", func(c *quicktest.C) {
			err := safe.Do(func() {})
			c.Assert(err, quicktest.IsNil)
		})

		c.Run("func() error", func(c *quicktest.C) {
			err := safe.Do(func() error { return nil })
			c.Assert(err, quicktest.IsNil)
		})
	})
}

func TestDoWithContext(t *testing.T) {
	c := quicktest.New(t)

	c.Run("DoWithContext", func(c *quicktest.C) {
		c.Run("nil", func(c *quicktest.C) {
			err := safe.DoWithContext(context.Background(), nil)
			c.Assert(err, quicktest.ErrorMatches, "panic: invalid function type: <nil>")
		})

		c.Run("func()", func(c *quicktest.C) {
			err := safe.DoWithContext(context.Background(), func() {})
			c.Assert(err, quicktest.IsNil)
		})

		c.Run("func() error", func(c *quicktest.C) {
			err := safe.DoWithContext(context.Background(), func() error { return nil })
			c.Assert(err, quicktest.IsNil)
		})
	})
}

func TestDoWithTimeout(t *testing.T) {
	c := quicktest.New(t)

	c.Run("DoWithTimeout", func(c *quicktest.C) {
		c.Run("nil", func(c *quicktest.C) {
			err := safe.DoWithTimeout(time.Second, nil)
			c.Assert(err, quicktest.ErrorMatches, "panic: invalid function type: <nil>")
		})

		c.Run("func()", func(c *quicktest.C) {
			err := safe.DoWithTimeout(time.Second, func() {})
			c.Assert(err, quicktest.IsNil)
		})

		c.Run("func() error", func(c *quicktest.C) {
			err := safe.DoWithTimeout(time.Second, func() error { return nil })
			c.Assert(err, quicktest.IsNil)
		})

		c.Run("timeout", func(c *quicktest.C) {
			err := safe.DoWithTimeout(time.Nanosecond, func() {
				time.Sleep(time.Second)
			})
			c.Assert(err, quicktest.ErrorMatches, "context deadline exceeded")
		})

		c.Run("panic", func(c *quicktest.C) {
			err := safe.DoWithTimeout(time.Second, func() {
				panic("foo")
			})
			c.Assert(err, quicktest.ErrorMatches, "panic: foo")
		})

		c.Run("error", func(c *quicktest.C) {
			err := safe.DoWithTimeout(time.Second, func() error {
				return errors.New("foo")
			})
			c.Assert(err, quicktest.ErrorMatches, "foo")
		})
	})
}
