package util

import (
	"reflect"
	"testing"
)

func TestClone1(t *testing.T) {
	type args struct {
		ori []int
	}

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"s1",
			args{s1},
			s1,
		},
		{
			"s2",
			args{s2},
			s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clone1(tt.args.ori); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClone2(t *testing.T) {
	type args struct {
		ori []int
	}

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"s1",
			args{s1},
			s1,
		},
		{
			"s2",
			args{s2},
			s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clone2(tt.args.ori); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelInOrder(t *testing.T) {
	type args struct {
		s []int
		i int
	}
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 3, 4, 5}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"s1",
			args{s1, 1},
			s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelInOrder(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelOutOfOrder(t *testing.T) {
	type args struct {
		s []int
		i int
	}
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 5, 3, 4}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"s1",
			args{s1, 1},
			s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelOutOfOrder(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelOutOfOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelRangeInOrder(t *testing.T) {
	type args struct {
		s       []int
		from    int
		to      int
		release bool
	}
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 4, 5}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"s1",
			args{s1, 1, 3, false},
			s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelRangeInOrder(tt.args.s, tt.args.from, tt.args.to, tt.args.release); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelRangeInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelRangeOutOfOrder(t *testing.T) {
	type args struct {
		s       []int
		from    int
		to      int
		release bool
	}
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{1, 5, 6, 4}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"s1",
			args{s1, 1, 3, false},
			s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelRangeOutOfOrder(tt.args.s, tt.args.from, tt.args.to, tt.args.release); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelRangeOutOfOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelEleOnCondition(t *testing.T) {
	type args struct {
		s     []int
		keep  func(int) bool
		clear bool
	}
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 3, 5}
	isOdd := func(i int) bool { return i%2 == 1 }
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"s1",
			args{s1, isOdd, true},
			s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelEleOnCondition(tt.args.s, tt.args.keep, tt.args.clear); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelEleOnCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertSlice1(t *testing.T) {
	type args struct {
		s        []int
		elements []int
		i        int
	}
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 3, 5}
	result := []int{1, 2, 1, 3, 5, 3, 4, 5}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"s1",
			args{s1, s2, 2},
			result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSlice1(tt.args.s, tt.args.elements, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSlice1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertSlice2(t *testing.T) {
	type args struct {
		s        []int
		elements []int
		i        int
	}
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 3, 5}
	result := []int{1, 2, 1, 3, 5, 3, 4, 5}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"s1",
			args{s1, s2, 2},
			result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSlice2(tt.args.s, tt.args.elements, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSlice2() = %v, want %v", got, tt.want)
			}
		})
	}
}
