package safe_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cloudlibraries/safe"
	qt "github.com/frankban/quicktest"
)

func TestDo(t *testing.T) {
	c := qt.New(t)

	c.Run("Do", func(c *qt.C) {
		c.Run("nil", func(c *qt.C) {
			err := safe.Do(nil)
			c.Assert(err, qt.ErrorMatches, "panic: invalid function type: <nil>")
		})

		c.Run("func()", func(c *qt.C) {
			err := safe.Do(func() {})
			c.Assert(err, qt.IsNil)
		})

		c.Run("func() error", func(c *qt.C) {
			err := safe.Do(func() error { return nil })
			c.Assert(err, qt.IsNil)
		})
	})
}

func TestDoWithContext(t *testing.T) {
	c := qt.New(t)

	c.Run("DoWithContext", func(c *qt.C) {
		c.Run("nil", func(c *qt.C) {
			err := safe.DoWithContext(context.Background(), nil)
			c.Assert(err, qt.ErrorMatches, "panic: invalid function type: <nil>")
		})

		c.Run("func()", func(c *qt.C) {
			err := safe.DoWithContext(context.Background(), func() {})
			c.Assert(err, qt.IsNil)
		})

		c.Run("func() error", func(c *qt.C) {
			err := safe.DoWithContext(context.Background(), func() error { return nil })
			c.Assert(err, qt.IsNil)
		})
	})
}

func TestDoWithTimeout(t *testing.T) {
	c := qt.New(t)

	c.Run("DoWithTimeout", func(c *qt.C) {
		c.Run("nil", func(c *qt.C) {
			err := safe.DoWithTimeout(time.Second, nil)
			c.Assert(err, qt.ErrorMatches, "panic: invalid function type: <nil>")
		})

		c.Run("func()", func(c *qt.C) {
			err := safe.DoWithTimeout(time.Second, func() {})
			c.Assert(err, qt.IsNil)
		})

		c.Run("func() error", func(c *qt.C) {
			err := safe.DoWithTimeout(time.Second, func() error { return nil })
			c.Assert(err, qt.IsNil)
		})

		c.Run("timeout", func(c *qt.C) {
			err := safe.DoWithTimeout(time.Nanosecond, func() {
				time.Sleep(time.Second)
			})
			c.Assert(err, qt.ErrorMatches, "context deadline exceeded")
		})

		c.Run("panic", func(c *qt.C) {
			err := safe.DoWithTimeout(time.Second, func() {
				panic("foo")
			})
			c.Assert(err, qt.ErrorMatches, "panic: foo")
		})

		c.Run("error", func(c *qt.C) {
			err := safe.DoWithTimeout(time.Second, func() error {
				return errors.New("foo")
			})
			c.Assert(err, qt.ErrorMatches, "foo")
		})
	})
}
