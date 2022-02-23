package main

import (
	"fmt"
	"unicode"
)

func HanCounter(s string) int {
	var count int //0
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	return count
}
func main() {
	s := "hello沙河小王子"
	count := HanCounter(s)
	fmt.Println(count)
}
