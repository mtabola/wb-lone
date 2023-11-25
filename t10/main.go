package main

import (
	"fmt"
	"math"
)

func main() {
	temperature := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0, -0}
	tempGroups := make(map[int][]float64)
	for _, v := range temperature {
		groupInd := (int(math.Floor(v)) / 10) * 10
		tempGroups[groupInd] = append(tempGroups[groupInd], v)
	}

	for k, v := range tempGroups {
		fmt.Printf("Group %d: %.1f\n", k, v)
	}
}
