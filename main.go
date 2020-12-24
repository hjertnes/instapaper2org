package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		panic("You have to specify a filename")
	}

	filename := filepath.Clean(args[1])

	f, _ := os.Open(filename)

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range records {
		url := line[0]
		title := line[1]
		folder := line[3]

		if title == "" {
			title = url
		}

		if folder != "Starred" {
			continue
		}

		fmt.Printf("- [%s](%s)\n", title, url)
	}
}
