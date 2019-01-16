package tokenbucket

import (
	"sync/atomic"
	"time"
)

type simpleTokenBucketAtomic struct {
	cur  int64
	qps  int64
	buff int64
}

func NewSimpleTokenBucketAtomic(qps int64, buff int64) (s *simpleTokenBucketAtomic) {
	s = &simpleTokenBucketAtomic{qps: qps, buff: buff}
	s.start()
	return s
}

func (s *simpleTokenBucketAtomic) start() {
	go func() {
		count := int64(1)
		if s.qps > tcount {
			count = s.qps / freq
		}
		t := time.NewTicker(time.Second * time.Duration(count) / time.Duration(s.qps))
		for range t.C {
			v := atomic.LoadInt64(&s.cur)
			if v < s.buff {
				atomic.AddInt64(&s.cur, count)
			}
		}
	}()
}

func (s *simpleTokenBucketAtomic) GetToken() (ok bool) {
	v := atomic.AddInt64(&s.cur, -1)
	if v >= 0 {
		return true
	}
	atomic.AddInt64(&s.cur, 1)
	return false
}
