// Пример синхронизации работы двух горутин "Пинг-понг"
//
// Пример использует канал для синхронизации работы. Каждому игроку соответствует горутина.
// Подача начинается командой "begin", дальше произвольный игрок подаёт и пишет в канал "ping".
// Второй игрок отвечает "pong".
// Случайным образом игроку может повезти и он загасит шарик и выиграет очко.
// В этом случае в канал подаётся сигнал "stop", чтобы вторая горутина завершила работу и не возникло клинча.
// Вместе с сигналом Стоп горутина завершает работу, и, поскольку он обрабатывается первым, второй игрок тоже останавливается.
// Карта (map) используется для ведения счёта. Количество геймов определяется счётчиком цикла.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	score := make(map[string]int) // счёт встречи

	var wg sync.WaitGroup // ожидаем завершения горутин

	// функция, имитирующая игрока
	player := func(ch chan string, name string) {
		defer wg.Done()
		for {
			cmd := <-ch
			if cmd == "stop" {
				return
			}
			time.Sleep(time.Duration(rand.Int63n(300)) * time.Millisecond)
			if rand.Intn(100) > 80 {
				fmt.Println(name + " smashes the ball and wins the point!\n")
				score[name] += 1
				ch <- "stop"
				return
			}
			switch cmd {
			case "begin":
				{
					fmt.Println(name + " is serving")
					ch <- "ping"
				}
			case "ping":
				{
					fmt.Println(name + " pong")
					ch <- "pong"
				}
			case "pong":
				{
					fmt.Println(name + " ping")
					ch <- "ping"
				}
			}
		}
	}

	ch := make(chan string) // канал для коммуникаций и синхронизации
	// запускаем матч из нескольких геймов, порядок подачи случайный
	for i := 0; i < 3; i++ {
		wg.Add(2)
		go player(ch, "Player1")
		go player(ch, "Player2")
		ch <- "begin"
		wg.Wait()
	}

	// печатаем счёт
	fmt.Println("\nScore is:", score["Player1"], " : ", score["Player2"])

	// Output:
	// Player2 is serving
	// Player1 pong
	// Player2 ping
	// Player1 pong
	// Player2 ping
	// Player1 smashes the ball and wins the point!
	//
	// Player2 is serving
	// Player1 pong
	// Player2 smashes the ball and wins the point!
	//
	// Player2 is serving
	// Player1 pong
	// Player2 ping
	// Player1 pong
	// Player2 smashes the ball and wins the point!
	//
	//
	// Score is: 1  :  2

}
