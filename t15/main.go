package main

import (
	"crypto/rand"
	"fmt"
	_ "net/http/pprof"
)

// Генерируем случайную строку
func RandStringBytes(n int) []rune {
	const letterRune = "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	//Генерируем пустой массив случайных строк и разбиваем строку по элементам
	s, r := make([]rune, n), []rune(letterRune)

	for i := range s {
		//Получаем случайное число
		p, _ := rand.Prime(rand.Reader, len(r))
		//Конвертируем и получаем длину наших рун
		x, y := p.Uint64(), uint64(len(r))
		//Вносим в возращаемый слайс
		s[i] = r[x%y]
	}
	return s
}

var justString []rune

func someFunc() {
	v := RandStringBytes(1 << 10)
	//Копируем значение из слайса в нашу глобальную строку
	justString = append(justString, v[:100]...)
}

func main() {
	someFunc()
	fmt.Println(string(justString))
}
