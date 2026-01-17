package iomust_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Jumpaku/go-mustd/iomust"
)

func TestCopy(t *testing.T) {
	src := iomust.ReaderOf(strings.NewReader("hello world"))
	dst := &bytes.Buffer{}
	writer := iomust.WriterOf(dst)

	written := iomust.Copy(writer, src)
	if written != 11 {
		t.Errorf("expected 11 bytes written, got %d", written)
	}
	if dst.String() != "hello world" {
		t.Errorf("expected 'hello world', got %s", dst.String())
	}
}

func TestCopyN(t *testing.T) {
	src := iomust.ReaderOf(strings.NewReader("hello world"))
	dst := &bytes.Buffer{}
	writer := iomust.WriterOf(dst)

	written := iomust.CopyN(writer, src, 5)
	if written != 5 {
		t.Errorf("expected 5 bytes written, got %d", written)
	}
	if dst.String() != "hello" {
		t.Errorf("expected 'hello', got %s", dst.String())
	}
}

func TestReadAll(t *testing.T) {
	r := iomust.ReaderOf(strings.NewReader("hello world"))
	data := iomust.ReadAll(r)
	if string(data) != "hello world" {
		t.Errorf("expected 'hello world', got %s", data)
	}
}

func TestReadFull(t *testing.T) {
	r := iomust.ReaderOf(strings.NewReader("hello world"))
	buf := make([]byte, 5)
	n := iomust.ReadFull(r, buf)
	if n != 5 {
		t.Errorf("expected 5 bytes read, got %d", n)
	}
	if string(buf) != "hello" {
		t.Errorf("expected 'hello', got %s", buf)
	}
}

func TestWriteString(t *testing.T) {
	dst := &bytes.Buffer{}
	writer := iomust.WriterOf(dst)

	n := iomust.WriteString(writer, "hello")
	if n != 5 {
		t.Errorf("expected 5 bytes written, got %d", n)
	}
	if dst.String() != "hello" {
		t.Errorf("expected 'hello', got %s", dst.String())
	}
}

func TestPipe(t *testing.T) {
	r, w := iomust.Pipe()
	if r == nil || w == nil {
		t.Error("Pipe returned nil")
	}

	go func() {
		w.Write([]byte("hello"))
		w.Close()
	}()

	buf := make([]byte, 5)
	n := r.Read(buf)
	if n != 5 {
		t.Errorf("expected 5 bytes read, got %d", n)
	}
	if string(buf) != "hello" {
		t.Errorf("expected 'hello', got %s", buf)
	}
}

func TestReaderOf(t *testing.T) {
	r := iomust.ReaderOf(strings.NewReader("hello"))
	if r == nil {
		t.Error("ReaderOf returned nil")
	}
	buf := make([]byte, 5)
	n := r.Read(buf)
	if n != 5 {
		t.Errorf("expected 5 bytes read, got %d", n)
	}
}

func TestWriterOf(t *testing.T) {
	buf := &bytes.Buffer{}
	w := iomust.WriterOf(buf)
	if w == nil {
		t.Error("WriterOf returned nil")
	}
	n := w.Write([]byte("hello"))
	if n != 5 {
		t.Errorf("expected 5 bytes written, got %d", n)
	}
}

func TestLimitReader(t *testing.T) {
	r := iomust.ReaderOf(strings.NewReader("hello world"))
	limited := iomust.LimitReader(r, 5)
	data := iomust.ReadAll(limited)
	if string(data) != "hello" {
		t.Errorf("expected 'hello', got %s", data)
	}
}

func TestMultiReader(t *testing.T) {
	r1 := iomust.ReaderOf(strings.NewReader("hello "))
	r2 := iomust.ReaderOf(strings.NewReader("world"))
	multi := iomust.MultiReader(r1, r2)
	data := iomust.ReadAll(multi)
	if string(data) != "hello world" {
		t.Errorf("expected 'hello world', got %s", data)
	}
}

func TestMultiWriter(t *testing.T) {
	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}
	w1 := iomust.WriterOf(buf1)
	w2 := iomust.WriterOf(buf2)
	multi := iomust.MultiWriter(w1, w2)

	n := multi.Write([]byte("hello"))
	if n != 5 {
		t.Errorf("expected 5 bytes written, got %d", n)
	}
	if buf1.String() != "hello" || buf2.String() != "hello" {
		t.Errorf("expected 'hello' in both buffers, got '%s' and '%s'", buf1.String(), buf2.String())
	}
}

func TestTeeReader(t *testing.T) {
	src := iomust.ReaderOf(strings.NewReader("hello"))
	buf := &bytes.Buffer{}
	w := iomust.WriterOf(buf)
	tee := iomust.TeeReader(src, w)

	data := iomust.ReadAll(tee)
	if string(data) != "hello" {
		t.Errorf("expected 'hello' from reader, got %s", data)
	}
	if buf.String() != "hello" {
		t.Errorf("expected 'hello' in buffer, got %s", buf.String())
	}
}
