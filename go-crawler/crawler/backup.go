package crawler

import (
	"bufio"
	"log"
	"os"
)

func Backup() {
	SetupRabbit()
	defer CleanupRabbit()
	file, err := os.Open("static/haymora.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		publishToRabbit(scanner.Text(), "emailsvc-domain-waiting-to-crawl")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
