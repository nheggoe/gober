package util

import (
	"fmt"
)

// WrapError is a helper function to be used with named return with defers.
//
// Example:
//
//	func test() (err error) {
//		defer WrapError(&err, "test")
//		return errors.New("operation failed")
//	}
//
// This will result in: `Error{"test: operation failed"}`
func WrapError(err *error, src string) {
	if err != nil && *err != nil {
		*err = fmt.Errorf("%s: %w", src, *err)
	}
}

type Errors []error

func (es *Errors) Append(err error) {
	*es = append(*es, err)
}

func (es *Errors) Error() string {
	if len(*es) == 0 {
		return ""
	}
	err := (*es)[0]
	for _, e := range (*es)[1:] {
		err = fmt.Errorf("%w, %w", err, e)
	}
	return err.Error()
}

func (es *Errors) Unwrap() []error {
	return *es
}

func (es *Errors) IsEmpty() bool {
	return len(*es) == 0
}
