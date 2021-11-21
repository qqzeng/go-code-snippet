package bench

import (
	"github.com/qqzeng/sync-map-intro/maputil"
	"sync"
	"testing"
)

func BenchmarkDeleteRWMap(b *testing.B) {
	nums := nrand(b.N)
	rwm := maputil.NewRawRWIntMap()
	for _, v := range nums {
		rwm.Store(v, v)
	}

	b.ResetTimer()
	for _, v := range nums {
		rwm.Delete(v)
	}
}

func BenchmarkDeleteMap(b *testing.B) {
	nums := nrand(b.N)
	m := maputil.NewRawIntMap()
	for _, v := range nums {
		m.Store(v, v)
	}

	b.ResetTimer()
	for _, v := range nums {
		m.Delete(v)
	}
}

func BenchmarkDeleteSyncMap(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}

	b.ResetTimer()
	for _, v := range nums {
		sm.Delete(v)
	}
}

//BenchmarkDeleteRWMap
//BenchmarkDeleteRWMap-12      	 9221569	       155 ns/op
//BenchmarkDeleteMap
//BenchmarkDeleteMap-12        	 7741298	       150 ns/op
//BenchmarkDeleteSyncMap
//BenchmarkDeleteSyncMap-12    	 6175486	       222 ns/op
//PASS