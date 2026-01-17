// Package timemust provides wrappers for the time package with panicking error handling.
package timemust

import (
	"time"

	"github.com/Jumpaku/go-mustd"
)

// LoadLocation returns the Location with the given name. Panics if an error occurs.
func LoadLocation(name string) *time.Location {
	return mustd.Must1(time.LoadLocation(name))
}

// LoadLocationFromTZData returns a Location with the given name from data. Panics if an error occurs.
func LoadLocationFromTZData(name string, data []byte) *time.Location {
	return mustd.Must1(time.LoadLocationFromTZData(name, data))
}

// Parse parses a formatted string and returns the time value it represents. Panics if an error occurs.
func Parse(layout, value string) time.Time {
	return mustd.Must1(time.Parse(layout, value))
}

// ParseInLocation parses a formatted string and returns the time value it represents in the given location. Panics if an error occurs.
func ParseInLocation(layout, value string, loc *time.Location) time.Time {
	return mustd.Must1(time.ParseInLocation(layout, value, loc))
}
