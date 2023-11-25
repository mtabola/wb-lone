package main

import (
	"fmt"
	"t24/point"
)

func main() {
	p1 := point.New(10, -4)
	p2 := point.New(-20, -10)

	fmt.Println(p1.DistanceCalculate(p2))
	fmt.Println(p2.DistanceCalculate(p1))
}
