// Package base64must provides wrappers for the encoding/base64 package with panicking error handling.
package base64must

import (
	"encoding/base64"

	"github.com/Jumpaku/go-mustd"
	"github.com/Jumpaku/go-mustd/iomust"
)

// RawStdEncoding is the standard base64 encoding without padding.
var RawStdEncoding = StdEncoding.WithPadding(base64.NoPadding)

// RawURLEncoding is the URL-safe base64 encoding without padding.
var RawURLEncoding = URLEncoding.WithPadding(base64.NoPadding)

// StdEncoding is the standard base64 encoding.
var StdEncoding = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

// URLEncoding is the URL-safe base64 encoding.
var URLEncoding = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

// NewDecoder returns a new base64 stream decoder.
func NewDecoder(enc *base64.Encoding, r iomust.Reader) iomust.Reader {
	return iomust.ReaderOf(base64.NewDecoder(enc, r.Reader()))
}

// NewEncoder returns a new base64 stream encoder.
func NewEncoder(enc *base64.Encoding, w iomust.Writer) iomust.WriteCloser {
	return iomust.WriteCloserOf(base64.NewEncoder(enc, w.Writer()))
}

// Encoding wraps base64.Encoding and provides panicking error handling.
type Encoding struct {
	encoding *base64.Encoding
}

// NewEncoding returns a new Encoding defined by the given alphabet.
func NewEncoding(encoder string) *Encoding {
	return &Encoding{encoding: base64.NewEncoding(encoder)}
}

// AppendDecode appends the base64 decoded src to dst. Panics if an error occurs.
func (enc *Encoding) AppendDecode(dst, src []byte) []byte {
	return mustd.Must1(enc.encoding.AppendDecode(dst, src))
}

// AppendEncode appends the base64 encoded src to dst.
func (enc *Encoding) AppendEncode(dst, src []byte) []byte {
	return enc.encoding.AppendEncode(dst, src)
}

// Decode decodes src into dst. Panics if an error occurs.
func (enc *Encoding) Decode(dst, src []byte) (n int) {
	return mustd.Must1(enc.encoding.Decode(dst, src))
}

// DecodeString returns the bytes represented by the base64 string s. Panics if an error occurs.
func (enc *Encoding) DecodeString(s string) []byte {
	return mustd.Must1(enc.encoding.DecodeString(s))
}

// DecodedLen returns the maximum length in bytes of the decoded data.
func (enc *Encoding) DecodedLen(n int) int {
	return enc.encoding.DecodedLen(n)
}

// Encode encodes src into dst.
func (enc *Encoding) Encode(dst, src []byte) {
	enc.encoding.Encode(dst, src)
}

// EncodeToString returns the base64 encoding of src.
func (enc *Encoding) EncodeToString(src []byte) string {
	return enc.encoding.EncodeToString(src)
}

// EncodedLen returns the length in bytes of the base64 encoding of n bytes.
func (enc *Encoding) EncodedLen(n int) int {
	return enc.encoding.EncodedLen(n)
}

// Strict creates a new encoding identical to enc except with strict decoding enabled.
func (enc Encoding) Strict() *Encoding {
	return &Encoding{encoding: enc.encoding.Strict()}
}

// WithPadding creates a new encoding identical to enc except with specified padding character.
func (enc Encoding) WithPadding(padding rune) *Encoding {
	return &Encoding{encoding: enc.encoding.WithPadding(padding)}
}
