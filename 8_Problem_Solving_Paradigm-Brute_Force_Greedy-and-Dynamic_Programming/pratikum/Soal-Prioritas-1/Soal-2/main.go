package main

import "fmt"

func generatePascalTriangle(numRows int) (result [][]int) {
	if numRows >= 1 {
		result = append(result, []int{1})
	}

	if numRows >= 2 {
		result = append(result, []int{1, 1})
	}

	if numRows > 2 {
		for i := 2; i < numRows; i++ {
			// [1, ...]
			row := []int{1}

			for j := 1; j <= i-1; j++ {
				row = append(row, result[i-1][j-1]+result[i-1][j])
			}

			// [..., 1]
			row = append(row, 1)
			result = append(result, row)
		}
	}

	return result
}

func main() {
	var numRows int
	fmt.Print("numRows = ")
	fmt.Scanln(&numRows)

	fmt.Println(generatePascalTriangle(numRows))
}