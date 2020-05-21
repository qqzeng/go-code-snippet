package main

import (
	"fmt"

	"github.com/pkg/profile"
)

// MyInt represents a tested type.
type MyInt int

const (
	arraySize = 1 << 26
	loopLimit = 100
)

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

func testMemclr() {
	sli := make([]MyInt, arraySize)
	for i := 0; i < loopLimit; i++ {
		memclr(sli)
	}
}

func testMemsetLoop() {
	sli := make([]MyInt, arraySize)
	for i := 0; i < loopLimit; i++ {
		memsetLoop(sli, 0)
	}
}

// MemclrPprofCases run a benchmark between memclr and memsetLoop
// with pporf to demostrate the optimization of memclr.
// See https://blog.golang.org/pprof and https://github.com/pkg/profile
func main() {
	// CPU profiling by default
	defer profile.Start().Stop()
	// p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).stop()
	fmt.Println("benchmark between memclr and memsetLoop with pporf.")
	testMemclr()
	testMemsetLoop()
}
