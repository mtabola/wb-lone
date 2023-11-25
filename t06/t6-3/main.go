// Прекращение выполнения горутины с помощью interrupt сигнала
// (Graceful shutdown)
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan os.Signal, 1)

	// При выполнении горутины мы проверяем канал done. Если в него notify записал данные, то это значит, что программа была прервана
	go func(c chan int) {
		for {
			select {
			case c <- 1:
			case <-done:
				//закрываем поток и горутину
				close(c)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}(ch)

	//Cоздаем отдельную горутину для того, чтобы слушать interrupt сигнал
	go func(c chan os.Signal) {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}(done)

	for i := range ch {
		fmt.Println("Received: ", i)
	}

	fmt.Println("Finish")
}
