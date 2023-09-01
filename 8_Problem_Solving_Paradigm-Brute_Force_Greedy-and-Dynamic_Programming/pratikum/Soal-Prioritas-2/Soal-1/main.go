package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}

func Frog(jumps []int) int {
	/*
		Gunakan DP bottom up approach
		Kita gunakan konsep:
		Optimal[i] = minimum(optimal[i-1] + abs(jarak[i] - jarak[i-1]), optimal[i-2] + abs(jarak[i] - jarak[i-2]))
		karena katak punya pilihan untuk melakukan lompatan i+1 atau i+2
	*/
	optimal := []int{0, abs(jumps[1] - jumps[0])}

	for i := 2; i < len(jumps); i++ {
		optimal = append(optimal, min(optimal[i-1]+abs(jumps[i]-jumps[i-1]), optimal[i-2]+abs(jumps[i]-jumps[i-2])))
	}

	return optimal[len(jumps)-1]
}

func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 10})) // 40

}
