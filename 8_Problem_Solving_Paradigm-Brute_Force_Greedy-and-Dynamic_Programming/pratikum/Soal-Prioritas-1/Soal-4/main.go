package main

import "fmt"

var memo map[int]bool

/* Time Complexity O(m) */
func faktorBilangan(num int) (result []int) {
	batas := num / 2
	for i := 1; i <= batas; i++ {
		if num%i == 0 {
			memo[i] = true // catat semua faktor dalam map agar bisa diakses dengan time complexity O(1)
			result = append(result, i)
		}
	}
	return result
}

/*
	Pendekatan bruteforce (optimized)
	Time Complexity O(m + n^2)
*/
func SimpleEquations(a, b, c int) {
	var answer []int
	// Cari faktor bilangan b
	faktor := faktorBilangan(b)  // x*y*z = b
	lenFaktor := len(faktor)

	answerFound := false
	for i := 0; i < lenFaktor; i++ {
		for j := i; j < lenFaktor; j++ {
			x := faktor[i]
			y := faktor[j]

			// prediksi nilai z
			z := a - (x + y) // x + y + z = a
			_, exist := memo[z]
			if !exist {
				// jika z bukan faktor b
				continue
			}

			if x*y*z != b {
				continue
			}

			if x*x+y*y+z*z != c {
				continue
			}

			answer = []int{x, y, z}
			answerFound = true
			break
		}

		if answerFound {
			break
		}
	}

	if answerFound {
		fmt.Println(answer)
	} else {
		fmt.Println("No Solution.")
	}

}

func main() {
	memo = make(map[int]bool)

	SimpleEquations(1, 2, 3) // no solution

	SimpleEquations(6, 6, 14) // 1 2 3
	SimpleEquations(5, 3, 11) // 1 1 3

}
