package errorutil

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// WrappedError A Wrapped Error
type WrappedError struct {
	frame runtime.Frame
	Err   error
	Msg   string
}

// Wrap Wrap the error with a stack frame
func wrap(skip int, err error, msg ...string) error {
	var f [3]uintptr
	runtime.Callers(skip, f[:])
	frames := runtime.CallersFrames(f[:])
	fr, _ := frames.Next()

	return WrappedError{frame: fr, Err: err, Msg: strings.Join(msg, " ")}
}

func (w WrappedError) Error() string {
	if w.Err != nil {
		prefix := ""
		if w.Msg != "" {
			prefix = w.Msg + ": "
		}
		return prefix + w.Err.Error() +
			fmt.Sprintf("\n    %s\n\t%s:%d", w.frame.Function, w.frame.File,
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

func Wrap(err error, msg ...string) error {
	return wrap(3, err, msg...)

}
