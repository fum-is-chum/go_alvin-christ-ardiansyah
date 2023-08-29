package main

import "fmt"

var memo map[int]int

func fibonacci(number int) int {
	if number < 0 {
		return 0
	}

	if number == 2 || number == 1 {
		return 1
	}

	if memo[number] > 0 {
		return memo[number]
	}

	f1 := fibonacci(number - 1)
	f2 := fibonacci(number - 2)

	memo[number-1] = f1
	memo[number-2] = f2
	memo[number] = f1+f2

	return f1 + f2
}

func main() {
	// use memoization to memoize result
	memo = make(map[int]int)

	fmt.Println(fibonacci(0)) // 0

	fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144

}
