// Package jsonmust provides wrappers for the encoding/json package with panicking error handling.
package jsonmust

import (
	"bytes"
	"encoding/json"

	"github.com/Jumpaku/go-mustd"
	"github.com/Jumpaku/go-mustd/iomust"
)

// Compact appends the JSON-encoded src to dst, eliminating insignificant whitespace. Panics if an error occurs.
func Compact(dst *bytes.Buffer, src []byte) {
	mustd.Must0(json.Compact(dst, src))
}

// Indent appends the indented JSON-encoded src to dst. Panics if an error occurs.
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) {
	mustd.Must0(json.Indent(dst, src, prefix, indent))
}

// Marshal returns the JSON encoding of v. Panics if an error occurs.
func Marshal(v any) []byte {
	return mustd.Must1(json.Marshal(v))
}

// MarshalIndent returns the indented JSON encoding of v. Panics if an error occurs.
func MarshalIndent(v any, prefix, indent string) []byte {
	return mustd.Must1(json.MarshalIndent(v, prefix, indent))
}

// Unmarshal parses the JSON-encoded data and stores the result in v. Panics if an error occurs.
func Unmarshal(data []byte, v any) {
	mustd.Must0(json.Unmarshal(data, v))
}

// Decoder wraps json.Decoder and provides panicking error handling.
type Decoder struct {
	decoder *json.Decoder
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r iomust.Reader) *Decoder {
	return &Decoder{decoder: json.NewDecoder(r.Reader())}
}

// Buffered returns a reader of the data remaining in the Decoder's buffer.
func (dec *Decoder) Buffered() iomust.Reader {
	return iomust.ReaderOf(dec.decoder.Buffered())
}

// Decode reads the next JSON-encoded value from its input and stores it in v. Panics if an error occurs.
func (dec *Decoder) Decode(v any) {
	mustd.Must0(dec.decoder.Decode(v))
}

// DisallowUnknownFields causes the Decoder to return an error when the destination is a struct and the input contains object keys which do not match any non-ignored, exported fields.
func (dec *Decoder) DisallowUnknownFields() {
	dec.decoder.DisallowUnknownFields()
}

// InputOffset returns the input stream byte offset of the current decoder position.
func (dec *Decoder) InputOffset() int64 {
	return dec.decoder.InputOffset()
}

// More reports whether there is another element in the current array or object being parsed.
func (dec *Decoder) More() bool {
	return dec.decoder.More()
}

// Token returns the next JSON token in the input stream. Panics if an error occurs.
func (dec *Decoder) Token() json.Token {
	return mustd.Must1(dec.decoder.Token())
}

// UseNumber causes the Decoder to unmarshal a number into an interface{} as a Number instead of as a float64.
func (dec *Decoder) UseNumber() {
	dec.decoder.UseNumber()
}

// Encoder wraps json.Encoder and provides panicking error handling.
type Encoder struct {
	encoder *json.Encoder
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w iomust.Writer) *Encoder {
	return &Encoder{encoder: json.NewEncoder(w.Writer())}
}

// Encode writes the JSON encoding of v. Panics if an error occurs.
func (enc *Encoder) Encode(v any) {
	mustd.Must0(enc.encoder.Encode(v))
}

// SetEscapeHTML specifies whether problematic HTML characters should be escaped inside JSON quoted strings.
func (enc *Encoder) SetEscapeHTML(on bool) {
	enc.encoder.SetEscapeHTML(on)
}

// SetIndent instructs the encoder to format each subsequent encoded value as if indented by the package-level function Indent.
func (enc *Encoder) SetIndent(prefix, indent string) {
	enc.encoder.SetIndent(prefix, indent)
}
