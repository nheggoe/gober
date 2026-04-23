package util

import (
	"errors"
	"testing"

	"github.com/shoenig/test/must"
)

func TestWrapError_WrapsAndPreservesCause(t *testing.T) {
	{
		original := errors.New("operation failed")
		run := func() (err error) {
			defer WrapError(&err, "test")
			return original
		}
		err := run()
		must.Error(t, err)
		must.Eq(t, "test: operation failed", err.Error())
		must.ErrorIs(t, err, original)
	}
	{
		original1, original2 := errors.New("operation1 failed"), errors.New("operation2 failed")
		run := func() (err error) {
			defer WrapError(&err, "test")
			return &Errors{original1, original2}
		}
		err := run()
		must.Error(t, err)
		must.Eq(t, "test: operation1 failed, operation2 failed", err.Error())
		must.ErrorIs(t, err, original1)
		must.ErrorIs(t, err, original2)
	}
}

func TestWrapError_DoesNothingWhenNil(t *testing.T) {
	run := func() (err error) {
		defer WrapError(&err, "test")
		return nil
	}
	err := run()
	must.NoError(t, err)
}

func TestErrors_ToString(t *testing.T) {
	errs := Errors{errors.New("err1"), errors.New("err2")}
	must.Eq(t, "err1, err2", errs.Error())
}
