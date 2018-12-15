// Пример использования агрегатора каналов select
//
// Пример демонстрирует два аспекта работы конструкции select:
// 1. Чтение и запись данных при работе с несколькими каналами, что можно использовать, к примеру,
//    для обработки параллельных вычислений.
//    В примере несколько горутин выдают некоторый результат с переменной задержкой. Можно считать эти
//    горутины репликами некоторых вычислительных процессов, которые выдают результаты вычислений аггрегатору
//    (грубо - так работает поиск в Интернет).
// 2. Обработка управляющих сигналов (в примере закрытие канала done завершает работу горутин),
//    пока канал открыт горутины продолжают работать
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generator(done <-chan bool, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("Goroutine stopped")
			return
		case ch <- rand.Intn(1000):
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}
}

func main() {

	rand.Seed(time.Now().Unix())

	channels := make([]chan int, 3)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		channels[i] = make(chan int)
		defer close(channels[i])
		go generator(channels[i], done, &wg)
	}

	deadline := time.After(2 * time.Second)

	out := false
	for {
		if out {
			break
		}
		select {
		case n := <-channels[0]:
			fmt.Println("Recived from channel 1: ", n)
		case n := <-channels[1]:
			fmt.Println("Recived from channel 2: ", n)
		case n := <-channels[2]:
			fmt.Println("Recived from channel 3: ", n)
		case <-deadline:
			close(done)
			fmt.Println("Exit on timeout")
			out = true
		}
	}

	wg.Wait()

	// Output:
	// ...
	// Recived from channel 3:  691
	// Recived from channel 3:  151
	// Recived from channel 1:  338
	// Recived from channel 2:  96
	// Recived from channel 2:  686
	// Exit on timeout
	// Goroutine stopped
	// Goroutine stopped
	// Goroutine stopped

}
