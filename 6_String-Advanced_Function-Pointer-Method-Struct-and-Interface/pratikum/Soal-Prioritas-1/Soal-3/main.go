package main

import (
	"fmt"
	"strings"
)

// Worst Time Complexity: O(m + n*m)
func Compare(a, b string) string {
	var result string
	memo := make(map[rune][]int)

	lenA := len(a)
	lenB := len(b)

	// iterate string a and save its char position to memo
	for pos,char := range a {
		memo[char] = append(memo[char], pos)
	}
	
	// iterate string b and find substring that matches with string a
	for pos,char := range b {
		if len(memo[char]) > 0 { // char found inside memo

			// iterate through all pos of char in memo
			for _,startIdx := range memo[char] {
				idx := 0
				var localResult strings.Builder

				// compare char in string a and b, stop if char is not the same or index is out of bound
				for {
					if (startIdx + idx == lenA) || (pos + idx == lenB) {
						break
					}

					if a[startIdx + idx] == b[pos + idx] {
						localResult.WriteByte(a[startIdx + idx])
					} else {
						break
					}
					idx++
				}

				// fmt.Println(localResult.String())
				if len(localResult.String()) > len(result) {
					result = localResult.String()
				}
			}
		}
	}
	return result
}


func main() {

	fmt.Println(Compare("AKA", "AKASHI")) //AKA

	fmt.Println(Compare("KANGOORO", "KANG")) //KANG

	fmt.Println(Compare("KI", "KIJANG")) //KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU

	fmt.Println(Compare("ILALANG", "ILA")) //ILA

	// Corner Case
	fmt.Println(Compare("AKAABCDASHI", "AKASHI")) // ASHI

	fmt.Println(Compare("KANGOORO", "KANGABCDEFGOORABCDGOORO")) //GOORO

}