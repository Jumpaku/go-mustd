package strconvmust_test

import (
	"testing"

	"github.com/Jumpaku/go-mustd/strconvmust"
)

func TestAtoi(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		result := strconvmust.Atoi("123")
		if result != 123 {
			t.Errorf("expected 123, got %d", result)
		}
	})

	t.Run("invalid string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Atoi did not panic with invalid string")
			}
		}()
		strconvmust.Atoi("abc")
	})
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1", true},
		{"0", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := strconvmust.ParseBool(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}

	t.Run("invalid string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("ParseBool did not panic with invalid string")
			}
		}()
		strconvmust.ParseBool("invalid")
	})
}

func TestParseFloat(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		result := strconvmust.ParseFloat("3.14", 64)
		if result != 3.14 {
			t.Errorf("expected 3.14, got %f", result)
		}
	})

	t.Run("invalid string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("ParseFloat did not panic with invalid string")
			}
		}()
		strconvmust.ParseFloat("abc", 64)
	})
}

func TestParseInt(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		result := strconvmust.ParseInt("123", 10, 64)
		if result != 123 {
			t.Errorf("expected 123, got %d", result)
		}
	})

	t.Run("invalid string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("ParseInt did not panic with invalid string")
			}
		}()
		strconvmust.ParseInt("abc", 10, 64)
	})
}

func TestParseUint(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		result := strconvmust.ParseUint("123", 10, 64)
		if result != 123 {
			t.Errorf("expected 123, got %d", result)
		}
	})

	t.Run("invalid string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("ParseUint did not panic with invalid string")
			}
		}()
		strconvmust.ParseUint("abc", 10, 64)
	})
}

func TestUnquote(t *testing.T) {
	t.Run("valid quoted string", func(t *testing.T) {
		result := strconvmust.Unquote(`"hello"`)
		if result != "hello" {
			t.Errorf("expected hello, got %s", result)
		}
	})

	t.Run("invalid string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Unquote did not panic with invalid string")
			}
		}()
		strconvmust.Unquote("not quoted")
	})
}

func TestParseComplex(t *testing.T) {
	t.Run("valid complex string", func(t *testing.T) {
		result := strconvmust.ParseComplex("1+2i", 128)
		expected := complex(1, 2)
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("invalid string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("ParseComplex did not panic with invalid string")
			}
		}()
		strconvmust.ParseComplex("abc", 128)
	})
}

func TestQuotedPrefix(t *testing.T) {
	t.Run("valid quoted prefix", func(t *testing.T) {
		result := strconvmust.QuotedPrefix(`"hello" world`)
		if result != `"hello"` {
			t.Errorf("expected \"hello\", got %s", result)
		}
	})

	t.Run("invalid string panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("QuotedPrefix did not panic with invalid string")
			}
		}()
		strconvmust.QuotedPrefix("not quoted")
	})
}

func TestUnquoteChar(t *testing.T) {
	t.Run("valid character", func(t *testing.T) {
		value, multibyte, tail := strconvmust.UnquoteChar(`a`, '"')
		if value != 'a' || multibyte != false || tail != "" {
			t.Errorf("unexpected result: %v, %v, %s", value, multibyte, tail)
		}
	})

	t.Run("invalid escape sequence panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("UnquoteChar did not panic with invalid escape sequence")
			}
		}()
		strconvmust.UnquoteChar(`\x`, '"')
	})
}
