package rate_limiter

import "time"

type LimiterMateData struct {
	PerTime     time.Duration
	Id          string
	limiterType string
}

const (
	TokenBucket = "Token bucket"
	//LeakingBucket = "Leaking bucket"
)
