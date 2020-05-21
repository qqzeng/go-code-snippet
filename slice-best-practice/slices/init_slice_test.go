package slices

import (
	"fmt"
	"testing"
	"unsafe"
)

// Test case used to learn three special states of slice.
// See https://juejin.im/post/5bea58df6fb9a049f153bca8
func TestThreeSpecialStateOfSlice(t *testing.T) {
	t.Run("TestTSSOS", func(t *testing.T) {
		var s1 []int            // nil slice
		var s2 = []int{}        // empty slice
		var s3 = make([]int, 0) // empty slice
		var s4 = *new([]int)    // nil slice

		fmt.Println(len(s1), len(s2), len(s3), len(s4))
		fmt.Println(cap(s1), cap(s2), cap(s3), cap(s4))
		fmt.Println(s1, s2, s3, s4)

		var a1 = *(*[3]int)(unsafe.Pointer(&s1))
		var a2 = *(*[3]int)(unsafe.Pointer(&s2))
		var a3 = *(*[3]int)(unsafe.Pointer(&s3))
		var a4 = *(*[3]int)(unsafe.Pointer(&s4))

		fmt.Println(a1)
		fmt.Println(a2)
		fmt.Println(a3)
		fmt.Println(a4)

		var s5 = make([]struct{ x, y, z int }, 0)
		var a5 = *(*[3]int)(unsafe.Pointer(&s5))
		fmt.Println(a5)
	})
}
