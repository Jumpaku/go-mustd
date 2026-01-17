package iomust

import (
	"io"

	"github.com/Jumpaku/go-mustd"
)

// ioWrapper is an internal type that wraps any io type and implements various interfaces for panicking error handling.
type ioWrapper struct {
	wrap any
}

var _ Closer = (*ioWrapper)(nil)
var _ Seeker = (*ioWrapper)(nil)

var _ Reader = (*ioWrapper)(nil)
var _ ReadCloser = (*ioWrapper)(nil)
var _ ReadSeeker = (*ioWrapper)(nil)
var _ ReadSeekCloser = (*ioWrapper)(nil)

var _ Writer = (*ioWrapper)(nil)
var _ WriteCloser = (*ioWrapper)(nil)
var _ WriteSeeker = (*ioWrapper)(nil)

var _ ReadWriter = (*ioWrapper)(nil)
var _ ReadWriteCloser = (*ioWrapper)(nil)
var _ ReadWriteSeeker = (*ioWrapper)(nil)

// ReaderOf returns a Reader that wraps the provided io.Reader.
func ReaderOf(r io.Reader) Reader {
	return &ioWrapper{wrap: r}
}

// ReadCloserOf returns a ReadCloser that wraps the provided io.ReadCloser.
func ReadCloserOf(r io.ReadCloser) ReadCloser {
	return &ioWrapper{wrap: r}
}

// WriterOf returns a Writer that wraps the provided io.Writer.
func WriterOf(w io.Writer) Writer {
	return &ioWrapper{wrap: w}
}

// WriteCloserOf returns a WriteCloser that wraps the provided io.WriteCloser.
func WriteCloserOf(w io.WriteCloser) WriteCloser {
	return &ioWrapper{wrap: w}
}

// Closer returns the underlying io.Closer, panicking if not implemented.
func (w *ioWrapper) Closer() io.Closer {
	return mustd.MustAs[io.Closer](w.wrap)
}

// Seeker returns the underlying io.Seeker, panicking if not implemented.
func (w *ioWrapper) Seeker() io.Seeker {
	return mustd.MustAs[io.Seeker](w.wrap)
}

// Reader returns the underlying io.Reader, panicking if not implemented.
func (w *ioWrapper) Reader() io.Reader {
	return mustd.MustAs[io.Reader](w.wrap)
}

// ReadCloser returns the underlying io.ReadCloser, panicking if not implemented.
func (w *ioWrapper) ReadCloser() io.ReadCloser {
	return mustd.MustAs[io.ReadCloser](w.wrap)
}

// ReadSeeker returns the underlying io.ReadSeeker, panicking if not implemented.
func (w *ioWrapper) ReadSeeker() io.ReadSeeker {
	return mustd.MustAs[io.ReadSeeker](w.wrap)
}

// ReadSeekCloser returns the underlying io.ReadSeekCloser, panicking if not implemented.
func (w *ioWrapper) ReadSeekCloser() io.ReadSeekCloser {
	return mustd.MustAs[io.ReadSeekCloser](w.wrap)
}

// Writer returns the underlying io.Writer, panicking if not implemented.
func (w *ioWrapper) Writer() io.Writer {
	return mustd.MustAs[io.Writer](w.wrap)
}

// WriteCloser returns the underlying io.WriteCloser, panicking if not implemented.
func (w *ioWrapper) WriteCloser() io.WriteCloser {
	return mustd.MustAs[io.WriteCloser](w.wrap)
}

// WriteSeeker returns the underlying io.WriteSeeker, panicking if not implemented.
func (w *ioWrapper) WriteSeeker() io.WriteSeeker {
	return mustd.MustAs[io.WriteSeeker](w.wrap)
}

// ReadWriter returns the underlying io.ReadWriter, panicking if not implemented.
func (w *ioWrapper) ReadWriter() io.ReadWriter {
	return mustd.MustAs[io.ReadWriter](w.wrap)
}

// ReadWriteCloser returns the underlying io.ReadWriteCloser, panicking if not implemented.
func (w *ioWrapper) ReadWriteCloser() io.ReadWriteCloser {
	return mustd.MustAs[io.ReadWriteCloser](w.wrap)
}

// ReadWriteSeeker returns the underlying io.ReadWriteSeeker, panicking if not implemented.
func (w *ioWrapper) ReadWriteSeeker() io.ReadWriteSeeker {
	return mustd.MustAs[io.ReadWriteSeeker](w.wrap)
}

// Close closes the underlying Closer, panicking if an error occurs.
func (w *ioWrapper) Close() {
	mustd.Must0(w.Closer().Close())
}

// Seek seeks to the given offset and whence in the underlying io.Seeker, panicking if an error occurs.
func (w *ioWrapper) Seek(offset int64, whence int) int64 {
	return mustd.Must1(w.Seeker().Seek(offset, whence))
}

// Read reads data into p from the underlying io.Reader, panicking if an error occurs.
func (w *ioWrapper) Read(p []byte) (n int) {
	return mustd.Must1(w.Reader().Read(p))
}

// Write writes data from p to the underlying io.Writer, panicking if an error occurs.
func (w *ioWrapper) Write(p []byte) (n int) {
	return mustd.Must1(w.Writer().Write(p))
}
