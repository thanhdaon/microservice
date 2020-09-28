package start

import (
	"fmt"
	"redis-rate-limit-queue/ratelimit"
)

func Prepare(r *ratelimit.RateLimit) {
	err := r.Prune()
	failOnError(err)

	err = r.DefineRateLimits(ARENA, []ratelimit.Limit{
		{ID: TENANT_ID_1, Limit: 10},
		{ID: TENANT_ID_2, Limit: 15},
	})
	failOnError(err)

	for i := 0; i < 100; i++ {
		err = r.Enqueue(ARENA, TENANT_ID_1, fmt.Sprintf("okok_%d", i))
		failOnError(err)
	}
}
