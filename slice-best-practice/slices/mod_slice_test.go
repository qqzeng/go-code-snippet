package slices

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// Test cases for the  modification operations library of MyInt slice.

func benchmarkCloneSlice(b *testing.B, f func(b *testing.B, sz int, cloner func(ori []MyInt) (result []MyInt))) {
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		sliceSize := 1 << uint(initExp+j*incrUint)
		b.Run("BenchmarkCS-1-"+strconv.Itoa(j), func(b *testing.B) {
			f(b, sliceSize, Clone)
		})

		b.Run("BenchmarkCS-2-"+strconv.Itoa(j), func(b *testing.B) {
			f(b, sliceSize, Clone2)
		})

		b.Run("BenchmarkCS-3-"+strconv.Itoa(j), func(b *testing.B) {
			f(b, sliceSize, Clone3)
		})
	}
}

// Shallow benchmark used to demostrate the effeiciency comparasion between
// three ways of opeartion on cloning an Slice.
func BenchmarkCloneSlice(b *testing.B) {
	benchmarkCloneSlice(b, func(b *testing.B, sz int, cloner func(ori []MyInt) (result []MyInt)) {
		sli := make([]MyInt, sz)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = cloner(sli)
		}
	})
}

func benchmarkDeleteSingleOfSlice(b *testing.B, f func(b *testing.B, sz int, deleter func(ori []MyInt, delIndex int) (result []MyInt))) {
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		sliceSize := 1 << uint(initExp+j*incrUint)
		b.Run("BenchmarkDSOS-1-"+strconv.Itoa(j), func(b *testing.B) {
			f(b, sliceSize, DelInOrder)
		})

		b.Run("BenchmarkDSOS-2-"+strconv.Itoa(j), func(b *testing.B) {
			f(b, sliceSize, DelInOrder2)
		})

		b.Run("BenchmarkDSOS-3-"+strconv.Itoa(j), func(b *testing.B) {
			f(b, sliceSize, DelOutOfOrder)
		})
	}
}

// Shallow benchmark used to demostrate the effeiciency comparasion between
// three ways of opeartion on deleting single element from an Slice.
func BenchmarkDeleteSingleOfSlice(b *testing.B) {
	benchmarkDeleteSingleOfSlice(b, func(b *testing.B, sz int, deleter func(ori []MyInt, delIndex int) (result []MyInt)) {
		sli := make([]MyInt, sz)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			delIndex := r.Intn(sz)
			_ = deleter(sli, delIndex)
		}
	})
}

func generateInterval(r *rand.Rand, max int) (int, int) {
	delIndexL := r.Intn(max)
	delIndexH := r.Intn(max)
	if delIndexL > delIndexH {
		return delIndexH, delIndexL
	}
	return delIndexL, delIndexH
}

func benchmarkDeleteRangeOfSlice(b *testing.B, f func(b *testing.B, sz int,
	deleter func(sli []MyInt, dl int, dh int, release bool) (result []MyInt))) {
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		sliceSize := 1 << uint(initExp+j*incrUint)
		b.Run("BenchmarkDROS-1-"+strconv.Itoa(j), func(b *testing.B) {
			f(b, sliceSize, DelRangeInOrder)
		})

		b.Run("BenchmarkDROS-2-"+strconv.Itoa(j), func(b *testing.B) {
			f(b, sliceSize, DelRangeInOrder2)
		})

		b.Run("BenchmarkDROS-3-"+strconv.Itoa(j), func(b *testing.B) {
			f(b, sliceSize, DelRangeOutOfOrder)
		})
	}
}

// Shallow benchmark used to demostrate the effeiciency comparasion between
// three ways of opeartion on deleting a range of elements from an Slice.
func BenchmarkDeleteRangeOfSlice2(b *testing.B) {
	benchmarkDeleteRangeOfSlice(b, func(b *testing.B, sz int,
		deleter func(sli []MyInt, dl int, dh int, release bool) (result []MyInt)) {
		sli := make([]MyInt, sz)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dl, dh := generateInterval(r, sz)
			_ = deleter(sli, dl, dh, false)
		}
	})
}

// Shallow benchmark used to demostrate the effeiciency comparasion between
// two ways of opeartion on inserting a another Slice into current Slice.
func BenchmarkInsertRangeOfSlice(b *testing.B) {
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		sliceSize := 1 << uint(initExp+j*incrUint)
		b.Run("BenchmarkIROS-1-"+strconv.Itoa(j), func(b *testing.B) {
			sli := make([]MyInt, sliceSize)
			sli2 := make([]MyInt, sliceSize/2)
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				insertIndex := r.Intn(sliceSize)
				_ = InsertSlice(sli, sli2, insertIndex)
			}
		})

		b.Run("BenchmarkIROS-2-"+strconv.Itoa(j), func(b *testing.B) {
			sli := make([]MyInt, sliceSize)
			sli2 := make([]MyInt, sliceSize/2)
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				insertIndex := r.Intn(sliceSize)
				_ = InsertSlice2(sli, sli2, insertIndex)
			}
		})
	}
}

// Shallow benchmark used to demostrate the effeiciency comparasion between
// three ways of opeartion on growing a Slice.
func BenchmarkGrowSlice(b *testing.B) {
	const innerLoops = 20
	const preAllocSize = innerLoops * 5
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		b.Run("BenchmarkGS-1-"+strconv.Itoa(j), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s []int
				for k := 0; k < innerLoops; k++ {
					s = append(s, []int{j, j + 1, j + 2, j + 3, j + 4}...)
				}
			}
		})

		b.Run("BenchmarkGS-2-"+strconv.Itoa(j), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := make([]int, 0, preAllocSize)
				for j := 0; j < innerLoops; j++ {
					s = append(s, []int{j, j + 1, j + 2, j + 3, j + 4}...)
				}
			}
		})

		b.Run("BenchmarkGS-3-"+strconv.Itoa(j), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := make([]int, preAllocSize)
				n := 0
				for j := 0; j < innerLoops; j++ {
					n += copy(s[n:], []int{j, j + 1, j + 2, j + 3, j + 4})
				}
			}
		})
	}
}
