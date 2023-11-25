package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ch chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		//роботяга проверяет, закрылся ли контекст или нет
		case <-ctx.Done():
			fmt.Printf("Worker at number %d are finished\n", id)
			return
		case <-ch:
			//Если нет, то запускаем еще 1 доп проверку, чтобы после закрытия канала у нас не шли "мусорные" данные
			//т.к без этой проверки у нас роботяга будет читать закрытый поток
			if _, ok := <-ch; ok {
				fmt.Println("Worker at number ", id, " : ", <-ch)
			}
		}
		time.Sleep(time.Second)
	}

}

func main() {
	var workers int
	flag.IntVar(&workers, "w", 3, "Quantity of workers")
	flag.Parse()

	ch := make(chan int)
	//Создаем контекст для прерывания (нужен он для того, чтобы завершить работу всех воркеров)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)

	// ВГ нужна нам для того, чтоб синхронно закончить работу всех роботяг
	wg := sync.WaitGroup{}

	//Запускаем наших роботяг, которые будут отдавать последние данные из канала (Которые они захватили на тот момент)
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, ch, ctx, &wg)
	}

	//Создаем бесконечный цикл для записи в канал
	for i := 0; ; i += 10 {
		select {
		case <-ctx.Done(): // Прослушиваем Done метод для того, чтобы отслеживать нажатие Ctrl+C
			cancel()  //Вызывем функцию для закрытия контекста
			close(ch) // Закрываем наш канал
			wg.Wait() // Дожидаемся когда роботяги закончат работу
			fmt.Println("Programm is interrupted")
			return
		default:
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
	Даннный метод закрытия всех воркеров был выбран потому-что:
	* позволяет безопасно выйти из всех роботяг (корректно завершается выполнение функции);
	* он прост в освоении;
	* все каналы закрываются явно
*/
