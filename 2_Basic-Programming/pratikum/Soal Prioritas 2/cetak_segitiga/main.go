package main

import "fmt"

func main() {
	var baris int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&baris)

	baris++
	for i := 1; i <= baris; i++ {
		for space := 1; space <= baris-i; space++ {
			fmt.Print(" ")
		}
		for j:= 0; j < 2*i-1; j++ {
			if j % 2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		} 
		fmt.Println("")
	}
}