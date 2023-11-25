package main

import "fmt"

func binarySearchLoop(arr []int, searchElem int, begin int, end int) int {
	for begin <= end {
		mid := begin + (end-begin)/2
		if arr[mid] == searchElem {
			return mid
		} else if arr[mid] < searchElem {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func binarySearchRecursive(arr []int, searchElem int, begin int, end int) int {
	if begin <= end {
		mid := begin + (end-begin)/2
		if arr[mid] == searchElem {
			return mid
		} else if arr[mid] > searchElem {
			return binarySearchRecursive(arr, searchElem, begin, mid-1)
		} else {
			return binarySearchRecursive(arr, searchElem, mid+1, end)
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 12, 16, 92, 101}
	fmt.Println(binarySearchRecursive(arr, 10, 0, len(arr)-1))
	fmt.Println(binarySearchLoop(arr, 16, 0, len(arr)-1))
}
