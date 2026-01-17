package jsonmust_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Jumpaku/go-mustd/encodingmust/jsonmust"
	"github.com/Jumpaku/go-mustd/iomust"
)

func TestMarshalUnmarshal(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		data := map[string]interface{}{
			"name": "test",
			"age":  42,
		}

		jsonData := jsonmust.Marshal(data)
		if len(jsonData) == 0 {
			t.Error("Marshal returned empty data")
		}

		var result map[string]interface{}
		jsonmust.Unmarshal(jsonData, &result)

		if result["name"] != "test" {
			t.Errorf("expected 'test', got %v", result["name"])
		}
		if result["age"].(float64) != 42 {
			t.Errorf("expected 42, got %v", result["age"])
		}
	})

	t.Run("unmarshal invalid JSON panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Unmarshal did not panic with invalid JSON")
			}
		}()
		var result map[string]interface{}
		jsonmust.Unmarshal([]byte("invalid json"), &result)
	})
}

func TestMarshalIndent(t *testing.T) {
	data := map[string]string{"key": "value"}

	jsonData := jsonmust.MarshalIndent(data, "", "  ")
	if !bytes.Contains(jsonData, []byte("\n")) {
		t.Error("MarshalIndent did not add indentation")
	}
}

func TestCompact(t *testing.T) {
	src := []byte(`{
		"key": "value",
		"num": 42
	}`)
	dst := &bytes.Buffer{}

	jsonmust.Compact(dst, src)

	if bytes.Contains(dst.Bytes(), []byte("\n")) {
		t.Error("Compact did not remove whitespace")
	}
}

func TestIndent(t *testing.T) {
	src := []byte(`{"key":"value","num":42}`)
	dst := &bytes.Buffer{}

	jsonmust.Indent(dst, src, "", "  ")

	if !bytes.Contains(dst.Bytes(), []byte("\n")) {
		t.Error("Indent did not add whitespace")
	}
}

func TestDecoder(t *testing.T) {
	t.Run("decode valid JSON", func(t *testing.T) {
		jsonStr := `{"name":"test","age":42}`
		r := iomust.ReaderOf(strings.NewReader(jsonStr))
		dec := jsonmust.NewDecoder(r)

		var result map[string]interface{}
		dec.Decode(&result)

		if result["name"] != "test" {
			t.Errorf("expected 'test', got %v", result["name"])
		}
	})

	t.Run("decode invalid JSON panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Decode did not panic with invalid JSON")
			}
		}()
		r := iomust.ReaderOf(strings.NewReader("invalid json"))
		dec := jsonmust.NewDecoder(r)
		var result map[string]interface{}
		dec.Decode(&result)
	})

	t.Run("Token", func(t *testing.T) {
		jsonStr := `{"key":"value"}`
		r := iomust.ReaderOf(strings.NewReader(jsonStr))
		dec := jsonmust.NewDecoder(r)

		token := dec.Token()
		if token == nil {
			t.Error("Token returned nil")
		}
	})

	t.Run("More", func(t *testing.T) {
		jsonStr := `[1,2,3]`
		r := iomust.ReaderOf(strings.NewReader(jsonStr))
		dec := jsonmust.NewDecoder(r)

		dec.Token() // Read opening bracket
		if !dec.More() {
			t.Error("More should return true")
		}
	})
}

func TestEncoder(t *testing.T) {
	t.Run("encode valid data", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := iomust.WriterOf(buf)
		enc := jsonmust.NewEncoder(w)

		data := map[string]string{"key": "value"}
		enc.Encode(data)

		if buf.Len() == 0 {
			t.Error("Encoder did not write data")
		}
	})

	t.Run("SetIndent", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := iomust.WriterOf(buf)
		enc := jsonmust.NewEncoder(w)

		enc.SetIndent("", "  ")
		data := map[string]string{"key": "value"}
		enc.Encode(data)

		if !bytes.Contains(buf.Bytes(), []byte("\n")) {
			t.Error("SetIndent did not affect output")
		}
	})

	t.Run("SetEscapeHTML", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := iomust.WriterOf(buf)
		enc := jsonmust.NewEncoder(w)

		enc.SetEscapeHTML(false)
		data := map[string]string{"key": "<html>"}
		enc.Encode(data)

		// Just verify it doesn't panic
		if buf.Len() == 0 {
			t.Error("Encoder did not write data")
		}
	})
}
