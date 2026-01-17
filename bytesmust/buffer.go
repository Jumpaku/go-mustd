package bytesmust

import (
	"bytes"
	"io"

	"github.com/Jumpaku/go-mustd"
)

// Buffer wraps bytes.Buffer and provides panicking error handling for buffer operations.
type Buffer struct {
	buffer *bytes.Buffer
}

// NewBuffer returns a new Buffer initialized with buf.
func NewBuffer(buf []byte) *Buffer {
	return &Buffer{buffer: bytes.NewBuffer(buf)}
}

// NewBufferString returns a new Buffer initialized with string s.
func NewBufferString(s string) *Buffer {
	return &Buffer{buffer: bytes.NewBufferString(s)}
}

// Available returns the number of bytes of available buffer.
func (b *Buffer) Available() int {
	return b.buffer.Available()
}

// AvailableBuffer returns the available buffer as a slice.
func (b *Buffer) AvailableBuffer() []byte {
	return b.buffer.AvailableBuffer()
}

// Bytes returns the contents of the unread portion of the buffer.
func (b *Buffer) Bytes() []byte {
	return b.buffer.Bytes()
}

// Cap returns the capacity of the buffer.
func (b *Buffer) Cap() int {
	return b.buffer.Cap()
}

// Grow increases the buffer's capacity.
func (b *Buffer) Grow(n int) {
	b.buffer.Grow(n)
}

// Len returns the number of bytes of the unread portion of the buffer.
func (b *Buffer) Len() int {
	return b.buffer.Len()
}

// Next returns the next n bytes from the buffer.
func (b *Buffer) Next(n int) []byte {
	return b.buffer.Next(n)
}

// Read reads up to len(p) bytes into p. Panics if an error occurs.
func (b *Buffer) Read(p []byte) (n int) {
	return mustd.Must1(b.buffer.Read(p))
}

// ReadByte reads and returns a single byte. Panics if an error occurs.
func (b *Buffer) ReadByte() byte {
	return mustd.Must1(b.buffer.ReadByte())
}

// ReadBytes reads until the first occurrence of delim. Panics if an error occurs.
func (b *Buffer) ReadBytes(delim byte) (line []byte) {
	return mustd.Must1(b.buffer.ReadBytes(delim))
}

// ReadFrom reads data from r into the buffer. Panics if an error occurs.
func (b *Buffer) ReadFrom(r io.Reader) (n int64) {
	return mustd.Must1(b.buffer.ReadFrom(r))
}

// ReadRune reads a Unicode character and its byte size. Panics if an error occurs.
func (b *Buffer) ReadRune() (r rune, size int) {
	return mustd.Must2(b.buffer.ReadRune())
}

// ReadString reads until the first occurrence of delim and returns the data as a string. Panics if an error occurs.
func (b *Buffer) ReadString(delim byte) (line string) {
	return mustd.Must1(b.buffer.ReadString(delim))
}

// Reset clears the buffer, discarding any unread data.
func (b *Buffer) Reset() {
	b.buffer.Reset()
}

// String returns the contents of the buffer as a string.
func (b *Buffer) String() string {
	return b.buffer.String()
}

// Truncate truncates the buffer so it no longer contains the given number of bytes.
func (b *Buffer) Truncate(n int) {
	b.buffer.Truncate(n)
}

// UnreadByte unreads the last byte read. Panics if the buffer is empty.
func (b *Buffer) UnreadByte() {
	mustd.Must0(b.buffer.UnreadByte())
}

// Write writes the contents of p to the buffer. Panics if an error occurs.
func (b *Buffer) Write(p []byte) (n int) {
	return mustd.Must1(b.buffer.Write(p))
}

// WriteByte writes a single byte to the buffer. Panics if an error occurs.
func (b *Buffer) WriteByte(c byte) {
	mustd.Must0(b.buffer.WriteByte(c))
}

// WriteRune writes a Unicode character to the buffer. Panics if an error occurs.
func (b *Buffer) WriteRune(r rune) (n int) {
	return mustd.Must1(b.buffer.WriteRune(r))
}

// WriteString writes the string s to the buffer. Panics if an error occurs.
func (b *Buffer) WriteString(s string) (n int) {
	return mustd.Must1(b.buffer.WriteString(s))
}

// WriteTo writes the contents of the buffer to w. Panics if an error occurs.
func (b *Buffer) WriteTo(w io.Writer) (n int64) {
	return mustd.Must1(b.buffer.WriteTo(w))
}
