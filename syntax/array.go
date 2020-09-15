package syntax

import "fmt"

func BasicArrayFuncs() {
	arr := []string{"imgomi", "korea", "student", "gopher"}
	arr = append(arr, "work")
	fmt.Println(arr)
}

func BasicMapStruct() {
	m := map[string]string{"imgomi": "student"}
	m["level"] = "MASTER"
	for k, v := range m {
		fmt.Println(k, v)
	}
}
