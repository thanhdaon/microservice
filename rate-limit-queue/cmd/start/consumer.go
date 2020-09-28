package start

import (
	"log"
	"redis-rate-limit-queue/ratelimit"
	"time"
)

func Consume(r *ratelimit.RateLimit) {
	delay := 0

	for {
		p, txid, err := r.Dequeue(ARENA)
		if err != nil {
			if err == ratelimit.ERR_QUEUE_EMPTY {
				time.Sleep(100 * time.Millisecond)
				delay += 100
				continue
			}
			log.Fatalf("[ERROR] %v", err)
		}
		log.Printf("[INFO] delay:%d | p:%s", delay, p)

		if err := r.Commit(txid); err != nil {
			log.Fatalf("[ERROR] %v", err)
		}
		delay = 0
	}
}
