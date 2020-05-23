package slices

import (
	"fmt"
	"strconv"
	"testing"
)

// Test cases to explore the effeiciency comparision of
// range operation between Array, Slice and Pointer of Array.

// printSeperateLine prints a speperated line with the given title.
func printSeperateLine(title string) {
	const kwidth = 20
	// fmt.Println(strings.Repeat("-", kwidth) + title + strings.Repeat("-", kwidth))
}

// Test cases used to explore the properties of range operations
// for Array, Pointer of Array and Slice.
func TestRangeOfContainer(t *testing.T) {
	printSeperateLine("TestROC1")
	t.Run("TestROC1", func(t *testing.T) {
		arr := []MyInt{1, 2, 3}
		for _, v := range arr {
			arr = append(arr, v)
		}
		fmt.Println(arr) // [1 2 3 1 2 3]
	})

	printSeperateLine("TestROC2")
	t.Run("TestROC2", func(t *testing.T) {
		arr := []MyInt{1, 2, 3}
		newArr := []*MyInt{}
		for _, v := range arr {
			newArr = append(newArr, &v)
		}
		for _, v := range newArr {
			fmt.Print(*v, " ")
		} // [3 3 3]
		fmt.Println()
	})

	printSeperateLine("TestROC3")
	t.Run("TestROC3", func(t *testing.T) {
		type Person struct {
			name string
			age  MyInt
		}
		persons := [2]Person{{"Alice", 28}, {"Bob", 25}}
		for i, p := range persons {
			fmt.Println(i, p)
			// fail to update the element of original array because of its duplicate <persons> is provided.
			persons[1].name = "Jack"
			// fail to update the field of original array because of its being a element in its duplicate.
			p.age = 31
		}
		fmt.Println("persons:", &persons)
	})

	printSeperateLine("TestROC4")
	t.Run("TestROC4", func(t *testing.T) {
		type Person struct {
			name string
			age  MyInt
		}
		persons := [2]Person{{"Alice", 28}, {"Bob", 25}}
		pp := &persons
		for i, p := range pp {
			fmt.Println(i, p)
			// this modification has effects on the iteration.
			pp[1].name = "Jack"
			// fail to update the field of original array pointer
			p.age = 31
		}
		fmt.Println("persons:", &persons)
	})

	printSeperateLine("TestROC5")
	t.Run("TestROC5", func(t *testing.T) {
		type Person struct {
			name string
			age  MyInt
		}
		persons := []Person{{"Alice", 28}, {"Bob", 25}}
		for i, p := range persons {
			fmt.Println(i, p)
			// this modification has effects on the iteration.
			persons[1].name = "Jack"
			// fail to update the field of original slice
			p.age = 31
		}
		fmt.Println("persons:", &persons)
	})
}

const (
	arraySize   = 1 << 15
	arrayD2Size = 1 << 3
)

// Some benchmarks used to demostrate the effeciency comparasion
// of Range operation between Array, Array Pointer and Slice.
// Another two ways of iteration of array are also tested.
func BenchmarkRangeOfArray(b *testing.B) {
	b.Run("BenchmarkROA1", func(b *testing.B) {
		arr := [arraySize][arrayD2Size]MyInt{}
		for i := 0; i < b.N; i++ {
			for i, v := range arr {
				_, _ = i, v
			}
		}
	})

	b.Run("BenchmarkROA2", func(b *testing.B) {
		arr := [arraySize][arrayD2Size]MyInt{}
		parr := &arr
		for i := 0; i < b.N; i++ {
			for i, v := range parr {
				_, _ = i, v
			}
		}
	})

	b.Run("BenchmarkROA3", func(b *testing.B) {
		arr := [arraySize][arrayD2Size]MyInt{}
		sarr := arr[:]
		for i := 0; i < b.N; i++ {
			for i, v := range sarr {
				_, _ = i, v
			}
		}
	})

	b.Run("BenchmarkROA4", func(b *testing.B) {
		arr := [arraySize][arrayD2Size]MyInt{}
		for i := 0; i < b.N; i++ {
			for i := range arr {
				_ = i
			}
		}
	})

	b.Run("BenchmarkROA5", func(b *testing.B) {
		arr := [arraySize][arrayD2Size]MyInt{}
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(arr); j++ {
				_, _ = j, arr[j]
			}
		}
	})
}

// Some benchmarks used to illustrate the memclr optimization
// when emptying an array.
func BenchmarkClearArray(b *testing.B) {
	b.Run("BenchmarkCA1", func(b *testing.B) {
		arr := [arraySize]MyInt{}
		for i := 0; i < b.N; i++ {
			for i := range arr {
				arr[i] = 0
			}
		}
	})

	b.Run("BenchmarkCA2", func(b *testing.B) {
		arr := [arraySize]MyInt{}
		b.StopTimer()
		sarr := arr[:]
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			for i := range sarr {
				sarr[i] = 0
			}
		}
	})

	b.Run("BenchmarkCA3", func(b *testing.B) {
		arr := [arraySize]MyInt{}
		parr := &arr
		for i := 0; i < b.N; i++ {
			for i := range parr {
				parr[i] = 0
			}
		}
	})
}

const (
	expLimit = 30
	incrUint = 2
	initExp  = 6
)

// Addition detailed benchmarks used to illustrate the memclr optimization
// when emptying an slice.

func memclr(s []MyInt) {
	for i := range s {
		s[i] = 0
	}
}
func memsetLoop(s []MyInt, v MyInt) {
	for i := 0; i < len(s); i++ {
		s[i] = v
	}
}
func BenchmarkClearSlice(b *testing.B) {
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		sliceSize := 1 << uint(initExp+j*incrUint)
		b.Run("BenchmarkCS-1-"+strconv.Itoa(sliceSize), func(b *testing.B) {
			sli := make([]MyInt, sliceSize)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				memclr(sli)
			}
		})

		b.Run("BenchmarkCS-2-"+strconv.Itoa(sliceSize), func(b *testing.B) {
			sli := make([]MyInt, sliceSize)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				memsetLoop(sli, 0)
			}
		})
	}
}

// Shallow benchmark used to demostrate the effeciency comparision between Array and Slice.
func BenchmarkBasicSlice(b *testing.B) {
	const incrUint = 3
	const expLimit = 28
	for j := 0; initExp+j*incrUint < expLimit; j++ {
		sliceSize := 1 << uint(initExp+j*incrUint)
		b.Run("BenchmarkS-"+strconv.Itoa(sliceSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := make([]MyInt, sliceSize)
				for i, v := range s {
					s[i] = MyInt(1 + i)
					_ = v
				}
			}
		})
	}
}
func BenchmarkBasicArray(b *testing.B) {
	const incrUint = 3
	const (
		arraySize6  = 1 << 6
		arraySize9  = 1 << (6 + incrUint)
		arraySize12 = 1 << (6 + incrUint*2)
		arraySize15 = 1 << (6 + incrUint*3)
		arraySize18 = 1 << (6 + incrUint*4)
		arraySize21 = 1 << (6 + incrUint*5)
		arraySize24 = 1 << (6 + incrUint*6)
		arraySize27 = 1 << (6 + incrUint*7)
	)

	b.Run("BenchmarkA-"+strconv.Itoa(arraySize6), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := [arraySize6]MyInt{}
			for i, v := range a {
				a[i] = MyInt(1 + i)
				_ = v
			}
		}
	})

	b.Run("BenchmarkA-"+strconv.Itoa(arraySize9), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := [arraySize9]MyInt{}
			for i, v := range a {
				a[i] = MyInt(1 + i)
				_ = v
			}
		}
	})

	b.Run("BenchmarkA-"+strconv.Itoa(arraySize12), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := [arraySize12]MyInt{}
			for i, v := range a {
				a[i] = MyInt(1 + i)
				_ = v
			}
		}
	})

	b.Run("BenchmarkA-"+strconv.Itoa(arraySize15), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := [arraySize15]MyInt{}
			for i, v := range a {
				a[i] = MyInt(1 + i)
				_ = v
			}
		}
	})

	b.Run("BenchmarkA-"+strconv.Itoa(arraySize18), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := [arraySize18]MyInt{}
			for i, v := range a {
				a[i] = MyInt(1 + i)
				_ = v
			}
		}
	})

	b.Run("BenchmarkA-"+strconv.Itoa(arraySize21), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := [arraySize21]MyInt{}
			for i, v := range a {
				a[i] = MyInt(1 + i)
				_ = v
			}
		}
	})

	b.Run("BenchmarkA-"+strconv.Itoa(arraySize24), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := [arraySize24]MyInt{}
			for i, v := range a {
				a[i] = MyInt(1 + i)
				_ = v
			}
		}
	})

	b.Run("BenchmarkA-"+strconv.Itoa(arraySize27), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := [arraySize27]MyInt{}
			for i, v := range a {
				a[i] = MyInt(1 + i)
				_ = v
			}
		}
	})
}
