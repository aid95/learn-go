package syntax

import "fmt"

func SuperAdd(numbers ...int) (result int) {
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return
}

func SuperAddRange(numbers ...int) (result int) {
	// '_' means a value that is not used.
	// Index does not use values here.
	for _, number := range numbers {
		result += number
	}
	return
}

// Private function, because it starts with a lower case.
func infinityLoop() {
	for {
		fmt.Println("DO NOT CALL THIS FUNCTION.")
	}
}
