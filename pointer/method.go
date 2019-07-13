package pointer

import (
	"fmt"
)

// ViaPointer tbd
func ViaPointer() {
	fmt.Println("\nViaPointer")

	x := 42
	fmt.Println(x)
	foo(&x)
	fmt.Println(x)
}

func foo(x *int) {
	fmt.Println(*x)
	*x = 43
	fmt.Println(*x)
}
