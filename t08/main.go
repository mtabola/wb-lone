package main

import "fmt"

// Для того, чтобы нам устанавливать n-ый бит в 0 или 1, нам нужна битовая маска
func maskCreating(k uint8) int64 {
	var mask int64 = 1
	shift := 64 - k
	//создаем маску путем смещения битов на 64-k позиций
	return mask << shift
}

func main() {
	var val int64
	var pos uint8
	fmt.Print("Please enter a number: ")
	fmt.Scan(&val)
	fmt.Printf("Number now: %b - %d\n", val, val)
	//проверка на валидность введенного числа (макс 64 бита)
	fmt.Print("Please enter edited position (from 1 to 64): ")
	fmt.Scan(&pos)
	if pos < 1 || pos > 64 {
		fmt.Println("Position must be between 1 and 64")
		return
	}
	//создаем маску
	mask := maskCreating(pos)
	//с помощью побитовго xor ставим n бит в противоположное положение
	val = val ^ mask
	fmt.Printf("Result is: %b - %d\n", val, val)
}
