package another

import (
	"fmt"
)

// JustCase tbd
func JustCase() {
	fmt.Println("\nJustCase")

	x := 42
	fmt.Println(x)
	foo(x)
	fmt.Println(x)
}

func foo(x int) {
	fmt.Println(x)
	x = 43
	fmt.Println(x)
}
