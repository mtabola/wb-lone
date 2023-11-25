package main

import (
	"fmt"
	"sync"
)

func concurrentPow(v int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- v * v
}

func main() {
	vals := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	for i := 0; i < len(vals); i++ {
		wg.Add(1)
		go concurrentPow(vals[i], wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	//основной код был взят из предыдущего задания
	//здесь только добавлен цикл прохода по каналу и суммирование квадратов в результатирующую переменную
	res := 0
	for val := range ch {
		res += val
	}
	fmt.Println("Result is:", res)
}
