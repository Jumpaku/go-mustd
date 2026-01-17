// Package fmtmust provides wrappers for the fmt package with panicking error handling.
package fmtmust

import (
	"fmt"

	"github.com/Jumpaku/go-mustd"
	"github.com/Jumpaku/go-mustd/iomust"
)

// Fprint formats using the default formats for its operands and writes to w. Panics if an error occurs.
func Fprint(w iomust.Writer, a ...any) (n int) {
	return mustd.Must1(fmt.Fprint(w.Writer(), a...))
}

// Fprintf formats according to a format specifier and writes to w. Panics if an error occurs.
func Fprintf(w iomust.Writer, format string, a ...any) (n int) {
	return mustd.Must1(fmt.Fprintf(w.Writer(), format, a...))
}

// Fprintln formats using the default formats for its operands and writes to w. Panics if an error occurs.
func Fprintln(w iomust.Writer, a ...any) (n int) {
	return mustd.Must1(fmt.Fprintln(w.Writer(), a...))
}

// Fscan scans text read from r. Panics if an error occurs.
func Fscan(r iomust.Reader, a ...any) (n int) {
	return mustd.Must1(fmt.Fscan(r.Reader(), a...))
}

// Fscanf scans text read from r according to a format string. Panics if an error occurs.
func Fscanf(r iomust.Reader, format string, a ...any) (n int) {
	return mustd.Must1(fmt.Fscanf(r.Reader(), format, a...))
}

// Fscanln is similar to Fscan, but stops scanning at a newline. Panics if an error occurs.
func Fscanln(r iomust.Reader, a ...any) (n int) {
	return mustd.Must1(fmt.Fscanln(r.Reader(), a...))
}

// Scan scans text read from standard input. Panics if an error occurs.
func Scan(a ...any) (n int) {
	return mustd.Must1(fmt.Scan(a...))
}

// Scanf scans text read from standard input according to a format string. Panics if an error occurs.
func Scanf(format string, a ...any) (n int) {
	return mustd.Must1(fmt.Scanf(format, a...))
}

// Scanln is similar to Scan, but stops scanning at a newline. Panics if an error occurs.
func Scanln(a ...any) (n int) {
	return mustd.Must1(fmt.Scanln(a...))
}

// Sscan scans the string str. Panics if an error occurs.
func Sscan(str string, a ...any) (n int) {
	return mustd.Must1(fmt.Sscan(str, a...))
}

// Sscanf scans the string str according to a format string. Panics if an error occurs.
func Sscanf(str string, format string, a ...any) (n int) {
	return mustd.Must1(fmt.Sscanf(str, format, a...))
}

// Sscanln is similar to Sscan, but stops scanning at a newline. Panics if an error occurs.
func Sscanln(str string, a ...any) (n int) {
	return mustd.Must1(fmt.Sscanln(str, a...))
}
