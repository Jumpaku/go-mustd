// Package iomust provides wrappers for the io package with panicking error handling and convenient interfaces.
package iomust

import (
	"io"

	"github.com/Jumpaku/go-mustd"
)

// Closer is an interface that wraps the basic Close method and provides access to the underlying io.Closer.
type Closer interface {
	Close()
	Closer() io.Closer
}

// Seeker is an interface that wraps the basic Seek method and provides access to the underlying io.Seeker.
type Seeker interface {
	Seek(offset int64, whence int) int64
	Seeker() io.Seeker
}

// Reader is an interface that wraps the basic Read method and provides access to the underlying io.Reader.
type Reader interface {
	Read(p []byte) (n int)
	Reader() io.Reader
}

// ReadCloser is an interface that groups Reader and Closer, and provides access to the underlying io.ReadCloser.
type ReadCloser interface {
	Reader
	Closer
	ReadCloser() io.ReadCloser
}

// ReadSeeker is an interface that groups Reader and Seeker, and provides access to the underlying io.ReadSeeker.
type ReadSeeker interface {
	Reader
	Seeker
	ReadSeeker() io.ReadSeeker
}

// ReadSeekCloser is an interface that groups Reader, Seeker, and Closer, and provides access to the underlying io.ReadSeekCloser.
type ReadSeekCloser interface {
	Reader
	Seeker
	Closer
	ReadSeekCloser() io.ReadSeekCloser
}

// Writer is an interface that wraps the basic Write method and provides access to the underlying io.Writer.
type Writer interface {
	Write(p []byte) (n int)
	Writer() io.Writer
}

// WriteCloser is an interface that groups Writer and Closer, and provides access to the underlying io.WriteCloser.
type WriteCloser interface {
	Writer
	Closer
	WriteCloser() io.WriteCloser
}

// WriteSeeker is an interface that groups Writer and Seeker, and provides access to the underlying io.WriteSeeker.
type WriteSeeker interface {
	Writer
	Seeker
	WriteSeeker() io.WriteSeeker
}

// ReadWriter is an interface that groups Reader and Writer, and provides access to the underlying io.ReadWriter.
type ReadWriter interface {
	Reader
	Writer
	ReadWriter() io.ReadWriter
}

// ReadWriteCloser is an interface that groups Reader, Writer, and Closer, and provides access to the underlying io.ReadWriteCloser.
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
	ReadWriteCloser() io.ReadWriteCloser
}

// ReadWriteSeeker is an interface that groups Reader, Writer, and Seeker, and provides access to the underlying io.ReadWriteSeeker.
type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
	ReadWriteSeeker() io.ReadWriteSeeker
}

// NopCloser returns a ReadCloser wrapping r, using io.NopCloser. Panics on error.
func NopCloser(r Reader) ReadCloser {
	return ReadCloserOf(io.NopCloser(r.Reader()))
}

// Copy copies from src to dst using io.Copy, panicking if an error occurs.
func Copy(dst Writer, src Reader) (written int64) {
	return mustd.Must1(io.Copy(dst.Writer(), src.Reader()))
}

// CopyBuffer copies from src to dst using io.CopyBuffer, panicking if an error occurs.
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64) {
	return mustd.Must1(io.CopyBuffer(dst.Writer(), src.Reader(), buf))
}

// CopyN copies n bytes from src to dst using io.CopyN, panicking if an error occurs.
func CopyN(dst Writer, src Reader, n int64) (written int64) {
	return mustd.Must1(io.CopyN(dst.Writer(), src.Reader(), n))
}

// Pipe returns a connected pair of PipeReader and PipeWriter using io.Pipe.
func Pipe() (*PipeReader, *PipeWriter) {
	r, w := io.Pipe()
	return &PipeReader{pipeReader: r}, &PipeWriter{pipeWriter: w}
}

// ReadAll reads from r until EOF using io.ReadAll, panicking if an error occurs.
func ReadAll(r Reader) []byte {
	return mustd.Must1(io.ReadAll(r.Reader()))
}

// ReadAtLeast reads at least min bytes using io.ReadAtLeast, panicking if an error occurs.
func ReadAtLeast(r Reader, buf []byte, min int) (n int) {
	return mustd.Must1(io.ReadAtLeast(r.Reader(), buf, min))
}

// ReadFull reads exactly len(buf) bytes using io.ReadFull, panicking if an error occurs.
func ReadFull(r Reader, buf []byte) (n int) {
	return mustd.Must1(io.ReadFull(r.Reader(), buf))
}

// WriteString writes string s to w using io.WriteString, panicking if an error occurs.
func WriteString(w Writer, s string) (n int) {
	return mustd.Must1(io.WriteString(w.Writer(), s))
}

// LimitReader returns a Reader that reads from r but stops with EOF after n bytes using io.LimitReader.
func LimitReader(r Reader, n int64) Reader {
	return ReaderOf(io.LimitReader(r.Reader(), n))
}

// MultiReader returns a Reader that's the logical concatenation of the provided input readers using io.MultiReader.
func MultiReader(readers ...Reader) Reader {
	rs := []io.Reader{}
	for _, r := range readers {
		rs = append(rs, r.Reader())
	}
	return ReaderOf(io.MultiReader(rs...))
}

// TeeReader returns a Reader that reads from r and writes to w using io.TeeReader.
func TeeReader(r Reader, w Writer) Reader {
	return ReaderOf(io.TeeReader(r.Reader(), w.Writer()))
}

// MultiWriter returns a Writer that duplicates its writes to all provided writers using io.MultiWriter.
func MultiWriter(writers ...Writer) Writer {
	ws := []io.Writer{}
	for _, w := range writers {
		ws = append(ws, w.Writer())
	}
	return WriterOf(io.MultiWriter(ws...))
}
