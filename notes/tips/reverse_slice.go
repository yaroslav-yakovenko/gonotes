package main

import (
	"fmt"
)

func main() {

	data := []int{1, 2, 3, 5, 7, 9, 11, 13}

	reverse := func(input []int) {
		for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
			input[i], input[j] = input[j], input[i]
		}
	}

	fmt.Println("straight:", data)
	reverse(data)
	fmt.Println("reversed:", data)

	// Output:
	// straight: [1 2 3 5 7 9 11 13]
	// reversed: [13 11 9 7 5 3 2 1]

}
