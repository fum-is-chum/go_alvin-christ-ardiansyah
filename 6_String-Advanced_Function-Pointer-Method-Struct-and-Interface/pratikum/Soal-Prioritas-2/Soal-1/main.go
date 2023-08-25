package main

import (
	"fmt"
)

// Time Complexity: O(n)
func caesar(offset int, input string) string {
	result := make([]byte, len(input))

	for i,char := range input {
		if char == 32 {
			result[i] = 32
		} else {
			result[i] = 'a' + byte((int(char-'a')+offset)%26)
		}
	}

	return string(result)
}

func main() {

	fmt.Println(caesar(3, "abc def")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}
