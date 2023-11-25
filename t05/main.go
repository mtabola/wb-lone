package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandStringBytesRmndr(n int) string { // генератор случайной строки
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func main() {
	ch := make(chan string)
	// создаем таймер, который после завершения работы возвращает канал
	timer := time.After(5 * time.Second)

	//запускаем горутину, которая будет отправлять данные в канал
	go func() {
		for i := 1; ; i++ {
			select {
			//Пока таймер не истек записываем данные в основной канал
			case ch <- RandStringBytesRmndr(i):
				time.Sleep(1 * time.Second)
			//По истечению времени закрываем канал и горутину
			case <-timer:
				fmt.Println("Times up")
				close(ch)
				return
			}
		}
	}()

	// Читаем поступающие данные
	for v := range ch {
		fmt.Println(v)
	}
}
