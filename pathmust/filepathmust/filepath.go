// Package filepathmust provides wrappers for the path/filepath package with panicking error handling.
package filepathmust

import (
	"io/fs"
	"path/filepath"

	"github.com/Jumpaku/go-mustd"
)

// Abs returns an absolute representation of path. Panics if an error occurs.
func Abs(path string) string {
	return mustd.Must1(filepath.Abs(path))
}

// EvalSymlinks returns the path name after the evaluation of any symbolic links. Panics if an error occurs.
func EvalSymlinks(path string) string {
	return mustd.Must1(filepath.EvalSymlinks(path))
}

// Glob returns the names of all files matching pattern. Panics if an error occurs.
func Glob(pattern string) (matches []string) {
	return mustd.Must1(filepath.Glob(pattern))
}

// Localize converts a slash-separated path into an operating system path. Panics if an error occurs.
func Localize(path string) string {
	return mustd.Must1(filepath.Localize(path))
}

// Match reports whether name matches the shell pattern. Panics if an error occurs.
func Match(pattern, name string) (matched bool) {
	return mustd.Must1(filepath.Match(pattern, name))
}

// Rel returns a relative path that is lexically equivalent to targpath when joined to basepath. Panics if an error occurs.
func Rel(basepath, targpath string) string {
	return mustd.Must1(filepath.Rel(basepath, targpath))
}

// Walk walks the file tree rooted at root, calling fn for each file or directory. Panics if an error occurs.
func Walk(root string, fn filepath.WalkFunc) {
	mustd.Must0(filepath.Walk(root, fn))
}

// WalkDir walks the file tree rooted at root, calling fn for each file or directory. Panics if an error occurs.
func WalkDir(root string, fn fs.WalkDirFunc) {
	mustd.Must0(filepath.WalkDir(root, fn))
}
