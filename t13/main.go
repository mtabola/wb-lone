package main

import "fmt"

func pointersSwap(a, b *int8) (int8, int8) {
	return *b, *a
}

func arifmeticSwap(a, b int8) (int8, int8) {
	a += b
	b = a - b
	a -= b
	return a, b
}

// Об этом способе узнал из источника Tproger (https://tproger.ru/problems/popular-ways-to-swap)
func xorSwap(a, b int8) (int8, int8) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

func main() {
	var firv, secv int8
	fmt.Print("Enter the first number: ")
	fmt.Scan(&firv)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&secv)

	//Здесь показано 4 способа смены переменной
	//firv, secv = secv, firv // Самый простой вариант - с помощью встроенных возможностей языка
	//firv, secv = pointerSwap(&firv, &secv) //с помощью указателей
	//firv, secv = arifmeticSwap(firv, secv) //с помощью арифметики
	//firv, secv = xorSwap(firv, secv) // с помощью бинарной логики

	fmt.Printf("First method (through the function): firv = %d, secv = %d\n", firv, secv)

}
