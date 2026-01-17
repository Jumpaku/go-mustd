package bytesmust_test

import (
	"bytes"
	"testing"

	"github.com/Jumpaku/go-mustd/bytesmust"
)

func TestBuffer(t *testing.T) {
	t.Run("NewBuffer", func(t *testing.T) {
		buf := bytesmust.NewBuffer([]byte("hello"))
		if buf == nil {
			t.Error("NewBuffer returned nil")
		}
	})

	t.Run("NewBufferString", func(t *testing.T) {
		buf := bytesmust.NewBufferString("hello")
		if buf == nil {
			t.Error("NewBufferString returned nil")
		}
	})

	t.Run("ReadBytes", func(t *testing.T) {
		buf := bytesmust.NewBufferString("hello\nworld")
		line := buf.ReadBytes('\n')
		if !bytes.Equal(line, []byte("hello\n")) {
			t.Errorf("expected 'hello\\n', got %s", line)
		}
	})

	t.Run("WriteString", func(t *testing.T) {
		buf := bytesmust.NewBuffer(nil)
		n := buf.WriteString("hello")
		if n != 5 {
			t.Errorf("expected 5 bytes written, got %d", n)
		}
		if buf.String() != "hello" {
			t.Errorf("expected 'hello', got %s", buf.String())
		}
	})

	t.Run("ReadByte", func(t *testing.T) {
		buf := bytesmust.NewBufferString("a")
		b := buf.ReadByte()
		if b != 'a' {
			t.Errorf("expected 'a', got %c", b)
		}
	})

	t.Run("ReadString", func(t *testing.T) {
		buf := bytesmust.NewBufferString("hello\nworld")
		line := buf.ReadString('\n')
		if line != "hello\n" {
			t.Errorf("expected 'hello\\n', got %s", line)
		}
	})
}

func TestReader(t *testing.T) {
	t.Run("NewReader", func(t *testing.T) {
		r := bytesmust.NewReader([]byte("hello"))
		if r == nil {
			t.Error("NewReader returned nil")
		}
	})

	t.Run("Read", func(t *testing.T) {
		r := bytesmust.NewReader([]byte("hello"))
		buf := make([]byte, 5)
		n := r.Read(buf)
		if n != 5 {
			t.Errorf("expected 5 bytes read, got %d", n)
		}
		if !bytes.Equal(buf, []byte("hello")) {
			t.Errorf("expected 'hello', got %s", buf)
		}
	})

	t.Run("ReadAt", func(t *testing.T) {
		r := bytesmust.NewReader([]byte("hello world"))
		buf := make([]byte, 5)
		n := r.ReadAt(buf, 6)
		if n != 5 {
			t.Errorf("expected 5 bytes read, got %d", n)
		}
		if !bytes.Equal(buf, []byte("world")) {
			t.Errorf("expected 'world', got %s", buf)
		}
	})

	t.Run("ReadByte", func(t *testing.T) {
		r := bytesmust.NewReader([]byte("a"))
		b := r.ReadByte()
		if b != 'a' {
			t.Errorf("expected 'a', got %c", b)
		}
	})

	t.Run("Seek", func(t *testing.T) {
		r := bytesmust.NewReader([]byte("hello"))
		pos := r.Seek(2, 0)
		if pos != 2 {
			t.Errorf("expected position 2, got %d", pos)
		}
	})

	t.Run("Len and Size", func(t *testing.T) {
		r := bytesmust.NewReader([]byte("hello"))
		if r.Len() != 5 {
			t.Errorf("expected Len 5, got %d", r.Len())
		}
		if r.Size() != 5 {
			t.Errorf("expected Size 5, got %d", r.Size())
		}
	})
}
