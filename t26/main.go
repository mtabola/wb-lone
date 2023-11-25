package main

import (
	"fmt"
	"strings"
)

func uniqueString(s string) bool {
	chars := make(map[rune]bool)

	for _, v := range s {
		_, ok := chars[rune(v)]
		if ok {
			return false
		}
		chars[v] = true
	}
	return true
}

func main() {
	var in string
	fmt.Print("Please enter your string: ")
	fmt.Scan(&in)
	in = strings.ToLower(in)

	fmt.Println(uniqueString(in))
}
