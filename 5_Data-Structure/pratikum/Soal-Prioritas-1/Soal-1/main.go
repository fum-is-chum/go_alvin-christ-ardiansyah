package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// your code here
	seen := make(map[string]bool)
	var result []string

	// iterate arrayA
	for _ ,element := range arrayA {
		if !seen[element] {
			seen[element] = true
			result = append(result, element)
		}
	}

	// iterate arrayB and skip seen element
	for _, element := range arrayB {
		if !seen[element] {
			result = append(result, element)
		}
	}

	return result
}


func main() {

// Test cases

fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

// ["sergei", "jin", "steve", "bryan"]

fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

// ["alisa", "yoshimitsu", "devil jin", "law"]

fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

// ["devil jin", "sergei"]

fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

// ["hwoarang"]

fmt.Println(ArrayMerge([]string{}, []string{}))

// []

}