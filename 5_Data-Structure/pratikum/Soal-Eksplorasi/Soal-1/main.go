package main

import (
	"fmt"
	"math"
)

func Selisih(matriks [3][3]int) float64 {
	/*
		Hanya perlu menghitung selisih kolom 1 dan kolom 3 dari baris ke 1 dan ke 3 matriks
	*/
	return math.Abs(float64(matriks[0][0] - matriks[0][2])) + math.Abs(float64(matriks[2][0] - matriks[2][2]))
}

func main() {
	/*
	1 2 3
	4 5 6
	9 8 9
	*/
	var matriks [3][3]int

	for i:= 0; i < 3; i++ {
		for j:= 0; j < 3; j++ {
			var input int
			fmt.Scan(&input)
			matriks[i][j] = input
		}
	}

	fmt.Println(Selisih(matriks))
}