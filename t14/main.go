package main

import "fmt"

func typeCheck(param interface{}) {
	//У каждого интерфейса под копотом есть такое поле, как type
	//Когда мы вносим туда данные, то он сразу же автоматом определяет тип, который туда внесен
	switch v := param.(type) {
	default:
		fmt.Printf("Unexpected type %T\n", v)
	case uint64, string, chan interface{}, bool, int:
		fmt.Printf("Valid type: %T - %v\n", v, v)
	}

}

func main() {
	arr := make([]interface{}, 0)

	arr = append(arr, uint64(21))
	arr = append(arr, "string")
	arr = append(arr, 23)
	arr = append(arr, true)
	arr = append(arr, make(chan interface{}))
	arr = append(arr, make(chan int))
	arr = append(arr, uint32(20))

	for _, v := range arr {
		typeCheck(v)
	}
}
