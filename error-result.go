package result

import (
	"errors"
	"fmt"
)

type (
	Error interface {
		// Err returns the error of result
		Err() error
		// IsError returns true if result has an error
		IsError() bool
		// PanicIfError panics if result has an error
		PanicIfError()
	}
	errorResult struct {
		err error
	}
)

// NewError creates an error result
func NewError(err any) Error {
	var res error = nil
	if err != nil {
		switch v := err.(type) {
		case error:
			res = v
		case fmt.Stringer:
			res = errors.New(v.String())
		case string:
			res = errors.New(v)
		default:
			panic("unsupported error content type")
		}
	}
	return &errorResult{
		err: res,
	}
}

func (r *errorResult) Err() error {
	return r.err
}

func (r *errorResult) IsError() bool {
	return r.err != nil
}

func (r *errorResult) PanicIfError() {
	if r.IsError() {
		panic(r.Err())
	}
}
