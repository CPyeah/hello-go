package rate_limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTokenBucketLimiter(t *testing.T) {
	limiter := NewTokenBucketLimiter(time.Second, "tokenBucketLimiter1", 5)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 200)
		result := limiter.Handle()
		fmt.Println(i, "th request, result is", result)
	}
}
