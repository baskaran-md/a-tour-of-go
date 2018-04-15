package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount - Returns map of counts of each word in a given string
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, word := range strings.Fields(s) {
		result[word]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
