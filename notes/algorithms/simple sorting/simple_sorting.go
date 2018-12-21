// Пример реализации простейших алгоритмов сортировки со сравнением времени выполнения.
//
// В программе представлены реализации двух простейших алгоритмов сортировки - пузырьковой и слиянием.
// Сортировка производится над большим масссивом случайных чисел. Во время каждой сортировки производится замер времени.
// В результате хорошо видна разница в скорости алгоритмов.
// *** Асимптотическая сложность: Пузырьковая сортировка - O(n^2), Сортировка слиянием - O(n*log2(n))
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Инициализируем исходный мсходный массив случайными числами [0 - 10000]
func initData() (data []int) {

	rand.Seed(time.Now().Unix())
	for i := 0; i < 10000; i++ {
		data = append(data, rand.Intn(10000))
	}

	return data

}

// Пузырьковая сортировка. Асимптотическая сложность - O(n^2))
func bubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if data[j] < data[i] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

// Сортировка слиянием. Асимптотическая сложность - O(n*log2(n))
func mergeSort(data []int) {
	copy(data, msort(data))
}

// рекурсивная сортировка
func msort(data []int) (sorted []int) {
	if len(data) == 1 {
		return data
	}
	index := len(data) / 2
	left := make([]int, index)
	right := make([]int, len(data)-index)
	copy(left, data[:index])
	copy(right, data[index:])
	left = msort(left)
	right = msort(right)
	return mmerge(left, right)
}

// слияние
func mmerge(a, b []int) (c []int) {

	for (len(a) > 0) && (len(b) > 0) {
		if a[0] < b[0] {
			c = append(c, a[0])
			a = a[1:]
		} else {
			c = append(c, b[0])
			b = b[1:]
		}
	}

	if len(a) > 0 {
		c = append(c, a...)
	}

	if len(b) > 0 {
		c = append(c, b...)
	}

	return c
}

func main() {
	data := initData()
	d1 := make([]int, len(data))
	d2 := make([]int, len(data))
	copy(d1, data)
	copy(d2, data)

	fmt.Println("First 30 elements of source array:")
	for i := 0; i < 30; i++ {
		fmt.Print(d2[i], " ")
	}
	fmt.Println()

	start := time.Now()
	bubbleSort(d1)
	microsecs := time.Since(start).Nanoseconds() / 1000
	fmt.Println("Bubble Sort. Spent time in microseconds:", microsecs, "\nFirst 30 elements:")
	for i := 0; i < 30; i++ {
		fmt.Print(d1[i], " ")
	}
	fmt.Println()

	start = time.Now()
	mergeSort(d2)
	microsecs = time.Since(start).Nanoseconds() / 1000
	fmt.Println("Merge Sort. Spent time in microseconds:", microsecs, "\nFirst 30 elements:")
	for i := 0; i < 30; i++ {
		fmt.Print(d2[i], " ")
	}
	fmt.Println()

	// Output
	// First 30 elements of source array:
	// 9684 8053 197 5928 7685 4457 8675 2275 5676 39 8403 3172 3665 7097 2048 5417 7144 886 9194 3680 7038 3506 7190 7559 7482 8425 8263 3054 3284 400
	// Bubble Sort. Spent time in microseconds: 157045
	// First 30 elements:
	// 1 1 3 5 7 8 8 13 14 14 16 17 18 18 21 22 23 24 24 25 25 26 27 27 28 29 31 31 34 34
	// Merge Sort. Spent time in microseconds: 4990
	// First 30 elements:
	// 1 1 3 5 7 8 8 13 14 14 16 17 18 18 21 22 23 24 24 25 25 26 27 27 28 29 31 31 34 34

}
