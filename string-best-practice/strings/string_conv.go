package strings

import (
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// Simulation implementation of convertion between byte or rune of slice and string.

// Runes2Bytes converts Slice of rune to slice of byte.
func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

// ForOnString simulates string iteration (for ... := range string)
func ForOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

// Str2Bytes convets string to byte of slice ([]byte <- string)
func Str2Bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}

// Bytes2Str convets byte of slice to string (string <- []byte)
func Bytes2Str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)
	return p
}

// Str2Runes convets string to rune of slice ([]rune <- string)
func Str2Runes(s string) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		p = append(p, r)
		s = s[size:]
	}
	return []rune(p)
}

// Runes2Str convets rune of slice to string (string <- []rune)
func Runes2Str(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return Bytes2Str(p)
}
