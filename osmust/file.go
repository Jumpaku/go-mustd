// Package osmust provides wrappers for the os package with panicking error handling.
package osmust

import (
	"io"
	"os"

	"github.com/Jumpaku/go-mustd"
	"github.com/Jumpaku/go-mustd/iomust"
)

// File wraps os.File and provides panicking error handling for file operations.
type File struct {
	file *os.File
}

// FileOf returns a File wrapping the provided os.File.
func FileOf(f *os.File) *File {
	return &File{file: f}
}

// File returns the underlying os.File.
func (f *File) File() *os.File {
	return f.file
}

// Chdir changes the current working directory to the file. Panics if an error occurs.
func (f *File) Chdir() {
	mustd.Must0(f.file.Chdir())
}

// Chmod changes the mode of the file. Panics if an error occurs.
func (f *File) Chmod(mode os.FileMode) {
	mustd.Must0(f.file.Chmod(mode))
}

// Chown changes the numeric uid and gid of the file. Panics if an error occurs.
func (f *File) Chown(uid, gid int) {
	mustd.Must0(f.file.Chown(uid, gid))
}

// Close closes the file. Panics if an error occurs.
func (f *File) Close() {
	mustd.Must0(f.file.Close())
}

// Fd returns the file descriptor.
func (f *File) Fd() uintptr {
	return f.file.Fd()
}

// Name returns the name of the file.
func (f *File) Name() string {
	return f.file.Name()
}

// Read reads up to len(b) bytes from the file. Panics if an error occurs, except for io.EOF which is treated as a normal condition.
func (f *File) Read(b []byte) (n int) {
	n, err := f.file.Read(b)
	if err != nil && err != io.EOF {
		mustd.Must0(err)
	}
	return n
}

// ReadAt reads len(b) bytes from the file starting at byte offset off. Panics if an error occurs.
func (f *File) ReadAt(b []byte, off int64) (n int) {
	return mustd.Must1(f.file.ReadAt(b, off))
}

// ReadFrom reads data from r until EOF. Panics if an error occurs.
func (f *File) ReadFrom(r iomust.Reader) (n int64) {
	return mustd.Must1(f.file.ReadFrom(r.Reader()))
}

// Readdir reads the contents of the directory and returns n entries. Panics if an error occurs.
func (f *File) Readdir(n int) []os.FileInfo {
	return mustd.Must1(f.file.Readdir(n))
}

// Readdirnames reads the contents of the directory and returns n names. Panics if an error occurs.
func (f *File) Readdirnames(n int) (names []string) {
	return mustd.Must1(f.file.Readdirnames(n))
}

// Stat returns the FileInfo structure describing the file. Panics if an error occurs.
func (f *File) Stat() os.FileInfo {
	return mustd.Must1(f.file.Stat())
}

// Write writes len(b) bytes to the file. Panics if an error occurs.
func (f *File) Write(b []byte) (n int) {
	return mustd.Must1(f.file.Write(b))
}

// WriteAt writes len(b) bytes to the file starting at byte offset off. Panics if an error occurs.
func (f *File) WriteAt(b []byte, off int64) (n int) {
	return mustd.Must1(f.file.WriteAt(b, off))
}

// WriteString writes the string s to the file. Panics if an error occurs.
func (f *File) WriteString(s string) (n int) {
	return mustd.Must1(f.file.WriteString(s))
}

// WriteTo writes data to w. Panics if an error occurs.
func (f *File) WriteTo(w iomust.Writer) (n int64) {
	return mustd.Must1(f.file.WriteTo(w.Writer()))
}
