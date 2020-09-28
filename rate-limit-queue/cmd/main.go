package main

import (
	"log"
	"os"
	"redis-rate-limit-queue/cmd/start"
	"redis-rate-limit-queue/ratelimit"
)

func main() {
	r, err := ratelimit.New("localhost:6379")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	defer r.Close()

	switch os.Args[1] {
	case "prepare":
		start.Prepare(r)
	case "todo":
		start.Todo(r)
	case "consume":
		start.Consume(r)
	default:
		log.Fatal("Invalid arg!")
	}
}
