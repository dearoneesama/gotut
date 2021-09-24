package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := map[string]int{}
	for _, word := range strings.Fields(s) {
		res[word] += 1
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
