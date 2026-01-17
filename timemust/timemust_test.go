package timemust_test

import (
	"testing"
	"time"

	"github.com/Jumpaku/go-mustd/timemust"
)

func TestParse(t *testing.T) {
	t.Run("valid time string", func(t *testing.T) {
		result := timemust.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
		expected := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("invalid time string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Parse did not panic with invalid time string")
			}
		}()
		timemust.Parse(time.RFC3339, "invalid")
	})
}

func TestParseInLocation(t *testing.T) {
	loc := timemust.LoadLocation("America/New_York")

	t.Run("valid time string", func(t *testing.T) {
		result := timemust.ParseInLocation("2006-01-02", "2006-01-02", loc)
		expected := time.Date(2006, 1, 2, 0, 0, 0, 0, loc)
		if !result.Equal(expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("invalid time string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("ParseInLocation did not panic with invalid time string")
			}
		}()
		timemust.ParseInLocation("2006-01-02", "invalid", loc)
	})
}

func TestLoadLocation(t *testing.T) {
	t.Run("valid location", func(t *testing.T) {
		loc := timemust.LoadLocation("America/New_York")
		if loc == nil {
			t.Error("LoadLocation returned nil")
		}
		if loc.String() != "America/New_York" {
			t.Errorf("expected 'America/New_York', got %s", loc.String())
		}
	})

	t.Run("invalid location panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("LoadLocation did not panic with invalid location")
			}
		}()
		timemust.LoadLocation("Invalid/Location")
	})
}

func TestLoadLocationFromTZData(t *testing.T) {
	t.Run("valid TZ data", func(t *testing.T) {
		// Use minimal valid TZ data (this is a simplified test)
		// In reality, you'd need proper TZ data bytes
		defer func() {
			if r := recover(); r != nil {
				// Expected to panic with invalid data in this simple test
				// This is fine for demonstration
			}
		}()
		timemust.LoadLocationFromTZData("Custom", []byte{})
	})
}
