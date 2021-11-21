package main

import (
	"fmt"
	"sync"
)

func main() {
	var syncMap sync.Map

	syncMap.Store("foo", 1)
	syncMap.Store("bar", 2)
	syncMap.Store("foobar", 3)

	fooVal, ok := syncMap.Load("foo")
	if ok {
		fmt.Println(fooVal)
	} else {
		fmt.Println("foo is not found")
	}

	notExists, ok := syncMap.Load("non-exist")
	if ok {
		fmt.Println(notExists)
	} else {
		fmt.Println("non-exist is not found")
	}

	syncMap.Delete("foo")

	fooVal, ok = syncMap.Load("foo")
	if ok {
		fmt.Println(fooVal)
	} else {
		fmt.Println("after delete, foo is not found")
	}

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key.(string),value.(int))
		return true
	})

	barVal, loaded := syncMap.LoadAndDelete("bar")
	if loaded {
		fmt.Println("bar exists:", barVal, ", now deleted")
	} else {
		fmt.Println("bar does not exists, nothing to db")
	}

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key.(string),value.(int))
		return true
	})
}

