package strings

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

// Shallow bechmark to demostrate the effecifiency comparsion between
// two ways of concate operation on string.
func BenchmarkConcateString(b *testing.B) {
	const expLimit = 10
	const incrUint = 1
	const initExp = 2
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		strSize := 1 << uint(initExp+j*incrUint)
		b.Run("BenchmarkCS-1-"+strconv.Itoa(strSize), func(b *testing.B) {
			s := randStringBytes(strSize)
			b.ResetTimer()
			var resultStr string
			for i := 0; i < b.N; i++ {
				resultStr += s
			}
		})

		b.Run("BenchmarkCS-2-"+strconv.Itoa(strSize), func(b *testing.B) {
			s := randStringBytes(strSize)
			b.ResetTimer()
			var buf bytes.Buffer
			for i := 0; i < b.N; i++ {
				fmt.Fprint(&buf, s)
				_ = buf.String()
			}
		})

		b.Run("BenchmarkCS-3-"+strconv.Itoa(strSize), func(b *testing.B) {
			s := randStringBytes(strSize)
			b.ResetTimer()
			var buf bytes.Buffer
			for i := 0; i < b.N; i++ {
				fmt.Fprint(&buf, s)
				bs := buf.Bytes()
				_ = *(*string)(unsafe.Pointer(&bs))
			}
		})

		b.Run("BenchmarkCS-4-"+strconv.Itoa(strSize), func(b *testing.B) {
			s := randStringBytes(strSize)
			b.ResetTimer()
			var builder strings.Builder
			for i := 0; i < b.N; i++ {
				fmt.Fprint(&builder, s)
				_ = builder.String()
			}
		})
	}
}
