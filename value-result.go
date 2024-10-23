package result

type (
	Value[T any] interface {
		Error
		// IsOk returns true if there is no error in result
		IsOk() bool
		// Value return the value of result
		Value() T
		// Unwrap returns the value and error of result
		Unwrap() (T, error)
		// Must panics if result with error, otherwise returns the value of result
		Must() T
	}

	value[T any] struct {
		Error
		value T
	}
)

// NewValue creates result with value
func NewValue[T any](v T, err any) Value[T] {
	return &value[T]{
		value: v,
		Error: NewError(err),
	}
}

func (r *value[T]) IsOk() bool {
	return !r.IsError()
}

func (r *value[T]) Value() T {
	return r.value
}

func (r *value[T]) Unwrap() (T, error) {
	return r.Value(), r.Err()
}
func (r *value[T]) Must() T {
	r.PanicIfError()
	return r.Value()
}

// Wrap wraps input to result with value
func Wrap[T any](v T, e error) Value[T] {
	return NewValue(v, e)
}
