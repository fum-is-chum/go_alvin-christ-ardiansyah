package main

import (
	"fmt"
)


func munculSekali(angka string) []int {

	count := make(map[rune]int)
	var result []int
	 
	for _,chr := range angka {
		count[chr]++
	}

	for key, value := range count {
		if value == 1 {
			result = append(result, int(key - '0')) // konversi ascii ke angka 0 - 9
		}
	}

	return result
}


func main() {

// Test cases

fmt.Println(munculSekali("1234123")) // [4]

fmt.Println(munculSekali("76523752")) // [6 3]

fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

fmt.Println(munculSekali("1122334455")) // []

fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

}