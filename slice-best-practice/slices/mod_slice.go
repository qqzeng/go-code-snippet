package slices

// MyInt slice modification operations library.

// MyInt represents a tested type.
type MyInt int

// Clone clone a slice efficiently.
func Clone(ori []MyInt) []MyInt {
	oriClone := append(ori[:0:0], ori...)
	return oriClone
}

// Clone2 clone a slice using `make` and `copy`.
// if ori is nil, returns a non-nil slice.
func Clone2(ori []MyInt) []MyInt {
	oriClone := make([]MyInt, len(ori))
	copy(oriClone, ori)
	return oriClone
}

// Clone3 clone a slice using `append`.
// returns nil even if the source slice a is a non-nil blank slice.
func Clone3(ori []MyInt) []MyInt {
	oriClone := append([]MyInt(nil), ori...)
	return oriClone
}

// DelInOrder remove an elemen from a slice, keeping remain elements in order
// and relesing discarded elements (TODO).
// It panics if s is nil.
func DelInOrder(s []MyInt, i int) []MyInt {
	s = append(s[:i], s[i+1:]...)
	return s
}

// DelInOrder2 remove an elemen from a slice, keeping remain elements in order
// and relesing discarded elements (TODO).
// It panics if s is nil.
func DelInOrder2(s []MyInt, i int) []MyInt {
	s = s[:i+copy(s[i:], s[i+1:])]
	return s
}

// DelOutOfOrder remove an element from a slice.
// It panics if s is nil.
func DelOutOfOrder(s []MyInt, i int) []MyInt {
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	return s
}

// DelRangeInOrder remove a range of elements from a slice, keeping remain elements in order [)
// and relesing discarded elements
func DelRangeInOrder(s []MyInt, from int, to int, release bool) []MyInt {
	tmp := append(s[:from], s[to:]...)
	if release {
		for i := len(tmp); i < len(s); i++ {
			s[i] = 0
		}
	}
	return tmp
}

// DelRangeInOrder2 remove a range of elements from a slice, keeping remain elements in order [)
// and relesing discarded elements
func DelRangeInOrder2(s []MyInt, from int, to int, release bool) []MyInt {
	tmp := s[:from+copy(s[from:], s[to:])]
	if release {
		for i := len(tmp); i < len(s); i++ {
			s[i] = 0
		}
	}
	return tmp
}

// DelRangeOutOfOrder remove a range of elements from a slice [)
func DelRangeOutOfOrder(s []MyInt, from int, to int, release bool) []MyInt {
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
func DelEleOnCondition(s []MyInt, keep func(MyInt) bool, clear bool) []MyInt {
	// result := make([]MyInt, 0, len(s))
	result := s[:0] //  no need to allocate memory.
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

// InsertSlice insert a slice called elements into the ith index of another slice
func InsertSlice(s []MyInt, elements []MyInt, i int) []MyInt {
	s = append(s[:i], append(elements, s[i:]...)...)
	return s
}

// InsertSlice2 insert a slice called elements into the ith index of another slice (more efficient but tedious)
func InsertSlice2(s []MyInt, elements []MyInt, i int) []MyInt {
	if cap(s)-len(s) >= len(elements) {
		s = s[:len(s)+len(elements)]
		copy(s[i+len(elements):], s[i:])
		copy(s[i:], elements)
	} else {
		x := make([]MyInt, 0, len(elements)+len(s))
		x = append(x, s[:i]...)
		x = append(x, elements...)
		x = append(x, s[i:]...)
		s = x
	}
	return s
}
