package bench

import (
	"github.com/qqzeng/sync-map-intro/maputil"
	"runtime"
	"sync"
	"testing"
)

var resultHolder chan int

func benchmarkRWMapStableKeys(b *testing.B, numGroutine int) {
	runtime.GOMAXPROCS(numGroutine)

	rwm := maputil.NewRawRWIntMap()
	nums := nrand(b.N)
	for _, v := range nums {
		rwm.Store(v, v)
	}

	var wg sync.WaitGroup
	wg.Add(numGroutine)

	// Holds our final results, to prevent compiler optimizations.
	resultHolder = make(chan int, numGroutine)

	b.ResetTimer()

	for i := 0; i < numGroutine; i++ {
		go func(n int) {
			result := 0
			for j := 0; j < n; j++ {
				result, _ = rwm.Load(5)
			}
			resultHolder <- result
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func benchmarkMapStableKeys(b *testing.B, numGroutine int) {
	runtime.GOMAXPROCS(numGroutine)

	m := maputil.NewRawIntMap()
	nums := nrand(b.N)
	for _, v := range nums {
		m.Store(v, v)
	}

	var wg sync.WaitGroup
	wg.Add(numGroutine)

	// Holds our final results, to prevent compiler optimizations.
	resultHolder = make(chan int, numGroutine)

	b.ResetTimer()

	for i := 0; i < numGroutine; i++ {
		go func(n int) {
			result := 0
			for j := 0; j < n; j++ {
				result, _ = m.Load(5)
			}
			resultHolder <- result
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func benchmarkSyncMapStableKeys(b *testing.B, numGroutine int) {
	runtime.GOMAXPROCS(numGroutine)

	var sm sync.Map
	nums := nrand(b.N)
	for _, v := range nums {
		sm.Store(v, v)
	}

	var wg sync.WaitGroup
	wg.Add(numGroutine)

	// Holds our final results, to prevent compiler optimizations.
	resultHolder = make(chan int, numGroutine)

	b.ResetTimer()

	for i := 0; i < numGroutine; i++ {
		go func(n int) {
			result := 0
			for j := 0; j < n; j++ {
				r, ok := sm.Load(5)
				if ok {
					result = r.(int)
				}
			}
			resultHolder <- result
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func benchmarkRWMapStableKeysFound(b *testing.B, numGroutine int) {
	runtime.GOMAXPROCS(numGroutine)

	rwm := maputil.NewRawRWIntMap()
	nums := nrand(b.N)
	for _, v := range nums {
		rwm.Store(v, v)
	}

	var wg sync.WaitGroup
	wg.Add(numGroutine)

	// Holds our final results, to prevent compiler optimizations.
	resultHolder = make(chan int, numGroutine)

	b.ResetTimer()

	for i := 0; i < numGroutine; i++ {
		go func(n int) {
			result := 0
			for j := 0; j < n; j++ {
				result, _ = rwm.Load(nums[j])
			}
			resultHolder <- result
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func benchmarkMapStableKeysFound(b *testing.B, numGroutine int) {
	runtime.GOMAXPROCS(numGroutine)

	m := maputil.NewRawIntMap()
	nums := nrand(b.N)
	for _, v := range nums {
		m.Store(v, v)
	}

	var wg sync.WaitGroup
	wg.Add(numGroutine)

	// Holds our final results, to prevent compiler optimizations.
	resultHolder = make(chan int, numGroutine)

	b.ResetTimer()

	for i := 0; i < numGroutine; i++ {
		go func(n int) {
			result := 0
			for j := 0; j < n; j++ {
				result, _ = m.Load(nums[j])
			}
			resultHolder <- result
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func benchmarkSyncMapStableKeysFound(b *testing.B, numGroutine int) {
	runtime.GOMAXPROCS(numGroutine)

	var sm sync.Map
	nums := nrand(b.N)
	for _, v := range nums {
		sm.Store(v, v)
	}

	var wg sync.WaitGroup
	wg.Add(numGroutine)

	// Holds our final results, to prevent compiler optimizations.
	resultHolder = make(chan int, numGroutine)

	b.ResetTimer()

	for i := 0; i < numGroutine; i++ {
		go func(n int) {
			result := 0
			for j := 0; j < n; j++ {
				r, ok := sm.Load(nums[j])
				if ok {
					result = r.(int)
				}
			}
			resultHolder <- result
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func BenchmarkRWMapStableKeys1(b *testing.B) {
	benchmarkRWMapStableKeys(b, 1)
}

func BenchmarkRWMapStableKeys2(b *testing.B) {
	benchmarkRWMapStableKeys(b, 2)
}

func BenchmarkRWMapStableKeys4(b *testing.B) {
	benchmarkRWMapStableKeys(b, 4)
}

func BenchmarkRWMapStableKeys8(b *testing.B) {
	benchmarkRWMapStableKeys(b, 8)
}

func BenchmarkRWMapStableKeys16(b *testing.B) {
	benchmarkRWMapStableKeys(b, 16)
}

func BenchmarkRWMapStableKeys32(b *testing.B) {
	benchmarkRWMapStableKeys(b, 32)
}

func BenchmarkMapStableKeys1(b *testing.B) {
	benchmarkMapStableKeys(b, 1)
}

func BenchmarkMapStableKeys2(b *testing.B) {
	benchmarkMapStableKeys(b, 2)
}

func BenchmarkMapStableKeys4(b *testing.B) {
	benchmarkMapStableKeys(b, 4)
}

func BenchmarkMapStableKeys8(b *testing.B) {
	benchmarkMapStableKeys(b, 8)
}

func BenchmarkMapStableKeys16(b *testing.B) {
	benchmarkMapStableKeys(b, 16)
}

func BenchmarkMapStableKeys32(b *testing.B) {
	benchmarkMapStableKeys(b, 32)
}


func BenchmarkSyncMapStableKeys1(b *testing.B) {
	benchmarkSyncMapStableKeys(b, 1)
}

func BenchmarkSyncMapStableKeys2(b *testing.B) {
	benchmarkSyncMapStableKeys(b, 2)
}

func BenchmarkSyncMapStableKeys4(b *testing.B) {
	benchmarkSyncMapStableKeys(b, 4)
}

func BenchmarkSyncMapStableKeys8(b *testing.B) {
	benchmarkSyncMapStableKeys(b, 8)
}

func BenchmarkSyncMapStableKeys16(b *testing.B) {
	benchmarkSyncMapStableKeys(b, 16)
}

func BenchmarkSyncMapStableKeys32(b *testing.B) {
	benchmarkSyncMapStableKeys(b, 32)
}

//BenchmarkRWMapStableKeys1
//BenchmarkRWMapStableKeys1-12       	55306490	        23.9 ns/op
//testing: BenchmarkRWMapStableKeys1-12 left GOMAXPROCS set to 1
//BenchmarkRWMapStableKeys2
//BenchmarkRWMapStableKeys2-12       	 9586473	       117 ns/op
//testing: BenchmarkRWMapStableKeys2-12 left GOMAXPROCS set to 2
//BenchmarkRWMapStableKeys4
//BenchmarkRWMapStableKeys4-12       	 5369010	       217 ns/op
//testing: BenchmarkRWMapStableKeys4-12 left GOMAXPROCS set to 4
//BenchmarkRWMapStableKeys8
//BenchmarkRWMapStableKeys8-12       	 2808339	       428 ns/op
//testing: BenchmarkRWMapStableKeys8-12 left GOMAXPROCS set to 8
//BenchmarkRWMapStableKeys16
//BenchmarkRWMapStableKeys16-12      	 1501873	       801 ns/op
//testing: BenchmarkRWMapStableKeys16-12 left GOMAXPROCS set to 16
//BenchmarkRWMapStableKeys32
//BenchmarkRWMapStableKeys32-12      	  735696	      1587 ns/op
//testing: BenchmarkRWMapStableKeys32-12 left GOMAXPROCS set to 32
//BenchmarkMapStableKeys1
//BenchmarkMapStableKeys1-12         	54971413	        22.5 ns/op
//testing: BenchmarkMapStableKeys1-12 left GOMAXPROCS set to 1
//BenchmarkMapStableKeys2
//BenchmarkMapStableKeys2-12         	 9369276	       118 ns/op
//testing: BenchmarkMapStableKeys2-12 left GOMAXPROCS set to 2
//BenchmarkMapStableKeys4
//BenchmarkMapStableKeys4-12         	 5327538	       217 ns/op
//testing: BenchmarkMapStableKeys4-12 left GOMAXPROCS set to 4
//BenchmarkMapStableKeys8
//BenchmarkMapStableKeys8-12         	 2780664	       440 ns/op
//testing: BenchmarkMapStableKeys8-12 left GOMAXPROCS set to 8
//BenchmarkMapStableKeys16
//BenchmarkMapStableKeys16-12        	 1460196	       833 ns/op
//testing: BenchmarkMapStableKeys16-12 left GOMAXPROCS set to 16
//BenchmarkMapStableKeys32
//BenchmarkMapStableKeys32-12        	  687639	      1657 ns/op
//testing: BenchmarkMapStableKeys32-12 left GOMAXPROCS set to 32
//BenchmarkSyncMapStableKeys1
//BenchmarkSyncMapStableKeys1-12     	16321682	        75.9 ns/op
//testing: BenchmarkSyncMapStableKeys1-12 left GOMAXPROCS set to 1
//BenchmarkSyncMapStableKeys2
//BenchmarkSyncMapStableKeys2-12     	16944424	        62.4 ns/op
//testing: BenchmarkSyncMapStableKeys2-12 left GOMAXPROCS set to 2
//BenchmarkSyncMapStableKeys4
//BenchmarkSyncMapStableKeys4-12     	12722924	        92.4 ns/op
//testing: BenchmarkSyncMapStableKeys4-12 left GOMAXPROCS set to 4
//BenchmarkSyncMapStableKeys8
//BenchmarkSyncMapStableKeys8-12     	 9469438	       132 ns/op
//testing: BenchmarkSyncMapStableKeys8-12 left GOMAXPROCS set to 8
//BenchmarkSyncMapStableKeys16
//BenchmarkSyncMapStableKeys16-12    	 6746767	       180 ns/op
//testing: BenchmarkSyncMapStableKeys16-12 left GOMAXPROCS set to 16
//BenchmarkSyncMapStableKeys32
//BenchmarkSyncMapStableKeys32-12    	 4980276	       263 ns/op
//testing: BenchmarkSyncMapStableKeys32-12 left GOMAXPROCS set to 32
//PASS

//

//func BenchmarkRWMapStableKeysFound1(b *testing.B) {
//	benchmarkRWMapStableKeysFound(b, 1)
//}
//
//func BenchmarkRWMapStableKeysFound2(b *testing.B) {
//	benchmarkRWMapStableKeysFound(b, 2)
//}
//
//func BenchmarkRWMapStableKeysFound4(b *testing.B) {
//	benchmarkRWMapStableKeysFound(b, 4)
//}
//
//func BenchmarkRWMapStableKeysFound8(b *testing.B) {
//	benchmarkRWMapStableKeysFound(b, 8)
//}
//
//func BenchmarkRWMapStableKeysFound16(b *testing.B) {
//	benchmarkRWMapStableKeysFound(b, 16)
//}
//
//func BenchmarkRWMapStableKeysFound32(b *testing.B) {
//	benchmarkRWMapStableKeysFound(b, 32)
//}
//
//func BenchmarkRWMapStableKeysFound64(b *testing.B) {
//	benchmarkRWMapStableKeysFound(b, 64)
//}
//
//func BenchmarkMapStableKeysFound1(b *testing.B) {
//	benchmarkMapStableKeysFound(b, 1)
//}
//
//func BenchmarkMapStableKeysFound2(b *testing.B) {
//	benchmarkMapStableKeysFound(b, 2)
//}
//
//func BenchmarkMapStableKeysFound4(b *testing.B) {
//	benchmarkMapStableKeysFound(b, 4)
//}
//
//func BenchmarkMapStableKeysFound8(b *testing.B) {
//	benchmarkMapStableKeysFound(b, 8)
//}
//
//func BenchmarkMapStableKeysFound16(b *testing.B) {
//	benchmarkMapStableKeysFound(b, 16)
//}
//
//func BenchmarkMapStableKeysFound32(b *testing.B) {
//	benchmarkMapStableKeysFound(b, 32)
//}
//
//func BenchmarkMapStableKeysFound64(b *testing.B) {
//	benchmarkMapStableKeysFound(b, 64)
//}
//
//func BenchmarkSyncMapStableKeysFound1(b *testing.B) {
//	benchmarkSyncMapStableKeysFound(b, 1)
//}
//
//func BenchmarkSyncMapStableKeysFound2(b *testing.B) {
//	benchmarkSyncMapStableKeysFound(b, 2)
//}
//
//func BenchmarkSyncMapStableKeysFound4(b *testing.B) {
//	benchmarkSyncMapStableKeysFound(b, 4)
//}
//
//func BenchmarkSyncMapStableKeysFound8(b *testing.B) {
//	benchmarkSyncMapStableKeysFound(b, 8)
//}
//
//func BenchmarkSyncMapStableKeysFound16(b *testing.B) {
//	benchmarkSyncMapStableKeysFound(b, 16)
//}
//
//func BenchmarkSyncMapStableKeysFound32(b *testing.B) {
//	benchmarkSyncMapStableKeysFound(b, 32)
//}
//
//func BenchmarkSyncMapStableKeysFound64(b *testing.B) {
//	benchmarkSyncMapStableKeysFound(b, 64)
//}

//BenchmarkRWMapStableKeysFound1
//BenchmarkRWMapStableKeysFound1-12       	14810488	       119 ns/op
//testing: BenchmarkRWMapStableKeysFound1-12 left GOMAXPROCS set to 1
//BenchmarkRWMapStableKeysFound2
//BenchmarkRWMapStableKeysFound2-12       	 7763245	       163 ns/op
//testing: BenchmarkRWMapStableKeysFound2-12 left GOMAXPROCS set to 2
//BenchmarkRWMapStableKeysFound4
//BenchmarkRWMapStableKeysFound4-12       	 4780857	       274 ns/op
//testing: BenchmarkRWMapStableKeysFound4-12 left GOMAXPROCS set to 4
//BenchmarkRWMapStableKeysFound8
//BenchmarkRWMapStableKeysFound8-12       	 2630872	       481 ns/op
//testing: BenchmarkRWMapStableKeysFound8-12 left GOMAXPROCS set to 8
//BenchmarkRWMapStableKeysFound16
//BenchmarkRWMapStableKeysFound16-12      	 1472761	       823 ns/op
//testing: BenchmarkRWMapStableKeysFound16-12 left GOMAXPROCS set to 16
//BenchmarkRWMapStableKeysFound32
//BenchmarkRWMapStableKeysFound32-12      	  707697	      1624 ns/op
//testing: BenchmarkRWMapStableKeysFound32-12 left GOMAXPROCS set to 32
//BenchmarkRWMapStableKeysFound64
//BenchmarkRWMapStableKeysFound64-12      	  367813	      3262 ns/op
//testing: BenchmarkRWMapStableKeysFound64-12 left GOMAXPROCS set to 64
//BenchmarkMapStableKeysFound1
//BenchmarkMapStableKeysFound1-12         	14872831	       113 ns/op
//testing: BenchmarkMapStableKeysFound1-12 left GOMAXPROCS set to 1
//BenchmarkMapStableKeysFound2
//BenchmarkMapStableKeysFound2-12         	 8059984	       159 ns/op
//testing: BenchmarkMapStableKeysFound2-12 left GOMAXPROCS set to 2
//BenchmarkMapStableKeysFound4
//BenchmarkMapStableKeysFound4-12         	 4688964	       265 ns/op
//testing: BenchmarkMapStableKeysFound4-12 left GOMAXPROCS set to 4
//BenchmarkMapStableKeysFound8
//BenchmarkMapStableKeysFound8-12         	 2608935	       478 ns/op
//testing: BenchmarkMapStableKeysFound8-12 left GOMAXPROCS set to 8
//BenchmarkMapStableKeysFound16
//BenchmarkMapStableKeysFound16-12        	 1439472	       843 ns/op
//testing: BenchmarkMapStableKeysFound16-12 left GOMAXPROCS set to 16
//BenchmarkMapStableKeysFound32
//BenchmarkMapStableKeysFound32-12        	  731728	      1662 ns/op
//testing: BenchmarkMapStableKeysFound32-12 left GOMAXPROCS set to 32
//BenchmarkMapStableKeysFound64
//BenchmarkMapStableKeysFound64-12        	  366547	      3336 ns/op
//testing: BenchmarkMapStableKeysFound64-12 left GOMAXPROCS set to 64
//BenchmarkSyncMapStableKeysFound1
//BenchmarkSyncMapStableKeysFound1-12     	 4635644	       264 ns/op
//testing: BenchmarkSyncMapStableKeysFound1-12 left GOMAXPROCS set to 1
//BenchmarkSyncMapStableKeysFound2
//BenchmarkSyncMapStableKeysFound2-12     	 4420430	       269 ns/op
//testing: BenchmarkSyncMapStableKeysFound2-12 left GOMAXPROCS set to 2
//BenchmarkSyncMapStableKeysFound4
//BenchmarkSyncMapStableKeysFound4-12     	 4502130	       325 ns/op
//testing: BenchmarkSyncMapStableKeysFound4-12 left GOMAXPROCS set to 4
//BenchmarkSyncMapStableKeysFound8
//BenchmarkSyncMapStableKeysFound8-12     	 4123311	       338 ns/op
//testing: BenchmarkSyncMapStableKeysFound8-12 left GOMAXPROCS set to 8
//BenchmarkSyncMapStableKeysFound16
//BenchmarkSyncMapStableKeysFound16-12    	 2282990	       505 ns/op
//testing: BenchmarkSyncMapStableKeysFound16-12 left GOMAXPROCS set to 16
//BenchmarkSyncMapStableKeysFound32
//BenchmarkSyncMapStableKeysFound32-12    	 1513104	       821 ns/op
//testing: BenchmarkSyncMapStableKeysFound32-12 left GOMAXPROCS set to 32
//BenchmarkSyncMapStableKeysFound64
//BenchmarkSyncMapStableKeysFound64-12    	 1000000	      1377 ns/op
//testing: BenchmarkSyncMapStableKeysFound64-12 left GOMAXPROCS set to 64
//PASS


