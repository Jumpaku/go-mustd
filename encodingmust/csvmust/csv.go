// Package csvmust provides wrappers for the encoding/csv package with panicking error handling.
package csvmust

import (
	"encoding/csv"

	"github.com/Jumpaku/go-mustd"
	"github.com/Jumpaku/go-mustd/iomust"
)

// Reader wraps encoding/csv.Reader and provides panicking error handling for CSV reading operations.
type Reader struct {
	csv.Reader
}

// NewReader returns a new Reader that reads from r.
func NewReader(r iomust.Reader) *Reader {
	return &Reader{Reader: *csv.NewReader(r.Reader())}
}

// FieldPos returns the line and column of the specified field.
func (r *Reader) FieldPos(field int) (line, column int) {
	return r.Reader.FieldPos(field)
}

// InputOffset returns the input offset of the reader.
func (r *Reader) InputOffset() int64 {
	return r.Reader.InputOffset()
}

// Read reads one record from the CSV input. Panics if an error occurs.
func (r *Reader) Read() (record []string) {
	return mustd.Must1(r.Reader.Read())
}

// ReadAll reads all records from the CSV input. Panics if an error occurs.
func (r *Reader) ReadAll() (records [][]string) {
	return mustd.Must1(r.Reader.ReadAll())
}

// Writer wraps encoding/csv.Writer and provides panicking error handling for CSV writing operations.
type Writer struct {
	csv.Writer
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w iomust.Writer) *Writer {
	return &Writer{Writer: *csv.NewWriter(w.Writer())}
}

// Flush writes any buffered data to the underlying writer.
func (w *Writer) Flush() {
	w.Writer.Flush()
}

// Write writes a single CSV record. Panics if an error occurs.
func (w *Writer) Write(record []string) {
	mustd.Must0(w.Writer.Write(record))
}

// WriteAll writes multiple CSV records. Panics if an error occurs.
func (w *Writer) WriteAll(records [][]string) {
	mustd.Must0(w.Writer.WriteAll(records))
}
