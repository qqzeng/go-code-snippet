package slices

import (
	"fmt"
	"testing"
	"unsafe"
)

// Some test cases to understand  the capacity expansion mechanism
// of Slice in depth.
func TestGrowOfsli(t *testing.T) {
	t.Run("TestGOS-1", func(t *testing.T) {
		sli := []int{10, 20, 30, 40}
		newSli := append(sli, 50)
		fmt.Printf("Before append sli = %v, Pointer = %p, len = %d, cap = %d\n", sli, &sli, len(sli), cap(sli))
		fmt.Printf("Before append newSli = %v, Pointer = %p, len = %d, cap = %d\n", newSli, &newSli, len(newSli), cap(newSli))
		newSli[1] += 10
		fmt.Printf("After append sli = %v, Pointer = %p, len = %d, cap = %d\n", sli, &sli, len(sli), cap(sli))
		fmt.Printf("After append newSli = %v, Pointer = %p, len = %d, cap = %d\n", newSli, &newSli, len(newSli), cap(newSli))
	})

	t.Run("TestGOS-2", func(t *testing.T) {
		array := [4]int{10, 20, 30, 40}
		sli := array[0:2]
		newSli := append(sli, 50)
		var pArrayOfSli = (*[3]int)(unsafe.Pointer(&array))
		var pArrOfnewSli = (*[3]int)(unsafe.Pointer(&array))
		fmt.Printf("Before sli = %v, Pointer Slice = %p, Pointer Array Of Slice = %p, Pointer Array= %p, len = %d, cap = %d\n",
			sli, &sli, pArrayOfSli, &array, len(sli), cap(sli))
		fmt.Printf("Before newSli = %v, Pointer Slice = %p, Pointer Array Of NewSlice = %p, Pointer Array= %p, len = %d, cap = %d\n",
			newSli, &newSli, pArrOfnewSli, &array, len(newSli), cap(newSli))
		newSli[1] += 10
		fmt.Printf("After sli = %v, Pointer = %p, len = %d, cap = %d\n", sli, &sli, len(sli), cap(sli))
		fmt.Printf("After newSli = %v, Pointer = %p, len = %d, cap = %d\n", newSli, &newSli, len(newSli), cap(newSli))
		fmt.Printf("After array = %v\n", array)
	})
}
