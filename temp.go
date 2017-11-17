package main

import (
	"github.com/peterbourgon/diskv"
	"bufio"
    "log"
	"strings"
    "os"
)

func main() {
	// Simplest transform function: put all the data files into the base dir.
	flatTransform := func(s string) []string { return []string{} }

	// Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
	d := diskv.New(diskv.Options{
		BasePath:     "my-data-dir",
		Transform:    flatTransform,
		CacheSizeMax: 100 * 1024 * 1024,
	})

	file, err := os.Open("data/only_tamil_uniq_sorted_words.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		d.Write(strings.TrimSpace(scanner.Text()), []byte{'1'})
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
