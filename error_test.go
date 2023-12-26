package errorutil

import (
	"strings"
	"testing"
)

func TestErr(t *testing.T) {
	err := New("test")
	if err == nil {
		t.Errorf("New() returned nil")
	}
	if !strings.Contains(err.Error(), "errorutil.TestErr") ||
		!strings.Contains(err.Error(), "error_test.go") ||
		!strings.Contains(err.Error(), "test") {
		t.Fatalf("New() returned unexpected error: %s", err.Error())
	}

	err = Wrap(err, "test2")
	if !strings.Contains(err.Error(), "errorutil.TestErr") ||
		!strings.Contains(err.Error(), "error_test.go") ||
		!strings.Contains(err.Error(), "test") ||
		!strings.Contains(err.Error(), "test2") {
		t.Fatalf("New() returned unexpected error: %s", err.Error())
	}

	err2 := New("test3")
	err = Wrap(err2)
	if !strings.Contains(err.Error(), "errorutil.TestErr") ||
		!strings.Contains(err.Error(), "error_test.go") ||
		!strings.Contains(err.Error(), "test") ||
		strings.Contains(err.Error(), ": ") {
		t.Fatalf("New() returned unexpected error: %s", err.Error())
	}
}
