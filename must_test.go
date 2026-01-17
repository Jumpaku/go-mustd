package mustd_test

import (
	"errors"
	"io"
	"testing"

	"github.com/Jumpaku/go-mustd"
)

func TestMust0(t *testing.T) {
	t.Run("nil error does not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Must0 panicked with nil error: %v", r)
			}
		}()
		mustd.Must0(nil)
	})

	t.Run("non-nil error panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Must0 did not panic with non-nil error")
			}
		}()
		mustd.Must0(errors.New("test error"))
	})
}

func TestMust1(t *testing.T) {
	t.Run("nil error returns value", func(t *testing.T) {
		result := mustd.Must1(42, nil)
		if result != 42 {
			t.Errorf("expected 42, got %d", result)
		}
	})

	t.Run("non-nil error panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Must1 did not panic with non-nil error")
			}
		}()
		mustd.Must1(42, errors.New("test error"))
	})
}

func TestMust2(t *testing.T) {
	t.Run("nil error returns values", func(t *testing.T) {
		v1, v2 := mustd.Must2("hello", 42, nil)
		if v1 != "hello" || v2 != 42 {
			t.Errorf("expected (hello, 42), got (%s, %d)", v1, v2)
		}
	})

	t.Run("non-nil error panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Must2 did not panic with non-nil error")
			}
		}()
		mustd.Must2("hello", 42, errors.New("test error"))
	})
}

func TestMust3(t *testing.T) {
	t.Run("nil error returns values", func(t *testing.T) {
		v1, v2, v3 := mustd.Must3("hello", 42, true, nil)
		if v1 != "hello" || v2 != 42 || v3 != true {
			t.Errorf("expected (hello, 42, true), got (%s, %d, %v)", v1, v2, v3)
		}
	})

	t.Run("non-nil error panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Must3 did not panic with non-nil error")
			}
		}()
		mustd.Must3("hello", 42, true, errors.New("test error"))
	})
}

func TestMustImplement(t *testing.T) {
	t.Run("valid type assertion returns value", func(t *testing.T) {
		var r io.Reader = &testReader{}
		result := mustd.MustImplement[io.Reader](r)
		if result == nil {
			t.Error("MustImplement returned nil")
		}
	})

	t.Run("invalid type assertion panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustImplement did not panic with invalid type")
			}
		}()
		mustd.MustImplement[io.Reader]("not a reader")
	})
}

type testReader struct{}

func (tr *testReader) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}
