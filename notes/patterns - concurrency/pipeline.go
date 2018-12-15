// Шаблон "Конвейер"
//
// Шаблон демонстрирует выполнение набора последовательных операций над потоком данных.
// Этапы конвейера принимают и отдают потоки одинакого типа.
// Это позволяет комбинировать шаги в произвольном порядке.
// Кроме того в примере продемонстрировано использование контекста для завершения работы горутин по таймауту.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// генератор - преобразует массив входных данных в поток одиночных значений
func generator(ctx context.Context, input []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for { // бесконечно повторяем исходные данные
			for _, i := range input {
				select {
				case <-ctx.Done():
					fmt.Println("Generator. Canceled by context")
					return
				default:
					ch <- i
				}
			}
		}
	}()
	return ch
}

// обработчик - этап конвейера, который обрабатывет данные, полученные от генератора
func processor(ctx context.Context, input <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range input {
			select {
			case <-ctx.Done():
				fmt.Println("Processor. Canceled by context")
				return
			default:
				// симуляция продолжительных вычислений
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- i * i // возвращаем квадрат числа
			}
		}
	}()
	return ch
}

func main() {
	// инициализируем исходные данные
	var in []int
	for i := 0; i < 100; i++ {
		in = append(in, i)
	}

	// таймаут выполнения конвейера
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	// запускаем конвейер и печатаем выходные данные
	for out := range processor(ctx, generator(ctx, in)) {
		fmt.Println(out)
	}

	// Output:
	// 0
	// 1
	// 4
	// 9
	// 16
	// 25
	// 36
	// 49
	// 64
	// 81
	// 100
	// Processor. Canceled by context
	// 121
	// Generator. Canceled by context

}
