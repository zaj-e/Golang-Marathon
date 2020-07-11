package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	divided := strings.Fields(s)
	m := make(map[string]int)
	for _, w := range divided {
		_, ok := m[w]
		if ok == false {
			m[w] = 1
		} else {
			m[w] = m[w] +1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
