// Пример демонстрирует использование интерфейсов для объединения данных различных типов.
// Интерфейсы в Go в моём понимании преследуют две цели:
// 1. Объединение - объединение типов на основе поведения (а не схожести данных)
// 2. Разделение - разделение объявления (контракта) и реализации (методов)
//
// В данном примере демонстрируется первое свойство интерфейсов.
// В примере объявляются два типа данных - связанный список и массив.
// Далее создаётся интерфейс с контрактом на метод печати.
// Для каждого типа реализуется интерфейс.
// В функции main переменные разных типов добавляются в массив типа интерфейс,
// таким образом объединяя их по поведению (метод print).
// После чего с переменными обоих типов работа идёт однотипно в цикле.
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

const maxElements = 10 // размер элементов в массиве

type linkedList struct {
	value int
	next  *linkedList
}

type intSlice []int

var myList *linkedList
var mySlice intSlice

type printer interface {
	print()
}

func (l *linkedList) print() {
	if l.next != nil {
		l.next.print()
	}
	fmt.Printf("%d : ", l.value)
}

func (i intSlice) print() {
	for j := 0; j < len(i); j++ {
		fmt.Printf("%d : ", i[j])
	}
}

func init() {

	rand.Seed(time.Now().Unix())
	var temp *linkedList

	for i := 0; i < maxElements; i++ {

		mySlice = append(mySlice, rand.Intn(100))

		var l linkedList
		l.value = rand.Intn(100)
		if i == 0 {
			temp = &l
			myList = &l
			continue
		}
		temp.next = &l
		temp = &l
	}

}

func main() {

	var dataSet []printer
	dataSet = append(dataSet, myList)
	dataSet = append(dataSet, mySlice)

	for _, data := range dataSet {
		fmt.Printf("\n")
		fmt.Println(reflect.TypeOf(data))
		data.print()
		fmt.Printf("\n")
	}

	// Output
	// *main.linkedList
	// 26 : 27 : 11 : 11 : 9 : 1 : 89 : 69 : 73 : 86 :
	//
	// main.intSlice
	// 53 : 79 : 93 : 5 : 78 : 70 : 82 : 63 : 4 : 53 :

}
