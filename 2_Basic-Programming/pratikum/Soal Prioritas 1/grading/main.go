package main

import "fmt"

func main() {
	var nilai int32;
	var grade rune;

    fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	switch {
		case nilai >= 80 && nilai <= 100:
			grade = 'A'
		case nilai >= 65 && nilai <= 79:
			grade = 'B'
		case nilai >= 50 && nilai <= 64:
			grade = 'C'
		case nilai >= 35 && nilai <= 49:
			grade = 'D'
		case nilai >= 0 && nilai <= 34:
			grade = 'E'
		default:
			grade = '-'
	}

	if grade == '-' {
		fmt.Println("Nilai Invalid")
	} else {
		fmt.Println("Grade anda: " + string(grade))
	}
}