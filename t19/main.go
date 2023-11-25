package main

import "fmt"

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func main() {
	var in string
	fmt.Print("Please enter your string: ")
	fmt.Scan(&in)

	fmt.Println("Output: ", reverse(in))
}
