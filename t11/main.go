package main

import "fmt"

func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var res []int
	//Проходимся по первому массиву и заносим в мапу сам элемент (по ключу) и сколько раз оно встречается
	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem]++
		}
	}
	//Проходим по второму массиву и проверяем входят ли элементы в мапу
	for _, elem := range b {
		//если элемент есть в мапе, то уменьшаем счетчик и запиисывам в результатирующий массив
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem]--
			res = append(res, elem)
		}
		//иначе, пропускаем
	}
	return res
}

func main() {
	a := []int{23, 3, 1, 2, 1}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}
