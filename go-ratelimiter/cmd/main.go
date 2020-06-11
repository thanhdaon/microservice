package main

import (
	"fmt"
	"log"
	"math/rand"
	"ratelimiter/ratelimit"
	"sync"
	"time"
)

func main() {
	throttleRateLimiterTest()
}

func throttleRateLimiterTest() {
	r, err := ratelimit.NewThrottleRateLimiter(&ratelimit.Config{
		Throttle: 1 * time.Second,
	})

	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	doWork := func(id int) {
		token, err := r.Acquire()
		fmt.Printf("Rate Limit Token %s acquired at %s\n", token.ID, time.Now().UTC())
		if err != nil {
			panic(err)
		}
		n := rand.Intn(5)
		fmt.Printf("Worker %d Sleeping %d seconds...\n", id, n)
		time.Sleep(time.Duration(n) * time.Second)
		fmt.Printf("Worker %d Done\n", id)
		wg.Done()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doWork(i)
	}

	wg.Wait()
}
