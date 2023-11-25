//Конурентную запись в мапу можно реализовать 2 способами

package main

import (
	"fmt"
	"sync"
	"time"
)

// 1) С помощью RWMutex
func RWMutexRecord() {
	//Создаем наш мьютекс
	var mu sync.RWMutex
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		//записываем данные
		go func(i int) {
			//блокируем мьютекс, чтобы записать данные
			mu.Lock()
			//ставим его в очередь на разблокировку после завершения горутины
			defer mu.Unlock()
			//производим запись
			m[i] = i + 1
		}(i)
	}

	//Слип тут нужен для того, чтобы все горутины успели отработать
	time.Sleep(5 * time.Second)
	// делам тоже самое, что и на запись, только использумем мьютекс для чтения
	for i := 0; i < 10; i++ {
		func(i int) {
			mu.RLock()
			defer mu.RUnlock()
			fmt.Printf("%d - %d\n", i, m[i])
		}(i)
	}

}

// 2) C помощью функции Map из пакета sync
func RWMapRecord() {
	//создаем нашу мапу
	syncMap := sync.Map{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			//и использумем уже встроенные методы записи и чтения для того, чтобы работать с нашей мапой
			syncMap.Store(i, i+1)
		}(i)
	}
	//Слип тут нужен для того, чтобы все горутины успели отработать
	time.Sleep(5 * time.Second)

	for i := 0; i < 10; i++ {
		func(i int) {
			v, _ := syncMap.Load(i)
			fmt.Printf("%d - %d\n", i, v)
		}(i)
	}
}

func main() {
	fmt.Println("RwMutex")
	RWMutexRecord()
	fmt.Printf("\nsyncMap\n")
	RWMapRecord()
}
