package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	array := strings.Fields(s)
	for i := 0; i < len(array); i++ {
		t := array[i]
		value := m[t]
		if value == 0 {
			m[t] = 1
		} else {
			m[t] = value + 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
