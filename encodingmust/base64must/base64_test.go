package base64must_test

import (
	"bytes"
	"encoding/base64"
	"strings"
	"testing"

	"github.com/Jumpaku/go-mustd"
	"github.com/Jumpaku/go-mustd/encodingmust/base64must"
	"github.com/Jumpaku/go-mustd/iomust"
)

func TestEncoding(t *testing.T) {
	t.Run("EncodeToString and DecodeString", func(t *testing.T) {
		data := []byte("hello world")
		encoded := base64must.StdEncoding.EncodeToString(data)
		decoded := base64must.StdEncoding.DecodeString(encoded)

		if !bytes.Equal(decoded, data) {
			t.Errorf("expected %s, got %s", data, decoded)
		}
	})

	t.Run("DecodeString invalid data panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("DecodeString did not panic with invalid data")
			}
		}()
		base64must.StdEncoding.DecodeString("!!!invalid!!!")
	})

	t.Run("Encode and Decode", func(t *testing.T) {
		data := []byte("hello")
		dst := make([]byte, base64must.StdEncoding.EncodedLen(len(data)))
		base64must.StdEncoding.Encode(dst, data)

		decodeDst := make([]byte, base64must.StdEncoding.DecodedLen(len(dst)))
		n := base64must.StdEncoding.Decode(decodeDst, dst)

		if !bytes.Equal(decodeDst[:n], data) {
			t.Errorf("expected %s, got %s", data, decodeDst[:n])
		}
	})

	t.Run("AppendEncode", func(t *testing.T) {
		dst := []byte("prefix:")
		data := []byte("hello")
		result := base64must.StdEncoding.AppendEncode(dst, data)

		if !bytes.HasPrefix(result, []byte("prefix:")) {
			t.Error("AppendEncode did not preserve prefix")
		}
	})

	t.Run("AppendDecode", func(t *testing.T) {
		dst := []byte("prefix:")
		encoded := base64must.StdEncoding.EncodeToString([]byte("hello"))
		result := base64must.StdEncoding.AppendDecode(dst, []byte(encoded))

		if !bytes.HasPrefix(result, []byte("prefix:")) {
			t.Error("AppendDecode did not preserve prefix")
		}
		if !bytes.HasSuffix(result, []byte("hello")) {
			t.Error("AppendDecode did not decode correctly")
		}
	})

	t.Run("EncodedLen and DecodedLen", func(t *testing.T) {
		data := []byte("hello")
		encodedLen := base64must.StdEncoding.EncodedLen(len(data))
		if encodedLen == 0 {
			t.Error("EncodedLen returned 0")
		}

		decodedLen := base64must.StdEncoding.DecodedLen(encodedLen)
		if decodedLen == 0 {
			t.Error("DecodedLen returned 0")
		}
	})
}

func TestNewEncoding(t *testing.T) {
	t.Run("custom encoding", func(t *testing.T) {
		enc := base64must.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
		if enc == nil {
			t.Error("NewEncoding returned nil")
		}

		data := []byte("test")
		encoded := enc.EncodeToString(data)
		decoded := enc.DecodeString(encoded)

		if !bytes.Equal(decoded, data) {
			t.Errorf("expected %s, got %s", data, decoded)
		}
	})
}

func TestEncodingStrict(t *testing.T) {
	t.Run("strict encoding", func(t *testing.T) {
		strict := base64must.StdEncoding.Strict()
		if strict == nil {
			t.Error("Strict returned nil")
		}

		data := []byte("hello")
		encoded := strict.EncodeToString(data)
		decoded := strict.DecodeString(encoded)

		if !bytes.Equal(decoded, data) {
			t.Errorf("expected %s, got %s", data, decoded)
		}
	})
}

func TestEncodingWithPadding(t *testing.T) {
	t.Run("custom padding", func(t *testing.T) {
		enc := base64must.StdEncoding.WithPadding(base64.NoPadding)
		if enc == nil {
			t.Error("WithPadding returned nil")
		}

		data := []byte("hello")
		encoded := enc.EncodeToString(data)

		// No padding should mean no '=' characters
		if strings.Contains(encoded, "=") {
			t.Error("encoded string contains padding")
		}
	})
}

func TestPredefinedEncodings(t *testing.T) {
	t.Run("URLEncoding", func(t *testing.T) {
		data := []byte("hello world")
		encoded := base64must.URLEncoding.EncodeToString(data)
		decoded := base64must.URLEncoding.DecodeString(encoded)

		if !bytes.Equal(decoded, data) {
			t.Errorf("expected %s, got %s", data, decoded)
		}
	})

	t.Run("RawStdEncoding", func(t *testing.T) {
		data := []byte("hello")
		encoded := base64must.RawStdEncoding.EncodeToString(data)

		if strings.Contains(encoded, "=") {
			t.Error("RawStdEncoding should not have padding")
		}

		decoded := base64must.RawStdEncoding.DecodeString(encoded)
		if !bytes.Equal(decoded, data) {
			t.Errorf("expected %s, got %s", data, decoded)
		}
	})

	t.Run("RawURLEncoding", func(t *testing.T) {
		data := []byte("hello")
		encoded := base64must.RawURLEncoding.EncodeToString(data)

		if strings.Contains(encoded, "=") {
			t.Error("RawURLEncoding should not have padding")
		}

		decoded := base64must.RawURLEncoding.DecodeString(encoded)
		if !bytes.Equal(decoded, data) {
			t.Errorf("expected %s, got %s", data, decoded)
		}
	})
}

func TestNewEncoder(t *testing.T) {
	t.Run("encode stream", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := iomust.WriterOf(buf)
		enc := base64must.NewEncoder(base64.StdEncoding, w)

		n := enc.Write([]byte("hello"))
		if n != 5 {
			t.Errorf("expected 5 bytes written, got %d", n)
		}

		// Close the underlying WriteCloser to flush
		mustd.Must0(enc.WriteCloser().Close())

		if buf.Len() == 0 {
			t.Error("encoder did not write data")
		}
	})
}

func TestNewDecoder(t *testing.T) {
	t.Run("decode stream", func(t *testing.T) {
		encoded := base64must.StdEncoding.EncodeToString([]byte("hello world"))
		r := iomust.ReaderOf(strings.NewReader(encoded))
		dec := base64must.NewDecoder(base64.StdEncoding, r)

		decoded := iomust.ReadAll(dec)
		if !bytes.Equal(decoded, []byte("hello world")) {
			t.Errorf("expected 'hello world', got %s", decoded)
		}
	})
}
