package iomust

import (
	"io"

	"github.com/Jumpaku/go-mustd"
)

// PipeReader wraps io.PipeReader and provides panicking error handling for pipe reading operations.
type PipeReader struct {
	pipeReader *io.PipeReader
}

// Close closes the reader end of the pipe. Panics if an error occurs.
func (r *PipeReader) Close() {
	mustd.Must0(r.pipeReader.Close())
}

// CloseWithError closes the reader end of the pipe with an error. Panics if an error occurs.
func (r *PipeReader) CloseWithError(err error) {
	mustd.Must0(r.pipeReader.CloseWithError(err))
}

// Read reads data into the provided slice. Panics if an error occurs, except for io.EOF which is treated as a normal condition.
func (r *PipeReader) Read(data []byte) (n int) {
	n, err := r.pipeReader.Read(data)
	if err != nil && err != io.EOF {
		mustd.Must0(err)
	}
	return n
}

// PipeWriter wraps io.PipeWriter and provides panicking error handling for pipe writing operations.
type PipeWriter struct {
	pipeWriter *io.PipeWriter
}

// Close closes the writer end of the pipe. Panics if an error occurs.
func (w *PipeWriter) Close() {
	mustd.Must0(w.pipeWriter.Close())
}

// CloseWithError closes the writer end of the pipe with an error. Panics if an error occurs.
func (w *PipeWriter) CloseWithError(err error) {
	mustd.Must0(w.pipeWriter.CloseWithError(err))
}

// Write writes data from the provided slice. Panics if an error occurs.
func (w *PipeWriter) Write(data []byte) (n int) {
	return mustd.Must1(w.pipeWriter.Write(data))
}
