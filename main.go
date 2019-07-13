package main

import (
	"fmt"

	another "github.com/solairerove/go-pointers-playground/another"
	point "github.com/solairerove/go-pointers-playground/pointer"
	basic "github.com/solairerove/go-pointers-playground/simple"
)

func main() {
	fmt.Println("Pointers")

	basic.BasicPointer()
	another.JustCase()
	point.ViaPointer()
}
