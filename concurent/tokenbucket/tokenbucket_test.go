package tokenbucket

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	var v int64
	s := NewSimpleTokenBucket(90, 130)
	s.start()
	var wg sync.WaitGroup
	num := 300
	wg.Add(num)
	start := time.Now()
	for i := 0; i < num; i++ {
		go func(i int) {
			if s.GetToken() {
				fmt.Println(i, atomic.AddInt64(&v, 1), time.Since(start))
			}
			wg.Done()
		}(i)
		time.Sleep(time.Second / 90)
	}
	wg.Wait()
	fmt.Println(v)
}
