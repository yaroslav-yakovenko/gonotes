// Параллельные HTTP GET запросы
//
// В примере демонстрируется использование клиента HTTP, типов данных "запрос", "ответ" для создания простого
// парсера данных сайта букмекерской конторы 1ХСтавка.
// В примере сначала загружается список событий, идущих в режиме лайв.
// После чего для каждого события стартует отдельная горутина с запросом информации по данному событию.
// Полученные данные выводятся на экран.
// ** При увеличении числа одновременных запросов срабатывает защита от DDOS от фирмы Qrator Labs.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// ссылки получены опытным путём
const eventsListURL = "https://1xstavka.ru/LiveFeed/Get1x2_VZip?count=10&mode=4&country=1"
const eventURL = "https://1xstavka.ru/LiveFeed/GetGameZip?lng=ru&id="

// структуры данных для десериализации данных о текущих событиях с сайта 1ХСтавка
// структуры содержат лишь небольшую часть полей для демонстрации (id события, турнир, участники)
// состав структур данных и типы полей получены на  основе анализа JSON-файлов
type liveEvents struct {
	Value []value
}
type value struct {
	I int
	L string
}

type eventData struct {
	Value eventValue
}
type eventValue struct {
	L  string
	O1 string
	O2 string
}

func main() {
	var events liveEvents
	resp, err := http.Get(eventsListURL)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		log.Fatal(err)
	}

	// The Client's Transport typically has internal state (cached TCP connections),
	// so Clients should be reused instead of created as needed.
	// Clients are safe for concurrent use by multiple goroutines.
	client := &http.Client{}

	var wg sync.WaitGroup
	wg.Add(len(events.Value))
	for i, v := range events.Value {
		go func(num, id int) {
			defer wg.Done()
			var event eventData
			req, err := http.NewRequest("GET", eventURL+strconv.Itoa(id), nil)
			if err != nil {
				log.Printf("%d: %#v\n", id, err)
				return
			}

			req.Header.Add("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				log.Printf("%d: %#v\n", id, err)
				return
			}

			err = json.NewDecoder(resp.Body).Decode(&event)
			if err != nil {
				log.Println(resp)
				log.Printf("%d: %#v\n", id, err)
				return
			}

			log.Printf("goroutine %d\t%s. %s : %s\n", num, event.Value.L, event.Value.O1, event.Value.O2)
		}(i, v.I)
	}

	wg.Wait()

	//  Output:
	//  2018/12/15 15:48:48 goroutine 0	Чемпионат Англии. Премьер-лига. Манчестер Сити : Эвертон
	//  2018/12/15 15:48:48 goroutine 5	Открытый чемпионат России. ВХЛ. Рубин : Саров
	//  2018/12/15 15:48:48 goroutine 4	Открытый чемпионат России. ВХЛ. Ермак : Зауралье
	//  2018/12/15 15:48:48 goroutine 8	Открытый чемпионат России. МХЛ. Алмаз : СКА-1946
	//  2018/12/15 15:48:48 goroutine 1	Чемпионат Испании. Примера. Хетафе : Реал Сосьедад
	//  2018/12/15 15:48:48 goroutine 2	Евротур. Кубок Первого Канала. Россия : Чехия
	//  2018/12/15 15:48:48 goroutine 6	Открытый чемпионат России. ВХЛ. Торос : Звезда
	//  2018/12/15 15:48:48 goroutine 7	Чемпионат Турции. Суперлига. Бахчешехир Колежи : Бешикташ
	//  2018/12/15 15:48:48 goroutine 3	Открытый чемпионат России. ВХЛ. Югра : Тамбов
	//  2018/12/15 15:48:48 goroutine 9	Чемпионат Германии. 2-я Бундеслига. Падерборн : Динамо Дрезден

}
