package syntax

import "fmt"

// base := 20 <--- ERROR! Only available in function.
var base int = 20 // Cool

func VariableSyntax() int {
	a := 0			// (1)
	var b = 0		// (2)
	var c int = 0	// (3)
	return base + a + b + c
}

func ConstantSyntax() {
	const name string = "I love golang..<3"
	fmt.Println(name)
}
