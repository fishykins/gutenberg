package main

import "github.com/fishykins/gutenberg/internal/webbers"

func main() {
	_, err := webbers.ParseDictionary("")
	if err != nil {
		panic(err)
	}
}
