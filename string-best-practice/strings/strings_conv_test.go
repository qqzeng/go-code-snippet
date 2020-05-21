package strings

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"
)

// Test case used to convert rune slice to byte slice.
func TestRunes2Bytes(t *testing.T) {
	t.Run("TestR2B", func(t *testing.T) {
		var text = "hello world, 你好世界"
		runes := []rune(text)
		bytes := Runes2Bytes(runes)
		fmt.Println(bytes)
	})
}

const (
	expLimit = 30
	incrUint = 2
	initExp  = 6
)

func randBytes(n int) []byte {
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return b
}

func randInts(n int) []int {
	rand.Seed(time.Now().UnixNano())
	ints := make([]int, n)
	for i := range ints {
		ints[i] = rand.Intn(n)
	}
	return ints
}

func randStringBytes(n int) string {
	return string(randBytes(n))
}

// Shallow bechmark to demostrate the effecifiency comparsion between
// two ways of range operation on string.
func BenchmarkRangeString(b *testing.B) {
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		strSize := 1 << uint(initExp+j*incrUint)
		b.Run("BenchmarkRS-1-"+strconv.Itoa(strSize), func(b *testing.B) {
			s := randStringBytes(strSize)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for i := 0; i < len(s); i++ {
					_, _ = i, s[i]
				}
			}
		})

		b.Run("BenchmarkRS-2-"+strconv.Itoa(strSize), func(b *testing.B) {
			s := randStringBytes(strSize)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for i, v := range []byte(s) {
					_, _ = i, v
				}
			}
		})
	}
}

// sliceInt2String converts a slice of int to a string
// based on + operator.
func sliceInt2String(s []int) string {
	if len(s) < 1 {
		return ""
	}
	ss := strconv.Itoa(s[0])
	for i := 1; i < len(s); i++ {
		ss += "," + strconv.Itoa(s[i])
	}
	return ss
}

// sliceInt2String2 converts a slice of int to a string
// based on append.
func sliceInt2String2(s []int) string {
	if len(s) < 1 {
		return ""
	}
	b := make([]byte, 0, 256)
	b = append(b, strconv.Itoa(s[0])...)
	for i := 1; i < len(s); i++ {
		b = append(b, ',')
		b = append(b, strconv.Itoa(s[i])...)
	}
	return string(b)
}

// sliceInt2String3 converts a slice of int to a string
// based on append and unsafe.
func sliceInt2String3(s []int) string {
	if len(s) < 1 {
		return ""
	}
	b := make([]byte, 0, 256)
	b = append(b, strconv.Itoa(s[0])...)
	for i := 1; i < len(s); i++ {
		b = append(b, ',')
		b = append(b, strconv.Itoa(s[i])...)
	}
	return *(*string)(unsafe.Pointer(&b))
}

// sliceInt2String4 converts a slice of int to a string
// based on bytes.Buffer.
func sliceInt2String4(s []int) string {
	if len(s) < 1 {
		return ""
	}
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		buf.WriteString(strconv.Itoa(s[i]))
		buf.WriteString(",")
	}
	return buf.String()
}

// sliceInt2String5 converts a slice of int to a string
// based on strings.Builder.
func sliceInt2String5(s []int) string {
	if len(s) < 1 {
		return ""
	}
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		builder.WriteString(strconv.Itoa(s[i]))
		builder.WriteString(",")
	}
	return builder.String()
}

func benchmarkSliceInt2String(b *testing.B, f func(b *testing.B, strSize int, converter func(ints []int) (result string))) {
	const expLimit = 15
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		strSize := 1 << uint(initExp+j*incrUint)
		b.Run("BenchmarkSI2S-1-"+strconv.Itoa(strSize), func(b *testing.B) {
			b.ReportAllocs()
			f(b, strSize, sliceInt2String)
		})
		b.Run("BenchmarkSI2S-2-"+strconv.Itoa(strSize), func(b *testing.B) {
			b.ReportAllocs()
			f(b, strSize, sliceInt2String2)
		})
		b.Run("BenchmarkSI2S-3-"+strconv.Itoa(strSize), func(b *testing.B) {
			b.ReportAllocs()
			f(b, strSize, sliceInt2String3)
		})
		b.Run("BenchmarkSI2S-4-"+strconv.Itoa(strSize), func(b *testing.B) {
			b.ReportAllocs()
			f(b, strSize, sliceInt2String4)
		})
		b.Run("BenchmarkSI2S-5-"+strconv.Itoa(strSize), func(b *testing.B) {
			b.ReportAllocs()
			f(b, strSize, sliceInt2String5)
		})
	}
}

// Shallow bechmark to demostrate the efficiency comparsion between
// serveral ways of operation of int of slice to string.
func BenchmarkSliceInt2String(b *testing.B) {
	benchmarkSliceInt2String(b, func(b *testing.B, strSize int, converter func(ints []int) (result string)) {
		ints := randInts(strSize)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = converter(ints)
		}
	})
}

func sliceInt2Strings(s [][]int, reuse bool) []string {
	res := make([]string, len(s))
	b := make([]byte, 0, 256)
	for i, v := range s {
		if len(v) < 1 {
			res[i] = ""
			continue
		}
		b = b[:0]
		if reuse {
			b = b[:0]
		} else {
			b = make([]byte, 0, 5110242)
		}
		b = append(b, strconv.Itoa(v[0])...)
		for j := 1; j < len(v); j++ {
			b = append(b, ',')
			b = append(b, strconv.Itoa(v[j])...)
		}
		res[i] = string(b)
	}
	return res
}

// Shallow bechmark used to demostrate the efficiency comparsion between
// serveral ways of operation of int of slice to string.
func BenchmarkSliceInt2Strings(b *testing.B) {
	const arrayd1 = 1000
	const expLimit = 16
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		intSize := 1 << uint(initExp+j*incrUint)
		dIntsArr := make([][]int, intSize)
		for i := range dIntsArr {
			dIntsArr[i] = randInts(arrayd1)
		}

		b.Run("BenchmarkSI2S-1-"+strconv.Itoa(intSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sliceInt2Strings(dIntsArr, false)
			}
		})

		b.Run("BenchmarkSI2S-2-"+strconv.Itoa(intSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sliceInt2Strings(dIntsArr, true)
			}
		})
	}
}

// Benchmark used to explose the efficiency of strings.Builder().Grow().
var someBytes = []byte("some bytes sdljlk jsklj3lkjlk djlkjw")

var sinkS string

func benchmarkBuilder(b *testing.B, f func(b *testing.B, numWrite int, grow bool)) {
	const expLimit = 15
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		times := 1 << uint(initExp+j*incrUint)
		b.Run("Write-NoGrow-"+strconv.Itoa(times), func(b *testing.B) {
			b.ReportAllocs()
			f(b, times, false)
		})
		b.Run("Write-Grow-"+strconv.Itoa(times), func(b *testing.B) {
			b.ReportAllocs()
			f(b, times, true)
		})
	}
}

func BenchmarkBuildString_Builder(b *testing.B) {
	benchmarkBuilder(b, func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf strings.Builder
			if grow {
				buf.Grow(len(someBytes) * numWrite)
			}
			for i := 0; i < numWrite; i++ {
				buf.Write(someBytes)
			}
			sinkS = buf.String()
		}
	})
}
