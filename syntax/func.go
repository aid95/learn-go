package syntax

import (
	"fmt"
	"strings"
)

// func Multiply(a int, b int) int {
func Multiply(a, b int) int {
	return a * b
}

// Multiple return values
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// Naked return, defer example
func LenAndUpperV2(name string) (length int, uppercase string) {
	defer fmt.Println("The function ended successfully.")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func Swap(a, b int) (int, int) {
	return b, a
}

// Multiple argument values
func RepeatMe(words ...string) {
	fmt.Println(words)
}
