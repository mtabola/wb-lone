//Закрытие горутин с помощью context

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//Если не создадим этот канал, то не сможем увидеть внутреннюю работу горутины
	ch := make(chan int)
	//Создаем контекст с Done каналом и cancel для вызова это done канала
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			// проверяем, не послала ли cancel функция или родительские контекст данные в этот канал
			case <-ctx.Done():
				//если все же послала, то в наш канал данные о закрытии и выходим из функции
				ch <- 1
				return
			default:
				//В остальных случаях просто слушаем
				fmt.Println("Worked")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func(c context.CancelFunc) { // вызвываем параллельно функцию закрытия контекста
		time.Sleep(5 * time.Second)
		c()
	}(cancel)

	fmt.Println(<-ch)
	fmt.Println("Finish")

}
