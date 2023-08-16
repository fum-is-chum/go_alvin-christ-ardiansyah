package main

import "fmt"

func main() {
	var a int64;
	var res string;

	fmt.Print("Masukkan bilangan: ");
	fmt.Scan(&a);
	
	if a % 2 == 0 {
		res = "genap"
	} else {
		res = "ganjil"
	}
	fmt.Println(fmt.Sprintf("Bilangan %d adalah %s", a, res))
}