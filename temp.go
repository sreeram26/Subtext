package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
	"bufio"
    "log"
	"strings"
    "os"
)

func main() {

	c := cache.New(5*time.Hour, 10*time.Hour)

	file, err := os.Open("data/only_tamil_uniq_sorted_words.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		c.Set(strings.TrimSpace(scanner.Text()), true, cache.NoExpiration)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	fmt.Println(len(c.Items()))
}
