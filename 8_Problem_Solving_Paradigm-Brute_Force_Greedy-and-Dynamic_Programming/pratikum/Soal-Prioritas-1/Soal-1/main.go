package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Input: n = ")
	fmt.Scanln(&n)

	result := []string{"0"}
	for i:= 1; i <= n; i++ {
		result = append(result, fmt.Sprintf("%b", i))
	}

	fmt.Println(result)
}
