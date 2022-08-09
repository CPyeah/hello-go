package rate_limiter

import (
	"fmt"
	"sync/atomic"
	"time"
)

type TokenBucketLimiter struct {
	metaData   LimiterMateData
	tokenCount int64
	tokenCap   int64
}

func NewTokenBucketLimiter(per time.Duration, id string, count int64) *TokenBucketLimiter {
	limiter := TokenBucketLimiter{
		tokenCap:   count,
		tokenCount: count,
		metaData: LimiterMateData{
			PerTime:     per,
			Id:          id,
			limiterType: TokenBucket,
		},
	}
	go putToken(&limiter)
	return &limiter
}

func (t *TokenBucketLimiter) Handle() bool {

	if t.tokenCount > 0 {
		atomic.AddInt64(&t.tokenCount, -1)
		return true
	} else {
		return false
	}
}

// 定时往桶里面添加token
func putToken(limiter *TokenBucketLimiter) {
	ticker := time.NewTicker(limiter.metaData.PerTime)
	for {
		select {
		case <-ticker.C:
			if limiter.tokenCount < limiter.tokenCap {
				current := atomic.AddInt64(&limiter.tokenCount, 1)
				fmt.Println("add a token, current count is", current)
			}
		}
	}

}
