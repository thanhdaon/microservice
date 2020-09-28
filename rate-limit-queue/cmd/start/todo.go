package start

import (
	"fmt"
	"redis-rate-limit-queue/ratelimit"
)

func Todo(r *ratelimit.RateLimit) {
	fmt.Println()
	// p, txid, err := r.Dequeue(ARENA)
	// failOnError(err)

	// log.Println(p, txid)

	// err = r.Commit(txid)
	// failOnError(err)
}
