package main

import (
	"fmt"
	"github.com/aid95/learn-go/syntax"
)

func main() {
	fmt.Println(syntax.Multiply(10, 10))

	l, n := syntax.LenAndUpper("imgomi")
	fmt.Println(l, n)

	fmt.Println(syntax.LenAndUpperV2("imgomi"))

	syntax.RepeatMe("Hello", "goland", "imgomi")
}
