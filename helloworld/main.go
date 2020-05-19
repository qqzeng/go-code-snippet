package main

import (
	"fmt"
	util "github.com/qqzeng/helloworld/util"
)


func deleteRangeInOrder() {
	s := []int{1, 2, 3, 4, 5, 6}
	s1 := util.DelRangeInOrder(s, 1, 3, true)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s1, len(s1), cap(s1))
}

func deleteRangeOutOfOrder() {
	s := []int{1, 2, 3, 4, 5, 6}
	s1 := util.DelRangeOutOfOrder(s, 1, 3, true)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s1, len(s1), cap(s1))
}

func main() {
	fmt.Println("Here are some int slice test..")
	deleteRangeInOrder()
	deleteRangeOutOfOrder()
}
