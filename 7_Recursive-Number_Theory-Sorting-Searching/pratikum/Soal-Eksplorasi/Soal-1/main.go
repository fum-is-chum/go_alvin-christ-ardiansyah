package main

import (
	"fmt"
	"math"
)

func MaxSequence(arr []int) int {
	/*
		Iterate arr, jika ada arr[i] >= sum, maka reset sum = arr[i]
		Dan selalu bandingkan nilai sum dengan ans
	*/
	sum := 0
	ans := -math.MaxInt

	for i:= 0; i < len(arr); i++ {
		sum += arr[i]
		if arr[i] >= sum {
			sum = arr[i]
		}

		if sum > ans {
			ans = sum
		}

	}

	return ans
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}
