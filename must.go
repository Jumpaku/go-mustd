// Package mustd provides utility functions for converting error-returning functions into panicking functions.
// This is useful for scenarios where errors are unexpected and should cause immediate program termination,
// such as in initialization code or test utilities.
package mustd

import "fmt"

// Must0 panics if err is not nil.
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// Must1 returns v if err is nil, otherwise panics with err.
func Must1[T any](v T, err error) (r T) {
	if err != nil {
		panic(err)
	}
	return v
}

// Must2 returns v0 and v1 if err is nil, otherwise panics with err.
func Must2[T0, T1 any](v0 T0, v1 T1, err error) (r0 T0, r1 T1) {
	if err != nil {
		panic(err)
	}
	return v0, v1
}

// Must3 returns v0, v1, and v2 if err is nil, otherwise panics with err.
func Must3[T0, T1, T2 any](v0 T0, v1 T1, v2 T2, err error) (r0 T0, r1 T1, r2 T2) {
	if err != nil {
		panic(err)
	}
	return v0, v1, v2
}

// MustImplement asserts that v implements type T and returns v cast to T.
// Panics if v does not implement T.
func MustImplement[T any](v any) T {
	val, ok := v.(T)
	if !ok {
		var t T
		panic(fmt.Sprintf("value must implement %T", t))
	}
	return val
}
