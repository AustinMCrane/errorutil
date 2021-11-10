package errorutil

import (
	"errors"
	"fmt"
	"runtime"
)

// WrappedError A Wrapped Error
type WrappedError struct {
	frame runtime.Frame
	Err   error
}

// Wrap Wrap the error with a stack frame
func wrap(skip int, err error) error {
	var f [3]uintptr
	runtime.Callers(skip, f[:])
	frames := runtime.CallersFrames(f[:])
	fr, _ := frames.Next()

	return WrappedError{frame: fr, Err: err}
}

func (w WrappedError) Error() string {
	if w.Err != nil {
		return w.Err.Error() + fmt.Sprintf("\n    %s\n\t%s:%d", w.frame.Function, w.frame.File,
			w.frame.Line)
	}

	return fmt.Sprintf("\n    %s\n\t%s:%d", w.frame.Function, w.frame.File,
		w.frame.Line)
}

func (w WrappedError) Unwrap() error {
	return w.Err
}

func New(msg string) error {
	err := errors.New(msg)
	return wrap(3, err)
}

func Wrap(err error) error {
	return wrap(3, err)

}
