package bench

import (
	"github.com/qqzeng/sync-map-intro/maputil"
	"runtime"
	"sync"
	"testing"
)

func benchmarkRWMapStore(b *testing.B, numGroutine int) {
	runtime.GOMAXPROCS(numGroutine)

	rwm := maputil.NewRawRWIntMap()
	nums := nrand(b.N)
	for _, v := range nums {
		rwm.Store(v, v)
	}

	var wg sync.WaitGroup
	wg.Add(numGroutine)

	b.ResetTimer()

	for i := 0; i < numGroutine; i++ {
		go func(n int) {
			for j := 0; j < n; j++ {
				rwm.Delete(nums[j])
			}
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func benchmarkMapStore(b *testing.B, numGroutine int) {
	runtime.GOMAXPROCS(numGroutine)

	m := maputil.NewRawIntMap()
	nums := nrand(b.N)
	for _, v := range nums {
		m.Store(v, v)
	}

	var wg sync.WaitGroup
	wg.Add(numGroutine)

	b.ResetTimer()

	for i := 0; i < numGroutine; i++ {
		go func(n int) {
			for j := 0; j < n; j++ {
				m.Delete(nums[j])
			}
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func benchmarkSyncMapStore(b *testing.B, numGroutine int) {
	runtime.GOMAXPROCS(numGroutine)

	var sm sync.Map
	nums := nrand(b.N)
	for _, v := range nums {
		sm.Store(v, v)
	}

	var wg sync.WaitGroup
	wg.Add(numGroutine)

	b.ResetTimer()

	for i := 0; i < numGroutine; i++ {
		go func(n int) {
			for j := 0; j < n; j++ {
				sm.Delete(nums[j])
			}
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func BenchmarkRWMapStore1(b *testing.B) {
	benchmarkRWMapStore(b, 1)
}

func BenchmarkRWMapStore2(b *testing.B) {
	benchmarkRWMapStore(b, 2)
}

func BenchmarkRWMapStore4(b *testing.B) {
	benchmarkRWMapStore(b, 4)
}

func BenchmarkRWMapStore8(b *testing.B) {
	benchmarkRWMapStore(b, 8)
}

func BenchmarkRWMapStore16(b *testing.B) {
	benchmarkRWMapStore(b, 16)
}

func BenchmarkRWMapStore32(b *testing.B) {
	benchmarkRWMapStore(b, 32)
}

func BenchmarkRWMapStore64(b *testing.B) {
	benchmarkRWMapStore(b, 64)
}

func BenchmarkMapStore1(b *testing.B) {
	benchmarkMapStore(b, 1)
}

func BenchmarkMapStore2(b *testing.B) {
	benchmarkMapStore(b, 2)
}

func BenchmarkMapStore4(b *testing.B) {
	benchmarkMapStore(b, 4)
}

func BenchmarkMapStore8(b *testing.B) {
	benchmarkMapStore(b, 8)
}

func BenchmarkMapStore16(b *testing.B) {
	benchmarkMapStore(b, 16)
}

func BenchmarkMapStore32(b *testing.B) {
	benchmarkMapStore(b, 32)
}

func BenchmarkMapStore64(b *testing.B) {
	benchmarkMapStore(b, 64)
}

func BenchmarkSyncMapStore1(b *testing.B) {
	benchmarkSyncMapStore(b, 1)
}

func BenchmarkSyncMapStore2(b *testing.B) {
	benchmarkSyncMapStore(b, 2)
}

func BenchmarkSyncMapStore4(b *testing.B) {
	benchmarkSyncMapStore(b, 4)
}

func BenchmarkSyncMapStore8(b *testing.B) {
	benchmarkSyncMapStore(b, 8)
}

func BenchmarkSyncMapStore16(b *testing.B) {
	benchmarkSyncMapStore(b, 16)
}

func BenchmarkSyncMapStore32(b *testing.B) {
	benchmarkSyncMapStore(b, 32)
}

func BenchmarkSyncMapStore64(b *testing.B) {
	benchmarkSyncMapStore(b, 64)
}

//BenchmarkRWMapStore1
//BenchmarkRWMapStore1-12       	 8702679	       170 ns/op
//testing: BenchmarkRWMapStore1-12 left GOMAXPROCS set to 1
//BenchmarkRWMapStore2
//BenchmarkRWMapStore2-12       	 3810754	       339 ns/op
//testing: BenchmarkRWMapStore2-12 left GOMAXPROCS set to 2
//BenchmarkRWMapStore4
//BenchmarkRWMapStore4-12       	 1886948	       636 ns/op
//testing: BenchmarkRWMapStore4-12 left GOMAXPROCS set to 4
//BenchmarkRWMapStore8
//BenchmarkRWMapStore8-12       	 1000000	      1241 ns/op
//testing: BenchmarkRWMapStore8-12 left GOMAXPROCS set to 8
//BenchmarkRWMapStore16
//BenchmarkRWMapStore16-12      	  750942	      2792 ns/op
//testing: BenchmarkRWMapStore16-12 left GOMAXPROCS set to 16
//BenchmarkRWMapStore32
//BenchmarkRWMapStore32-12      	  367998	      4948 ns/op
//testing: BenchmarkRWMapStore32-12 left GOMAXPROCS set to 32
//BenchmarkRWMapStore64
//BenchmarkRWMapStore64-12      	  175092	     10653 ns/op
//testing: BenchmarkRWMapStore64-12 left GOMAXPROCS set to 64
//BenchmarkMapStore1
//BenchmarkMapStore1-12         	 7926660	       148 ns/op
//testing: BenchmarkMapStore1-12 left GOMAXPROCS set to 1
//BenchmarkMapStore2
//BenchmarkMapStore2-12         	 4333833	       359 ns/op
//testing: BenchmarkMapStore2-12 left GOMAXPROCS set to 2
//BenchmarkMapStore4
//BenchmarkMapStore4-12         	 1947223	       618 ns/op
//testing: BenchmarkMapStore4-12 left GOMAXPROCS set to 4
//BenchmarkMapStore8
//BenchmarkMapStore8-12         	 1000000	      1242 ns/op
//testing: BenchmarkMapStore8-12 left GOMAXPROCS set to 8
//BenchmarkMapStore16
//BenchmarkMapStore16-12        	  763926	      2726 ns/op
//testing: BenchmarkMapStore16-12 left GOMAXPROCS set to 16
//BenchmarkMapStore32
//BenchmarkMapStore32-12        	  335679	      4707 ns/op
//testing: BenchmarkMapStore32-12 left GOMAXPROCS set to 32
//BenchmarkMapStore64
//BenchmarkMapStore64-12        	  174895	      9158 ns/op
//testing: BenchmarkMapStore64-12 left GOMAXPROCS set to 64
//BenchmarkSyncMapStore1
//BenchmarkSyncMapStore1-12     	 4157820	       364 ns/op
//testing: BenchmarkSyncMapStore1-12 left GOMAXPROCS set to 1
//BenchmarkSyncMapStore2
//BenchmarkSyncMapStore2-12     	 3975222	       278 ns/op
//testing: BenchmarkSyncMapStore2-12 left GOMAXPROCS set to 2
//BenchmarkSyncMapStore4
//BenchmarkSyncMapStore4-12     	 4279014	       329 ns/op
//testing: BenchmarkSyncMapStore4-12 left GOMAXPROCS set to 4
//BenchmarkSyncMapStore8
//BenchmarkSyncMapStore8-12     	 3353834	       422 ns/op
//testing: BenchmarkSyncMapStore8-12 left GOMAXPROCS set to 8
//BenchmarkSyncMapStore16
//BenchmarkSyncMapStore16-12    	 2027538	       556 ns/op
//testing: BenchmarkSyncMapStore16-12 left GOMAXPROCS set to 16
//BenchmarkSyncMapStore32
//BenchmarkSyncMapStore32-12    	 1441564	       819 ns/op
//testing: BenchmarkSyncMapStore32-12 left GOMAXPROCS set to 32
//BenchmarkSyncMapStore64
//BenchmarkSyncMapStore64-12    	 1000000	      1491 ns/op
//testing: BenchmarkSyncMapStore64-12 left GOMAXPROCS set to 64
//PASS

