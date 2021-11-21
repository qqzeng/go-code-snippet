package bench

import (
	"github.com/qqzeng/sync-map-intro/maputil"
	"sync"
	"testing"
)

var globalResult int

func BenchmarkLoadRWMapFound(b *testing.B) {
	nums := nrand(b.N)
	rm := maputil.NewRawRWIntMap()
	for _, v := range nums {
		rm.Store(v, v)
	}

	currentResult := 0
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		currentResult, _ = rm.Load(nums[i])
	}

	globalResult = currentResult
}

func BenchmarkLoadRWMapNotFound(b *testing.B) {
	nums := nrand(b.N)
	rm := maputil.NewRawIntMap()
	for _, v := range nums {
		rm.Store(v, v)
	}
	currentResult := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		currentResult, _ = rm.Load(i)
	}
	globalResult = currentResult
}

func BenchmarkLoadMapFound(b *testing.B) {
	nums := nrand(b.N)
	m := maputil.NewRawIntMap()
	for _, v := range nums {
		m.Store(v, v)
	}

	currentResult := 0
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		currentResult, _ = m.Load(nums[i])
	}

	globalResult = currentResult
}

func BenchmarkLoadMapNotFound(b *testing.B) {
	nums := nrand(b.N)
	m := maputil.NewRawIntMap()
	for _, v := range nums {
		m.Store(v, v)
	}
	currentResult := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		currentResult, _ = m.Load(i)
	}
	globalResult = currentResult
}

func BenchmarkLoadSyncMapFound(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}
	currentResult := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, ok := sm.Load(nums[i])
		if ok {
			currentResult = r.(int)
		}
	}
	globalResult = currentResult
}

func BenchmarkLoadSyncMapNotFound(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}
	currentResult := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, ok := sm.Load(i)
		if ok {
			currentResult = r.(int)
		}
	}
	globalResult = currentResult
}

//BenchmarkLoadRWMapFound
//BenchmarkLoadRWMapFound-12         	14907756	       123 ns/op
//BenchmarkLoadRWMapNotFound
//BenchmarkLoadRWMapNotFound-12      	14490526	       135 ns/op
//BenchmarkLoadMapFound
//BenchmarkLoadMapFound-12           	15135213	       101 ns/op
//BenchmarkLoadMapNotFound
//BenchmarkLoadMapNotFound-12        	16142097	        94.8 ns/op
//BenchmarkLoadSyncMapFound
//BenchmarkLoadSyncMapFound-12       	 6341990	       206 ns/op
//BenchmarkLoadSyncMapNotFound
//BenchmarkLoadSyncMapNotFound-12    	 7330227	       215 ns/op
//PASS