package main

import (
	"fmt"
	"math"
)

/*
	Pengecekan bilangan prima dengan kompleksitas O(sqrt(N))
	Disini kita hanya perlu mengecek bilangan N terhadap bilangan dari 2 sampai sqrt(N)
	Jika ditemukan faktor dari bilangan N pada interval 2 sampai sqrt(N), return false
	Jika tidak, return true

	Alasan kenapa kita boleh mengecek hanya dari 2 sampai sqrt(N) dikarenakan faktor
	bilangan dari N pada suatu saat akan mencapai titik dimana faktor itu akan terbalik
	Contoh:
		Faktor dari 18 adalah:
		1 18
		2 9
		3 6
		6 3 -> ini adalah kebalikan dari 3 6, mulai dari sini semua adalah kebalikan
		9 2
		18 1

	Maka kita hanya perlu mengecek bilangan dari 2 sampai sqrt(18) = 4,24 -> 4
*/
func isPrimeNumber(number int) bool {
	if number < 2 {
		return false
	} else {
		akar := int(math.Sqrt(float64(number)))
		for i:= 2; i <= akar; i++ {
			if number % i == 0 {
				return false;
			}
		}
		return true;
	}
}

func main() {
	var number int;
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&number);

	if isPrimeNumber(number) {
		fmt.Println("Output: Bilangan Prima")
	} else {
		fmt.Println("Output: Bukan Bilangan Prima")
	}
}