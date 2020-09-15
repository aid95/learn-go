package main

import (
	"fmt"
	"github.com/aid95/learn-go/syntax"
)

func main() {
	a := 0
	syntax.BasicPointerConcept(&a)
	fmt.Println(a)
}
