package tokenbucket

import (
	"time"
)

type simpleTokenBucket struct {
	ch  chan struct{}
	qps int64
}

var obj struct{}

func NewSimpleTokenBucket(qps int64, buff int64) (s *simpleTokenBucket) {
	s = &simpleTokenBucket{
		qps: qps,
		ch:  make(chan struct{}, buff),
	}
	s.start()
	return s
}

func (s *simpleTokenBucket) start() {
	go func() {
		count := int64(1)
		if s.qps > tcount {
			count = s.qps / freq
		}
		t := time.NewTicker(time.Second * time.Duration(count) / time.Duration(s.qps))
		for range t.C {
			for i := int64(0); i < count; i++ {
				s.ch <- obj
			}
		}
	}()
}

func (s *simpleTokenBucket) GetToken() (ok bool) {
	select {
	case <-s.ch:
		return true
	default:
		return false
	}
}
