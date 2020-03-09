package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	sStr := strings.Fields(s)
	wordCount := map[string]int{}
	for _, v := range sStr {
		wordCount[v]++
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
