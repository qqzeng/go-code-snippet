package bench

import (
	"github.com/qqzeng/sync-map-intro/maputil"
	"math/rand"
	"sync"
	"testing"
)

func nrand(n int) []int {
	arr := make([]int, n)
	for ind := range arr {
		arr[ind] = rand.Int()
	}
	return arr
}

func BenchmarkStoreRWMap(b *testing.B) {
	nums := nrand(b.N)
	rwm := maputil.NewRawRWIntMap()
	b.ResetTimer()
	for _, v := range nums {
		rwm.Store(v, v)
	}
}


func BenchmarkStoreMap(b *testing.B) {
	nums := nrand(b.N)
	m := maputil.NewRawIntMap()
	b.ResetTimer()
	for _, v := range nums {
		m.Store(v, v)
	}
}

func BenchmarkStoreSyncMap(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	b.ResetTimer()
	for _, v := range nums {
		sm.Store(v, v)
	}
}

//BenchmarkStoreRWMap
//BenchmarkStoreRWMap-12      	 4940822	       260 ns/op
//BenchmarkStoreMap
//BenchmarkStoreMap-12        	 6019843	       201 ns/op
//BenchmarkStoreSyncMap
//BenchmarkStoreSyncMap-12    	 1777538	       625 ns/op
//PASS

