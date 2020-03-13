package crawler

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func Backup() {
	SetupRabbit()
	defer CleanupRabbit()
	file, err := os.Open("backup.csv")
	failOnError(err, "can not open backup.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		failOnError(err, "fail when read backup.cvs")
		publishToRabbit(record[0], "emailsvc-bing-search-result")
		fmt.Println(record[0])
	}
}
