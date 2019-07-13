package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	a := 42
	fmt.Println(a)
	// address
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	// value stored at an address
	fmt.Println(*b)
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)
}
