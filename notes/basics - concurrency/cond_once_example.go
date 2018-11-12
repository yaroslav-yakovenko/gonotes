// Пример использования sync.Cond и sync.Once.
//
// sync.Cond представляет собой блок и событие, которое можно использовать как управляющий сигнал для горутин.
// Содержит мюьтекс, который блокиррует при вызове метода Wait до наступления события.
// При этом блок необходимо предварительно установить явно, метод Wait блок не устанавливает.
// sync.Once используется для выполнения функции-аргумента строго один раз.
// Считается только количество вызовов once.Do. Изменение аргумента результата не даёт.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var once sync.Once
	var condition = sync.NewCond(&sync.Mutex{})
	var counter int

	incr := func() { counter++ }
	decr := func() { counter-- }

	for i := 0; i < 5; i++ {
		go func(i int) {
			condition.L.Lock()
			defer condition.L.Unlock()
			once.Do(incr)
			condition.Wait()
			fmt.Printf("Recieved condition event. Goroutine #%d exits\n", i)
		}(i)
	}

	// Ожидаем выполнение всех горутин. Конечно, лучше использовать каналы
	time.Sleep(1 * time.Second)

	condition.Broadcast()

	// Ожидаем завершение всех горутин
	time.Sleep(1 * time.Second)

	fmt.Printf("Counter is: %d\n", counter)

	once.Do(decr)

	fmt.Printf("After 'decr' counter is: %d\n", counter)

	// Output:
	// Recieved condition event. Goroutine #3 exits
	// Recieved condition event. Goroutine #2 exits
	// Recieved condition event. Goroutine #4 exits
	// Recieved condition event. Goroutine #1 exits
	// Recieved condition event. Goroutine #0 exits
	// Counter is: 1
	// After 'decr' counter is: 1

}
