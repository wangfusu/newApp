package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimiterIface interface {
	Key(c *gin.Context) string                          //Key：获取对应的限流器的键值对名称。
	GetBucket(key string) (*ratelimit.Bucket, bool)     //GetBucket：获取令牌桶。
	AddBuckets(rules ...LimiterBucketRule) LimiterIface //AddBuckets：新增多个令牌桶。
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string        //Key：自定义键值对名称。
	FillInterval time.Duration //FillInterval：间隔多久时间放 N 个令牌。
	Capacity     int64         //Capacity：令牌桶的容量。
	Quantum      int64         //Quantum：每次到达间隔时间后所放的具体令牌数量。
}
