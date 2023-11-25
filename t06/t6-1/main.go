// С использованием проверки закрытости кананлов
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // создаем канал для общения с горутинами

	go func(c chan int) {
		for {
			v, ok := <-c // берем данные и статус из канала и смотрим
			if !ok {     // если наш канал закрыт, то он возвращает статус false
				fmt.Println("Goroutine closed")
				return // а заначит мы имеем полное право прервать горутину
			}
			fmt.Println(v)
		}
	}(ch)

	ch <- 1
	ch <- 2
	time.Sleep(100 * time.Microsecond)
	ch <- 3
	close(ch)
	time.Sleep(1 * time.Second)
}
