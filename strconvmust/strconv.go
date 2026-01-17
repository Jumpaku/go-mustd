// Package strconvmust provides wrappers for the strconv package with panicking error handling.
package strconvmust

import (
	"strconv"

	"github.com/Jumpaku/go-mustd"
)

// Atoi returns the result of strconv.Atoi(s), panicking if an error occurs.
func Atoi(s string) int {
	return mustd.Must1(strconv.Atoi(s))
}

// ParseBool returns the result of strconv.ParseBool(str), panicking if an error occurs.
func ParseBool(str string) bool {
	return mustd.Must1(strconv.ParseBool(str))
}

// ParseComplex returns the result of strconv.ParseComplex(s, bitSize), panicking if an error occurs.
func ParseComplex(s string, bitSize int) complex128 {
	return mustd.Must1(strconv.ParseComplex(s, bitSize))
}

// ParseFloat returns the result of strconv.ParseFloat(s, bitSize), panicking if an error occurs.
func ParseFloat(s string, bitSize int) float64 {
	return mustd.Must1(strconv.ParseFloat(s, bitSize))
}

// ParseInt returns the result of strconv.ParseInt(s, base, bitSize), panicking if an error occurs.
func ParseInt(s string, base int, bitSize int) (i int64) {
	return mustd.Must1(strconv.ParseInt(s, base, bitSize))
}

// ParseUint returns the result of strconv.ParseUint(s, base, bitSize), panicking if an error occurs.
func ParseUint(s string, base int, bitSize int) uint64 {
	return mustd.Must1(strconv.ParseUint(s, base, bitSize))
}

// QuotedPrefix returns the result of strconv.QuotedPrefix(s), panicking if an error occurs.
func QuotedPrefix(s string) string {
	return mustd.Must1(strconv.QuotedPrefix(s))
}

// Unquote returns the result of strconv.Unquote(s), panicking if an error occurs.
func Unquote(s string) string {
	return mustd.Must1(strconv.Unquote(s))
}

// UnquoteChar returns the result of strconv.UnquoteChar(s, quote), panicking if an error occurs.
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string) {
	return mustd.Must3(strconv.UnquoteChar(s, quote))
}
