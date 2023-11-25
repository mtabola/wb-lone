package main

import (
	"fmt"
	"time"
)

// Простое решение - использовать функию After, которая посылает сигнал в канал о завершении работы, после истечении t времени
func sleepWithTimeLib(t time.Duration) {
	<-time.After(t)
}

// Сложное решение - создвать цикл, который будет просто инкрементировать переменную пока она не достигнит t
func sleepWithLoop(t time.Duration) {
	var i int64
	delay := int64(t*3) + int64((18*int64(t))/100) //формула того, "точного" расчета задержки
	for i = 0; i < delay; i++ {
	}
}

func main() {

	fmt.Println("Start sleeping (sleepWithLoop)...")
	startTime := time.Now()
	sleepWithLoop(10 * time.Second)
	endTime := time.Now()

	executeTime := endTime.Sub(startTime)
	fmt.Println("Function executed by time: ", executeTime)

	fmt.Println("Start sleeping (sleepWithTimeLib)...")
	startTime = time.Now()
	sleepWithTimeLib(10 * time.Second)
	endTime = time.Now()

	executeTime = endTime.Sub(startTime)
	fmt.Println("Function executed by time: ", executeTime)
}
