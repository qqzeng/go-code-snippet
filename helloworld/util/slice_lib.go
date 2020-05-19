package util

// int slice operations library.

// Clone1 clone a slice efficiently
func Clone1(ori []int) []int {
	oriClone := append(ori[:0:0], ori...)
	return oriClone
}

// Clone2 clone a slice using `make` and `copy`
func Clone2(ori []int) []int {
	oriClone := make([]int, len(ori))
	copy(oriClone, ori)
	return oriClone
}

// DelInOrder remove an elemen from a slice, keeping remain elements in order
// and relesing discarded elements (TODO)
func DelInOrder(s []int, i int) []int {
	s = append(s[:i], s[i+1:]...)
	// s = s[:i+copy(s[i:], s[i+1:])]
	return s
}

// DelOutOfOrder remove an element from a slice
func DelOutOfOrder(s []int, i int) []int {
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	return s
}

// DelRangeInOrder remove a range of elements from a slice, keeping remain elements in order [)
// and relesing discarded elements
func DelRangeInOrder(s []int, from int, to int, release bool) []int {
	tmp := append(s[:from], s[to:]...)
	// s = s[:from+copy(s[from:], s[to:])]
	if release {
		for i := len(tmp); i < len(s); i++ {
			s[i] = 0
		}
	}
	return tmp
}

// DelRangeOutOfOrder remove a range of elements from a slice [)
func DelRangeOutOfOrder(s []int, from int, to int, release bool) []int {
	if n := to - from; len(s)-to < n {
		copy(s[from:to], s[to:])
	} else {
		copy(s[from:to], s[len(s)-n:])
	}
	tmp := s[:len(s)-(to-from)]
	if release {
		for i := len(tmp); i < len(s); i++ {
			s[i] = 0
		}
	}
	return tmp
}

// DelEleOnCondition remove elements from a slice based on condition
func DelEleOnCondition(s []int, keep func(int) bool, clear bool) []int {
	// result := make([]int, 0, len(s))
	result := s[:0] //  无须开辟内存
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	if clear {
		temp := s[len(result):]
		for i := range temp {
			temp[i] = 0
		}
	}
	return result
}

// InsertSlice1 insert a slice called elements into the ith index of another slice
func InsertSlice1(s []int, elements []int, i int) []int {
	s = append(s[:i], append(elements, s[i:]...)...)
	return s
}

// InsertSlice2 insert a slice called elements into the ith index of another slice (more efficient but tedious)
func InsertSlice2(s []int, elements []int, i int) []int {
	if cap(s)-len(s) >= len(elements) {
		s = s[:len(s)+len(elements)]
		copy(s[i+len(elements):], s[i:])
		copy(s[i:], elements)
	} else {
		x := make([]int, 0, len(elements)+len(s))
		x = append(x, s[:i]...)
		x = append(x, elements...)
		x = append(x, s[i:]...)
		s = x
	}
	return s
}
