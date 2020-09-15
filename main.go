package main

import (
	"fmt"
	"github.com/aid95/learn-go/syntax"
)

func main() {
	fmt.Println(syntax.SuperAdd(1, 2, 3, 4, 5))
	fmt.Println(syntax.SuperAddRange(1, 2, 3, 4, 5))
}
