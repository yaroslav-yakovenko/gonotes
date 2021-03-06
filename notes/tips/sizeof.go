package main

import "fmt"
import "unsafe"

func main() {
	a := int(123)
	a1 := 123
	b := int64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}
	e := struct{}{}
	f := struct{}{}
	g := []int{}
	var h []int
	var i []int = []int{1, 2, 3, 4, 5, 6} // размер как у "h"

	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("a1: %T, %d\n", a1, unsafe.Sizeof(a1))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d))
	fmt.Printf("e: %T, %p, %d\n", e, &e, unsafe.Sizeof(e))
	fmt.Printf("f: %T, %p, %d\n", f, &f, unsafe.Sizeof(f))
	fmt.Printf("g: %T, %d\n", g, unsafe.Sizeof(g))
	fmt.Printf("h: %T, %d\n", h, unsafe.Sizeof(h))
	fmt.Printf("i: %T, %d\n", i, unsafe.Sizeof(i))

	// Output:
	// a: int, 8
	// a1: int, 8
	// b: int64, 8
	// c: string, 16
	// d: struct { FieldA float32; FieldB string }, 24
	// e: struct {}, 0x5791c8, 0
	// f: struct {}, 0x5791c8, 0
	// g: []int, 24
	// h: []int, 24
	// i: []int, 24
}
