package main

import "github.com/fishykins/gutenberg/internal/webbers"

func main() {
	_, err := webbers.ParseDictionary("trevor")
	if err != nil {
		panic(err)
	}
}
