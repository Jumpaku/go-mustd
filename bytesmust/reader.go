// Package bytesmust provides wrappers for the bytes package with panicking error handling.
package bytesmust

import (
	"bytes"
	"io"

	"github.com/Jumpaku/go-mustd"
)

// Reader wraps bytes.Reader and provides panicking error handling for reader operations.
type Reader struct {
	reader *bytes.Reader
}

// NewReader returns a new Reader reading from b.
func NewReader(b []byte) *Reader {
	return &Reader{reader: bytes.NewReader(b)}
}

// Len returns the number of bytes of the unread portion of the slice.
func (r *Reader) Len() int {
	return r.reader.Len()
}

// Read reads data into b. Panics if an error occurs, except for io.EOF which is treated as a normal condition.
func (r *Reader) Read(b []byte) (n int) {
	n, err := r.reader.Read(b)
	if err != nil && err != io.EOF {
		mustd.Must0(err)
	}
	return n
}

// ReadAt reads len(b) bytes into b starting at offset off. Panics if an error occurs.
func (r *Reader) ReadAt(b []byte, off int64) (n int) {
	return mustd.Must1(r.reader.ReadAt(b, off))
}

// ReadByte reads and returns a single byte. Panics if an error occurs.
func (r *Reader) ReadByte() byte {
	return mustd.Must1(r.reader.ReadByte())
}

// ReadRune reads and returns a single UTF-8 encoded Unicode character. Panics if an error occurs.
func (r *Reader) ReadRune() (ch rune, size int) {
	return mustd.Must2(r.reader.ReadRune())
}

// Reset resets the Reader to be reading from b.
func (r *Reader) Reset(b []byte) {
	r.reader.Reset(b)
}

// Seek sets the offset for the next Read or Write. Panics if an error occurs.
func (r *Reader) Seek(offset int64, whence int) int64 {
	return mustd.Must1(r.reader.Seek(offset, whence))
}

// Size returns the original length of the underlying byte slice.
func (r *Reader) Size() int64 {
	return r.reader.Size()
}

// UnreadByte unreads the last byte. Panics if an error occurs.
func (r *Reader) UnreadByte() {
	mustd.Must0(r.reader.UnreadByte())
}

// UnreadRune unreads the last rune. Panics if an error occurs.
func (r *Reader) UnreadRune() {
	mustd.Must0(r.reader.UnreadRune())
}

// WriteTo writes data to w. Panics if an error occurs.
func (r *Reader) WriteTo(w io.Writer) (n int64) {
	return mustd.Must1(r.reader.WriteTo(w))
}
