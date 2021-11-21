package bench

import (
	"github.com/qqzeng/sync-map-intro/maputil"
	"runtime"
	"sync"
	"testing"
)

func benchmarkRWMapDelete(b *testing.B, numGroutine int) {
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

func benchmarkMapDelete(b *testing.B, numGroutine int) {
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

func benchmarkSyncMapDelete(b *testing.B, numGroutine int) {
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

func BenchmarkRWMapDelete1(b *testing.B) {
	benchmarkRWMapDelete(b, 1)
}

func BenchmarkRWMapDelete2(b *testing.B) {
	benchmarkRWMapDelete(b, 2)
}

func BenchmarkRWMapDelete4(b *testing.B) {
	benchmarkRWMapDelete(b, 4)
}

func BenchmarkRWMapDelete8(b *testing.B) {
	benchmarkRWMapDelete(b, 8)
}

func BenchmarkRWMapDelete16(b *testing.B) {
	benchmarkRWMapDelete(b, 16)
}

func BenchmarkRWMapDelete32(b *testing.B) {
	benchmarkRWMapDelete(b, 32)
}

func BenchmarkRWMapDelete64(b *testing.B) {
	benchmarkRWMapDelete(b, 64)
}

func BenchmarkMapDelete1(b *testing.B) {
	benchmarkMapDelete(b, 1)
}

func BenchmarkMapDelete2(b *testing.B) {
	benchmarkMapDelete(b, 2)
}

func BenchmarkMapDelete4(b *testing.B) {
	benchmarkMapDelete(b, 4)
}

func BenchmarkMapDelete8(b *testing.B) {
	benchmarkMapDelete(b, 8)
}

func BenchmarkMapDelete16(b *testing.B) {
	benchmarkMapDelete(b, 16)
}

func BenchmarkMapDelete32(b *testing.B) {
	benchmarkMapDelete(b, 32)
}

func BenchmarkMapDelete64(b *testing.B) {
	benchmarkMapDelete(b, 64)
}

func BenchmarkSyncMapDelete1(b *testing.B) {
	benchmarkSyncMapDelete(b, 1)
}

func BenchmarkSyncMapDelete2(b *testing.B) {
	benchmarkSyncMapDelete(b, 2)
}

func BenchmarkSyncMapDelete4(b *testing.B) {
	benchmarkSyncMapDelete(b, 4)
}

func BenchmarkSyncMapDelete8(b *testing.B) {
	benchmarkSyncMapDelete(b, 8)
}

func BenchmarkSyncMapDelete16(b *testing.B) {
	benchmarkSyncMapDelete(b, 16)
}

func BenchmarkSyncMapDelete32(b *testing.B) {
	benchmarkSyncMapDelete(b, 32)
}

func BenchmarkSyncMapDelete64(b *testing.B) {
	benchmarkSyncMapDelete(b, 64)
}

//BenchmarkRWMapDelete1
//BenchmarkRWMapDelete1-12       	 8589968	       157 ns/op
//testing: BenchmarkRWMapDelete1-12 left GOMAXPROCS set to 1
//BenchmarkRWMapDelete2
//BenchmarkRWMapDelete2-12       	 4066572	       320 ns/op
//testing: BenchmarkRWMapDelete2-12 left GOMAXPROCS set to 2
//BenchmarkRWMapDelete4
//BenchmarkRWMapDelete4-12       	 2158764	       610 ns/op
//testing: BenchmarkRWMapDelete4-12 left GOMAXPROCS set to 4
//BenchmarkRWMapDelete8
//BenchmarkRWMapDelete8-12       	 1000000	      1111 ns/op
//testing: BenchmarkRWMapDelete8-12 left GOMAXPROCS set to 8
//BenchmarkRWMapDelete16
//BenchmarkRWMapDelete16-12      	  754124	      2436 ns/op
//testing: BenchmarkRWMapDelete16-12 left GOMAXPROCS set to 16
//BenchmarkRWMapDelete32
//BenchmarkRWMapDelete32-12      	  390938	      4697 ns/op
//testing: BenchmarkRWMapDelete32-12 left GOMAXPROCS set to 32
//BenchmarkRWMapDelete64
//BenchmarkRWMapDelete64-12      	  187551	      8693 ns/op
//testing: BenchmarkRWMapDelete64-12 left GOMAXPROCS set to 64
//BenchmarkMapDelete1
//BenchmarkMapDelete1-12         	 8714604	       127 ns/op
//testing: BenchmarkMapDelete1-12 left GOMAXPROCS set to 1
//BenchmarkMapDelete2
//BenchmarkMapDelete2-12         	 4815802	       331 ns/op
//testing: BenchmarkMapDelete2-12 left GOMAXPROCS set to 2
//BenchmarkMapDelete4
//BenchmarkMapDelete4-12         	 2202747	       554 ns/op
//testing: BenchmarkMapDelete4-12 left GOMAXPROCS set to 4
//BenchmarkMapDelete8
//BenchmarkMapDelete8-12         	 1000000	      1145 ns/op
//testing: BenchmarkMapDelete8-12 left GOMAXPROCS set to 8
//BenchmarkMapDelete16
//BenchmarkMapDelete16-12        	  743526	      2582 ns/op
//testing: BenchmarkMapDelete16-12 left GOMAXPROCS set to 16
//BenchmarkMapDelete32
//BenchmarkMapDelete32-12        	  338246	      4446 ns/op
//testing: BenchmarkMapDelete32-12 left GOMAXPROCS set to 32
//BenchmarkMapDelete64
//BenchmarkMapDelete64-12        	  189139	      9141 ns/op
//testing: BenchmarkMapDelete64-12 left GOMAXPROCS set to 64
//BenchmarkSyncMapDelete1
//BenchmarkSyncMapDelete1-12     	 4049092	       353 ns/op
//testing: BenchmarkSyncMapDelete1-12 left GOMAXPROCS set to 1
//BenchmarkSyncMapDelete2
//BenchmarkSyncMapDelete2-12     	 4940329	       267 ns/op
//testing: BenchmarkSyncMapDelete2-12 left GOMAXPROCS set to 2
//BenchmarkSyncMapDelete4
//BenchmarkSyncMapDelete4-12     	 4361911	       311 ns/op
//testing: BenchmarkSyncMapDelete4-12 left GOMAXPROCS set to 4
//BenchmarkSyncMapDelete8
//BenchmarkSyncMapDelete8-12     	 3670501	       381 ns/op
//testing: BenchmarkSyncMapDelete8-12 left GOMAXPROCS set to 8
//BenchmarkSyncMapDelete16
//BenchmarkSyncMapDelete16-12    	 2569736	       479 ns/op
//testing: BenchmarkSyncMapDelete16-12 left GOMAXPROCS set to 16
//BenchmarkSyncMapDelete32
//BenchmarkSyncMapDelete32-12    	 1633868	       762 ns/op
//testing: BenchmarkSyncMapDelete32-12 left GOMAXPROCS set to 32
//BenchmarkSyncMapDelete64
//BenchmarkSyncMapDelete64-12    	 1000000	      1255 ns/op
//testing: BenchmarkSyncMapDelete64-12 left GOMAXPROCS set to 64
//PASS

