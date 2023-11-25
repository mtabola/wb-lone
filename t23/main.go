package main

import "fmt"

// достигает своей скорости путем использования уже выделенной памяти, однако меняет порядок
func fastVersion(a []string, pos int) []string {
	if len(a) < pos {
		return nil
	}
	a[pos] = a[len(a)-1]
	a[len(a)-1] = "" // записываем нулевое значение для того, чтоб дальше эту область памяти можно было использовать
	a = a[:len(a)-1]
	return a
}

// медленен из-за того, что внутри функции происходит копирование массива от pos + 1
func slowVersion(a []string, pos int) []string {
	if len(a) < pos {
		return nil
	}
	copy(a[pos:], a[pos+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}

func main() {
	arr := []string{"dog", "cat", "mouse", "rat", "bear", "fox"}
	fmt.Println(arr)
	arr = fastVersion(arr, 2)
	fmt.Println("Fast version", arr)

	arr = slowVersion(arr, 1)
	fmt.Println("Slowed versoion: ", arr)
}
