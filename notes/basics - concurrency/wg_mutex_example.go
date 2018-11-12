// Пример использования sync.WaitGroup и sync.Mutex.
//
// sync.WaitGroup позволяет дождаться завершения произворльного количества конкурентных функций.
// В даннном примере видно, что порядок выполнения конкурентных функций отличается от порядка их создания.
// Кроме того в примере использован мьютекс для обеспечения атомарного доступа к разделяемой памяти (сумма чисел) и
// избежания ситуации "race condition".
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	var sum struct {
		val int
		mu  sync.Mutex
	}

	var wg sync.WaitGroup

	f := func(i int) {
		defer wg.Done()

		n := rand.Intn(100)
		fmt.Printf("Goroutine №%d generated random num: %d\n", i, n)

		sum.mu.Lock()
		// код из блока defer сработает даже при панике в функции, иначе в случае паники - deadlock
		defer sum.mu.Unlock()
		sum.val += n
	}

	const MAXFUNCS = 10
	wg.Add(MAXFUNCS)

	for i := 0; i < MAXFUNCS; i++ {
		// значение переменной вычисляеся во время выполнения (а не создания) горутины и не детерменированно
		// поэтому передаём копию счётчика как аргумент, а не используем перменную из котекста замыкания
		go f(i)
	}

	wg.Wait()

	fmt.Println("Sum is: ", sum.val)

	// Output:
	// Goroutine №2 generated random num: 89
	// Goroutine №6 generated random num: 73
	// Goroutine №8 generated random num: 29
	// Goroutine №3 generated random num: 53
	// Goroutine №9 generated random num: 34
	// Goroutine №7 generated random num: 24
	// Goroutine №0 generated random num: 94
	// Goroutine №1 generated random num: 23
	// Goroutine №4 generated random num: 82
	// Goroutine №5 generated random num: 7
	// Sum is:  508

}
