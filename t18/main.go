package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count atomic.Uint64 // создаем атомарный счетчик

	var wg sync.WaitGroup // создаем wg для того, чтобы дождаться всех инкрементов

	for i := 0; i < 15; i++ {
		wg.Add(1) // инкрементируем счетчик ожидания выполнения горутины

		go func(inc uint64) { // сама горутина
			fmt.Println("Increase for goroutine", inc)
			count.Add(1) // инкрементируем наш общий атомарный счетчик
			wg.Done()    // декрементируем счетчик ожидания
		}(uint64(i))
	}

	wg.Wait() // дожидаемся, пока все горутины не закончат свое выполенение (пока счетчик wg не станет 0)
	fmt.Println("Done. Result = ", count.Load())
}
