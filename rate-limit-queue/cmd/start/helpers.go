package start

import "log"

func failOnError(err error) {
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
}
