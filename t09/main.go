// Описание конвеера и понимание того как оно работает было взято в основном из статьи Артема Ковардина
// https://kovardin.ru/articles/go/modeli-konkurentnosti-v-go/
package main

import "fmt"

func main() {

	arr := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

	fromArr := make(chan int)
	modified := make(chan int)

	go func() {
		//запускаем цикл, где перекидываем данные из массива в канал
		for _, v := range arr {
			fromArr <- v
		}
		//закрываем каналы для того, чтобы ничего не попало лишнего
		close(fromArr)
	}()
	go func() {
		//пока канал заполняется, мы уже начинаем читать из него и умножать эти значения на 2
		for v := range fromArr {
			modified <- v * 2
		}
		close(modified)
	}()
	//"слушаем" этот канал и по мере того, как туда заносятся данные, мы их выводим
	for v := range modified {
		fmt.Println(v)
	}

}
