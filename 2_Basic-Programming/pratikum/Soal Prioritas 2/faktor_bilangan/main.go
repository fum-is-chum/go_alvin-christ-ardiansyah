package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		fmt.Print("Bilangan tidak boleh negatif")
	} else {
		batas:= bilangan / 2
		for i:= 1; i <= batas; i++ {
			if bilangan % i == 0 {
				fmt.Println(i)
			}
		}
		fmt.Println(bilangan)
	}

}