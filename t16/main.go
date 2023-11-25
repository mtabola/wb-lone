// Код основан на алгоритме, описанном в книге Томаса Н Кормана "Алгоритмы и структуры данных"
// В данном примере алгоритм работает рекурсивно и построен по философии "Divide and concure"
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func generateSlice(size int, minVal int, maxVal int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(int(math.Abs(float64(maxVal)))) - rand.Intn(int(math.Abs(float64(minVal))))
	}
	return slice
}

func partitionClassic(arr []int, min int, max int) int {
	x := arr[max]
	i := min - 1
	for j := min; j <= max-1; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[max] = arr[max], arr[i+1]
	return i + 1
}

func quicksortClassic(arr []int, min int, max int) {
	if min < max {
		mid := partitionClassic(arr, min, max)
		quicksortClassic(arr, min, mid-1)
		quicksortClassic(arr, mid+1, max)
	}
}

func main() {
	vals := generateSlice(20, -100, 100)
	fmt.Println("Before sorting: ", vals)
	quicksortClassic(vals, 0, len(vals)-1)

	fmt.Println("After sorting: ", vals)
}
