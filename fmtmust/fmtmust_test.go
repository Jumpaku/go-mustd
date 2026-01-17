package fmtmust_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Jumpaku/go-mustd/fmtmust"
	"github.com/Jumpaku/go-mustd/iomust"
)

func TestFprint(t *testing.T) {
	buf := &bytes.Buffer{}
	w := iomust.WriterOf(buf)

	n := fmtmust.Fprint(w, "hello", " ", "world")
	if n != 11 {
		t.Errorf("expected 11 bytes written, got %d", n)
	}
	if buf.String() != "hello world" {
		t.Errorf("expected 'hello world', got %s", buf.String())
	}
}

func TestFprintf(t *testing.T) {
	buf := &bytes.Buffer{}
	w := iomust.WriterOf(buf)

	n := fmtmust.Fprintf(w, "hello %s %d", "world", 42)
	if buf.String() != "hello world 42" {
		t.Errorf("expected 'hello world 42', got %s", buf.String())
	}
	if n != 14 {
		t.Errorf("expected 14 bytes written, got %d", n)
	}
}

func TestFprintln(t *testing.T) {
	buf := &bytes.Buffer{}
	w := iomust.WriterOf(buf)

	n := fmtmust.Fprintln(w, "hello", "world")
	if buf.String() != "hello world\n" {
		t.Errorf("expected 'hello world\\n', got %s", buf.String())
	}
	if n != 12 {
		t.Errorf("expected 12 bytes written, got %d", n)
	}
}

func TestFscan(t *testing.T) {
	r := iomust.ReaderOf(strings.NewReader("hello world"))
	var s1, s2 string

	n := fmtmust.Fscan(r, &s1, &s2)
	if n != 2 {
		t.Errorf("expected 2 items scanned, got %d", n)
	}
	if s1 != "hello" || s2 != "world" {
		t.Errorf("expected 'hello' and 'world', got '%s' and '%s'", s1, s2)
	}
}

func TestFscanf(t *testing.T) {
	r := iomust.ReaderOf(strings.NewReader("hello 42"))
	var s string
	var i int

	n := fmtmust.Fscanf(r, "%s %d", &s, &i)
	if n != 2 {
		t.Errorf("expected 2 items scanned, got %d", n)
	}
	if s != "hello" || i != 42 {
		t.Errorf("expected 'hello' and 42, got '%s' and %d", s, i)
	}
}

func TestFscanln(t *testing.T) {
	r := iomust.ReaderOf(strings.NewReader("hello world\n"))
	var s1, s2 string

	n := fmtmust.Fscanln(r, &s1, &s2)
	if n != 2 {
		t.Errorf("expected 2 items scanned, got %d", n)
	}
	if s1 != "hello" || s2 != "world" {
		t.Errorf("expected 'hello' and 'world', got '%s' and '%s'", s1, s2)
	}
}

func TestSscan(t *testing.T) {
	var s1, s2 string

	n := fmtmust.Sscan("hello world", &s1, &s2)
	if n != 2 {
		t.Errorf("expected 2 items scanned, got %d", n)
	}
	if s1 != "hello" || s2 != "world" {
		t.Errorf("expected 'hello' and 'world', got '%s' and '%s'", s1, s2)
	}
}

func TestSscanf(t *testing.T) {
	var s string
	var i int

	n := fmtmust.Sscanf("hello 42", "%s %d", &s, &i)
	if n != 2 {
		t.Errorf("expected 2 items scanned, got %d", n)
	}
	if s != "hello" || i != 42 {
		t.Errorf("expected 'hello' and 42, got '%s' and %d", s, i)
	}
}

func TestSscanln(t *testing.T) {
	var s1, s2 string

	n := fmtmust.Sscanln("hello world", &s1, &s2)
	if n != 2 {
		t.Errorf("expected 2 items scanned, got %d", n)
	}
	if s1 != "hello" || s2 != "world" {
		t.Errorf("expected 'hello' and 'world', got '%s' and '%s'", s1, s2)
	}
}
